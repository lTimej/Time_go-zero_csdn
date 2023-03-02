package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ UserSearchModel = (*customUserSearchModel)(nil)

type (
	// UserSearchModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserSearchModel.
	UserSearchModel interface {
		userSearchModel
	}

	customUserSearchModel struct {
		*defaultUserSearchModel
	}
)

// NewUserSearchModel returns a model for the database table.
func NewUserSearchModel(conn sqlx.SqlConn, c cache.CacheConf) UserSearchModel {
	return &customUserSearchModel{
		defaultUserSearchModel: newUserSearchModel(conn, c),
	}
}
