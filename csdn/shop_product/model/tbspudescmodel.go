package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ TbSpuDescModel = (*customTbSpuDescModel)(nil)

type (
	// TbSpuDescModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTbSpuDescModel.
	TbSpuDescModel interface {
		tbSpuDescModel
	}

	customTbSpuDescModel struct {
		*defaultTbSpuDescModel
	}
)

// NewTbSpuDescModel returns a model for the database table.
func NewTbSpuDescModel(conn sqlx.SqlConn, c cache.CacheConf) TbSpuDescModel {
	return &customTbSpuDescModel{
		defaultTbSpuDescModel: newTbSpuDescModel(conn, c),
	}
}
