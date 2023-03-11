package logic

import (
	"context"

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
	
	return &channel.ArticlestatusResponse{}, nil
}
