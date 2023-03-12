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
	builder := l.svcCtx.ArticleAttitudeModel.AllArticleAttitudeBuilder().Where("article_id = ?", in.ArticleId)
	nas, err := l.svcCtx.ArticleAttitudeModel.FindAllByArticleId(l.ctx, builder)
	if err != nil {
		return nil, err
	}
	data := []*channel.ArticleLikeResponse_UserInfo{}
	for _, na := range nas {
		data = append(data, &channel.ArticleLikeResponse_UserInfo{
			HeadPhoto: na.HeadPhoto,
			Aid:       na.Aid,
		})
	}
	return &channel.ArticleLikeResponse{
		UsersInfo: data,
	}, nil
}
