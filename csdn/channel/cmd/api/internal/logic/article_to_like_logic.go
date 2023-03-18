package logic

import (
	"context"
	"fmt"
	"liujun/Time_go-zero_csdn/common/ctxdata"
	"liujun/Time_go-zero_csdn/common/globalkey"
	"liujun/Time_go-zero_csdn/common/utils"
	"liujun/Time_go-zero_csdn/csdn/channel/cmd/rpc/channelclient"

	"liujun/Time_go-zero_csdn/csdn/channel/cmd/api/internal/svc"
	"liujun/Time_go-zero_csdn/csdn/channel/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ArticleToLikeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewArticleToLikeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ArticleToLikeLogic {
	return &ArticleToLikeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ArticleToLikeLogic) ArticleToLike(req *types.ArticleToLikeRequest) (resp *types.ArticleToLikeResponse, err error) {
	// todo: add your logic here and delete this line
	uid := ctxdata.GetUidFromCtx(l.ctx)
	aid := req.ArticleId
	res, err := l.svcCtx.ChannelRpc.ArticleToLike(l.ctx, &channelclient.ArticleToLikeRequest{UserId: uid, ArticleId: aid})
	if err != nil {
		return nil, err
	}
	key := fmt.Sprintf(globalkey.ArticleStatus, utils.Int64ToString(req.ArticleId))
	field := globalkey.ArticleLikeNum
	fmt.Println(key, ":", field)
	l.svcCtx.RedisClient.Hincrby(key, field, 1)
	return &types.ArticleToLikeResponse{ArticleId: res.Aid}, nil
}
