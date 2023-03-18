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
	newsCollectionFieldNames          = builder.RawFieldNames(&NewsCollection{})
	newsCollectionRows                = strings.Join(newsCollectionFieldNames, ",")
	newsCollectionNum = "count(*) as c"
	newsCollectionRowsExpectAutoSet   = strings.Join(stringx.Remove(newsCollectionFieldNames, "`collection_id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	newsCollectionRowsWithPlaceHolder = strings.Join(stringx.Remove(newsCollectionFieldNames, "`collection_id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"
	cacheNewsArticleCollectionNumPrefix = "cache:newsCollection:collectionNum:"
	cacheNewsCollectionCollectionIdPrefix    = "cache:newsCollection:collectionId:"
	cacheNewsCollectionUserIdArticleIdPrefix = "cache:newsCollection:userId:articleId:"
)

type (
	newsCollectionModel interface {
		Insert(ctx context.Context, data *NewsCollection) (sql.Result, error)
		FindOne(ctx context.Context, collectionId int64) (*NewsCollection, error)
		FindArticleCollectionNum(ctx context.Context, ArticleId int64)(int64,error)
		FindOneByUserIdArticleId(ctx context.Context, userId string, articleId int64) (*NewsCollection, error)
		Update(ctx context.Context, data *NewsCollection) error
		Delete(ctx context.Context, collectionId int64) error
	}

	defaultNewsCollectionModel struct {
		sqlc.CachedConn
		table string
	}

	NewsCollection struct {
		CollectionId int64     `db:"collection_id"` // 主键id
		UserId       string     `db:"user_id"`       // 用户ID
		ArticleId    int64     `db:"article_id"`    // 文章ID
		CreateTime   time.Time `db:"create_time"`   // 创建时间
		IsDeleted    int64     `db:"is_deleted"`    // 是否取消收藏, 0-未取消, 1-已取消
		UpdateTime   time.Time `db:"update_time"`   // 更新时间
	}
	NewsCollectionNum struct{
		Count int64 `db:"c"`
	}
)

func newNewsCollectionModel(conn sqlx.SqlConn, c cache.CacheConf) *defaultNewsCollectionModel {
	return &defaultNewsCollectionModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`news_collection`",
	}
}

func (m *defaultNewsCollectionModel) Delete(ctx context.Context, collectionId int64) error {
	data, err := m.FindOne(ctx, collectionId)
	if err != nil {
		return err
	}

	newsCollectionCollectionIdKey := fmt.Sprintf("%s%v", cacheNewsCollectionCollectionIdPrefix, collectionId)
	newsCollectionUserIdArticleIdKey := fmt.Sprintf("%s%v:%v", cacheNewsCollectionUserIdArticleIdPrefix, data.UserId, data.ArticleId)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `collection_id` = ?", m.table)
		return conn.ExecCtx(ctx, query, collectionId)
	}, newsCollectionCollectionIdKey, newsCollectionUserIdArticleIdKey)
	return err
}

func (m *defaultNewsCollectionModel) FindOne(ctx context.Context, collectionId int64) (*NewsCollection, error) {
	newsCollectionCollectionIdKey := fmt.Sprintf("%s%v", cacheNewsCollectionCollectionIdPrefix, collectionId)
	var resp NewsCollection
	err := m.QueryRowCtx(ctx, &resp, newsCollectionCollectionIdKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
		query := fmt.Sprintf("select %s from %s where `collection_id` = ? limit 1", newsCollectionRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, collectionId)
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
func (m *defaultNewsCollectionModel)FindArticleCollectionNum(ctx context.Context, ArticleId int64)(int64,error){
	newsArticleCollectionNumKey := fmt.Sprintf("%s%v", cacheNewsArticleCollectionNumPrefix, ArticleId)
	var resp NewsCollectionNum
	err := m.QueryRowCtx(ctx, &resp, newsArticleCollectionNumKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
		query := fmt.Sprintf("select %s from %s where `article_id` = ? limit 1", newsCollectionNum, m.table)
		return conn.QueryRowCtx(ctx, v, query, ArticleId)
	})
	switch err {
	case nil:
		return resp.Count, nil
	case sqlc.ErrNotFound:
		return 0, nil
	default:
		return 0, err
	}
}
func (m *defaultNewsCollectionModel) FindOneByUserIdArticleId(ctx context.Context, userId string, articleId int64) (*NewsCollection, error) {
	newsCollectionUserIdArticleIdKey := fmt.Sprintf("%s%v:%v", cacheNewsCollectionUserIdArticleIdPrefix, userId, articleId)
	var resp NewsCollection
	err := m.QueryRowIndexCtx(ctx, &resp, newsCollectionUserIdArticleIdKey, m.formatPrimary, func(ctx context.Context, conn sqlx.SqlConn, v any) (i any, e error) {
		query := fmt.Sprintf("select %s from %s where `user_id` = ? and `article_id` = ? limit 1", newsCollectionRows, m.table)
		if err := conn.QueryRowCtx(ctx, &resp, query, userId, articleId); err != nil {
			return nil, err
		}
		return resp.CollectionId, nil
	}, m.queryPrimary)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, nil
	default:
		return nil, err
	}
}

func (m *defaultNewsCollectionModel) Insert(ctx context.Context, data *NewsCollection) (sql.Result, error) {
	newsCollectionCollectionIdKey := fmt.Sprintf("%s%v", cacheNewsCollectionCollectionIdPrefix, data.CollectionId)
	newsCollectionUserIdArticleIdKey := fmt.Sprintf("%s%v:%v", cacheNewsCollectionUserIdArticleIdPrefix, data.UserId, data.ArticleId)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?)", m.table, newsCollectionRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.UserId, data.ArticleId, data.IsDeleted)
	}, newsCollectionCollectionIdKey, newsCollectionUserIdArticleIdKey)
	if err == nil{
		m.FindOneByUserIdArticleId(ctx,data.UserId,data.ArticleId)
	}
	return ret, err
}

func (m *defaultNewsCollectionModel) Update(ctx context.Context, newData *NewsCollection) error {
	data, err := m.FindOneByUserIdArticleId(ctx, newData.UserId,newData.ArticleId)
	if err != nil {
		return err
	}

	newsCollectionCollectionIdKey := fmt.Sprintf("%s%v", cacheNewsCollectionCollectionIdPrefix, data.CollectionId)
	newsCollectionUserIdArticleIdKey := fmt.Sprintf("%s%v:%v", cacheNewsCollectionUserIdArticleIdPrefix, data.UserId, data.ArticleId)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `article_id` = ? and user_id = ?", m.table, "is_deleted = ?")
		return conn.ExecCtx(ctx, query, newData.IsDeleted, newData.ArticleId, newData.UserId)
	}, newsCollectionCollectionIdKey, newsCollectionUserIdArticleIdKey)
	return err
}

func (m *defaultNewsCollectionModel) formatPrimary(primary any) string {
	return fmt.Sprintf("%s%v", cacheNewsCollectionCollectionIdPrefix, primary)
}

func (m *defaultNewsCollectionModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary any) error {
	query := fmt.Sprintf("select %s from %s where `collection_id` = ? limit 1", newsCollectionRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultNewsCollectionModel) tableName() string {
	return m.table
}
