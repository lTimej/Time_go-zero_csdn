package svc

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"liujun/Time_go-zero_csdn/csdn/channel/cmd/rpc/internal/config"
	"liujun/Time_go-zero_csdn/csdn/channel/model"
)

type ServiceContext struct {
	Config       config.Config
	RedisClient  *redis.Redis
	ChannelModel model.NewsChannelModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	sqlConn := sqlx.NewMysql(c.DB.DataSource)
	return &ServiceContext{
		Config: c,
		RedisClient: redis.New(c.Redis.Host, func(r *redis.Redis) {
			r.Type = c.Redis.Type
			r.Pass = c.Redis.Pass
		}),
		ChannelModel: model.NewNewsChannelModel(sqlConn, c.Cache),
	}
}
