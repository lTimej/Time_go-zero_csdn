package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ TbSkuImageModel = (*customTbSkuImageModel)(nil)

type (
	// TbSkuImageModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTbSkuImageModel.
	TbSkuImageModel interface {
		tbSkuImageModel
	}

	customTbSkuImageModel struct {
		*defaultTbSkuImageModel
	}
)

// NewTbSkuImageModel returns a model for the database table.
func NewTbSkuImageModel(conn sqlx.SqlConn, c cache.CacheConf) TbSkuImageModel {
	return &customTbSkuImageModel{
		defaultTbSkuImageModel: newTbSkuImageModel(conn, c),
	}
}
