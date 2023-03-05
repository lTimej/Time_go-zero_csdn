package svc

import (
	"liujun/Time_go-zero_csdn/csdn/user/cmd/api/internal/config"
	"liujun/Time_go-zero_csdn/csdn/user/cmd/api/internal/middleware"
	"liujun/Time_go-zero_csdn/csdn/user/cmd/rpc/userclient"

	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config                config.Config
	UserRpc               userclient.User
	SetUidToCtxMiddleware rest.Middleware
	RedisClient           *redis.Redis
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:                c,
		UserRpc:               userclient.NewUser(zrpc.MustNewClient(c.UserRpc)),
		SetUidToCtxMiddleware: middleware.NewAuthMiddleWare().Handle,
		RedisClient: redis.New(c.Redis.Host, func(r *redis.Redis) {
			r.Type = c.Redis.Type
			r.Pass = c.Redis.Pass
		}),
	}
}
