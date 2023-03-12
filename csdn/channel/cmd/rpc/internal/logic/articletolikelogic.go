package logic

import (
	"context"
	"liujun/Time_go-zero_csdn/csdn/channel/model"

	"liujun/Time_go-zero_csdn/csdn/channel/cmd/rpc/internal/svc"
	"liujun/Time_go-zero_csdn/csdn/channel/cmd/rpc/types/channel"

	"github.com/zeromicro/go-zero/core/logx"
)

type ArticleToLikeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewArticleToLikeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ArticleToLikeLogic {
	return &ArticleToLikeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ArticleToLikeLogic) ArticleToLike(in *channel.ArticleToLikeRequest) (*channel.ArticleToLikeResponse, error) {
	// todo: add your logic here and delete this line
	na, err := l.svcCtx.ArticleAttitudeModel.FindOneByUserIdArticleId(l.ctx, in.UserId, in.ArticleId)
	if err != nil {
		return nil, err
	}
	news_attitude := model.NewsAttitude{
		UserId:    in.UserId,
		ArticleId: in.ArticleId,
		Attitude:  1,
	}
	if na == nil && err == nil {
		l.svcCtx.ArticleAttitudeModel.Insert(l.ctx, &news_attitude)
	}
	return &channel.ArticleToLikeResponse{
		Aid: in.ArticleId,
	}, nil
}
