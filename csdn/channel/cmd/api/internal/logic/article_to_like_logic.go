package logic

import (
	"context"
	"liujun/Time_go-zero_csdn/common/ctxdata"
	"liujun/Time_go-zero_csdn/csdn/channel/cmd/rpc/channelclient"
	"strconv"

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
	uid := strconv.FormatInt(ctxdata.GetUidFromCtx(l.ctx), 10)
	aid, _ := strconv.ParseInt(req.ArticleId, 10, 64)
	res, err := l.svcCtx.ChannelRpc.ArticleToLike(l.ctx, &channelclient.ArticleToLikeRequest{UserId: uid, ArticleId: aid})
	if err != nil {
		return nil, err
	}
	return &types.ArticleToLikeResponse{ArticleId: res.Aid}, nil
}
