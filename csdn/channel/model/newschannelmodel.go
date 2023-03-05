package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ NewsChannelModel = (*customNewsChannelModel)(nil)

type (
	// NewsChannelModel is an interface to be customized, add more methods here,
	// and implement the added methods in customNewsChannelModel.
	NewsChannelModel interface {
		newsChannelModel
	}

	customNewsChannelModel struct {
		*defaultNewsChannelModel
	}
)

// NewNewsChannelModel returns a model for the database table.
func NewNewsChannelModel(conn sqlx.SqlConn, c cache.CacheConf) NewsChannelModel {
	return &customNewsChannelModel{
		defaultNewsChannelModel: newNewsChannelModel(conn, c),
	}
}
