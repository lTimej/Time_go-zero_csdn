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
	newsCommentFieldNames          = builder.RawFieldNames(&NewsComment{})
	newsCommentRows                = strings.Join(newsCommentFieldNames, ",")
	newsCommentRowsExpectAutoSet   = strings.Join(stringx.Remove(newsCommentFieldNames, "`comment_id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	newsCommentRowsWithPlaceHolder = strings.Join(stringx.Remove(newsCommentFieldNames, "`comment_id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"

	cacheNewsCommentCommentIdPrefix = "cache:newsComment:commentId:"
	cacheNewsCommentArticleIdUserIdPrefix = "cache:newsComment:articleId:userId:"
)

type (
	newsCommentModel interface {
		RowBuilder() squirrel.SelectBuilder
		Insert(ctx context.Context, data *NewsComment) (sql.Result, error)
		FindOne(ctx context.Context, commentId int64) (*NewsComment, error)
		FindAll(ctx context.Context, builder squirrel.SelectBuilder) ([]*NewsComment, error)
		FindOneByArticleIdUserIdParentId(ctx context.Context, ArticleId int64,UserId string,ParentId int64) (*NewsComment, error)
		Update(ctx context.Context, data *NewsComment) error
		Delete(ctx context.Context, commentId int64) error
	}

	defaultNewsCommentModel struct {
		sqlc.CachedConn
		table string
	}

	NewsComment struct {
		CommentId  int64         `db:"comment_id"`  // 评论id
		UserId     string         `db:"user_id"`     // 用户ID
		ArticleId  int64         `db:"article_id"`  // 文章ID
		ParentId   int64 `db:"parent_id"`   // 评论ID
		LikeCount  int64         `db:"like_count"`  // 点赞数
		ReplyCount int64         `db:"reply_count"` // 回复数
		Content    string        `db:"content"`     // 评论内容
		IsTop      int64         `db:"is_top"`      // 是否置顶
		Status     int64         `db:"status"`      // 状态，0-待审核，1-审核通过，2-审核失败，3-已删除
		CreateTime time.Time     `db:"create_time"` // 创建时间
	}
)

func newNewsCommentModel(conn sqlx.SqlConn, c cache.CacheConf) *defaultNewsCommentModel {
	return &defaultNewsCommentModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`news_comment`",
	}
}

func (m *defaultNewsCommentModel) Delete(ctx context.Context, commentId int64) error {
	newsCommentCommentIdKey := fmt.Sprintf("%s%v", cacheNewsCommentCommentIdPrefix, commentId)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `comment_id` = ?", m.table)
		return conn.ExecCtx(ctx, query, commentId)
	}, newsCommentCommentIdKey)
	return err
}

func (m *defaultNewsCommentModel) FindOne(ctx context.Context, commentId int64) (*NewsComment, error) {
	newsCommentCommentIdKey := fmt.Sprintf("%s%v", cacheNewsCommentCommentIdPrefix, commentId)
	var resp NewsComment
	err := m.QueryRowCtx(ctx, &resp, newsCommentCommentIdKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
		query := fmt.Sprintf("select %s from %s where `comment_id` = ? limit 1", newsCommentRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, commentId)
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

func (m *defaultNewsCommentModel) FindAll(ctx context.Context, builder squirrel.SelectBuilder) ([]*NewsComment, error){
	query,values,err := builder.OrderBy("is_top desc,comment_id desc").ToSql()
	if err != nil {
		return nil, err
	}

	var resp []*NewsComment
	err = m.QueryRowsNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

func (m *defaultNewsCommentModel)FindOneByArticleIdUserIdParentId(ctx context.Context, ArticleId int64,UserId string,ParentId int64) (*NewsComment, error){
	newsCommentCommentIdKey := fmt.Sprintf("%s%v:%v:%v", cacheNewsCommentArticleIdUserIdPrefix, ArticleId,UserId,ParentId)
	var resp NewsComment
	err := m.QueryRowCtx(ctx, &resp, newsCommentCommentIdKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
		query := fmt.Sprintf("select %s from %s where `article_id` = ? and `user_id` = ? and `parent_id` = ? limit 1", newsCommentRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, ArticleId,UserId,ParentId)
	})
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, nil
	default:
		return nil, err
	}
}

func (m *defaultNewsCommentModel) Insert(ctx context.Context, data *NewsComment) (sql.Result, error) {
	newsCommentCommentIdKey := fmt.Sprintf("%s%v:%v:%v", cacheNewsCommentArticleIdUserIdPrefix, data.ArticleId,data.UserId,data.ParentId)
	//newsCommentCommentIdKey := fmt.Sprintf("%s%v", cacheNewsCommentCommentIdPrefix, data.CommentId)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?)", m.table, newsCommentRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.UserId, data.ArticleId, data.ParentId, data.LikeCount, data.ReplyCount, data.Content, data.IsTop, data.Status)
	}, newsCommentCommentIdKey)
	return ret, err
}

func (m *defaultNewsCommentModel) Update(ctx context.Context, data *NewsComment) error {
	newsCommentCommentIdKey := fmt.Sprintf("%s%v", cacheNewsCommentCommentIdPrefix, data.CommentId)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `comment_id` = ?", m.table, newsCommentRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.UserId, data.ArticleId, data.ParentId, data.LikeCount, data.ReplyCount, data.Content, data.IsTop, data.Status, data.CommentId)
	}, newsCommentCommentIdKey)
	return err
}

func (m *defaultNewsCommentModel) formatPrimary(primary any) string {
	return fmt.Sprintf("%s%v", cacheNewsCommentCommentIdPrefix, primary)
}

func (m *defaultNewsCommentModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary any) error {
	query := fmt.Sprintf("select %s from %s where `comment_id` = ? limit 1", newsCommentRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultNewsCommentModel) tableName() string {
	return m.table
}

func (m *defaultNewsCommentModel) RowBuilder() squirrel.SelectBuilder {
	return squirrel.Select(newsCommentRows).From(m.table)
}