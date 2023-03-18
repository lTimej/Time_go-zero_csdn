package logic

import (
	"context"
	"liujun/Time_go-zero_csdn/csdn/channel/model"

	"liujun/Time_go-zero_csdn/csdn/channel/cmd/rpc/internal/svc"
	"liujun/Time_go-zero_csdn/csdn/channel/cmd/rpc/types/channel"

	"github.com/zeromicro/go-zero/core/logx"
)

type ArticleToCollectionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewArticleToCollectionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ArticleToCollectionLogic {
	return &ArticleToCollectionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ArticleToCollectionLogic) ArticleToCollection(in *channel.ArticleToCollectionRequest) (*channel.ArticleToCollectionResponse, error) {
	// todo: add your logic here and delete this line
	article_collection := model.NewsCollection{
		ArticleId: in.Aid,
		UserId:    in.UserId,
		IsDeleted: 0,
	}
	res, err := l.svcCtx.ArticleCollectionModel.FindOneByUserIdArticleId(l.ctx, in.UserId, in.Aid)
	if res == nil && err == nil {
		_, err := l.svcCtx.ArticleCollectionModel.Insert(l.ctx, &article_collection)
		if err != nil {
			return nil, err
		}
	} else {
		res.IsDeleted = 0
		err := l.svcCtx.ArticleCollectionModel.Update(l.ctx, res)
		if err != nil {
			return nil, err
		}
	}
	return &channel.ArticleToCollectionResponse{Aid: in.Aid}, nil
}
