package svc

import (
	"liujun/Time_go-zero_csdn/csdn/im/cmd/api/internal/config"
	"liujun/Time_go-zero_csdn/csdn/im/cmd/api/internal/middleware"
	"liujun/Time_go-zero_csdn/csdn/im/cmd/rpc/imclient"

	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config                config.Config
	ImRpc                 imclient.Im
	SetUidToCtxMiddleware rest.Middleware
	RedisClient           *redis.Redis
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:                c,
		SetUidToCtxMiddleware: middleware.NewSetUidToCtxMiddleware(c).Handle,
		ImRpc:                 imclient.NewIm(zrpc.MustNewClient(c.ImRpc)),
		RedisClient: redis.New(c.Redis.Host, func(r *redis.Redis) {
			r.Type = c.Redis.Type
			r.Pass = c.Redis.Pass
		}),
	}
}
