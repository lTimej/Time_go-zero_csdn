package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ NewsCommentModel = (*customNewsCommentModel)(nil)

type (
	// NewsCommentModel is an interface to be customized, add more methods here,
	// and implement the added methods in customNewsCommentModel.
	NewsCommentModel interface {
		newsCommentModel
	}

	customNewsCommentModel struct {
		*defaultNewsCommentModel
	}
)

// NewNewsCommentModel returns a model for the database table.
func NewNewsCommentModel(conn sqlx.SqlConn, c cache.CacheConf) NewsCommentModel {
	return &customNewsCommentModel{
		defaultNewsCommentModel: newNewsCommentModel(conn, c),
	}
}
