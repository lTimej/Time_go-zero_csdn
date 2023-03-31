package redisIm

import (
	"liujun/Time_go-zero_csdn/csdn/im/cmd/api/internal/config"

	"github.com/go-redis/redis/v8"
)

var (
	RedisIm *redis.Client
)

func init() {
	var c config.Config
	RedisIm = redis.NewClient(&redis.Options{
		Addr:         c.RedisIm.Host,
		Password:     c.RedisIm.Pass,
		DB:           c.RedisIm.DB,
		PoolSize:     c.RedisIm.PoolSize,
		MinIdleConns: c.RedisIm.MinIdleConn,
	})
}
