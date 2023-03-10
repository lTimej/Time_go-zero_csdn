package config

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	ChannelRpc zrpc.RpcClientConf
	JwtAuth    struct {
		AccessSecret string
	}
	Redis struct {
		Host string
		Pass string
		Type string
	}
}
