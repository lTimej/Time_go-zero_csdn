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
	userProfileFieldNames          = builder.RawFieldNames(&UserProfile{})
	userProfileRows                = strings.Join(userProfileFieldNames, ",")
	userProfileRowsExpectAutoSet   = strings.Join(stringx.Remove(userProfileFieldNames, "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	userProfileRowsWithPlaceHolder = strings.Join(stringx.Remove(userProfileFieldNames, "`user_id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"

	cacheUserProfileUserIdPrefix = "cache:userProfile:userId:"
)

type (
	userProfileModel interface {
		Insert(ctx context.Context, data *UserProfile) (sql.Result, error)
		FindOne(ctx context.Context, userId int64) (*UserProfile, error)
		Update(ctx context.Context, data *UserProfile) error
		Delete(ctx context.Context, userId int64) error
	}

	defaultUserProfileModel struct {
		sqlc.CachedConn
		table string
	}

	UserProfile struct {
		UserId            int64          `db:"user_id"`             // 用户ID
		Gender            int64          `db:"gender"`              // 性别，0-男，1-女
		Birthday          sql.NullTime   `db:"birthday"`            // 生日
		RealName          sql.NullString `db:"real_name"`           // 真实姓名
		IdNumber          sql.NullString `db:"id_number"`           // 身份证号
		IdCardFront       sql.NullString `db:"id_card_front"`       // 身份证正面
		IdCardBack        sql.NullString `db:"id_card_back"`        // 身份证背面
		IdCardHandheld    sql.NullString `db:"id_card_handheld"`    // 手持身份证
		CreateTime        time.Time      `db:"create_time"`         // 创建时间
		UpdateTime        time.Time      `db:"update_time"`         // 更新时间
		RegisterMediaTime sql.NullTime   `db:"register_media_time"` // 注册自媒体时间
		Area              sql.NullString `db:"area"`                // 地区
		Company           sql.NullString `db:"company"`             // 公司
		Career            sql.NullString `db:"career"`              // 职业
		Tag               sql.NullString `db:"tag"`                 // 标签
	}
)

func newUserProfileModel(conn sqlx.SqlConn, c cache.CacheConf) *defaultUserProfileModel {
	return &defaultUserProfileModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`user_profile`",
	}
}

func (m *defaultUserProfileModel) Delete(ctx context.Context, userId int64) error {
	userProfileUserIdKey := fmt.Sprintf("%s%v", cacheUserProfileUserIdPrefix, userId)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `user_id` = ?", m.table)
		return conn.ExecCtx(ctx, query, userId)
	}, userProfileUserIdKey)
	return err
}

func (m *defaultUserProfileModel) FindOne(ctx context.Context, userId int64) (*UserProfile, error) {
	userProfileUserIdKey := fmt.Sprintf("%s%v", cacheUserProfileUserIdPrefix, userId)
	var resp UserProfile
	err := m.QueryRowCtx(ctx, &resp, userProfileUserIdKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
		query := fmt.Sprintf("select %s from %s where `user_id` = ? limit 1", userProfileRows, m.table)
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

func (m *defaultUserProfileModel) Insert(ctx context.Context, data *UserProfile) (sql.Result, error) {
	userProfileUserIdKey := fmt.Sprintf("%s%v", cacheUserProfileUserIdPrefix, data.UserId)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, userProfileRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.UserId, data.Gender, data.Birthday, data.RealName, data.IdNumber, data.IdCardFront, data.IdCardBack, data.IdCardHandheld, data.RegisterMediaTime, data.Area, data.Company, data.Career, data.Tag)
	}, userProfileUserIdKey)
	return ret, err
}

func (m *defaultUserProfileModel) Update(ctx context.Context, data *UserProfile) error {
	userProfileUserIdKey := fmt.Sprintf("%s%v", cacheUserProfileUserIdPrefix, data.UserId)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `user_id` = ?", m.table, userProfileRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.Gender, data.Birthday, data.RealName, data.IdNumber, data.IdCardFront, data.IdCardBack, data.IdCardHandheld, data.RegisterMediaTime, data.Area, data.Company, data.Career, data.Tag, data.UserId)
	}, userProfileUserIdKey)
	return err
}

func (m *defaultUserProfileModel) formatPrimary(primary any) string {
	return fmt.Sprintf("%s%v", cacheUserProfileUserIdPrefix, primary)
}

func (m *defaultUserProfileModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary any) error {
	query := fmt.Sprintf("select %s from %s where `user_id` = ? limit 1", userProfileRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultUserProfileModel) tableName() string {
	return m.table
}
