package cache

import (
	"context"
	stdsql "database/sql"
	"database/sql/driver"
	"errors"
	"fmt"
	"strings"
	"sync/atomic"
	"time"
	_ "unsafe"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"github.com/go-redis/redis/v8"
	"github.com/mitchellh/hashstructure"
)

type (
	// Options wraps the basic configuration cache options.
	Options struct {
		// TTL defines the period of time that an Entry
		// is valid in the cache.
		TTL time.Duration

		// Cache defines the GetAddDeleter (cache implementation)
		// for holding the cache entries. If no cache implementation
		// was provided, an LRU cache with no limit is used.
		Cache AddGetDeleter

		// Hash defines an optional Hash function for converting
		// a query and its arguments to a cache key. If no Hash
		// function was provided, the DefaultHash is used.
		Hash func(query string, args []interface{}) (Key, error)

		// Logf function. If provided, the Driver will call it with
		// errors that can not be handled.
		Log func(...interface{})
	}

	// Option allows configuring the cache
	// driver using functional options.
	Option func(*Options)

	// A Driver is an SQL cached client. Users should use the
	// constructor below for creating new driver.
	Driver struct {
		dialect.Driver
		*Options
		stats Stats
	}
)

// NewDriver returns a new Driver an existing driver and optional
func NewDriver(drv dialect.Driver, c CacheConfig) (*Driver, *Redis) {
	Levels := make([]AddGetDeleter, 0)
	if c.MaxEntries > -1 {
		Levels = append(Levels, NewLRU(c.MaxEntries))
	}
	redis := NewRedis(redis.NewClient(&redis.Options{
		DB:       c.Db,
		Addr:     c.Addr,
		Network:  c.Network,
		Username: c.Username,
		Password: c.Password,
		PoolSize: c.PoolSize,
	}).WithTimeout(time.Second * 2))
	_, err := redis.C.Ping(context.Background()).Result()
	if err == nil {
		Levels = append(Levels, redis)
	}
	return &Driver{
		Driver: drv,
		Options: &Options{
			Hash:  DefaultHash,
			Cache: getLevels(Levels...),
			TTL:   time.Duration(c.TTL) * time.Second,
		},
	}, redis
}

// Hash configures an optional Hash function for
// converting a query and its arguments to a cache key.
func Hash(hash func(query string, args []interface{}) (Key, error)) Option {
	return func(o *Options) {
		o.Hash = hash
	}
}

// Levels configures the Driver to work with the given cache levels.
// For example, in process LRU cache and a remote Redis cache.
func getLevels(levels ...AddGetDeleter) AddGetDeleter {
	if len(levels) == 1 {
		return levels[0]
	} else {
		return &multiLevel{levels: levels}
	}
}

// ContextLevel configures the driver to work with context/request level cache.
// Users that use this option, should wraps the *http.Request context with the
// cache value as follows:
//
//	ctx = entcache.NewContext(ctx)
//
//	ctx = entcache.NewContext(ctx, entcache.NewLRU(128))
//
func ContextLevel() Option {
	return func(o *Options) {
		o.Cache = &contextLevel{}
	}
}

// monitor start Tx then clear cache
func (d *Driver) Tx(ctx context.Context) (dialect.Tx, error) {
	go d.Cache.Clear(ctx)
	return d.Driver.Tx(ctx)
}

// monitor update and delete then clear cache
func (d *Driver) Exec(ctx context.Context, query string, args, v interface{}) error {
	if !strings.HasPrefix(query, "SELECT") && !strings.HasPrefix(query, "select") {
		go d.Cache.Clear(ctx)
	}
	return d.Driver.Exec(ctx, query, args, v)
}

// Query implements the Querier interface for the driver. It falls back to the
// underlying wrapped driver in case of caching error.
//
// Note that, the driver does not synchronize identical queries that are executed
// concurrently. Hence, if 2 identical queries are executed at the ~same time, and
// there is no cache entry for them, the driver will execute both of them and the
// last successful one will be stored in the cache.
func (d *Driver) Query(ctx context.Context, query string, args, v interface{}) error {
	// Check if the given statement looks like a standard Ent query (e.g. SELECT).
	// Custom queries (e.g. CTE) or statements that are prefixed with comments are
	// not supported. This check is mainly necessary, because PostgreSQL and SQLite
	// may execute insert statement like "INSERT ... RETURNING" using Driver.Query.
	if !strings.HasPrefix(query, "SELECT") && !strings.HasPrefix(query, "select") {
		go d.Cache.Clear(ctx)
		return d.Driver.Query(ctx, query, args, v)
	}
	vr, ok := v.(*sql.Rows)
	if !ok {
		return fmt.Errorf("entcache: invalid type %T. expect *sql.Rows", v)
	}
	argv, ok := args.([]interface{})
	if !ok {
		return fmt.Errorf("entcache: invalid type %T. expect []interface{} for args", args)
	}
	opts, err := d.optionsFromContext(ctx, query, argv)
	if err != nil {
		return d.Driver.Query(ctx, query, args, v)
	}
	atomic.AddUint64(&d.stats.Gets, 1)
	switch scanner, err := d.Cache.Get(ctx, opts.key); {
	case err == nil:
		atomic.AddUint64(&d.stats.Hits, 1)
		vr.ColumnScanner = &repeater{columns: scanner.Columns, values: scanner.Values}
	case err == ErrNotFound:
		if err := d.Driver.Query(ctx, query, args, vr); err != nil {
			return err
		}
		vr.ColumnScanner = &recorder{
			ColumnScanner: vr.ColumnScanner,
			onClose: func(columns []string, values [][]driver.Value) {
				err := d.Cache.Add(ctx, opts.key, Entry{Columns: columns, Values: values}, opts.ttl)
				if err != nil && d.Log != nil {
					atomic.AddUint64(&d.stats.Errors, 1)
					d.Log(fmt.Sprintf("entcache: failed storing entry %v in cache: %v", opts.key, err))
				}
			},
		}
	default:
		return d.Driver.Query(ctx, query, args, v)
	}
	return nil
}

