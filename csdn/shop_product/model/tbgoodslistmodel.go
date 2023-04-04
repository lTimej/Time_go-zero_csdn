package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ TbGoodsListModel = (*customTbGoodsListModel)(nil)

type (
	// TbGoodsListModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTbGoodsListModel.
	TbGoodsListModel interface {
		tbGoodsListModel
	}

	customTbGoodsListModel struct {
		*defaultTbGoodsListModel
	}
)

// NewTbGoodsListModel returns a model for the database table.
func NewTbGoodsListModel(conn sqlx.SqlConn, c cache.CacheConf) TbGoodsListModel {
	return &customTbGoodsListModel{
		defaultTbGoodsListModel: newTbGoodsListModel(conn, c),
	}
}
