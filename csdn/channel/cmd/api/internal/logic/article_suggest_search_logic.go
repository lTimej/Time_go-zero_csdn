package logic

import (
	"context"
	"fmt"
	"liujun/Time_go-zero_csdn/csdn/channel/cmd/rpc/channelclient"

	"liujun/Time_go-zero_csdn/csdn/channel/cmd/api/internal/svc"
	"liujun/Time_go-zero_csdn/csdn/channel/cmd/api/internal/types"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type ArticleSuggestSearchLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewArticleSuggestSearchLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ArticleSuggestSearchLogic {
	return &ArticleSuggestSearchLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ArticleSuggestSearchLogic) ArticleSuggestSearch(req *types.ArticleSuggestSearchRequest) (resp *types.ArticleSuggestSearchResponse, err error) {
	// todo: add your logic here and delete this line
	searchs, err := l.svcCtx.ChannelRpc.ArticleSuggestSearch(l.ctx, &channelclient.ArticleSuggestSearchRequest{Keyword: req.Keyword})
	if err != nil {
		fmt.Println(err, "))))))))))))))")
		return nil, err
	}
	resp = new(types.ArticleSuggestSearchResponse)
	_ = copier.Copy(&resp, searchs)
	return
}
