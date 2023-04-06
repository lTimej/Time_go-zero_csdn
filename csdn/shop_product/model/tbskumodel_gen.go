// Code generated by goctl. DO NOT EDIT.

package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/Masterminds/squirrel"
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

	cacheTbSkuIdPrefix         = "cache:tbSku:id:"
	cacheTbSkuCategoryIdPrefix = "cache:tbSku:categoryid:"
)

type (
	tbSkuModel interface {
		Builder() squirrel.SelectBuilder
		BuilderSpec() squirrel.SelectBuilder
		Insert(ctx context.Context, data *TbSku) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*TbSku, error)
		FindOneByCategoryId(ctx context.Context, category_id int64) (*TbSku, error)
		FindAllSkuBasicInfoBySpuId(ctx context.Context, builder squirrel.SelectBuilder) ([]*SkuBaseInfo, error)
		FindAllSkuSpecBySpuId(ctx context.Context, builder squirrel.SelectBuilder, spu_id int64) ([]*SkuSpecInfo, error)
		Update(ctx context.Context, data *TbSku) error
		Delete(ctx context.Context, id int64) error
	}

	defaultTbSkuModel struct {
		sqlc.CachedConn
		table string
	}

	TbSku struct {
		Id           int64     `db:"id"`
		Title        string    `db:"title"`         // 名称
		SpuId        int64     `db:"spu_id"`        // 商品
		CategoryId   int64     `db:"category_id"`   // 从属类别
		Price        float64   `db:"price"`         // 单价
		NowPrice     float64   `db:"now_price"`     // 进价
		Stock        int64     `db:"stock"`         // 库存
		Sales        int64     `db:"sales"`         // 销量
		Comments     int64     `db:"comments"`      // 评价数
		IsLaunched   int64     `db:"is_launched"`   // 是否上架销售
		DefaultImage string    `db:"default_image"` // 默认图片
		CreateTime   time.Time `db:"create_time"`
		UpdateTime   time.Time `db:"update_time"`
	}
	SkuBaseInfo struct {
		Name         string  `db:"title"`
		Price        float32 `db:"price"`
		NowPrice     float32 `db:"now_price"`
		DefaultImage string  `db:"default_image"`
		Stock        int64   `db:"stock"`
	}
	SkuSpecInfo struct {
		SkuId        int64   `db:"sku_id"`
		Title        string  `db:"title"`
		Stock        int64   `db:"stock"`
		Price        float32 `db:"price"`
		NowPrice     float32 `db:"now_price"`
		DefaultImage string  `db:"default_image"`
		Label        string  `db:"label"`
		Name         string  `db:"name"`
		SpecId       int64   `db:"spec_id"`
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

func (m *defaultTbSkuModel) FindAllSkuBasicInfoBySpuId(ctx context.Context, builder squirrel.SelectBuilder) ([]*SkuBaseInfo, error) {
	query, values, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	var resp []*SkuBaseInfo
	err = m.QueryRowsNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

func (m *defaultTbSkuModel) FindAllSkuSpecBySpuId(ctx context.Context, builder squirrel.SelectBuilder, spu_id int64) ([]*SkuSpecInfo, error) {
	query, values, err := builder.Join("tb_sku_specification as t2,tb_spu_specification as t3,tb_specification_option as t4 where tb_sku.id=t2.sku_id and t3.id=t4.spec_id and t2.option_id=t4.id").ToSql()
	query = fmt.Sprintf("%s and tb_sku.spu_id = %d", query, spu_id)
	fmt.Println(query, "********************")
	if err != nil {
		return nil, err
	}
	var resp []*SkuSpecInfo
	err = m.QueryRowsNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

func (m *defaultTbSkuModel) FindOneByCategoryId(ctx context.Context, category_id int64) (*TbSku, error) {
	tbSkuCategoryIdKey := fmt.Sprintf("%s%v", cacheTbSkuCategoryIdPrefix, category_id)
	var resp TbSku
	err := m.QueryRowCtx(ctx, &resp, tbSkuCategoryIdKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `category_id` = ? limit 1", tbSkuRows, m.table)
		fmt.Println(query, category_id, "777777777777")
		return conn.QueryRowCtx(ctx, v, query, category_id)
	})
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, nil
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

func (m *defaultTbSkuModel) Builder() squirrel.SelectBuilder {
	return squirrel.Select("title,price,now_price,default_image,stock").From(m.table)

}

func (m *defaultTbSkuModel) BuilderSpec() squirrel.SelectBuilder {
	return squirrel.Select("tb_sku.id as sku_id, tb_sku.title,tb_sku.stock,tb_sku.price,tb_sku.now_price,tb_sku.default_image,t3.name as label,t4.value as name,t3.id as spec_id").From(m.table)
}
