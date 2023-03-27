package logic

import (
	"context"
	"liujun/Time_go-zero_csdn/common/ctxdata"
	"liujun/Time_go-zero_csdn/csdn/channel/cmd/rpc/channelclient"

	"liujun/Time_go-zero_csdn/csdn/channel/cmd/api/internal/svc"
	"liujun/Time_go-zero_csdn/csdn/channel/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteArticleUserSearchHistoryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteArticleUserSearchHistoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteArticleUserSearchHistoryLogic {
	return &DeleteArticleUserSearchHistoryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteArticleUserSearchHistoryLogic) DeleteArticleUserSearchHistory(req *types.DeleteArticleUserSearchHistoryRequest) (resp *types.DeleteArticleUserSearchHistoryResponse, err error) {
	// todo: add your logic here and delete this line
	user_id := ctxdata.GetUidFromCtx(l.ctx)
	_, err = l.svcCtx.ChannelRpc.DeleteArticleUserSearchHistory(l.ctx, &channelclient.DeleteArticleUserSearchHistoryRequest{UserId: user_id})
	if err != nil {
		return nil, err
	}

	return &types.DeleteArticleUserSearchHistoryResponse{Message: "删除成功"}, nil
}
