package logic

import (
	"context"
	"fmt"
	"liujun/Time_go-zero_csdn/common/globalkey"
	"liujun/Time_go-zero_csdn/csdn/channel/cmd/rpc/internal/svc"
	"liujun/Time_go-zero_csdn/csdn/channel/cmd/rpc/types/channel"

	"github.com/zeromicro/go-zero/core/logx"
)

type ArticleUserSearchHistoryLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewArticleUserSearchHistoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ArticleUserSearchHistoryLogic {
	return &ArticleUserSearchHistoryLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ArticleUserSearchHistoryLogic) ArticleUserSearchHistory(in *channel.ArticleUserSearchHistoryRequest) (*channel.ArticleUserSearchHistoryResponse, error) {
	// todo: add your logic here and delete this line
	key := fmt.Sprintf(globalkey.UserArticleSearch, in.UserId)
	resp := new(channel.ArticleUserSearchHistoryResponse)
	n, _ := l.svcCtx.RedisClient.Zcard(key)
	if n > 5 {
		resp.Keywords, _ = l.svcCtx.RedisClient.Zrevrange(key, 0, 4)
	} else {
		resp.Keywords, _ = l.svcCtx.RedisClient.Zrevrange(key, 0, int64(n))
	}

	return resp, nil
}
