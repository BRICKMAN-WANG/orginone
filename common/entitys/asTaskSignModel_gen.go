// Code generated by goctl. DO NOT EDIT!

package entitys

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	asTaskSignFieldNames          = builder.RawFieldNames(&AsTaskSign{})
	asTaskSignRows                = strings.Join(asTaskSignFieldNames, ",")
	asTaskSignRowsExpectAutoSet   = strings.Join(stringx.Remove(asTaskSignFieldNames, "`create_time`", "`update_time`"), ",")
	asTaskSignRowsWithPlaceHolder = strings.Join(stringx.Remove(asTaskSignFieldNames, "`id`", "`create_time`", "`update_time`"), "=?,") + "=?"

	cacheAssetAsTaskSignIdPrefix = "cache:asset:asTaskSign:id:"
)

type (
	asTaskSignModel interface {
		Insert(ctx context.Context, data *AsTaskSign) (sql.Result, error)
		FindOne(ctx context.Context, id string) (*AsTaskSign, error)
		Update(ctx context.Context, data *AsTaskSign) error
		Delete(ctx context.Context, id string) error
	}

	defaultAsTaskSignModel struct {
		sqlc.CachedConn
		table string
	}

	AsTaskSign struct {
		Id                  string         `db:"id"`
		TaskId              string         `db:"task_id"`
		Approve             sql.NullInt64  `db:"approve"`
		ApproveExecutors    sql.NullString `db:"approve_executors"`
		Disapprove          sql.NullInt64  `db:"disapprove"`
		DisapproveExecutors sql.NullString `db:"disapprove_executors"`
		AllExecutors        sql.NullInt64  `db:"all_executors"`
	}
)

func newAsTaskSignModel(conn sqlx.SqlConn, c cache.CacheConf) *defaultAsTaskSignModel {
	return &defaultAsTaskSignModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`as_task_sign`",
	}
}

func (m *defaultAsTaskSignModel) Insert(ctx context.Context, data *AsTaskSign) (sql.Result, error) {
	assetAsTaskSignIdKey := fmt.Sprintf("%s%v", cacheAssetAsTaskSignIdPrefix, data.Id)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?)", m.table, asTaskSignRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.Id, data.TaskId, data.Approve, data.ApproveExecutors, data.Disapprove, data.DisapproveExecutors, data.AllExecutors)
	}, assetAsTaskSignIdKey)
	return ret, err
}

func (m *defaultAsTaskSignModel) FindOne(ctx context.Context, id string) (*AsTaskSign, error) {
	assetAsTaskSignIdKey := fmt.Sprintf("%s%v", cacheAssetAsTaskSignIdPrefix, id)
	var resp AsTaskSign
	err := m.QueryRowCtx(ctx, &resp, assetAsTaskSignIdKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", asTaskSignRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, id)
	})
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultAsTaskSignModel) Update(ctx context.Context, data *AsTaskSign) error {
	assetAsTaskSignIdKey := fmt.Sprintf("%s%v", cacheAssetAsTaskSignIdPrefix, data.Id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, asTaskSignRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.TaskId, data.Approve, data.ApproveExecutors, data.Disapprove, data.DisapproveExecutors, data.AllExecutors, data.Id)
	}, assetAsTaskSignIdKey)
	return err
}

func (m *defaultAsTaskSignModel) Delete(ctx context.Context, id string) error {
	assetAsTaskSignIdKey := fmt.Sprintf("%s%v", cacheAssetAsTaskSignIdPrefix, id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, assetAsTaskSignIdKey)
	return err
}

func (m *defaultAsTaskSignModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheAssetAsTaskSignIdPrefix, primary)
}

func (m *defaultAsTaskSignModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", asTaskSignRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultAsTaskSignModel) tableName() string {
	return m.table
}
