package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ TbGoodsCategoryModel = (*customTbGoodsCategoryModel)(nil)

type (
	// TbGoodsCategoryModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTbGoodsCategoryModel.
	TbGoodsCategoryModel interface {
		tbGoodsCategoryModel
	}

	customTbGoodsCategoryModel struct {
		*defaultTbGoodsCategoryModel
	}
)

// NewTbGoodsCategoryModel returns a model for the database table.
func NewTbGoodsCategoryModel(conn sqlx.SqlConn, c cache.CacheConf) TbGoodsCategoryModel {
	return &customTbGoodsCategoryModel{
		defaultTbGoodsCategoryModel: newTbGoodsCategoryModel(conn, c),
	}
}
