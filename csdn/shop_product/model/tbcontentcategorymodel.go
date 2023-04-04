package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ TbContentCategoryModel = (*customTbContentCategoryModel)(nil)

type (
	// TbContentCategoryModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTbContentCategoryModel.
	TbContentCategoryModel interface {
		tbContentCategoryModel
	}

	customTbContentCategoryModel struct {
		*defaultTbContentCategoryModel
	}
)

// NewTbContentCategoryModel returns a model for the database table.
func NewTbContentCategoryModel(conn sqlx.SqlConn, c cache.CacheConf) TbContentCategoryModel {
	return &customTbContentCategoryModel{
		defaultTbContentCategoryModel: newTbContentCategoryModel(conn, c),
	}
}
