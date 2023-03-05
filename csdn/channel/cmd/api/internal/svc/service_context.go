package svc

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
	"liujun/Time_go-zero_csdn/csdn/channel/cmd/api/internal/config"
	"liujun/Time_go-zero_csdn/csdn/channel/cmd/api/internal/middleware"
	"liujun/Time_go-zero_csdn/csdn/channel/cmd/rpc/channelclient"
)

type ServiceContext struct {
	Config                config.Config
	ChannelRpc            channelclient.Channel
	SetUidToCtxMiddleware rest.Middleware
	RedisClient           *redis.Redis
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:                c,
		SetUidToCtxMiddleware: middleware.NewSetUidToCtxMiddleware().Handle,
		ChannelRpc:            channelclient.NewChannel(zrpc.MustNewClient(c.ChannelRpc)),
		RedisClient: redis.New(c.Redis.Host, func(r *redis.Redis) {
			r.Type = c.Redis.Type
			r.Pass = c.Redis.Pass
		}),
	}
}
