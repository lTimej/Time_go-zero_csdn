package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	ProductRpc zrpc.RpcClientConf
	JwtAuth    struct {
		AccessSecret string
		AccessExpire int64
	}
	Redis struct {
		Host string
		Pass string
		Type string
	}
	DB struct {
		DataSource string
	}
	Cache cache.CacheConf
}
