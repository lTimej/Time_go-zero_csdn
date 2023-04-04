package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ TbSpecificationOptionModel = (*customTbSpecificationOptionModel)(nil)

type (
	// TbSpecificationOptionModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTbSpecificationOptionModel.
	TbSpecificationOptionModel interface {
		tbSpecificationOptionModel
	}

	customTbSpecificationOptionModel struct {
		*defaultTbSpecificationOptionModel
	}
)

// NewTbSpecificationOptionModel returns a model for the database table.
func NewTbSpecificationOptionModel(conn sqlx.SqlConn, c cache.CacheConf) TbSpecificationOptionModel {
	return &customTbSpecificationOptionModel{
		defaultTbSpecificationOptionModel: newTbSpecificationOptionModel(conn, c),
	}
}
