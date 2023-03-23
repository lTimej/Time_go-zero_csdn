package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	DB struct {
		DataSource string
	}
	Redis struct {
		Host string
		Pass string
		Type string
	}
	Es struct {
		Host string
	}
	JwtAuth struct {
		AccessSecret string
		AccessExpire int64
	}
	Cache cache.CacheConf
}
