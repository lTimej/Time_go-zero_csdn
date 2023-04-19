package svc

import (
	"liujun/Time_go-zero_csdn/csdn/city/cmd/api/internal/config"
        "liujun/Time_go-zero_csdn/csdn/city/model"

        "github.com/zeromicro/go-zero/core/stores/redis"
        "github.com/zeromicro/go-zero/core/stores/sqlx"

)

type ServiceContext struct {
	Config config.Config
        RedisClient           *redis.Redis
        CityModel       model.CityModel
}

func NewServiceContext(c config.Config) *ServiceContext {
        sqlConn := sqlx.NewMysql(c.DB.DataSource)
	return &ServiceContext{
		Config: c,
                CityModel:       model.NewCityModel(sqlConn, c.Cache),
                RedisClient: redis.New(c.Redis.Host, func(r *redis.Redis) {
                        r.Type = c.Redis.Type
                        r.Pass = c.Redis.Pass
                }),
	}
}
