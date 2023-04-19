package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ CityModel = (*customCityModel)(nil)

type (
	// CityModel is an interface to be customized, add more methods here,
	// and implement the added methods in customCityModel.
	CityModel interface {
		cityModel
	}

	customCityModel struct {
		*defaultCityModel
	}
)

// NewCityModel returns a model for the database table.
func NewCityModel(conn sqlx.SqlConn, c cache.CacheConf) CityModel {
	return &customCityModel{
		defaultCityModel: newCityModel(conn, c),
	}
}
