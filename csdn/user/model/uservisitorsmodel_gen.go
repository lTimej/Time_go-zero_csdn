// Code generated by goctl. DO NOT EDIT.

package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	userVisitorsFieldNames          = builder.RawFieldNames(&UserVisitors{})
	userVisitorsRows                = strings.Join(userVisitorsFieldNames, ",")
	userVisitorsRowsExpectAutoSet   = strings.Join(stringx.Remove(userVisitorsFieldNames, "`visit_id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	userVisitorsRowsWithPlaceHolder = strings.Join(stringx.Remove(userVisitorsFieldNames, "`visit_id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"

	cacheUserVisitorsVisitIdPrefix = "cache:userVisitors:visitId:"
)

type (
	userVisitorsModel interface {
		Insert(ctx context.Context, data *UserVisitors) (sql.Result, error)
		FindOne(ctx context.Context, visitId int64) (*UserVisitors, error)
		Update(ctx context.Context, data *UserVisitors) error
		Delete(ctx context.Context, visitId int64) error
	}

	defaultUserVisitorsModel struct {
		sqlc.CachedConn
		table string
	}

	UserVisitors struct {
		VisitId    int64     `db:"visit_id"`    // 主键id
		UserId     int64     `db:"user_id"`     // 用户ID
		CreateTime time.Time `db:"create_time"` // 创建时间
		Count      int64     `db:"count"`       // 用户访问量
	}
)

func newUserVisitorsModel(conn sqlx.SqlConn, c cache.CacheConf) *defaultUserVisitorsModel {
	return &defaultUserVisitorsModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`user_visitors`",
	}
}

func (m *defaultUserVisitorsModel) Delete(ctx context.Context, visitId int64) error {
	userVisitorsVisitIdKey := fmt.Sprintf("%s%v", cacheUserVisitorsVisitIdPrefix, visitId)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `visit_id` = ?", m.table)
		return conn.ExecCtx(ctx, query, visitId)
	}, userVisitorsVisitIdKey)
	return err
}

func (m *defaultUserVisitorsModel) FindOne(ctx context.Context, visitId int64) (*UserVisitors, error) {
	userVisitorsVisitIdKey := fmt.Sprintf("%s%v", cacheUserVisitorsVisitIdPrefix, visitId)
	var resp UserVisitors
	err := m.QueryRowCtx(ctx, &resp, userVisitorsVisitIdKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `visit_id` = ? limit 1", userVisitorsRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, visitId)
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

func (m *defaultUserVisitorsModel) Insert(ctx context.Context, data *UserVisitors) (sql.Result, error) {
	userVisitorsVisitIdKey := fmt.Sprintf("%s%v", cacheUserVisitorsVisitIdPrefix, data.VisitId)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?)", m.table, userVisitorsRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.UserId, data.Count)
	}, userVisitorsVisitIdKey)
	return ret, err
}

func (m *defaultUserVisitorsModel) Update(ctx context.Context, data *UserVisitors) error {
	userVisitorsVisitIdKey := fmt.Sprintf("%s%v", cacheUserVisitorsVisitIdPrefix, data.VisitId)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `visit_id` = ?", m.table, userVisitorsRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.UserId, data.Count, data.VisitId)
	}, userVisitorsVisitIdKey)
	return err
}

func (m *defaultUserVisitorsModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheUserVisitorsVisitIdPrefix, primary)
}

func (m *defaultUserVisitorsModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `visit_id` = ? limit 1", userVisitorsRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultUserVisitorsModel) tableName() string {
	return m.table
}
