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
	userOrderFieldNames          = builder.RawFieldNames(&UserOrder{})
	userOrderRows                = strings.Join(userOrderFieldNames, ",")
	userOrderRowsExpectAutoSet   = strings.Join(stringx.Remove(userOrderFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	userOrderRowsWithPlaceHolder = strings.Join(stringx.Remove(userOrderFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"
	orderInfoRows                = "default_image,price,count,sku_id,spec_id,specs,title"
	cacheUserOrderIdPrefix       = "cache:userOrder:id:"
	cacheOrderInfoOrderIdPrefix  = "cache:orderinfo:order:id:"
)

type (
	userOrderModel interface {
		Builder() squirrel.SelectBuilder
		Insert(ctx context.Context, data *UserOrder) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*UserOrder, error)
		FindOneByOrderId(ctx context.Context, builder squirrel.SelectBuilder) ([]*OrderInfo, error)
		Update(ctx context.Context, data *UserOrder) error
		Delete(ctx context.Context, id int64) error
	}

	defaultUserOrderModel struct {
		sqlc.CachedConn
		table string
	}

	UserOrder struct {
		Id          int64     `db:"id"`
		OrderId     int64     `db:"order_id"`     // 订单id
		SkuId       int64     `db:"sku_id"`       // 商品id
		SpecId      string    `db:"spec_id"`      // 商品属性id;"1,2"
		Specs       string    `db:"specs"`        // 商品属性
		Count       int64     `db:"count"`        //商品个数
		Comment     string    `db:"comment"`      // 商品评价
		Score       int64     `db:"score"`        // 商品评分
		CreateTime  time.Time `db:"create_time"`  // 支付创建时间
		UpdateTime  time.Time `db:"update_time"`  // 支付修改时间
		IsAnonymous int64     `db:"is_anonymous"` // 是否匿名
		IsCommented int64     `db:"is_commented"` // 是否评论
	}
	OrderInfo struct {
		DefaultImage string  `db:"default_image"`
		Price        float32 `db:"price"`
		Count        int64   `db:"count"`
		SkuId        int64   `db:"sku_id"`
		SpecId       string  `db:"spec_id"`
		Specs        string  `db:"specs"`
		Title        string  `db:"title"`
	}
)

func newUserOrderModel(conn sqlx.SqlConn, c cache.CacheConf) *defaultUserOrderModel {
	return &defaultUserOrderModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`user_order`",
	}
}

func (m *defaultUserOrderModel) Delete(ctx context.Context, id int64) error {
	userOrderIdKey := fmt.Sprintf("%s%v", cacheUserOrderIdPrefix, id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, userOrderIdKey)
	return err
}

func (m *defaultUserOrderModel) FindOne(ctx context.Context, id int64) (*UserOrder, error) {
	userOrderIdKey := fmt.Sprintf("%s%v", cacheUserOrderIdPrefix, id)
	var resp UserOrder
	err := m.QueryRowCtx(ctx, &resp, userOrderIdKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", userOrderRows, m.table)
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

func (m *defaultUserOrderModel) FindOneByOrderId(ctx context.Context, builder squirrel.SelectBuilder) ([]*OrderInfo, error) {
	fmt.Println("6666666666666666666666666666")
	query, values, err := builder.LeftJoin("tb_sku on user_order.sku_id = tb_sku.id").ToSql()
	fmt.Println(query, "55555555555555")
	if err != nil {
		return nil, err
	}
	var resp []*OrderInfo
	err = m.QueryRowsNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultUserOrderModel) Insert(ctx context.Context, data *UserOrder) (sql.Result, error) {
	userOrderIdKey := fmt.Sprintf("%s%v", cacheUserOrderIdPrefix, data.Id)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?,?)", m.table, userOrderRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.OrderId, data.SkuId, data.SpecId, data.Specs, data.Count, data.Comment, data.Score, data.IsAnonymous, data.IsCommented)
	}, userOrderIdKey)
	return ret, err
}

func (m *defaultUserOrderModel) Update(ctx context.Context, data *UserOrder) error {
	userOrderIdKey := fmt.Sprintf("%s%v", cacheUserOrderIdPrefix, data.Id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, userOrderRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.OrderId, data.SkuId, data.SpecId, data.Specs, data.Comment, data.Score, data.IsAnonymous, data.IsCommented, data.Id)
	}, userOrderIdKey)
	return err
}

func (m *defaultUserOrderModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheUserOrderIdPrefix, primary)
}

func (m *defaultUserOrderModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", userOrderRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultUserOrderModel) tableName() string {
	return m.table
}

func (m *defaultUserOrderModel) Builder() squirrel.SelectBuilder {
	return squirrel.Select(orderInfoRows).From(m.table)
}