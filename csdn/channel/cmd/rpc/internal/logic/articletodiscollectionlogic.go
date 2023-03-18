package logic

import (
	"context"
	"liujun/Time_go-zero_csdn/csdn/channel/model"

	"liujun/Time_go-zero_csdn/csdn/channel/cmd/rpc/internal/svc"
	"liujun/Time_go-zero_csdn/csdn/channel/cmd/rpc/types/channel"

	"github.com/zeromicro/go-zero/core/logx"
)

type ArticleToDisCollectionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewArticleToDisCollectionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ArticleToDisCollectionLogic {
	return &ArticleToDisCollectionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ArticleToDisCollectionLogic) ArticleToDisCollection(in *channel.ArticleToDisCollectionRequest) (*channel.ArticleToDisCollectionResponse, error) {
	// todo: add your logic here and delete this line
	article_collection := model.NewsCollection{
		ArticleId: in.Aid,
		UserId:    in.UserId,
		IsDeleted: 1,
	}
	err := l.svcCtx.ArticleCollectionModel.Update(l.ctx, &article_collection)
	if err != nil {
		return nil, err
	}
	return &channel.ArticleToDisCollectionResponse{}, nil
}
