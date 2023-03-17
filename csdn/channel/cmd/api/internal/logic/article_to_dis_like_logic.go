package logic

import (
	"context"
	"fmt"
	"liujun/Time_go-zero_csdn/common/ctxdata"
	"liujun/Time_go-zero_csdn/csdn/channel/cmd/rpc/channelclient"

	"liujun/Time_go-zero_csdn/csdn/channel/cmd/api/internal/svc"
	"liujun/Time_go-zero_csdn/csdn/channel/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ArticleToDisLikeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewArticleToDisLikeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ArticleToDisLikeLogic {
	return &ArticleToDisLikeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ArticleToDisLikeLogic) ArticleToDisLike(req *types.ArticleToDisLikeRequest) (resp *types.ArticleToDisLikeResponse, err error) {
	// todo: add your logic here and delete this line

	uid := ctxdata.GetUidFromCtx(l.ctx)
	fmt.Println(uid, "!!!!!!!!!!!!!!!!!!!!uid!!!!!!!!!!!!!!")
	aid := req.ArticleId
	_, err = l.svcCtx.ChannelRpc.ArticleToDisLike(l.ctx, &channelclient.ArticleToDisLikeRequest{UserId: uid, ArticleId: aid})
	if err != nil {
		return nil, err
	}
	return &types.ArticleToDisLikeResponse{Message: "success"}, nil
}
