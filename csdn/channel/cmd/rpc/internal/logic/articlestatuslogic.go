package logic

import (
	"context"
	"fmt"
	"liujun/Time_go-zero_csdn/common/globalkey"
	"liujun/Time_go-zero_csdn/common/utils"
	"liujun/Time_go-zero_csdn/csdn/channel/cmd/rpc/internal/svc"
	"liujun/Time_go-zero_csdn/csdn/channel/cmd/rpc/types/channel"

	"github.com/zeromicro/go-zero/core/logx"
)

type ArticleStatusLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewArticleStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ArticleStatusLogic {
	return &ArticleStatusLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ArticleStatusLogic) ArticleStatus(in *channel.ArticlestatusRequest) (*channel.ArticlestatusResponse, error) {
	// todo: add your logic here and delete this line
	newsAttitudeUserIdArticleIdKey := fmt.Sprintf("%s%v:%v", "cache:newsAttitude:userId:articleId:", in.UserId, in.ArticleId)
	is_like, err := l.svcCtx.RedisClient.Exists(newsAttitudeUserIdArticleIdKey)
	if err != nil {
		return nil, err
	}
	newsArticleStatisticArticleIdKey := fmt.Sprintf(globalkey.ArticleStatus, utils.Int64ToString(in.ArticleId))
	ok, err := l.svcCtx.RedisClient.Exists(newsArticleStatisticArticleIdKey)
	if err != nil {
		return nil, err
	}
	if !ok {
		build := l.svcCtx.ArticleStaticModel.RowBuilder().Where("article_id = ?", in.ArticleId)
		article_static, err := l.svcCtx.ArticleStaticModel.FindOneByArticle(l.ctx, build)
		if err != nil {
			return nil, err
		}
		l.svcCtx.RedisClient.Hset(newsArticleStatisticArticleIdKey, globalkey.ArticleLikeNum, utils.Int64ToString(article_static.LikeCount))
		l.svcCtx.RedisClient.Hset(newsArticleStatisticArticleIdKey, globalkey.ArticleReadNum, utils.Int64ToString(article_static.ReadCount))
		l.svcCtx.RedisClient.Hset(newsArticleStatisticArticleIdKey, globalkey.ArticleCollectionNum, utils.Int64ToString(article_static.CollectCount))
	}
	article_static, err := l.svcCtx.RedisClient.Hgetall(newsArticleStatisticArticleIdKey)
	if err != nil {
		return nil, err
	}
	return &channel.ArticlestatusResponse{
		CollectionNum: utils.StringToInt64(article_static[globalkey.ArticleCollectionNum]),
		LikeNum:       utils.StringToInt64(article_static[globalkey.ArticleLikeNum]),
		ReadNum:       utils.StringToInt64(article_static[globalkey.ArticleReadNum]),
		Aid:           in.ArticleId,
		Isfocus:       false,
		Iscollection:  false,
		Islike:        is_like,
	}, nil
}
