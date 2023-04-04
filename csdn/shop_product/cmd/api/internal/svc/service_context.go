package svc

import (
	"liujun/Time_go-zero_csdn/csdn/shop_product/cmd/api/internal/config"
	"liujun/Time_go-zero_csdn/csdn/shop_product/cmd/api/internal/middleware"
	"liujun/Time_go-zero_csdn/csdn/shop_product/cmd/rpc/productclient"

	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config                config.Config
	ProductRpc            productclient.Product
	SetUidToCtxMiddleware rest.Middleware
	RedisClient           *redis.Redis
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:                c,
		SetUidToCtxMiddleware: middleware.NewSetUidToCtxMiddleware(c).Handle,
		ProductRpc:            productclient.NewProduct(zrpc.MustNewClient(c.ProductRpc)),
		RedisClient: redis.New(c.Redis.Host, func(r *redis.Redis) {
			r.Type = c.Redis.Type
			r.Pass = c.Redis.Pass
		}),
	}
}
