package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ NewsAttitudeModel = (*customNewsAttitudeModel)(nil)

type (
	// NewsAttitudeModel is an interface to be customized, add more methods here,
	// and implement the added methods in customNewsAttitudeModel.
	NewsAttitudeModel interface {
		newsAttitudeModel
	}

	customNewsAttitudeModel struct {
		*defaultNewsAttitudeModel
	}
)

// NewNewsAttitudeModel returns a model for the database table.
func NewNewsAttitudeModel(conn sqlx.SqlConn, c cache.CacheConf) NewsAttitudeModel {
	return &customNewsAttitudeModel{
		defaultNewsAttitudeModel: newNewsAttitudeModel(conn, c),
	}
}
