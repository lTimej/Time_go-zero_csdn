package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ TbSpuSpecificationModel = (*customTbSpuSpecificationModel)(nil)

type (
	// TbSpuSpecificationModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTbSpuSpecificationModel.
	TbSpuSpecificationModel interface {
		tbSpuSpecificationModel
	}

	customTbSpuSpecificationModel struct {
		*defaultTbSpuSpecificationModel
	}
)

// NewTbSpuSpecificationModel returns a model for the database table.
func NewTbSpuSpecificationModel(conn sqlx.SqlConn, c cache.CacheConf) TbSpuSpecificationModel {
	return &customTbSpuSpecificationModel{
		defaultTbSpuSpecificationModel: newTbSpuSpecificationModel(conn, c),
	}
}
