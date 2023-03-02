package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ UserQualificationModel = (*customUserQualificationModel)(nil)

type (
	// UserQualificationModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserQualificationModel.
	UserQualificationModel interface {
		userQualificationModel
	}

	customUserQualificationModel struct {
		*defaultUserQualificationModel
	}
)

// NewUserQualificationModel returns a model for the database table.
func NewUserQualificationModel(conn sqlx.SqlConn, c cache.CacheConf) UserQualificationModel {
	return &customUserQualificationModel{
		defaultUserQualificationModel: newUserQualificationModel(conn, c),
	}
}
