package logic

import (
	"context"
	"fmt"
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
	ok, err := l.svcCtx.RedisClient.Exists(newsAttitudeUserIdArticleIdKey)
	if err != nil {
		return nil, err
	}
	fmt.Println(ok, "===&&&&&&&===", newsAttitudeUserIdArticleIdKey)
	article_static, err := l.svcCtx.ArticleStaticModel.FindOne(l.ctx, in.ArticleId)
	if err != nil {
		return nil, err
	}
	return &channel.ArticlestatusResponse{
		CollectionNum: article_static.CollectCount,
		LikeNum:       article_static.LikeCount,
		ReadNum:       article_static.ReadCount,
		Aid:           in.ArticleId,
		Isfocus:       false,
		Iscollection:  false,
		Islike:        ok,
	}, nil
}
