package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ UserMaterialModel = (*customUserMaterialModel)(nil)

type (
	// UserMaterialModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserMaterialModel.
	UserMaterialModel interface {
		userMaterialModel
	}

	customUserMaterialModel struct {
		*defaultUserMaterialModel
	}
)

// NewUserMaterialModel returns a model for the database table.
func NewUserMaterialModel(conn sqlx.SqlConn, c cache.CacheConf) UserMaterialModel {
	return &customUserMaterialModel{
		defaultUserMaterialModel: newUserMaterialModel(conn, c),
	}
}
