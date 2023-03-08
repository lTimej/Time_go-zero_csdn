package logic

import (
	"context"
	"fmt"
	"liujun/Time_go-zero_csdn/csdn/channel/cmd/rpc/channelclient"

	"liujun/Time_go-zero_csdn/csdn/channel/cmd/api/internal/svc"
	"liujun/Time_go-zero_csdn/csdn/channel/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AllArticleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAllArticleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AllArticleLogic {
	return &AllArticleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AllArticleLogic) AllArticle(req *types.AllArticleRequest, channel_id int64, page, page_num int32) (resp *types.AllArticleResponse, err error) {
	// todo: add your logic here and delete this line
	articles, err := l.svcCtx.ChannelRpc.ArticleChannel(l.ctx, &channelclient.ArticleChannelRequest{
		ChannelId: channel_id,
		Page:      page,
		PageNum:   page_num,
	})
	if err != nil {
		fmt.Println(err, "==============")
		return nil, err
	}
	data := []types.ArticleList{}
	for _, a := range articles.Articles {
		d := types.ArticleList{
			Title:         a.Title,
			UserId:        a.UserId,
			CreateTime:    a.CreateTime,
			ArtId:         a.ArtId,
			ChannelId:     a.ChannelId,
			Content:       a.Content,
			AllowComment:  a.AllowComment,
			UserName:      a.UserName,
			HeadPhoto:     a.HeadPhoto,
			Career:        a.Career,
			CodeYear:      a.CodeYear,
			ReadNum:       a.ReadNum,
			CommentNum:    a.CommentNum,
			LikeNum:       a.LikeNum,
			CollectionNum: a.CollectionNum,
		}
		data = append(data, d)
	}
	return &types.AllArticleResponse{
		Articles: data,
	}, nil
	return
}
