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

type ArticleUserSearchHistoryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewArticleUserSearchHistoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ArticleUserSearchHistoryLogic {
	return &ArticleUserSearchHistoryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ArticleUserSearchHistoryLogic) ArticleUserSearchHistory(req *types.ArticleUserSearchHistoryRequest) (resp *types.ArticleUserSearchHistoryResponse, err error) {
	// todo: add your logic here and delete this line
	user_id := ctxdata.GetUidFromCtx(l.ctx)
	keywords, err := l.svcCtx.ChannelRpc.ArticleUserSearchHistory(l.ctx, &channelclient.ArticleUserSearchHistoryRequest{UserId: user_id})
	if err != nil {
		return nil, err
	}
	resp = new(types.ArticleUserSearchHistoryResponse)
	_ = copier.Copy(resp, keywords)
	return
}
