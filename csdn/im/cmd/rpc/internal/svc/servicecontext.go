package svc

import (
	"liujun/Time_go-zero_csdn/csdn/im/cmd/rpc/internal/config"
	"liujun/Time_go-zero_csdn/csdn/im/model"
	usermodel "liujun/Time_go-zero_csdn/csdn/user/model"

	redisclient "github.com/go-redis/redis/v8"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config      config.Config
	UserContact model.ContactModel
	UserBasic   usermodel.UserBasicModel
	RedisIm     *redisclient.Client
	RedisClient *redis.Redis
}

func NewServiceContext(c config.Config) *ServiceContext {
	sqlConn := sqlx.NewMysql(c.DB.DataSource)
	return &ServiceContext{
		Config:      c,
		UserContact: model.NewContactModel(sqlConn, c.Cache),
		RedisIm: redisclient.NewClient(&redisclient.Options{
			Addr:         c.RedisIm.Host,
			Password:     c.RedisIm.Pass,
			DB:           c.RedisIm.DB,
			PoolSize:     c.RedisIm.PoolSize,
			MinIdleConns: c.RedisIm.MinIdleConn,
		}),
		RedisClient: redis.New(c.Redis.Host, func(r *redis.Redis) {
			r.Type = c.Redis.Type
			r.Pass = c.Redis.Pass
		}),
		UserBasic: usermodel.NewUserBasicModel(sqlConn, c.Cache),
	}
}
