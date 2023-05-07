// Code generated by goctl. DO NOT EDIT.

package model

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
	cityFieldNames          = builder.RawFieldNames(&City{})
	cityRows                = strings.Join(cityFieldNames, ",")
	cityRowsExpectAutoSet   = strings.Join(stringx.Remove(cityFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	cityRowsWithPlaceHolder = strings.Join(stringx.Remove(cityFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"

	cacheCityIdPrefix = "cache:city:id:"
)

type (
	cityModel interface {
		Insert(ctx context.Context, data *City) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*City, error)
		Update(ctx context.Context, data *City) error
		Delete(ctx context.Context, id int64) error
	}

	defaultCityModel struct {
		sqlc.CachedConn
		table string
	}

	City struct {
		Id   int64          `db:"id"`
		Code sql.NullInt64  `db:"code"` // 行政区划代码
		Name sql.NullString `db:"name"` // 名称
		Pid  sql.NullInt64  `db:"pid"`  // 上级id
		Type sql.NullString `db:"type"` // 类型
	}
)

func newCityModel(conn sqlx.SqlConn, c cache.CacheConf) *defaultCityModel {
	return &defaultCityModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`city`",
	}
}

func (m *defaultCityModel) Delete(ctx context.Context, id int64) error {
	cityIdKey := fmt.Sprintf("%s%v", cacheCityIdPrefix, id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, cityIdKey)
	return err
}

func (m *defaultCityModel) FindOne(ctx context.Context, id int64) (*City, error) {
	cityIdKey := fmt.Sprintf("%s%v", cacheCityIdPrefix, id)
	var resp City
	err := m.QueryRowCtx(ctx, &resp, cityIdKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", cityRows, m.table)
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

func (m *defaultCityModel) Insert(ctx context.Context, data *City) (sql.Result, error) {
	cityIdKey := fmt.Sprintf("%s%v", cacheCityIdPrefix, data.Id)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?)", m.table, cityRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.Code, data.Name, data.Pid, data.Type)
	}, cityIdKey)
	return ret, err
}

func (m *defaultCityModel) Update(ctx context.Context, data *City) error {
	cityIdKey := fmt.Sprintf("%s%v", cacheCityIdPrefix, data.Id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, cityRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.Code, data.Name, data.Pid, data.Type, data.Id)
	}, cityIdKey)
	return err
}

func (m *defaultCityModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheCityIdPrefix, primary)
}

func (m *defaultCityModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", cityRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultCityModel) tableName() string {
	return m.table
}