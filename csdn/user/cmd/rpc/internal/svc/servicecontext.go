package svc

import (
	"liujun/Time_go-zero_csdn/csdn/user/cmd/rpc/internal/config"
	"liujun/Time_go-zero_csdn/csdn/user/model"

	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config            config.Config
	RedisClient       *redis.Redis
	UserModel         model.UserBasicModel
	UserRelationModel model.UserRelationModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	sqlConn := sqlx.NewMysql(c.DB.DataSource)
	return &ServiceContext{
		Config: c,
		RedisClient: redis.New(c.Redis.Host, func(r *redis.Redis) {
			r.Type = c.Redis.Type
			r.Pass = c.Redis.Pass
		}),
		UserModel:         model.NewUserBasicModel(sqlConn, c.Cache),
		UserRelationModel: model.NewUserRelationModel(sqlConn, c.Cache),
	}
}
