package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ TbSkuSpecificationModel = (*customTbSkuSpecificationModel)(nil)

type (
	// TbSkuSpecificationModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTbSkuSpecificationModel.
	TbSkuSpecificationModel interface {
		tbSkuSpecificationModel
	}

	customTbSkuSpecificationModel struct {
		*defaultTbSkuSpecificationModel
	}
)

// NewTbSkuSpecificationModel returns a model for the database table.
func NewTbSkuSpecificationModel(conn sqlx.SqlConn, c cache.CacheConf) TbSkuSpecificationModel {
	return &customTbSkuSpecificationModel{
		defaultTbSkuSpecificationModel: newTbSkuSpecificationModel(conn, c),
	}
}
