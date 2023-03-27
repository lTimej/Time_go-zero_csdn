package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	rest.RestConf
	DB struct {
		DataSource string
	}
	RedisIm struct {
		Host        string
		Pass        string
		DB          int
		PoolSize    int
		MinIdleConn int
	}
	Redis struct {
		Host string
		Pass string
		Type string
	}
	JwtAuth struct {
		AccessSecret string
		AccessExpire int64
	}
	Cache cache.CacheConf
}
