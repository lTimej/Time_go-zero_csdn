package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ NewsReadModel = (*customNewsReadModel)(nil)

type (
	// NewsReadModel is an interface to be customized, add more methods here,
	// and implement the added methods in customNewsReadModel.
	NewsReadModel interface {
		newsReadModel
	}

	customNewsReadModel struct {
		*defaultNewsReadModel
	}
)

// NewNewsReadModel returns a model for the database table.
func NewNewsReadModel(conn sqlx.SqlConn, c cache.CacheConf) NewsReadModel {
	return &customNewsReadModel{
		defaultNewsReadModel: newNewsReadModel(conn, c),
	}
}
