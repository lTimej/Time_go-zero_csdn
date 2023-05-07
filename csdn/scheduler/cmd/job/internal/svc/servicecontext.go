package svc

import (
	"liujun/Time_go-zero_csdn/csdn/channel/cmd/rpc/channelclient"
	"liujun/Time_go-zero_csdn/csdn/order/cmd/rpc/orderclient"
	"liujun/Time_go-zero_csdn/csdn/scheduler/cmd/job/internal/config"
	"liujun/Time_go-zero_csdn/csdn/user/cmd/rpc/userclient"

	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"

	"github.com/hibiken/asynq"
)

type ServiceContext struct {
	Config      config.Config
	AsynqServer *asynq.Server
	ChannelRpc  channelclient.Channel
	UserRpc     userclient.User
	RedisClient *redis.Redis
	OrderRpc    orderclient.Order
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:      c,
		AsynqServer: newAsynqServer(c),
		ChannelRpc:  channelclient.NewChannel(zrpc.MustNewClient(c.ChannelRpcConf)),
		UserRpc:     userclient.NewUser(zrpc.MustNewClient(c.UserRpcConf)),
		RedisClient: redis.New(c.Redis.Host, func(r *redis.Redis) {
			r.Type = c.Redis.Type
			r.Pass = c.Redis.Pass
		}),
		OrderRpc: orderclient.NewOrder(zrpc.MustNewClient(c.OrderRpcConf)),
	}
}
