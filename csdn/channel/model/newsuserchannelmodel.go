package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ NewsUserChannelModel = (*customNewsUserChannelModel)(nil)

type (
	// NewsUserChannelModel is an interface to be customized, add more methods here,
	// and implement the added methods in customNewsUserChannelModel.
	NewsUserChannelModel interface {
		newsUserChannelModel
	}

	customNewsUserChannelModel struct {
		*defaultNewsUserChannelModel
	}
)

// NewNewsUserChannelModel returns a model for the database table.
func NewNewsUserChannelModel(conn sqlx.SqlConn, c cache.CacheConf) NewsUserChannelModel {
	return &customNewsUserChannelModel{
		defaultNewsUserChannelModel: newNewsUserChannelModel(conn, c),
	}
}
