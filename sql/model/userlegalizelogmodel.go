package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ UserLegalizeLogModel = (*customUserLegalizeLogModel)(nil)

type (
	// UserLegalizeLogModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserLegalizeLogModel.
	UserLegalizeLogModel interface {
		userLegalizeLogModel
	}

	customUserLegalizeLogModel struct {
		*defaultUserLegalizeLogModel
	}
)

// NewUserLegalizeLogModel returns a model for the database table.
func NewUserLegalizeLogModel(conn sqlx.SqlConn, c cache.CacheConf) UserLegalizeLogModel {
	return &customUserLegalizeLogModel{
		defaultUserLegalizeLogModel: newUserLegalizeLogModel(conn, c),
	}
}
