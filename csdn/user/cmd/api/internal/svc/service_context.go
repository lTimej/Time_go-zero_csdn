package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"liujun/Time_go-zero_csdn/csdn/user/cmd/api/internal/config"
	"liujun/Time_go-zero_csdn/csdn/user/cmd/rpc/userclient"
)

type ServiceContext struct {
	Config  config.Config
	UserRpc userclient.User
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:  c,
		UserRpc: userclient.NewUser(zrpc.MustNewClient(c.UserRpc)),
	}
}
