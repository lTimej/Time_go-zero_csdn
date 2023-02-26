package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ UserVisitorsModel = (*customUserVisitorsModel)(nil)

type (
	// UserVisitorsModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserVisitorsModel.
	UserVisitorsModel interface {
		userVisitorsModel
	}

	customUserVisitorsModel struct {
		*defaultUserVisitorsModel
	}
)

// NewUserVisitorsModel returns a model for the database table.
func NewUserVisitorsModel(conn sqlx.SqlConn, c cache.CacheConf) UserVisitorsModel {
	return &customUserVisitorsModel{
		defaultUserVisitorsModel: newUserVisitorsModel(conn, c),
	}
}
