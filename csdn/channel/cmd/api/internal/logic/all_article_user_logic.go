package logic

import (
	"context"
	"fmt"
	"liujun/Time_go-zero_csdn/common/ctxdata"
	"liujun/Time_go-zero_csdn/csdn/channel/cmd/rpc/channelclient"

	"liujun/Time_go-zero_csdn/csdn/channel/cmd/api/internal/svc"
	"liujun/Time_go-zero_csdn/csdn/channel/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AllArticleUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAllArticleUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AllArticleUserLogic {
	return &AllArticleUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AllArticleUserLogic) AllArticleUser(req *types.AllArticleUserRequest) (resp *types.AllArticleUserResponse, err error) {
	// todo: add your logic here and delete this line
	user_id := ctxdata.GetUidFromCtx(l.ctx)
	articles, err := l.svcCtx.ChannelRpc.ArticleUserList(l.ctx, &channelclient.ArticleUserRequest{UserId: user_id, Page: req.Page, PageNum: req.PageNum})
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
	return &types.AllArticleUserResponse{
		Articles: data,
	}, nil
}
