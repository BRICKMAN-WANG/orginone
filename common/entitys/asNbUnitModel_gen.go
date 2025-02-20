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
	asNbUnitFieldNames          = builder.RawFieldNames(&AsNbUnit{})
	asNbUnitRows                = strings.Join(asNbUnitFieldNames, ",")
	asNbUnitRowsExpectAutoSet   = strings.Join(stringx.Remove(asNbUnitFieldNames, "`id`", "`create_time`", "`update_time`"), ",")
	asNbUnitRowsWithPlaceHolder = strings.Join(stringx.Remove(asNbUnitFieldNames, "`id`", "`create_time`", "`update_time`"), "=?,") + "=?"

	cacheAssetAsNbUnitIdPrefix = "cache:asset:asNbUnit:id:"
)

type (
	asNbUnitModel interface {
		Insert(ctx context.Context, data *AsNbUnit) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*AsNbUnit, error)
		Update(ctx context.Context, data *AsNbUnit) error
		Delete(ctx context.Context, id int64) error
	}

	defaultAsNbUnitModel struct {
		sqlc.CachedConn
		table string
	}

	AsNbUnit struct {
		UnitName sql.NullString `db:"unit_name"`
		ZzjbCode sql.NullString `db:"zzjb_code"`
		Code     sql.NullString `db:"code"`
		ParentId sql.NullString `db:"parent_id"`
		Shifmjjd sql.NullString `db:"shifmjjd"`
		Jc       sql.NullString `db:"jc"`
		Id       int64          `db:"id"`
	}
)

func newAsNbUnitModel(conn sqlx.SqlConn, c cache.CacheConf) *defaultAsNbUnitModel {
	return &defaultAsNbUnitModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`as_nb_unit`",
	}
}

func (m *defaultAsNbUnitModel) Insert(ctx context.Context, data *AsNbUnit) (sql.Result, error) {
	assetAsNbUnitIdKey := fmt.Sprintf("%s%v", cacheAssetAsNbUnitIdPrefix, data.Id)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?)", m.table, asNbUnitRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.UnitName, data.ZzjbCode, data.Code, data.ParentId, data.Shifmjjd, data.Jc)
	}, assetAsNbUnitIdKey)
	return ret, err
}

func (m *defaultAsNbUnitModel) FindOne(ctx context.Context, id int64) (*AsNbUnit, error) {
	assetAsNbUnitIdKey := fmt.Sprintf("%s%v", cacheAssetAsNbUnitIdPrefix, id)
	var resp AsNbUnit
	err := m.QueryRowCtx(ctx, &resp, assetAsNbUnitIdKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", asNbUnitRows, m.table)
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

func (m *defaultAsNbUnitModel) Update(ctx context.Context, data *AsNbUnit) error {
	assetAsNbUnitIdKey := fmt.Sprintf("%s%v", cacheAssetAsNbUnitIdPrefix, data.Id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, asNbUnitRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.UnitName, data.ZzjbCode, data.Code, data.ParentId, data.Shifmjjd, data.Jc, data.Id)
	}, assetAsNbUnitIdKey)
	return err
}

func (m *defaultAsNbUnitModel) Delete(ctx context.Context, id int64) error {
	assetAsNbUnitIdKey := fmt.Sprintf("%s%v", cacheAssetAsNbUnitIdPrefix, id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, assetAsNbUnitIdKey)
	return err
}

func (m *defaultAsNbUnitModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheAssetAsNbUnitIdPrefix, primary)
}

func (m *defaultAsNbUnitModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", asNbUnitRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultAsNbUnitModel) tableName() string {
	return m.table
}
