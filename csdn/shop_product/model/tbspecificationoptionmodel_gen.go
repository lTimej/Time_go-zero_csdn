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
	tbSpecificationOptionFieldNames          = builder.RawFieldNames(&TbSpecificationOption{})
	tbSpecificationOptionRows                = strings.Join(tbSpecificationOptionFieldNames, ",")
	tbSpecificationOptionRowsExpectAutoSet   = strings.Join(stringx.Remove(tbSpecificationOptionFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	tbSpecificationOptionRowsWithPlaceHolder = strings.Join(stringx.Remove(tbSpecificationOptionFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"

	cacheTbSpecificationOptionIdPrefix = "cache:tbSpecificationOption:id:"
)

type (
	tbSpecificationOptionModel interface {
		Insert(ctx context.Context, data *TbSpecificationOption) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*TbSpecificationOption, error)
		Update(ctx context.Context, data *TbSpecificationOption) error
		Delete(ctx context.Context, id int64) error
	}

	defaultTbSpecificationOptionModel struct {
		sqlc.CachedConn
		table string
	}

	TbSpecificationOption struct {
		Id         int64          `db:"id"`
		SpecId     int64          `db:"spec_id"` // 规格
		Value      sql.NullString `db:"value"`   // 选项值
		CreateTime time.Time      `db:"create_time"`
		UpdateTime time.Time      `db:"update_time"`
	}
)

func newTbSpecificationOptionModel(conn sqlx.SqlConn, c cache.CacheConf) *defaultTbSpecificationOptionModel {
	return &defaultTbSpecificationOptionModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`tb_specification_option`",
	}
}

func (m *defaultTbSpecificationOptionModel) Delete(ctx context.Context, id int64) error {
	tbSpecificationOptionIdKey := fmt.Sprintf("%s%v", cacheTbSpecificationOptionIdPrefix, id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, tbSpecificationOptionIdKey)
	return err
}

func (m *defaultTbSpecificationOptionModel) FindOne(ctx context.Context, id int64) (*TbSpecificationOption, error) {
	tbSpecificationOptionIdKey := fmt.Sprintf("%s%v", cacheTbSpecificationOptionIdPrefix, id)
	var resp TbSpecificationOption
	err := m.QueryRowCtx(ctx, &resp, tbSpecificationOptionIdKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", tbSpecificationOptionRows, m.table)
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

func (m *defaultTbSpecificationOptionModel) Insert(ctx context.Context, data *TbSpecificationOption) (sql.Result, error) {
	tbSpecificationOptionIdKey := fmt.Sprintf("%s%v", cacheTbSpecificationOptionIdPrefix, data.Id)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?)", m.table, tbSpecificationOptionRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.SpecId, data.Value)
	}, tbSpecificationOptionIdKey)
	return ret, err
}

func (m *defaultTbSpecificationOptionModel) Update(ctx context.Context, data *TbSpecificationOption) error {
	tbSpecificationOptionIdKey := fmt.Sprintf("%s%v", cacheTbSpecificationOptionIdPrefix, data.Id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, tbSpecificationOptionRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.SpecId, data.Value, data.Id)
	}, tbSpecificationOptionIdKey)
	return err
}

func (m *defaultTbSpecificationOptionModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheTbSpecificationOptionIdPrefix, primary)
}

func (m *defaultTbSpecificationOptionModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", tbSpecificationOptionRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultTbSpecificationOptionModel) tableName() string {
	return m.table
}