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
	newsArticleAllInfo = "news_article_basic.article_id as article_id,news_article_basic.user_id as user_id,channel_id,title,news_article_basic.create_time as create_time,allow_comment,content,user_name,profile_photo,career,code_year"
	newsArticleBasicFieldNames          = builder.RawFieldNames(&NewsArticleBasic{})
	newsArticleBasicRows                = strings.Join(newsArticleBasicFieldNames, ",")
	newsArticleBasicRowsExpectAutoSet   = strings.Join(stringx.Remove(newsArticleBasicFieldNames, "`article_id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	newsArticleBasicRowsWithPlaceHolder = strings.Join(stringx.Remove(newsArticleBasicFieldNames, "`article_id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"

	cacheNewsArticleBasicArticleIdPrefix = "cache:newsArticleBasic:articleId:"
)

type (
	newsArticleBasicModel interface {
		RowBuilder() squirrel.SelectBuilder
		Insert(ctx context.Context, data *NewsArticleBasic) (sql.Result, error)
		FindOne(ctx context.Context, articleId int64) (*NewsArticleBasic, error)
		FindAllArticle(ctx context.Context,rowBuilder squirrel.SelectBuilder,channel_id int64,page,page_num int32)([]*AllArticleInfo,error)
		FindAllArticleByUserId(ctx context.Context,rowBuilder squirrel.SelectBuilder,user_id string,page,page_num int32)([]*AllArticleInfo,error)
		Update(ctx context.Context, data *NewsArticleBasic) error
		Delete(ctx context.Context, articleId int64) error
	}

	defaultNewsArticleBasicModel struct {
		sqlc.CachedConn
		table string
	}

	NewsArticleBasic struct {
		ArticleId     int64     `db:"article_id"`     // 文章ID
		UserId        string     `db:"user_id"`        // 用户ID
		ChannelId     int64     `db:"channel_id"`     // 频道ID
		Title         string    `db:"title"`          // 标题
		IsAdvertising int64     `db:"is_advertising"` // 是否投放广告，0-不投放，1-投放
		CreateTime    time.Time `db:"create_time"`    // 创建时间
		UpdateTime    time.Time `db:"update_time"`    // 更新时间
		Status        int64     `db:"status"`         // 贴文状态，0-草稿，1-待审核，2-审核通过，3-审核失败，4-已删除
		ReviewerId    int64     `db:"reviewer_id"`    // 审核人员ID
		ReviewTime    time.Time `db:"review_time"`    // 审核时间
		DeleteTime    time.Time `db:"delete_time"`    // 删除时间
		RejectReason  string    `db:"reject_reason"`  // 驳回原因
		CommentCount  int64     `db:"comment_count"`  // 累计评论数
		AllowComment  int64     `db:"allow_comment"`  // 是否允许评论，0-不允许，1-允许
	}
	AllArticleInfo struct {
		ArtId     int64     `db:"article_id"`     // 文章ID
		UserId        string     `db:"user_id"`        // 用户ID
		ChannelId     int64     `db:"channel_id"`     // 频道ID
		Title         string    `db:"title"`          // 标题
		CreateTime    string `db:"create_time"`    // 创建时间
		AllowComment  int32    `db:"allow_comment"`  // 是否允许评论，0-不允许，1-允许
		Content string `db:"content"`
		UserName string `db:"user_name"`
		HeadPhoto string `db:"profile_photo"`
		Career string `db:"career"`
		CodeYear int32 `db:"code_year"`
		//ReadNum int32 `db:"read_num"`
		//CommentNum int32 `db:"comment_num"`
		//LikeNum int32 `db:"like_num"`
		//CollectionNum int32 `db:"collection_num"`
	}
)

func newNewsArticleBasicModel(conn sqlx.SqlConn, c cache.CacheConf) *defaultNewsArticleBasicModel {
	return &defaultNewsArticleBasicModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`news_article_basic`",
	}
}

func (m *defaultNewsArticleBasicModel) Delete(ctx context.Context, articleId int64) error {
	newsArticleBasicArticleIdKey := fmt.Sprintf("%s%v", cacheNewsArticleBasicArticleIdPrefix, articleId)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `article_id` = ?", m.table)
		return conn.ExecCtx(ctx, query, articleId)
	}, newsArticleBasicArticleIdKey)
	return err
}

func (m *defaultNewsArticleBasicModel) FindOne(ctx context.Context, articleId int64) (*NewsArticleBasic, error) {
	newsArticleBasicArticleIdKey := fmt.Sprintf("%s%v", cacheNewsArticleBasicArticleIdPrefix, articleId)
	var resp NewsArticleBasic
	err := m.QueryRowCtx(ctx, &resp, newsArticleBasicArticleIdKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
		query := fmt.Sprintf("select %s from %s where `article_id` = ? limit 1", newsArticleBasicRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, articleId)
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
func (m *defaultNewsArticleBasicModel) FindAllArticle(ctx context.Context,rowBuilder squirrel.SelectBuilder,channel_id int64,page,page_num int32)([]*AllArticleInfo,error){
	if page <= 0{
		page = 1
	}
	offset := (page - 1) * page_num
	q,values,err := rowBuilder.Join("news_article_content,user_basic,user_profile where news_article_basic.article_id = news_article_content.article_id and news_article_basic.user_id = user_basic.user_id and user_basic.user_id = user_profile.user_id and news_article_basic.channel_id = %d limit %d,%d").ToSql()
	query := fmt.Sprintf(q,channel_id,offset,page_num)
	fmt.Println(query,"===============")
	if err != nil {
		return nil, err
	}

	var resp []*AllArticleInfo
	err = m.QueryRowsNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

func(m *defaultNewsArticleBasicModel) FindAllArticleByUserId(ctx context.Context,rowBuilder squirrel.SelectBuilder,user_id string,page,page_num int32)([]*AllArticleInfo,error){
	if page <= 0{
		page = 1
	}
	offset := (page - 1) * page_num
	q,values,err := rowBuilder.Join("news_article_content,user_basic,user_profile,news_collection where news_article_basic.article_id = news_article_content.article_id and news_article_basic.user_id = user_basic.user_id and user_basic.user_id = user_profile.user_id and news_collection.user_id = %s limit %d,%d").ToSql()
	query := fmt.Sprintf(q,user_id,offset,page_num)
	fmt.Println(query,"-------------")
	if err != nil {
		return nil, err
	}

	var resp []*AllArticleInfo
	err = m.QueryRowsNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

func (m *defaultNewsArticleBasicModel) Insert(ctx context.Context, data *NewsArticleBasic) (sql.Result, error) {
	newsArticleBasicArticleIdKey := fmt.Sprintf("%s%v", cacheNewsArticleBasicArticleIdPrefix, data.ArticleId)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, newsArticleBasicRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.UserId, data.ChannelId, data.Title, data.IsAdvertising, data.Status, data.ReviewerId, data.ReviewTime, data.DeleteTime, data.RejectReason, data.CommentCount, data.AllowComment)
	}, newsArticleBasicArticleIdKey)
	return ret, err
}

func (m *defaultNewsArticleBasicModel) Update(ctx context.Context, data *NewsArticleBasic) error {
	newsArticleBasicArticleIdKey := fmt.Sprintf("%s%v", cacheNewsArticleBasicArticleIdPrefix, data.ArticleId)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `article_id` = ?", m.table, newsArticleBasicRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.UserId, data.ChannelId, data.Title, data.IsAdvertising, data.Status, data.ReviewerId, data.ReviewTime, data.DeleteTime, data.RejectReason, data.CommentCount, data.AllowComment, data.ArticleId)
	}, newsArticleBasicArticleIdKey)
	return err
}

func (m *defaultNewsArticleBasicModel) formatPrimary(primary any) string {
	return fmt.Sprintf("%s%v", cacheNewsArticleBasicArticleIdPrefix, primary)
}

func (m *defaultNewsArticleBasicModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary any) error {
	query := fmt.Sprintf("select %s from %s where `article_id` = ? limit 1", newsArticleBasicRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultNewsArticleBasicModel) tableName() string {
	return m.table
}
func (m *defaultNewsArticleBasicModel) RowBuilder() squirrel.SelectBuilder {
	return squirrel.Select(newsArticleAllInfo).From(m.table)
}