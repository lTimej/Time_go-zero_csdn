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
	tbContentCategoryFieldNames          = builder.RawFieldNames(&TbContentCategory{})
	tbContentCategoryRows                = strings.Join(tbContentCategoryFieldNames, ",")
	tbContentCategoryRowsExpectAutoSet   = strings.Join(stringx.Remove(tbContentCategoryFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	tbContentCategoryRowsWithPlaceHolder = strings.Join(stringx.Remove(tbContentCategoryFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"

	cacheTbContentCategoryIdPrefix = "cache:tbContentCategory:id:"
)

type (
	tbContentCategoryModel interface {
		Insert(ctx context.Context, data *TbContentCategory) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*TbContentCategory, error)
		Update(ctx context.Context, data *TbContentCategory) error
		Delete(ctx context.Context, id int64) error
	}

	defaultTbContentCategoryModel struct {
		sqlc.CachedConn
		table string
	}

	TbContentCategory struct {
		Id           int64          `db:"id"`
		Cid          sql.NullString `db:"cid"`           // 内容id
		AdCategoryId int64          `db:"adCategory_id"` // 广告类别
		Title        sql.NullString `db:"title"`         // 内容标题
		CreateTime   time.Time      `db:"create_time"`
		UpdateTime   time.Time      `db:"update_time"`
	}
)

func newTbContentCategoryModel(conn sqlx.SqlConn, c cache.CacheConf) *defaultTbContentCategoryModel {
	return &defaultTbContentCategoryModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`tb_content_category`",
	}
}

func (m *defaultTbContentCategoryModel) Delete(ctx context.Context, id int64) error {
	tbContentCategoryIdKey := fmt.Sprintf("%s%v", cacheTbContentCategoryIdPrefix, id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, tbContentCategoryIdKey)
	return err
}

func (m *defaultTbContentCategoryModel) FindOne(ctx context.Context, id int64) (*TbContentCategory, error) {
	tbContentCategoryIdKey := fmt.Sprintf("%s%v", cacheTbContentCategoryIdPrefix, id)
	var resp TbContentCategory
	err := m.QueryRowCtx(ctx, &resp, tbContentCategoryIdKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", tbContentCategoryRows, m.table)
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

func (m *defaultTbContentCategoryModel) Insert(ctx context.Context, data *TbContentCategory) (sql.Result, error) {
	tbContentCategoryIdKey := fmt.Sprintf("%s%v", cacheTbContentCategoryIdPrefix, data.Id)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?)", m.table, tbContentCategoryRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.Cid, data.AdCategoryId, data.Title)
	}, tbContentCategoryIdKey)
	return ret, err
}

func (m *defaultTbContentCategoryModel) Update(ctx context.Context, data *TbContentCategory) error {
	tbContentCategoryIdKey := fmt.Sprintf("%s%v", cacheTbContentCategoryIdPrefix, data.Id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, tbContentCategoryRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.Cid, data.AdCategoryId, data.Title, data.Id)
	}, tbContentCategoryIdKey)
	return err
}

func (m *defaultTbContentCategoryModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheTbContentCategoryIdPrefix, primary)
}

func (m *defaultTbContentCategoryModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", tbContentCategoryRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultTbContentCategoryModel) tableName() string {
	return m.table
}