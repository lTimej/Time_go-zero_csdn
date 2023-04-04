package svc

import "liujun/Time_go-zero_csdn/csdn/shop_product/cmd/rpc/internal/config"

type ServiceContext struct {
	Config config.Config
	RedisClient            *redis.Redis
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		RedisClient: redis.New(c.Redis.Host, func(r *redis.Redis) {
			r.Type = c.Redis.Type
			r.Pass = c.Redis.Pass
		}),
	}
}
