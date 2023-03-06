package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ NewsArticleContentModel = (*customNewsArticleContentModel)(nil)

type (
	// NewsArticleContentModel is an interface to be customized, add more methods here,
	// and implement the added methods in customNewsArticleContentModel.
	NewsArticleContentModel interface {
		newsArticleContentModel
	}

	customNewsArticleContentModel struct {
		*defaultNewsArticleContentModel
	}
)

// NewNewsArticleContentModel returns a model for the database table.
func NewNewsArticleContentModel(conn sqlx.SqlConn, c cache.CacheConf) NewsArticleContentModel {
	return &customNewsArticleContentModel{
		defaultNewsArticleContentModel: newNewsArticleContentModel(conn, c),
	}
}
