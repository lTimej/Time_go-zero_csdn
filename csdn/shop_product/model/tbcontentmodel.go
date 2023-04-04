package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ TbContentModel = (*customTbContentModel)(nil)

type (
	// TbContentModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTbContentModel.
	TbContentModel interface {
		tbContentModel
	}

	customTbContentModel struct {
		*defaultTbContentModel
	}
)

// NewTbContentModel returns a model for the database table.
func NewTbContentModel(conn sqlx.SqlConn, c cache.CacheConf) TbContentModel {
	return &customTbContentModel{
		defaultTbContentModel: newTbContentModel(conn, c),
	}
}
