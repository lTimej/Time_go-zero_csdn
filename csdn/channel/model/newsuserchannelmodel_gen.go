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
	newsUserChannelFieldNames          = builder.RawFieldNames(&NewsUserChannel{})
	newsUserChannelRows                = strings.Join(newsUserChannelFieldNames, ",")
	UserChannelRowsExpectAutoSet = strings.Join(stringx.Remove(newsUserChannelFieldNames,  "`is_deleted`", "`sequence`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	newsUserChannelRowsExpectAutoSet   = strings.Join(stringx.Remove(newsUserChannelFieldNames, "`user_channel_id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	newsUserChannelRowsWithPlaceHolder = strings.Join(stringx.Remove(newsUserChannelFieldNames, "`user_channel_id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"

	cacheNewsUserChannelUserChannelIdPrefix   = "cache:newsUserChannel:userChannelId:"
	cacheNewsUserChannelUserIdChannelIdPrefix = "cache:newsUserChannel:userId:channelId:"
)

type (
	newsUserChannelModel interface {
		RowBuilder() squirrel.SelectBuilder
		Insert(ctx context.Context, data *NewsUserChannel) (sql.Result, error)
		FindOne(ctx context.Context, userChannelId int64) (*NewsUserChannel, error)
		FindAllByUserId(ctx context.Context,rowBuilder squirrel.SelectBuilder, user_id int64,orderBy string) ([]*UserChannel, error)
		FindOneByUserIdChannelId(ctx context.Context, userId int64, channelId int64) (*NewsUserChannel, error)
		Update(ctx context.Context, data *NewsUserChannel) error
		Delete(ctx context.Context, userChannelId int64) error
	}

	defaultNewsUserChannelModel struct {
		sqlc.CachedConn
		table string
	}

	NewsUserChannel struct {
		UserChannelId int64     `db:"user_channel_id"` // ??????id
		UserId        int64     `db:"user_id"`         // ??????ID
		ChannelId     int64     `db:"channel_id"`      // ??????ID
		CreateTime    time.Time `db:"create_time"`     // ????????????
		IsDeleted     int64     `db:"is_deleted"`      // ????????????, 0-?????????, 1-?????????
		UpdateTime    time.Time `db:"update_time"`     // ????????????
		Sequence      int64     `db:"sequence"`        // ??????
	}
	UserChannel struct {
		ChannelName string    `db:"channel_name"` // ????????????
		UserId        int64     `db:"user_id"`         // ??????ID
		ChannelId     int64     `db:"channel_id"`      // ??????ID
	}
)

func newNewsUserChannelModel(conn sqlx.SqlConn, c cache.CacheConf) *defaultNewsUserChannelModel {
	return &defaultNewsUserChannelModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`news_user_channel`",
	}
}

func (m *defaultNewsUserChannelModel) Delete(ctx context.Context, userChannelId int64) error {
	data, err := m.FindOne(ctx, userChannelId)
	if err != nil {
		return err
	}

	newsUserChannelUserChannelIdKey := fmt.Sprintf("%s%v", cacheNewsUserChannelUserChannelIdPrefix, userChannelId)
	newsUserChannelUserIdChannelIdKey := fmt.Sprintf("%s%v:%v", cacheNewsUserChannelUserIdChannelIdPrefix, data.UserId, data.ChannelId)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `user_channel_id` = ?", m.table)
		return conn.ExecCtx(ctx, query, userChannelId)
	}, newsUserChannelUserChannelIdKey, newsUserChannelUserIdChannelIdKey)
	return err
}

