package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ NewsArticleStatisticModel = (*customNewsArticleStatisticModel)(nil)

type (
	// NewsArticleStatisticModel is an interface to be customized, add more methods here,
	// and implement the added methods in customNewsArticleStatisticModel.
	NewsArticleStatisticModel interface {
		newsArticleStatisticModel
	}

	customNewsArticleStatisticModel struct {
		*defaultNewsArticleStatisticModel
	}
)

// NewNewsArticleStatisticModel returns a model for the database table.
func NewNewsArticleStatisticModel(conn sqlx.SqlConn, c cache.CacheConf) NewsArticleStatisticModel {
	return &customNewsArticleStatisticModel{
		defaultNewsArticleStatisticModel: newNewsArticleStatisticModel(conn, c),
	}
}
