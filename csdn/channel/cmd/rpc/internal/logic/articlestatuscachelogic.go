package logic

import (
	"context"
	"fmt"
	"liujun/Time_go-zero_csdn/common/utils"
	"liujun/Time_go-zero_csdn/csdn/channel/model"
	"strconv"
	"strings"

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
	status_key := globalkey.ArticleStatus
	status, err := l.svcCtx.RedisClient.Hgetall(status_key)
	if err != nil {
		fmt.Println("获取文章状态失败", err)
		return nil, err
	}
	// fmt.Println("文章状态:", status)
	ids := make(map[int64]bool)
	for key, _ := range status {
		sli := strings.Split(key, ":")
		aid, _ := strconv.ParseInt(sli[2], 10, 64)
		ids[aid] = true
	}
	for aid, _ := range ids {
		like_num, _ := l.svcCtx.RedisClient.Hget(status_key, fmt.Sprintf(globalkey.ArticleLikeNum, aid))
		read_num, _ := l.svcCtx.RedisClient.Hget(status_key, fmt.Sprintf(globalkey.ArticleReadNum, aid))
		collection_num, _ := l.svcCtx.RedisClient.Hget(status_key, fmt.Sprintf(globalkey.ArticleCollectionNum, aid))
		nas := model.NewsArticleStatistic{
			ArticleId:        aid,
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
	}
	return &channel.ArticleStatusCacheResponse{}, nil
}
