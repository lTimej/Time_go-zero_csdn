package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ TbSpuModel = (*customTbSpuModel)(nil)

type (
	// TbSpuModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTbSpuModel.
	TbSpuModel interface {
		tbSpuModel
	}

	customTbSpuModel struct {
		*defaultTbSpuModel
	}
)

// NewTbSpuModel returns a model for the database table.
func NewTbSpuModel(conn sqlx.SqlConn, c cache.CacheConf) TbSpuModel {
	return &customTbSpuModel{
		defaultTbSpuModel: newTbSpuModel(conn, c),
	}
}