// Stats returns a copy of the cache statistics.
func (d *Driver) Stats() Stats {
	return Stats{
		Gets:   atomic.LoadUint64(&d.stats.Gets),
		Hits:   atomic.LoadUint64(&d.stats.Hits),
		Errors: atomic.LoadUint64(&d.stats.Errors),
	}
}

// errSkip tells the driver to skip cache layer.
var errSkip = errors.New("entcache: skip cache")

// optionsFromContext returns the injected options from the context, or its default value.
func (d *Driver) optionsFromContext(ctx context.Context, query string, args []interface{}) (ctxOptions, error) {
	var opts ctxOptions
	if c, ok := ctx.Value(ctxOptionsKey).(*ctxOptions); ok {
		opts = *c
	}
	if opts.key == nil {
		key, err := d.Hash(query, args)
		if err != nil {
			return opts, errSkip
		}
		opts.key = key
	}
	if opts.ttl == 0 {
		opts.ttl = d.TTL
	}
	if opts.evict {
		if err := d.Cache.Del(ctx, opts.key); err != nil {
			return opts, err
		}
	}
	if opts.skip {
		return opts, errSkip
	}
	return opts, nil
}

// DefaultHash provides the default implementation for converting
// a query and its argument to a cache key.
func DefaultHash(query string, args []interface{}) (Key, error) {
	key, err := hashstructure.Hash(struct {
		Q string
		A []interface{}
	}{
		Q: query,
		A: args,
	}, nil)
	if err != nil {
		return nil, err
	}
	return key, nil
}

// Stats represents the cache statistics of the driver.
type Stats struct {
	Gets   uint64
	Hits   uint64
	Errors uint64
}

// rawCopy copies the driver values by implementing
// the sql.Scanner interface.
type rawCopy struct {
	values []driver.Value
}

func (c *rawCopy) Scan(src interface{}) error {
	c.values[0] = src
	c.values = c.values[1:]
	return nil
}

// recorder represents an sql.Rows recorder that implements
// the entgo.io/ent/dialect/sql.ColumnScanner interface.
type recorder struct {
	sql.ColumnScanner
	values  [][]driver.Value
	columns []string
	done    bool
	onClose func([]string, [][]driver.Value)
}

// Next wraps the underlying Next method
func (r *recorder) Next() bool {
	hasNext := r.ColumnScanner.Next()
	r.done = !hasNext
	return hasNext
}

// Scan copies database values for future use (by the repeater)
// and assign them to the given destinations using the standard
// database/sql.convertAssign function.
func (r *recorder) Scan(dest ...interface{}) error {
	values := make([]driver.Value, len(dest))
	args := make([]interface{}, len(dest))
	c := &rawCopy{values: values}
	for i := range args {
		args[i] = c
	}
	if err := r.ColumnScanner.Scan(args...); err != nil {
		return err
	}
	for i := range values {
		if err := convertAssign(dest[i], values[i]); err != nil {
			return err
		}
		//处理中文[]uint8乱码问题
		if temp, ok := values[i].([]uint8); ok {
			values[i] = string(temp)
		}
	}
	r.values = append(r.values, values)
	return nil
}

// Columns wraps the underlying Column method and stores it in the recorder state.
// The repeater.Columns cannot be called if the recorder method was not called before.
// That means, raw scanning should be identical for identical queries.
func (r *recorder) Columns() ([]string, error) {
	columns, err := r.ColumnScanner.Columns()
	if err != nil {
		return nil, err
	}
	r.columns = columns
	return columns, nil
}

func (r *recorder) Close() error {
	if err := r.ColumnScanner.Close(); err != nil {
		return err
	}
	// If we did not encounter any error during iteration,
	// and we scanned all rows, we store it on cache.
	if err := r.ColumnScanner.Err(); err == nil || r.done {
		r.onClose(r.columns, r.values)
	}
	return nil
}

// repeater repeats columns scanning from cache history.
type repeater struct {
	columns []string
	values  [][]driver.Value
}

func (*repeater) Close() error {
	return nil
}
func (*repeater) ColumnTypes() ([]*stdsql.ColumnType, error) {
	return nil, fmt.Errorf("entcache.ColumnTypes is not supported")
}
func (r *repeater) Columns() ([]string, error) {
	return r.columns, nil
}
func (*repeater) Err() error {
	return nil
}
func (r *repeater) Next() bool {
	return len(r.values) > 0
}
func (r *repeater) NextResultSet() bool {
	return len(r.values) > 0
}

func (r *repeater) Scan(dest ...interface{}) error {
	if !r.Next() {
		return stdsql.ErrNoRows
	}
	for i, src := range r.values[0] {
		if err := convertAssign(dest[i], src); err != nil {
			return err
		}
	}
	r.values = r.values[1:]
	return nil
}

//go:linkname convertAssign database/sql.convertAssign
func convertAssign(dest, src interface{}) error
