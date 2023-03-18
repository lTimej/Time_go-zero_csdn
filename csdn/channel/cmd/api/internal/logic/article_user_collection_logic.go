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
	user_id := ctxdata.GetUidFromCtx(l.ctx)
	articles, err := l.svcCtx.ChannelRpc.ArticleUserCollection(l.ctx, &channelclient.ArticleUserCollectionRequest{UserId: user_id, PageNum: req.PageNum, Page: req.Page})
	if err != nil {
		fmt.Println(err, "=======错误=======")
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
	return &types.ArticleUserCollectionResponse{
		Collections: data,
		PageNum:     req.PageNum,
		Page:        req.Page,
		TotalNum:    int32(len(data)),
	}, nil
}
