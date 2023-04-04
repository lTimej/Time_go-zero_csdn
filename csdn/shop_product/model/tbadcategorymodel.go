package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ TbAdCategoryModel = (*customTbAdCategoryModel)(nil)

type (
	// TbAdCategoryModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTbAdCategoryModel.
	TbAdCategoryModel interface {
		tbAdCategoryModel
	}

	customTbAdCategoryModel struct {
		*defaultTbAdCategoryModel
	}
)

// NewTbAdCategoryModel returns a model for the database table.
func NewTbAdCategoryModel(conn sqlx.SqlConn, c cache.CacheConf) TbAdCategoryModel {
	return &customTbAdCategoryModel{
		defaultTbAdCategoryModel: newTbAdCategoryModel(conn, c),
	}
}
