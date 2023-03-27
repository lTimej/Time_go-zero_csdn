package logic

import (
	"context"
	"liujun/Time_go-zero_csdn/common/ctxdata"
	"liujun/Time_go-zero_csdn/csdn/channel/cmd/rpc/channelclient"

	"github.com/jinzhu/copier"

	"liujun/Time_go-zero_csdn/csdn/channel/cmd/api/internal/svc"
	"liujun/Time_go-zero_csdn/csdn/channel/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ArticleUserSearchLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewArticleUserSearchLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ArticleUserSearchLogic {
	return &ArticleUserSearchLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ArticleUserSearchLogic) ArticleUserSearch(req *types.ArticleUserSearchRequest) (resp *types.ArticleUserSearchResponse, err error) {
	// todo: add your logic here and delete this line
	user_id := ctxdata.GetUidFromCtx(l.ctx)
	articles, err := l.svcCtx.ChannelRpc.ArticleUserSearch(l.ctx, &channelclient.ArticleUserSearchRequest{Keyword: req.Keyword, Page: req.Page, PageNum: req.PageNum, UserId: user_id})
	if err != nil {
		return nil, err
	}
	resp = new(types.ArticleUserSearchResponse)
	copier.Copy(resp, articles)
	return
}