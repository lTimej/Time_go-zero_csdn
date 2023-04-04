package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ TbSkuModel = (*customTbSkuModel)(nil)

type (
	// TbSkuModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTbSkuModel.
	TbSkuModel interface {
		tbSkuModel
	}

	customTbSkuModel struct {
		*defaultTbSkuModel
	}
)

// NewTbSkuModel returns a model for the database table.
func NewTbSkuModel(conn sqlx.SqlConn, c cache.CacheConf) TbSkuModel {
	return &customTbSkuModel{
		defaultTbSkuModel: newTbSkuModel(conn, c),
	}
}
