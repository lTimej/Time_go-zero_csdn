package logic

import (
	"context"
	"fmt"

	"liujun/Time_go-zero_csdn/common/globalkey"
	"liujun/Time_go-zero_csdn/csdn/channel/cmd/rpc/internal/svc"
	"liujun/Time_go-zero_csdn/csdn/channel/cmd/rpc/types/channel"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteArticleUserSearchHistoryLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteArticleUserSearchHistoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteArticleUserSearchHistoryLogic {
	return &DeleteArticleUserSearchHistoryLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteArticleUserSearchHistoryLogic) DeleteArticleUserSearchHistory(in *channel.DeleteArticleUserSearchHistoryRequest) (*channel.DeleteArticleUserSearchHistoryResponse, error) {
	// todo: add your logic here and delete this line
	key := fmt.Sprintf(globalkey.UserArticleSearch, in.UserId)
	_, err := l.svcCtx.RedisClient.Del(key)
	if err != nil {
		return nil, err
	}
	return &channel.DeleteArticleUserSearchHistoryResponse{Message: "删除成功"}, nil
}
