package logic

import (
	"context"

	"liujun/Time_go-zero_csdn/csdn/channel/cmd/rpc/internal/svc"
	"liujun/Time_go-zero_csdn/csdn/channel/cmd/rpc/types/channel"

	"github.com/zeromicro/go-zero/core/logx"
)

type ArticleChannelLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewArticleChannelLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ArticleChannelLogic {
	return &ArticleChannelLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ArticleChannelLogic) ArticleChannel(in *channel.ArticleChannelRequest) (*channel.ArticleChannelResponse, error) {
	// todo: add your logic here and delete this line
	builder := l.svcCtx.ArticleModel.RowBuilder()
	articles, err := l.svcCtx.ArticleModel.FindAllArticle(l.ctx, builder, in.ChannelId, in.Page, in.PageNum)
	if err != nil {
		return nil, err
	}
	datas := []*channel.ArticleList{}
	for _, article := range articles {
		datas = append(datas, &channel.ArticleList{
			ArtId:         article.ArtId,
			UserId:        article.UserId,
			ChannelId:     article.ChannelId,
			Title:         article.Title,
			CreateTime:    article.CreateTime,
			AllowComment:  article.AllowComment,
			Content:       article.Content,
			UserName:      article.UserName,
			HeadPhoto:     article.HeadPhoto,
			Career:        article.Career,
			ReadNum:       0,
			CodeYear:      article.CodeYear,
			CommentNum:    0,
			LikeNum:       0,
			CollectionNum: 0,
		})
	}
	return &channel.ArticleChannelResponse{Articles: datas}, nil
}
