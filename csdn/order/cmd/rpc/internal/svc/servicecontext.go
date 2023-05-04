package svc

import (
	"liujun/Time_go-zero_csdn/csdn/order/cmd/rpc/internal/config"
	"liujun/Time_go-zero_csdn/csdn/order/model"

	"github.com/hibiken/asynq"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config         config.Config
	RedisClient    *redis.Redis
	OrderModel     model.OrderModel
	OrderUserModel model.UserOrderModel
	AsynqClient    *asynq.Client
}

func NewServiceContext(c config.Config) *ServiceContext {
	sqlConn := sqlx.NewMysql(c.DB.DataSource)
	return &ServiceContext{
		Config:      c,
		AsynqClient: newAsynqClient(c),
		RedisClient: redis.New(c.Redis.Host, func(r *redis.Redis) {
			r.Type = c.Redis.Type
			r.Pass = c.Redis.Pass
		}),
		OrderModel:     model.NewOrderModel(sqlConn, c.Cache),
		OrderUserModel: model.NewUserOrderModel(sqlConn, c.Cache),
	}
}
