package logic

import (
	"context"
	"liujun/Time_go-zero_csdn/csdn/channel/cmd/rpc/internal/svc"
	"liujun/Time_go-zero_csdn/csdn/channel/cmd/rpc/types/channel"
	"liujun/Time_go-zero_csdn/csdn/channel/model"

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
	na, err := l.svcCtx.ArticleAttitudeModel.FindOneByUserIdArticleId(l.ctx, in.UserId, in.ArticleId)
	if err != nil {
		return nil, err
	}
	news_attitude := model.NewsAttitude{
		UserId:    in.UserId,
		ArticleId: in.ArticleId,
		Attitude:  0,
	}
	if na != nil {
		news_attitude.AttitudeId = na.AttitudeId
		l.svcCtx.ArticleAttitudeModel.Update(l.ctx, &news_attitude)
	}
	return &channel.ArticleToDisLikeResponse{
		Aid: in.ArticleId,
	}, nil
}
