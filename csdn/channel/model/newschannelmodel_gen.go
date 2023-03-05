// Code generated by goctl. DO NOT EDIT.

package model

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/Masterminds/squirrel"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	newsChannelFieldNames          = builder.RawFieldNames(&NewsChannel{})
	newsChannelRows                = strings.Join(newsChannelFieldNames, ",")
	newsChannelRowsExpectAutoSet   = strings.Join(stringx.Remove(newsChannelFieldNames, "`channel_id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	newsChannelRowsWithPlaceHolder = strings.Join(stringx.Remove(newsChannelFieldNames, "`channel_id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"

	cacheNewsChannelChannelIdPrefix   = "cache:newsChannel:channelId:"
	cacheNewsChannelChannelNamePrefix = "cache:newsChannel:channelName:"
)

type (
	newsChannelModel interface {
		RowBuilder() squirrel.SelectBuilder
		Insert(ctx context.Context, data *NewsChannel) (sql.Result, error)
		FindOne(ctx context.Context, channelId int64) (*NewsChannel, error)
		FindAll(ctx context.Context, rowBuilder squirrel.SelectBuilder, orderBy string) ([]*NewsChannel, error)
		FindOneByChannelName(ctx context.Context, channelName string) (*NewsChannel, error)
		Update(ctx context.Context, data *NewsChannel) error
		Delete(ctx context.Context, channelId int64) error
	}

	defaultNewsChannelModel struct {
		sqlc.CachedConn
		table string
	}

	NewsChannel struct {
		ChannelId   int64     `db:"channel_id"`   // 频道ID
		ChannelName string    `db:"channel_name"` // 频道名称
		CreateTime  time.Time `db:"create_time"`  // 创建时间
		UpdateTime  time.Time `db:"update_time"`  // 更新时间
		Sequence    int64     `db:"sequence"`     // 序号
		IsVisible   int64     `db:"is_visible"`   // 是否可见
		IsDefault   int64     `db:"is_default"`   // 是否默认
	}
)

func newNewsChannelModel(conn sqlx.SqlConn, c cache.CacheConf) *defaultNewsChannelModel {
	return &defaultNewsChannelModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`news_channel`",
	}
}

func (m *defaultNewsChannelModel) Delete(ctx context.Context, channelId int64) error {
	data, err := m.FindOne(ctx, channelId)
	if err != nil {
		return err
	}

	newsChannelChannelIdKey := fmt.Sprintf("%s%v", cacheNewsChannelChannelIdPrefix, channelId)
	newsChannelChannelNameKey := fmt.Sprintf("%s%v", cacheNewsChannelChannelNamePrefix, data.ChannelName)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `channel_id` = ?", m.table)
		return conn.ExecCtx(ctx, query, channelId)
	}, newsChannelChannelIdKey, newsChannelChannelNameKey)
	return err
}

func (m *defaultNewsChannelModel) FindOne(ctx context.Context, channelId int64) (*NewsChannel, error) {
	newsChannelChannelIdKey := fmt.Sprintf("%s%v", cacheNewsChannelChannelIdPrefix, channelId)
	var resp NewsChannel
	err := m.QueryRowCtx(ctx, &resp, newsChannelChannelIdKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
		query := fmt.Sprintf("select %s from %s where `channel_id` = ? limit 1", newsChannelRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, channelId)
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
func (m *defaultNewsChannelModel) FindAll(ctx context.Context, rowBuilder squirrel.SelectBuilder, orderBy string) ([]*NewsChannel, error) {
	if orderBy == "" {
		rowBuilder = rowBuilder.OrderBy("channel_id DESC")
	} else {
		rowBuilder = rowBuilder.OrderBy(orderBy)
	}
	query,values,err := rowBuilder.ToSql()
	//query, values, err := rowBuilder.Where("del_state = ?", globalkey.DelStateNo).ToSql()
	if err != nil {
		return nil, err
	}

	var resp []*NewsChannel
	err = m.QueryRowsNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}
func (m *defaultNewsChannelModel) FindOneByChannelName(ctx context.Context, channelName string) (*NewsChannel, error) {
	newsChannelChannelNameKey := fmt.Sprintf("%s%v", cacheNewsChannelChannelNamePrefix, channelName)
	var resp NewsChannel
	err := m.QueryRowIndexCtx(ctx, &resp, newsChannelChannelNameKey, m.formatPrimary, func(ctx context.Context, conn sqlx.SqlConn, v any) (i any, e error) {
		query := fmt.Sprintf("select %s from %s where `channel_name` = ? limit 1", newsChannelRows, m.table)
		if err := conn.QueryRowCtx(ctx, &resp, query, channelName); err != nil {
			return nil, err
		}
		return resp.ChannelId, nil
	}, m.queryPrimary)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultNewsChannelModel) Insert(ctx context.Context, data *NewsChannel) (sql.Result, error) {
	newsChannelChannelIdKey := fmt.Sprintf("%s%v", cacheNewsChannelChannelIdPrefix, data.ChannelId)
	newsChannelChannelNameKey := fmt.Sprintf("%s%v", cacheNewsChannelChannelNamePrefix, data.ChannelName)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?)", m.table, newsChannelRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.ChannelName, data.Sequence, data.IsVisible, data.IsDefault)
	}, newsChannelChannelIdKey, newsChannelChannelNameKey)
	return ret, err
}

func (m *defaultNewsChannelModel) Update(ctx context.Context, newData *NewsChannel) error {
	data, err := m.FindOne(ctx, newData.ChannelId)
	if err != nil {
		return err
	}

	newsChannelChannelIdKey := fmt.Sprintf("%s%v", cacheNewsChannelChannelIdPrefix, data.ChannelId)
	newsChannelChannelNameKey := fmt.Sprintf("%s%v", cacheNewsChannelChannelNamePrefix, data.ChannelName)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `channel_id` = ?", m.table, newsChannelRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, newData.ChannelName, newData.Sequence, newData.IsVisible, newData.IsDefault, newData.ChannelId)
	}, newsChannelChannelIdKey, newsChannelChannelNameKey)
	return err
}

func (m *defaultNewsChannelModel) formatPrimary(primary any) string {
	return fmt.Sprintf("%s%v", cacheNewsChannelChannelIdPrefix, primary)
}

func (m *defaultNewsChannelModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary any) error {
	query := fmt.Sprintf("select %s from %s where `channel_id` = ? limit 1", newsChannelRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultNewsChannelModel) tableName() string {
	return m.table
}
// export logic
func (m *defaultNewsChannelModel) RowBuilder() squirrel.SelectBuilder {
	return squirrel.Select(newsChannelRows).From(m.table)
}