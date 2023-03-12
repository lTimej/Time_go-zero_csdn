package logic

import (
	"context"

	"liujun/Time_go-zero_csdn/csdn/channel/cmd/rpc/internal/svc"
	"liujun/Time_go-zero_csdn/csdn/channel/cmd/rpc/types/channel"

	"github.com/zeromicro/go-zero/core/logx"
)

type ArticleLikeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewArticleLikeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ArticleLikeLogic {
	return &ArticleLikeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ArticleLikeLogic) ArticleLike(in *channel.ArticleLikeRequest) (*channel.ArticleLikeResponse, error) {
	// todo: add your logic here and delete this line

	return &channel.ArticleLikeResponse{}, nil
}
