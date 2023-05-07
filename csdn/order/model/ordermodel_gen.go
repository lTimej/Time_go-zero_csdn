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
	orderFieldNames          = builder.RawFieldNames(&Order{})
	orderRows                = strings.Join(orderFieldNames, ",")
	orderRowsExpectAutoSet   = strings.Join(stringx.Remove(orderFieldNames, "`id`", "`pay_status`", "`is_deleted`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	orderRowsWithPlaceHolder = strings.Join(stringx.Remove(orderFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"

	cacheOrderIdPrefix = "cache:order:id:"
	cacheOrderSnPrefix = "cache:order:sn:"
)

type (
	orderModel interface {
		Builder() squirrel.SelectBuilder
		Insert(ctx context.Context, data *Order) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*Order, error)
		FindOneBySn(ctx context.Context, sn string) (*Order, error)
		FindAllByUserId(ctx context.Context, builder squirrel.SelectBuilder) ([]*Order, error)
		Update(ctx context.Context, data *Order) error
		Delete(ctx context.Context, id int64) error
	}

	defaultOrderModel struct {
		sqlc.CachedConn
		table string
	}

	Order struct {
		Id         int64     `db:"id"`
		UserId     string    `db:"user_id"`     // 用户id
		AddressId  int64     `db:"address_id"`  // 地址
		TotalCount int64     `db:"total_count"` // 商品总数
		TotalPrice float32   `db:"total_price"` // 商品总金额
		Freight    float32   `db:"freight"`     // 运费
		Version    int64     `db:"version"`     // 乐观锁版本号
		Sn         string    `db:"sn"`          // 流水单号
		PayStatus  int64     `db:"pay_status"`  // 支付状态 0: 已取消 1:待支付 2:未使用 3:已使用  4:已过期  5:代发货  6:待收货  7:待评价  8:售后
		CreateTime time.Time `db:"create_time"` // 支付创建时间
		UpdateTime time.Time `db:"update_time"` // 支付修改时间
		IsDeleted  int64     `db:"is_deleted"`  // 逻辑删除
	}
)

func newOrderModel(conn sqlx.SqlConn, c cache.CacheConf) *defaultOrderModel {
	return &defaultOrderModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`orders`",
	}
}

func (m *defaultOrderModel) Delete(ctx context.Context, id int64) error {
	orderIdKey := fmt.Sprintf("%s%v", cacheOrderIdPrefix, id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, orderIdKey)
	return err
}

func (m *defaultOrderModel) FindOne(ctx context.Context, id int64) (*Order, error) {
	orderIdKey := fmt.Sprintf("%s%v", cacheOrderIdPrefix, id)
	var resp Order
	err := m.QueryRowCtx(ctx, &resp, orderIdKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", orderRows, m.table)
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

func (m *defaultOrderModel) FindOneBySn(ctx context.Context, sn string) (*Order, error) {
	orderSnKey := fmt.Sprintf("%s%v", cacheOrderSnPrefix, sn)
	var resp Order
	err := m.QueryRowCtx(ctx, &resp, orderSnKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `sn` = ? limit 1", orderRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, sn)
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

func (m *defaultOrderModel) FindAllByUserId(ctx context.Context, builder squirrel.SelectBuilder) ([]*Order, error) {
	query, values, err := builder.ToSql()
	if err != nil {
		return nil, err
	}
	var resp []*Order
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

func (m *defaultOrderModel) Insert(ctx context.Context, data *Order) (sql.Result, error) {
	orderIdKey := fmt.Sprintf("%s%v", cacheOrderIdPrefix, data.Id)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?)", m.table, orderRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.UserId, data.AddressId, data.TotalCount, data.TotalPrice, data.Freight, data.Version, data.Sn)
	}, orderIdKey)
	fmt.Println(ret, err, "***************&&&&&&&&&&&")
	return ret, err
}

func (m *defaultOrderModel) Update(ctx context.Context, data *Order) error {
	orderIdKey := fmt.Sprintf("%s%v", cacheOrderIdPrefix, data.Id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, orderRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.UserId, data.AddressId, data.TotalCount, data.TotalPrice, data.Freight, data.Version, data.Sn, data.PayStatus, data.IsDeleted, data.Id)
	}, orderIdKey)
	return err
}

func (m *defaultOrderModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheOrderIdPrefix, primary)
}

func (m *defaultOrderModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", orderRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultOrderModel) tableName() string {
	return m.table
}

func (m *defaultOrderModel) Builder() squirrel.SelectBuilder {
	return squirrel.Select(orderRows).From(m.table)
}
