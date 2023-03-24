package svc

import (
	"fmt"
	"liujun/Time_go-zero_csdn/csdn/channel/cmd/rpc/internal/config"
	"liujun/Time_go-zero_csdn/csdn/channel/model"

	"github.com/olivere/elastic/v7"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config                 config.Config
	RedisClient            *redis.Redis
	EsClient               *elastic.Client
	ArticleModel           model.NewsArticleBasicModel
	ChannelModel           model.NewsChannelModel
	UserChannelModel       model.NewsUserChannelModel
	ArticleCollectionModel model.NewsCollectionModel
	ArticleStaticModel     model.NewsArticleStatisticModel
	ArticleReadModel       model.NewsReadModel
	ArticleAttitudeModel   model.NewsAttitudeModel
	ArticleCommentModel    model.NewsCommentModel
	UserArticleSearchModel model.UserSearchModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	sqlConn := sqlx.NewMysql(c.DB.DataSource)
	EsClient, err := elastic.NewClient(elastic.SetURL(c.Es.Host))
	if err != nil {
		fmt.Println("es连接错误", err)
	}
	return &ServiceContext{
		Config: c,
		RedisClient: redis.New(c.Redis.Host, func(r *redis.Redis) {
			r.Type = c.Redis.Type
			r.Pass = c.Redis.Pass
		}),
		EsClient:               EsClient,
		ArticleModel:           model.NewNewsArticleBasicModel(sqlConn, c.Cache),
		ChannelModel:           model.NewNewsChannelModel(sqlConn, c.Cache),
		UserChannelModel:       model.NewNewsUserChannelModel(sqlConn, c.Cache),
		ArticleCollectionModel: model.NewNewsCollectionModel(sqlConn, c.Cache),
		ArticleStaticModel:     model.NewNewsArticleStatisticModel(sqlConn, c.Cache),
		ArticleReadModel:       model.NewNewsReadModel(sqlConn, c.Cache),
		ArticleAttitudeModel:   model.NewNewsAttitudeModel(sqlConn, c.Cache),
		ArticleCommentModel:    model.NewNewsCommentModel(sqlConn, c.Cache),
		UserArticleSearchModel: model.NewUserSearchModel(sqlConn, c.Cache),
	}
}
