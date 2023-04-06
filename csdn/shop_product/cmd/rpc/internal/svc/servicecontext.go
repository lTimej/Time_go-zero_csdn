package svc

import (
	"liujun/Time_go-zero_csdn/csdn/shop_product/cmd/rpc/internal/config"
	"liujun/Time_go-zero_csdn/csdn/shop_product/model"

	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config                       config.Config
	RedisClient                  *redis.Redis
	ProductCategoryModel         model.TbGoodsCategoryModel
	ProductSpuModel              model.TbSpuModel
	ProductSkuModel              model.TbSkuModel
	ProductSpuSpecificationModel model.TbSpuSpecificationModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	sqlConn := sqlx.NewMysql(c.DB.DataSource)
	return &ServiceContext{
		Config: c,
		RedisClient: redis.New(c.Redis.Host, func(r *redis.Redis) {
			r.Type = c.Redis.Type
			r.Pass = c.Redis.Pass
		}),
		ProductCategoryModel:         model.NewTbGoodsCategoryModel(sqlConn, c.Cache),
		ProductSpuModel:              model.NewTbSpuModel(sqlConn, c.Cache),
		ProductSkuModel:              model.NewTbSkuModel(sqlConn, c.Cache),
		ProductSpuSpecificationModel: model.NewTbSpuSpecificationModel(sqlConn, c.Cache),
	}
}
