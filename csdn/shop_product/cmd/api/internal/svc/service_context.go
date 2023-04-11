package svc

import (
	"liujun/Time_go-zero_csdn/csdn/shop_product/cmd/api/internal/config"
	"liujun/Time_go-zero_csdn/csdn/shop_product/cmd/api/internal/middleware"
	"liujun/Time_go-zero_csdn/csdn/shop_product/cmd/rpc/productclient"
	"liujun/Time_go-zero_csdn/csdn/shop_product/model"

	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config                config.Config
	ProductRpc            productclient.Product
	SetUidToCtxMiddleware rest.Middleware
	RedisClient           *redis.Redis
	ProductSkuModel       model.TbSkuModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	sqlConn := sqlx.NewMysql(c.DB.DataSource)
	return &ServiceContext{
		Config:                c,
		SetUidToCtxMiddleware: middleware.NewSetUidToCtxMiddleware(c).Handle,
		ProductRpc:            productclient.NewProduct(zrpc.MustNewClient(c.ProductRpc)),
		ProductSkuModel:       model.NewTbSkuModel(sqlConn, c.Cache),
		RedisClient: redis.New(c.Redis.Host, func(r *redis.Redis) {
			r.Type = c.Redis.Type
			r.Pass = c.Redis.Pass
		}),
	}
}
