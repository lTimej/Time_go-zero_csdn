package logic

import (
	"context"

	"liujun/Time_go-zero_csdn/csdn/channel/cmd/rpc/internal/svc"
	"liujun/Time_go-zero_csdn/csdn/channel/cmd/rpc/types/channel"

	"github.com/zeromicro/go-zero/core/logx"
)

type ArticleToDisLikeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewArticleToDisLikeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ArticleToDisLikeLogic {
	return &ArticleToDisLikeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ArticleToDisLikeLogic) ArticleToDisLike(in *channel.ArticleToDisLikeRequest) (*channel.ArticleToDisLikeResponse, error) {
	// todo: add your logic here and delete this line

	return &channel.ArticleToDisLikeResponse{}, nil
}
