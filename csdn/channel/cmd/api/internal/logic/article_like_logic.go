package logic

import (
	"context"
	"fmt"
	"liujun/Time_go-zero_csdn/csdn/channel/cmd/rpc/channelclient"

	"liujun/Time_go-zero_csdn/csdn/channel/cmd/api/internal/svc"
	"liujun/Time_go-zero_csdn/csdn/channel/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ArticleLikeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewArticleLikeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ArticleLikeLogic {
	return &ArticleLikeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ArticleLikeLogic) ArticleLike(req *types.ArticleLikeRequest) (resp *types.ArticleLikeResponse, err error) {
	// todo: add your logic here and delete this line
	als, err := l.svcCtx.ChannelRpc.ArticleLike(l.ctx, &channelclient.ArticleLikeRequest{ArticleId: req.ArticleId})
	if err != nil {
		fmt.Println(err, "&&&&&&&****************")
		return nil, err
	}
	data := []types.ArticleLikeList{}
	for _, al := range als.UsersInfo {
		data = append(data, types.ArticleLikeList{Aid: al.Aid, HeadPhoto: al.HeadPhoto})
	}
	return &types.ArticleLikeResponse{
		UsersInfo: data,
	}, nil
}
