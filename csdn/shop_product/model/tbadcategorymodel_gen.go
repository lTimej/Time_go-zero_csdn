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
	tbAdCategoryFieldNames          = builder.RawFieldNames(&TbAdCategory{})
	tbAdCategoryRows                = strings.Join(tbAdCategoryFieldNames, ",")
	tbAdCategoryRowsExpectAutoSet   = strings.Join(stringx.Remove(tbAdCategoryFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	tbAdCategoryRowsWithPlaceHolder = strings.Join(stringx.Remove(tbAdCategoryFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"

	cacheTbAdCategoryIdPrefix = "cache:tbAdCategory:id:"
)

type (
	tbAdCategoryModel interface {
		Insert(ctx context.Context, data *TbAdCategory) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*TbAdCategory, error)
		Update(ctx context.Context, data *TbAdCategory) error
		Delete(ctx context.Context, id int64) error
	}

	defaultTbAdCategoryModel struct {
		sqlc.CachedConn
		table string
	}

	TbAdCategory struct {
		Id         int64          `db:"id"`
		Title      sql.NullString `db:"title"` // 广告类别
		CreateTime time.Time      `db:"create_time"`
		UpdateTime time.Time      `db:"update_time"`
	}
)

func newTbAdCategoryModel(conn sqlx.SqlConn, c cache.CacheConf) *defaultTbAdCategoryModel {
	return &defaultTbAdCategoryModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`tb_ad_category`",
	}
}

func (m *defaultTbAdCategoryModel) Delete(ctx context.Context, id int64) error {
	tbAdCategoryIdKey := fmt.Sprintf("%s%v", cacheTbAdCategoryIdPrefix, id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, tbAdCategoryIdKey)
	return err
}

func (m *defaultTbAdCategoryModel) FindOne(ctx context.Context, id int64) (*TbAdCategory, error) {
	tbAdCategoryIdKey := fmt.Sprintf("%s%v", cacheTbAdCategoryIdPrefix, id)
	var resp TbAdCategory
	err := m.QueryRowCtx(ctx, &resp, tbAdCategoryIdKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", tbAdCategoryRows, m.table)
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

func (m *defaultTbAdCategoryModel) Insert(ctx context.Context, data *TbAdCategory) (sql.Result, error) {
	tbAdCategoryIdKey := fmt.Sprintf("%s%v", cacheTbAdCategoryIdPrefix, data.Id)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?)", m.table, tbAdCategoryRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.Title)
	}, tbAdCategoryIdKey)
	return ret, err
}

func (m *defaultTbAdCategoryModel) Update(ctx context.Context, data *TbAdCategory) error {
	tbAdCategoryIdKey := fmt.Sprintf("%s%v", cacheTbAdCategoryIdPrefix, data.Id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, tbAdCategoryRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.Title, data.Id)
	}, tbAdCategoryIdKey)
	return err
}

func (m *defaultTbAdCategoryModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheTbAdCategoryIdPrefix, primary)
}

func (m *defaultTbAdCategoryModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", tbAdCategoryRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultTbAdCategoryModel) tableName() string {
	return m.table
}
