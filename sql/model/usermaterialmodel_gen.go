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
	userMaterialFieldNames          = builder.RawFieldNames(&UserMaterial{})
	userMaterialRows                = strings.Join(userMaterialFieldNames, ",")
	userMaterialRowsExpectAutoSet   = strings.Join(stringx.Remove(userMaterialFieldNames, "`material_id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	userMaterialRowsWithPlaceHolder = strings.Join(stringx.Remove(userMaterialFieldNames, "`material_id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"

	cacheUserMaterialMaterialIdPrefix = "cache:userMaterial:materialId:"
	cacheUserMaterialUserIdHashPrefix = "cache:userMaterial:userId:hash:"
)

type (
	userMaterialModel interface {
		Insert(ctx context.Context, data *UserMaterial) (sql.Result, error)
		FindOne(ctx context.Context, materialId int64) (*UserMaterial, error)
		FindOneByUserIdHash(ctx context.Context, userId int64, hash sql.NullString) (*UserMaterial, error)
		Update(ctx context.Context, data *UserMaterial) error
		Delete(ctx context.Context, materialId int64) error
	}

	defaultUserMaterialModel struct {
		sqlc.CachedConn
		table string
	}

	UserMaterial struct {
		MaterialId  int64          `db:"material_id"`  // 素材id
		UserId      int64          `db:"user_id"`      // 用户ID
		Type        int64          `db:"type"`         // 素材类型，0-图片, 1-视频, 2-音频
		Hash        sql.NullString `db:"hash"`         // 素材指纹
		Url         string         `db:"url"`          // 素材链接地址
		CreateTime  time.Time      `db:"create_time"`  // 创建时间
		Status      int64          `db:"status"`       // 状态，0-待审核，1-审核通过，2-审核失败，3-已删除
		ReviewerId  sql.NullInt64  `db:"reviewer_id"`  // 审核人员ID
		ReviewTime  sql.NullTime   `db:"review_time"`  // 审核时间
		IsCollected int64          `db:"is_collected"` // 是否收藏，0-未收藏，1-已收藏
		UpdateTime  time.Time      `db:"update_time"`  // 更新时间
	}
)

func newUserMaterialModel(conn sqlx.SqlConn, c cache.CacheConf) *defaultUserMaterialModel {
	return &defaultUserMaterialModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`user_material`",
	}
}

func (m *defaultUserMaterialModel) Delete(ctx context.Context, materialId int64) error {
	data, err := m.FindOne(ctx, materialId)
	if err != nil {
		return err
	}

	userMaterialMaterialIdKey := fmt.Sprintf("%s%v", cacheUserMaterialMaterialIdPrefix, materialId)
	userMaterialUserIdHashKey := fmt.Sprintf("%s%v:%v", cacheUserMaterialUserIdHashPrefix, data.UserId, data.Hash)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `material_id` = ?", m.table)
		return conn.ExecCtx(ctx, query, materialId)
	}, userMaterialMaterialIdKey, userMaterialUserIdHashKey)
	return err
}

func (m *defaultUserMaterialModel) FindOne(ctx context.Context, materialId int64) (*UserMaterial, error) {
	userMaterialMaterialIdKey := fmt.Sprintf("%s%v", cacheUserMaterialMaterialIdPrefix, materialId)
	var resp UserMaterial
	err := m.QueryRowCtx(ctx, &resp, userMaterialMaterialIdKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
		query := fmt.Sprintf("select %s from %s where `material_id` = ? limit 1", userMaterialRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, materialId)
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

func (m *defaultUserMaterialModel) FindOneByUserIdHash(ctx context.Context, userId int64, hash sql.NullString) (*UserMaterial, error) {
	userMaterialUserIdHashKey := fmt.Sprintf("%s%v:%v", cacheUserMaterialUserIdHashPrefix, userId, hash)
	var resp UserMaterial
	err := m.QueryRowIndexCtx(ctx, &resp, userMaterialUserIdHashKey, m.formatPrimary, func(ctx context.Context, conn sqlx.SqlConn, v any) (i any, e error) {
		query := fmt.Sprintf("select %s from %s where `user_id` = ? and `hash` = ? limit 1", userMaterialRows, m.table)
		if err := conn.QueryRowCtx(ctx, &resp, query, userId, hash); err != nil {
			return nil, err
		}
		return resp.MaterialId, nil
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

func (m *defaultUserMaterialModel) Insert(ctx context.Context, data *UserMaterial) (sql.Result, error) {
	userMaterialMaterialIdKey := fmt.Sprintf("%s%v", cacheUserMaterialMaterialIdPrefix, data.MaterialId)
	userMaterialUserIdHashKey := fmt.Sprintf("%s%v:%v", cacheUserMaterialUserIdHashPrefix, data.UserId, data.Hash)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?)", m.table, userMaterialRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.UserId, data.Type, data.Hash, data.Url, data.Status, data.ReviewerId, data.ReviewTime, data.IsCollected)
	}, userMaterialMaterialIdKey, userMaterialUserIdHashKey)
	return ret, err
}

func (m *defaultUserMaterialModel) Update(ctx context.Context, newData *UserMaterial) error {
	data, err := m.FindOne(ctx, newData.MaterialId)
	if err != nil {
		return err
	}

	userMaterialMaterialIdKey := fmt.Sprintf("%s%v", cacheUserMaterialMaterialIdPrefix, data.MaterialId)
	userMaterialUserIdHashKey := fmt.Sprintf("%s%v:%v", cacheUserMaterialUserIdHashPrefix, data.UserId, data.Hash)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `material_id` = ?", m.table, userMaterialRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, newData.UserId, newData.Type, newData.Hash, newData.Url, newData.Status, newData.ReviewerId, newData.ReviewTime, newData.IsCollected, newData.MaterialId)
	}, userMaterialMaterialIdKey, userMaterialUserIdHashKey)
	return err
}

func (m *defaultUserMaterialModel) formatPrimary(primary any) string {
	return fmt.Sprintf("%s%v", cacheUserMaterialMaterialIdPrefix, primary)
}

func (m *defaultUserMaterialModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary any) error {
	query := fmt.Sprintf("select %s from %s where `material_id` = ? limit 1", userMaterialRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultUserMaterialModel) tableName() string {
	return m.table
}
