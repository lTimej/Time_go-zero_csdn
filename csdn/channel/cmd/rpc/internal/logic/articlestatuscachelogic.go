package logic

import (
	"context"
	"fmt"
	"liujun/Time_go-zero_csdn/common/utils"
	"liujun/Time_go-zero_csdn/csdn/channel/model"

	"liujun/Time_go-zero_csdn/common/globalkey"
	"liujun/Time_go-zero_csdn/csdn/channel/cmd/rpc/internal/svc"
	"liujun/Time_go-zero_csdn/csdn/channel/cmd/rpc/types/channel"

	"github.com/zeromicro/go-zero/core/logx"
)

type ArticleStatusCacheLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewArticleStatusCacheLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ArticleStatusCacheLogic {
	return &ArticleStatusCacheLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ArticleStatusCacheLogic) ArticleStatusCache(in *channel.ArticleStatusCacheRequest) (*channel.ArticleStatusCacheResponse, error) {
	// todo: add your logic here and delete this line
	article_ids_key := globalkey.ArticleIds
	article_status_key := globalkey.ArticleStatus
	aids, err := l.svcCtx.RedisClient.Smembers(article_ids_key)
	if err != nil {
		fmt.Println("获取文章状态失败", err)
		return nil, err
	}
	for _, aid := range aids {
		status_key := fmt.Sprintf(article_status_key, aid)
		like_num, _ := l.svcCtx.RedisClient.Hget(status_key, globalkey.ArticleLikeNum)
		read_num, _ := l.svcCtx.RedisClient.Hget(status_key, globalkey.ArticleReadNum)
		collection_num, _ := l.svcCtx.RedisClient.Hget(status_key, globalkey.ArticleCollectionNum)
		nas := model.NewsArticleStatistic{
			ArticleId:        utils.StringToInt64(aid),
			LikeCount:        utils.StringToInt64(like_num),
			ReadCount:        utils.StringToInt64(read_num),
			CollectCount:     utils.StringToInt64(collection_num),
			DislikeCount:     0,
			RepostCount:      0,
			FansCommentCount: 0,
		}
		err := l.svcCtx.ArticleStaticModel.UpdateCache(l.ctx, &nas)
		if err != nil {
			fmt.Println("文章状态缓存失败", err)
			return nil, err
		}
		// key := fmt.Sprintf("cache:newsArticleStatistic:articleId:%v", aid)
		// data, _ := json.Marshal(nas)
		// err = l.svcCtx.RedisClient.Set(key, string(data))
		// if err != nil {
		// 	fmt.Println("单个文章缓存状态失败")
		// }
	}
	return &channel.ArticleStatusCacheResponse{}, nil
}