func (m *defaultNewsUserChannelModel) FindOne(ctx context.Context, ChannelId int64) (*NewsUserChannel, error) {
	newsUserChannelUserChannelIdKey := fmt.Sprintf("%s%v", cacheNewsUserChannelUserChannelIdPrefix, ChannelId)
	var resp NewsUserChannel
	err := m.QueryRowCtx(ctx, &resp, newsUserChannelUserChannelIdKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
		query := fmt.Sprintf("select %s from %s where `channel_id` = ? limit 1", newsUserChannelRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, ChannelId)
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
func (m *defaultNewsUserChannelModel) FindAllByUserId(ctx context.Context,rowBuilder squirrel.SelectBuilder, user_id int64,orderBy string) ([]*UserChannel, error){
	if orderBy == "" {
		rowBuilder = rowBuilder.OrderBy("news_user_channel." + "sequence")
	} else {
		rowBuilder = rowBuilder.OrderBy("news_user_channel." + orderBy)
	}
	query,values,err := rowBuilder.Join("news_channel on news_user_channel.channel_id=news_channel.channel_id").Where("news_user_channel.user_id = ?",user_id).ToSql()
	if err != nil {
		return nil, err
	}

	var resp []*UserChannel
	err = m.QueryRowsNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}
func (m *defaultNewsUserChannelModel) FindOneByUserIdChannelId(ctx context.Context, userId int64, channelId int64) (*NewsUserChannel, error) {
	newsUserChannelUserIdChannelIdKey := fmt.Sprintf("%s%v:%v", cacheNewsUserChannelUserIdChannelIdPrefix, userId, channelId)
	var resp NewsUserChannel
	err := m.QueryRowIndexCtx(ctx, &resp, newsUserChannelUserIdChannelIdKey, m.formatPrimary, func(ctx context.Context, conn sqlx.SqlConn, v any) (i any, e error) {
		query := fmt.Sprintf("select %s from %s where `user_id` = ? and `channel_id` = ? limit 1", newsUserChannelRows, m.table)
		if err := conn.QueryRowCtx(ctx, &resp, query, userId, channelId); err != nil {
			return nil, err
		}
		return resp.UserChannelId, nil
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

func (m *defaultNewsUserChannelModel) Insert(ctx context.Context, data *NewsUserChannel) (sql.Result, error) {
	newsUserChannelUserChannelIdKey := fmt.Sprintf("%s%v", cacheNewsUserChannelUserChannelIdPrefix, data.UserChannelId)
	newsUserChannelUserIdChannelIdKey := fmt.Sprintf("%s%v:%v", cacheNewsUserChannelUserIdChannelIdPrefix, data.UserId, data.ChannelId)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?)", m.table, UserChannelRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query,data.UserChannelId, data.UserId, data.ChannelId)
	}, newsUserChannelUserChannelIdKey, newsUserChannelUserIdChannelIdKey)
	return ret, err
}

func (m *defaultNewsUserChannelModel) Update(ctx context.Context, newData *NewsUserChannel) error {
	data, err := m.FindOne(ctx, newData.ChannelId)
	fmt.Println(data,"????????????????????????")
	if err != nil {
		return err
	}

	newsUserChannelUserChannelIdKey := fmt.Sprintf("%s%v", cacheNewsUserChannelUserChannelIdPrefix, data.UserChannelId)
	newsUserChannelUserIdChannelIdKey := fmt.Sprintf("%s%v:%v", cacheNewsUserChannelUserIdChannelIdPrefix, data.UserId, data.ChannelId)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set is_deleted = ? where `user_channel_id` = ?", m.table)
		return conn.ExecCtx(ctx, query, newData.IsDeleted,data.UserChannelId)
	}, newsUserChannelUserChannelIdKey, newsUserChannelUserIdChannelIdKey)
	return err
}

func (m *defaultNewsUserChannelModel) formatPrimary(primary any) string {
	return fmt.Sprintf("%s%v", cacheNewsUserChannelUserChannelIdPrefix, primary)
}

func (m *defaultNewsUserChannelModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary any) error {
	query := fmt.Sprintf("select %s from %s where `user_channel_id` = ? limit 1", newsUserChannelRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultNewsUserChannelModel) tableName() string {
	return m.table
}
// export logic
func (m *defaultNewsUserChannelModel) RowBuilder() squirrel.SelectBuilder {
	return squirrel.Select("news_user_channel.user_id,news_user_channel.channel_id,news_channel.channel_name").From(m.table)
}