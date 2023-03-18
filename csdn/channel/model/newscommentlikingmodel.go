package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ NewsCommentLikingModel = (*customNewsCommentLikingModel)(nil)

type (
	// NewsCommentLikingModel is an interface to be customized, add more methods here,
	// and implement the added methods in customNewsCommentLikingModel.
	NewsCommentLikingModel interface {
		newsCommentLikingModel
	}

	customNewsCommentLikingModel struct {
		*defaultNewsCommentLikingModel
	}
)

// NewNewsCommentLikingModel returns a model for the database table.
func NewNewsCommentLikingModel(conn sqlx.SqlConn, c cache.CacheConf) NewsCommentLikingModel {
	return &customNewsCommentLikingModel{
		defaultNewsCommentLikingModel: newNewsCommentLikingModel(conn, c),
	}
}
