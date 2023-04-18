package svc

import (
	"liujun/Time_go-zero_csdn/csdn/order/cmd/api/internal/config"
	"liujun/Time_go-zero_csdn/csdn/order/cmd/api/internal/middleware"
	"liujun/Time_go-zero_csdn/csdn/order/cmd/rpc/orderclient"

	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config                config.Config
	OrderRpc              orderclient.Order
	RedisClient           *redis.Redis
	SetUidToCtxMiddleware rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:                c,
		OrderRpc:              orderclient.NewOrder(zrpc.MustNewClient(c.OrderRpc)),
		SetUidToCtxMiddleware: middleware.NewSetUidToCtxMiddleware(c).Handle,
		RedisClient: redis.New(c.Redis.Host, func(r *redis.Redis) {
			r.Type = c.Redis.Type
			r.Pass = c.Redis.Pass
		}),
	}
}
