package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ TbProductIdModel = (*customTbProductIdModel)(nil)

type (
	// TbProductIdModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTbProductIdModel.
	TbProductIdModel interface {
		tbProductIdModel
	}

	customTbProductIdModel struct {
		*defaultTbProductIdModel
	}
)

// NewTbProductIdModel returns a model for the database table.
func NewTbProductIdModel(conn sqlx.SqlConn, c cache.CacheConf) TbProductIdModel {
	return &customTbProductIdModel{
		defaultTbProductIdModel: newTbProductIdModel(conn, c),
	}
}
