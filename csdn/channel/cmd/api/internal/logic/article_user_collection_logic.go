package logic

import (
	"context"

	"liujun/Time_go-zero_csdn/csdn/channel/cmd/api/internal/svc"
	"liujun/Time_go-zero_csdn/csdn/channel/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ArticleUserCollectionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewArticleUserCollectionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ArticleUserCollectionLogic {
	return &ArticleUserCollectionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ArticleUserCollectionLogic) ArticleUserCollection(req *types.ArticleUserCollectionRequest) (resp *types.ArticleUserCollectionResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
