package logic

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"liujun/Time_go-zero_csdn/common/utils"
	"liujun/Time_go-zero_csdn/csdn/channel/cmd/rpc/internal/svc"
	"liujun/Time_go-zero_csdn/csdn/channel/cmd/rpc/types/channel"
	"liujun/Time_go-zero_csdn/csdn/channel/model"
)

type ArticleReadLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewArticleReadLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ArticleReadLogic {
	return &ArticleReadLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ArticleReadLogic) ArticleRead(in *channel.ArticleReadRequest) (*channel.ArticleReadResponse, error) {
	// todo: add your logic here and delete this line
	res, err := l.svcCtx.ArticleReadModel.FindOneByUserIdArticleId(l.ctx, in.UserId, in.ArticleId)
	if err == nil {
		return &channel.ArticleReadResponse{
			Aid: res.ArticleId,
		}, nil
	}
	news_read := model.NewsRead{
		ReadId:    utils.UUID(),
		UserId:    in.UserId,
		ArticleId: in.ArticleId,
	}
	_, err = l.svcCtx.ArticleReadModel.Insert(l.ctx, &news_read)
	if err != nil {
		return nil, err
	}
	return &channel.ArticleReadResponse{
		Aid: in.ArticleId,
	}, nil
}
