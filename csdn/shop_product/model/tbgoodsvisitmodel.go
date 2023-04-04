package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ TbGoodsVisitModel = (*customTbGoodsVisitModel)(nil)

type (
	// TbGoodsVisitModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTbGoodsVisitModel.
	TbGoodsVisitModel interface {
		tbGoodsVisitModel
	}

	customTbGoodsVisitModel struct {
		*defaultTbGoodsVisitModel
	}
)

// NewTbGoodsVisitModel returns a model for the database table.
func NewTbGoodsVisitModel(conn sqlx.SqlConn, c cache.CacheConf) TbGoodsVisitModel {
	return &customTbGoodsVisitModel{
		defaultTbGoodsVisitModel: newTbGoodsVisitModel(conn, c),
	}
}
