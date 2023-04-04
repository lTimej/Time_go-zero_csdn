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
	tbSkuFieldNames          = builder.RawFieldNames(&TbSku{})
	tbSkuRows                = strings.Join(tbSkuFieldNames, ",")
	tbSkuRowsExpectAutoSet   = strings.Join(stringx.Remove(tbSkuFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	tbSkuRowsWithPlaceHolder = strings.Join(stringx.Remove(tbSkuFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"

	cacheTbSkuIdPrefix = "cache:tbSku:id:"
)

type (
	tbSkuModel interface {
		Insert(ctx context.Context, data *TbSku) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*TbSku, error)
		Update(ctx context.Context, data *TbSku) error
		Delete(ctx context.Context, id int64) error
	}

	defaultTbSkuModel struct {
		sqlc.CachedConn
		table string
	}

	TbSku struct {
		Id           int64           `db:"id"`
		Title        sql.NullString  `db:"title"`         // 名称
		SpuId        int64           `db:"spu_id"`        // 商品
		CategoryId   int64           `db:"category_id"`   // 从属类别
		Price        sql.NullFloat64 `db:"price"`         // 单价
		NowPrice     sql.NullFloat64 `db:"now_price"`     // 进价
		Stock        sql.NullInt64   `db:"stock"`         // 库存
		Sales        sql.NullInt64   `db:"sales"`         // 销量
		Comments     sql.NullInt64   `db:"comments"`      // 评价数
		IsLaunched   sql.NullInt64   `db:"is_launched"`   // 是否上架销售
		DefaultImage sql.NullString  `db:"default_image"` // 默认图片
		CreateTime   time.Time       `db:"create_time"`
		UpdateTime   time.Time       `db:"update_time"`
	}
)

func newTbSkuModel(conn sqlx.SqlConn, c cache.CacheConf) *defaultTbSkuModel {
	return &defaultTbSkuModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`tb_sku`",
	}
}

func (m *defaultTbSkuModel) Delete(ctx context.Context, id int64) error {
	tbSkuIdKey := fmt.Sprintf("%s%v", cacheTbSkuIdPrefix, id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, tbSkuIdKey)
	return err
}

func (m *defaultTbSkuModel) FindOne(ctx context.Context, id int64) (*TbSku, error) {
	tbSkuIdKey := fmt.Sprintf("%s%v", cacheTbSkuIdPrefix, id)
	var resp TbSku
	err := m.QueryRowCtx(ctx, &resp, tbSkuIdKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", tbSkuRows, m.table)
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

func (m *defaultTbSkuModel) Insert(ctx context.Context, data *TbSku) (sql.Result, error) {
	tbSkuIdKey := fmt.Sprintf("%s%v", cacheTbSkuIdPrefix, data.Id)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, tbSkuRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.Title, data.SpuId, data.CategoryId, data.Price, data.NowPrice, data.Stock, data.Sales, data.Comments, data.IsLaunched, data.DefaultImage)
	}, tbSkuIdKey)
	return ret, err
}

func (m *defaultTbSkuModel) Update(ctx context.Context, data *TbSku) error {
	tbSkuIdKey := fmt.Sprintf("%s%v", cacheTbSkuIdPrefix, data.Id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, tbSkuRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.Title, data.SpuId, data.CategoryId, data.Price, data.NowPrice, data.Stock, data.Sales, data.Comments, data.IsLaunched, data.DefaultImage, data.Id)
	}, tbSkuIdKey)
	return err
}

func (m *defaultTbSkuModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheTbSkuIdPrefix, primary)
}

func (m *defaultTbSkuModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", tbSkuRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultTbSkuModel) tableName() string {
	return m.table
}