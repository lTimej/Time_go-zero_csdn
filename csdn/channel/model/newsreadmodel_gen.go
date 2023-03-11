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
	newsReadFieldNames          = builder.RawFieldNames(&NewsRead{})
	newsReadRows                = strings.Join(newsReadFieldNames, ",")
	newsReadRowsExpectAutoSet   = strings.Join(stringx.Remove(newsReadFieldNames, "`read_id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	newsReadRowsWithPlaceHolder = strings.Join(stringx.Remove(newsReadFieldNames, "`read_id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"

	cacheNewsReadReadIdPrefix          = "cache:newsRead:readId:"
	cacheNewsReadUserIdArticleIdPrefix = "cache:newsRead:userId:articleId:"
)

type (
	newsReadModel interface {
		Insert(ctx context.Context, data *NewsRead) (sql.Result, error)
		FindOne(ctx context.Context, readId int64) (*NewsRead, error)
		FindOneByUserIdArticleId(ctx context.Context, userId int64, articleId int64) (*NewsRead, error)
		Update(ctx context.Context, data *NewsRead) error
		Delete(ctx context.Context, readId int64) error
	}

	defaultNewsReadModel struct {
		sqlc.CachedConn
		table string
	}

	NewsRead struct {
		ReadId     int64     `db:"read_id"`     // 主键id
		UserId     int64     `db:"user_id"`     // 用户ID
		ArticleId  int64     `db:"article_id"`  // 文章ID
		CreateTime time.Time `db:"create_time"` // 创建时间
		UpdateTime time.Time `db:"update_time"` // 更新时间
	}
)

func newNewsReadModel(conn sqlx.SqlConn, c cache.CacheConf) *defaultNewsReadModel {
	return &defaultNewsReadModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`news_read`",
	}
}

func (m *defaultNewsReadModel) Delete(ctx context.Context, readId int64) error {
	data, err := m.FindOne(ctx, readId)
	if err != nil {
		return err
	}

	newsReadReadIdKey := fmt.Sprintf("%s%v", cacheNewsReadReadIdPrefix, readId)
	newsReadUserIdArticleIdKey := fmt.Sprintf("%s%v:%v", cacheNewsReadUserIdArticleIdPrefix, data.UserId, data.ArticleId)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `read_id` = ?", m.table)
		return conn.ExecCtx(ctx, query, readId)
	}, newsReadReadIdKey, newsReadUserIdArticleIdKey)
	return err
}

func (m *defaultNewsReadModel) FindOne(ctx context.Context, readId int64) (*NewsRead, error) {
	newsReadReadIdKey := fmt.Sprintf("%s%v", cacheNewsReadReadIdPrefix, readId)
	var resp NewsRead
	err := m.QueryRowCtx(ctx, &resp, newsReadReadIdKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
		query := fmt.Sprintf("select %s from %s where `read_id` = ? limit 1", newsReadRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, readId)
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

func (m *defaultNewsReadModel) FindOneByUserIdArticleId(ctx context.Context, userId int64, articleId int64) (*NewsRead, error) {
	newsReadUserIdArticleIdKey := fmt.Sprintf("%s%v:%v", cacheNewsReadUserIdArticleIdPrefix, userId, articleId)
	var resp NewsRead
	err := m.QueryRowIndexCtx(ctx, &resp, newsReadUserIdArticleIdKey, m.formatPrimary, func(ctx context.Context, conn sqlx.SqlConn, v any) (i any, e error) {
		query := fmt.Sprintf("select %s from %s where `user_id` = ? and `article_id` = ? limit 1", newsReadRows, m.table)
		if err := conn.QueryRowCtx(ctx, &resp, query, userId, articleId); err != nil {
			return nil, err
		}
		return resp.ReadId, nil
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

func (m *defaultNewsReadModel) Insert(ctx context.Context, data *NewsRead) (sql.Result, error) {
	newsReadReadIdKey := fmt.Sprintf("%s%v", cacheNewsReadReadIdPrefix, data.ReadId)
	newsReadUserIdArticleIdKey := fmt.Sprintf("%s%v:%v", cacheNewsReadUserIdArticleIdPrefix, data.UserId, data.ArticleId)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?)", m.table, newsReadRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.UserId, data.ArticleId)
	}, newsReadReadIdKey, newsReadUserIdArticleIdKey)
	return ret, err
}

func (m *defaultNewsReadModel) Update(ctx context.Context, newData *NewsRead) error {
	data, err := m.FindOne(ctx, newData.ReadId)
	if err != nil {
		return err
	}

	newsReadReadIdKey := fmt.Sprintf("%s%v", cacheNewsReadReadIdPrefix, data.ReadId)
	newsReadUserIdArticleIdKey := fmt.Sprintf("%s%v:%v", cacheNewsReadUserIdArticleIdPrefix, data.UserId, data.ArticleId)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `read_id` = ?", m.table, newsReadRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, newData.UserId, newData.ArticleId, newData.ReadId)
	}, newsReadReadIdKey, newsReadUserIdArticleIdKey)
	return err
}

func (m *defaultNewsReadModel) formatPrimary(primary any) string {
	return fmt.Sprintf("%s%v", cacheNewsReadReadIdPrefix, primary)
}

func (m *defaultNewsReadModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary any) error {
	query := fmt.Sprintf("select %s from %s where `read_id` = ? limit 1", newsReadRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultNewsReadModel) tableName() string {
	return m.table
}
