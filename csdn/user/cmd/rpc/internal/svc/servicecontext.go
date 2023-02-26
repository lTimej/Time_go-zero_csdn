package svc

import (
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
	"liujun/Time_go-zero_csdn/csdn/user/cmd/rpc/internal/config"
	"liujun/Time_go-zero_csdn/csdn/user/model"
)

type ServiceContext struct {
	Config    config.Config
	UserMysql *gorm.DB
	UserRedis *redis.Client
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:    c,
		UserMysql: model.InitMysql(c.Mysql.DataSource),
		UserRedis: model.InitRedis(c.Redis.Host, c.Redis.Password, c.Redis.DB),
	}
}
