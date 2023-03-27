package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ ContactModel = (*customContactModel)(nil)

type (
	// ContactModel is an interface to be customized, add more methods here,
	// and implement the added methods in customContactModel.
	ContactModel interface {
		contactModel
	}

	customContactModel struct {
		*defaultContactModel
	}
)

// NewContactModel returns a model for the database table.
func NewContactModel(conn sqlx.SqlConn, c cache.CacheConf) ContactModel {
	return &customContactModel{
		defaultContactModel: newContactModel(conn, c),
	}
}
