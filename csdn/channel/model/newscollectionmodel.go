package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ NewsCollectionModel = (*customNewsCollectionModel)(nil)

type (
	// NewsCollectionModel is an interface to be customized, add more methods here,
	// and implement the added methods in customNewsCollectionModel.
	NewsCollectionModel interface {
		newsCollectionModel
	}

	customNewsCollectionModel struct {
		*defaultNewsCollectionModel
	}
)

// NewNewsCollectionModel returns a model for the database table.
func NewNewsCollectionModel(conn sqlx.SqlConn, c cache.CacheConf) NewsCollectionModel {
	return &customNewsCollectionModel{
		defaultNewsCollectionModel: newNewsCollectionModel(conn, c),
	}
}
