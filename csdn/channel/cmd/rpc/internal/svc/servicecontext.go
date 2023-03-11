package svc

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"liujun/Time_go-zero_csdn/csdn/channel/cmd/rpc/internal/config"
	"liujun/Time_go-zero_csdn/csdn/channel/model"
)

type ServiceContext struct {
	Config                 config.Config
	RedisClient            *redis.Redis
	ArticleModel           model.NewsArticleBasicModel
	ChannelModel           model.NewsChannelModel
	UserChannelModel       model.NewsUserChannelModel
	ArticleCollectionModel model.NewsCollectionModel
	ArticleStaticModel     model.NewsArticleStatisticModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	sqlConn := sqlx.NewMysql(c.DB.DataSource)
	return &ServiceContext{
		Config: c,
		RedisClient: redis.New(c.Redis.Host, func(r *redis.Redis) {
			r.Type = c.Redis.Type
			r.Pass = c.Redis.Pass
		}),
		ArticleModel:           model.NewNewsArticleBasicModel(sqlConn, c.Cache),
		ChannelModel:           model.NewNewsChannelModel(sqlConn, c.Cache),
		UserChannelModel:       model.NewNewsUserChannelModel(sqlConn, c.Cache),
		ArticleCollectionModel: model.NewNewsCollectionModel(sqlConn, c.Cache),
		ArticleStaticModel:     model.NewNewsArticleStatisticModel(sqlConn, c.Cache),
	}
}
