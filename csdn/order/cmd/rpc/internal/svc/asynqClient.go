package svc

import (
	"liujun/Time_go-zero_csdn/csdn/order/cmd/rpc/internal/config"

	"github.com/hibiken/asynq"
)

// create asynq client.
func newAsynqClient(c config.Config) *asynq.Client {
	return asynq.NewClient(asynq.RedisClientOpt{Addr: c.Redis.Host, Password: c.Redis.Pass})
}
