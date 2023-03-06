package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ NewsArticleBasicModel = (*customNewsArticleBasicModel)(nil)

type (
	// NewsArticleBasicModel is an interface to be customized, add more methods here,
	// and implement the added methods in customNewsArticleBasicModel.
	NewsArticleBasicModel interface {
		newsArticleBasicModel
	}

	customNewsArticleBasicModel struct {
		*defaultNewsArticleBasicModel
	}
)

// NewNewsArticleBasicModel returns a model for the database table.
func NewNewsArticleBasicModel(conn sqlx.SqlConn, c cache.CacheConf) NewsArticleBasicModel {
	return &customNewsArticleBasicModel{
		defaultNewsArticleBasicModel: newNewsArticleBasicModel(conn, c),
	}
}
