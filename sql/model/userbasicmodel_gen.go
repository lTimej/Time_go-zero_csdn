// Code generated by goctl. DO NOT EDIT.

package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	userBasicFieldNames          = builder.RawFieldNames(&UserBasic{})
	userBasicRows                = strings.Join(userBasicFieldNames, ",")
	userBasicRowsExpectAutoSet   = strings.Join(stringx.Remove(userBasicFieldNames, "`user_id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	userBasicRowsWithPlaceHolder = strings.Join(stringx.Remove(userBasicFieldNames, "`user_id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"

	cacheUserBasicUserIdPrefix   = "cache:userBasic:userId:"
	cacheUserBasicMobilePrefix   = "cache:userBasic:mobile:"
	cacheUserBasicUserNamePrefix = "cache:userBasic:userName:"
)

type (
	userBasicModel interface {
		Insert(ctx context.Context, data *UserBasic) (sql.Result, error)
		FindOne(ctx context.Context, userId int64) (*UserBasic, error)
		FindOneByMobile(ctx context.Context, mobile string) (*UserBasic, error)
		FindOneByUserName(ctx context.Context, userName string) (*UserBasic, error)
		Update(ctx context.Context, data *UserBasic) error
		Delete(ctx context.Context, userId int64) error
	}

	defaultUserBasicModel struct {
		sqlc.CachedConn
		table string
	}

	UserBasic struct {
		UserId         int64          `db:"user_id"`         // 用户ID
		Account        sql.NullString `db:"account"`         // 账号
		Email          sql.NullString `db:"email"`           // 邮箱
		Status         int64          `db:"status"`          // 状态，是否可用，0-不可用，1-可用
		Mobile         string         `db:"mobile"`          // 手机号
		Password       sql.NullString `db:"password"`        // 密码
		UserName       string         `db:"user_name"`       // 昵称
		ProfilePhoto   sql.NullString `db:"profile_photo"`   // 头像
		LastLogin      sql.NullTime   `db:"last_login"`      // 最后登录时间
		IsMedia        int64          `db:"is_media"`        // 是否是自媒体，0-不是，1-是
		IsVerified     int64          `db:"is_verified"`     // 是否实名认证，0-不是，1-是
		Introduction   sql.NullString `db:"introduction"`    // 简介
		Certificate    sql.NullString `db:"certificate"`     // 认证
		ArticleCount   int64          `db:"article_count"`   // 发文章数
		FollowingCount int64          `db:"following_count"` // 关注的人数
		FansCount      int64          `db:"fans_count"`      // 被关注的人数
		LikeCount      int64          `db:"like_count"`      // 累计点赞人数
		ReadCount      int64          `db:"read_count"`      // 累计阅读人数
		CodeYear       int64          `db:"code_year"`       // 码龄
	}
)

func newUserBasicModel(conn sqlx.SqlConn, c cache.CacheConf) *defaultUserBasicModel {
	return &defaultUserBasicModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`user_basic`",
	}
}

func (m *defaultUserBasicModel) Delete(ctx context.Context, userId int64) error {
	data, err := m.FindOne(ctx, userId)
	if err != nil {
		return err
	}

	userBasicMobileKey := fmt.Sprintf("%s%v", cacheUserBasicMobilePrefix, data.Mobile)
	userBasicUserIdKey := fmt.Sprintf("%s%v", cacheUserBasicUserIdPrefix, userId)
	userBasicUserNameKey := fmt.Sprintf("%s%v", cacheUserBasicUserNamePrefix, data.UserName)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `user_id` = ?", m.table)
		return conn.ExecCtx(ctx, query, userId)
	}, userBasicMobileKey, userBasicUserIdKey, userBasicUserNameKey)
	return err
}

func (m *defaultUserBasicModel) FindOne(ctx context.Context, userId int64) (*UserBasic, error) {
	userBasicUserIdKey := fmt.Sprintf("%s%v", cacheUserBasicUserIdPrefix, userId)
	var resp UserBasic
	err := m.QueryRowCtx(ctx, &resp, userBasicUserIdKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
		query := fmt.Sprintf("select %s from %s where `user_id` = ? limit 1", userBasicRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, userId)
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

func (m *defaultUserBasicModel) FindOneByMobile(ctx context.Context, mobile string) (*UserBasic, error) {
	userBasicMobileKey := fmt.Sprintf("%s%v", cacheUserBasicMobilePrefix, mobile)
	var resp UserBasic
	err := m.QueryRowIndexCtx(ctx, &resp, userBasicMobileKey, m.formatPrimary, func(ctx context.Context, conn sqlx.SqlConn, v any) (i any, e error) {
		query := fmt.Sprintf("select %s from %s where `mobile` = ? limit 1", userBasicRows, m.table)
		if err := conn.QueryRowCtx(ctx, &resp, query, mobile); err != nil {
			return nil, err
		}
		return resp.UserId, nil
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

func (m *defaultUserBasicModel) FindOneByUserName(ctx context.Context, userName string) (*UserBasic, error) {
	userBasicUserNameKey := fmt.Sprintf("%s%v", cacheUserBasicUserNamePrefix, userName)
	var resp UserBasic
	err := m.QueryRowIndexCtx(ctx, &resp, userBasicUserNameKey, m.formatPrimary, func(ctx context.Context, conn sqlx.SqlConn, v any) (i any, e error) {
		query := fmt.Sprintf("select %s from %s where `user_name` = ? limit 1", userBasicRows, m.table)
		if err := conn.QueryRowCtx(ctx, &resp, query, userName); err != nil {
			return nil, err
		}
		return resp.UserId, nil
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

func (m *defaultUserBasicModel) Insert(ctx context.Context, data *UserBasic) (sql.Result, error) {
	userBasicMobileKey := fmt.Sprintf("%s%v", cacheUserBasicMobilePrefix, data.Mobile)
	userBasicUserIdKey := fmt.Sprintf("%s%v", cacheUserBasicUserIdPrefix, data.UserId)
	userBasicUserNameKey := fmt.Sprintf("%s%v", cacheUserBasicUserNamePrefix, data.UserName)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, userBasicRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.Account, data.Email, data.Status, data.Mobile, data.Password, data.UserName, data.ProfilePhoto, data.LastLogin, data.IsMedia, data.IsVerified, data.Introduction, data.Certificate, data.ArticleCount, data.FollowingCount, data.FansCount, data.LikeCount, data.ReadCount, data.CodeYear)
	}, userBasicMobileKey, userBasicUserIdKey, userBasicUserNameKey)
	return ret, err
}

func (m *defaultUserBasicModel) Update(ctx context.Context, newData *UserBasic) error {
	data, err := m.FindOne(ctx, newData.UserId)
	if err != nil {
		return err
	}

	userBasicMobileKey := fmt.Sprintf("%s%v", cacheUserBasicMobilePrefix, data.Mobile)
	userBasicUserIdKey := fmt.Sprintf("%s%v", cacheUserBasicUserIdPrefix, data.UserId)
	userBasicUserNameKey := fmt.Sprintf("%s%v", cacheUserBasicUserNamePrefix, data.UserName)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `user_id` = ?", m.table, userBasicRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, newData.Account, newData.Email, newData.Status, newData.Mobile, newData.Password, newData.UserName, newData.ProfilePhoto, newData.LastLogin, newData.IsMedia, newData.IsVerified, newData.Introduction, newData.Certificate, newData.ArticleCount, newData.FollowingCount, newData.FansCount, newData.LikeCount, newData.ReadCount, newData.CodeYear, newData.UserId)
	}, userBasicMobileKey, userBasicUserIdKey, userBasicUserNameKey)
	return err
}

func (m *defaultUserBasicModel) formatPrimary(primary any) string {
	return fmt.Sprintf("%s%v", cacheUserBasicUserIdPrefix, primary)
}

func (m *defaultUserBasicModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary any) error {
	query := fmt.Sprintf("select %s from %s where `user_id` = ? limit 1", userBasicRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultUserBasicModel) tableName() string {
	return m.table
}