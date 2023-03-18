package logic

import (
	"context"
	"fmt"
	"liujun/Time_go-zero_csdn/common/ctxdata"
	"liujun/Time_go-zero_csdn/common/globalkey"
	"liujun/Time_go-zero_csdn/common/utils"
	"liujun/Time_go-zero_csdn/csdn/channel/cmd/rpc/channelclient"

	"liujun/Time_go-zero_csdn/csdn/channel/cmd/api/internal/svc"
	"liujun/Time_go-zero_csdn/csdn/channel/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ArticleToCollectionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewArticleToCollectionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ArticleToCollectionLogic {
	return &ArticleToCollectionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ArticleToCollectionLogic) ArticleToCollection(req *types.ArticleToCollectionRequest) (resp *types.ArticleToCollectionResponse, err error) {
	// todo: add your logic here and delete this line
	user_id := ctxdata.GetUidFromCtx(l.ctx)
	article, err := l.svcCtx.ChannelRpc.ArticleToCollection(l.ctx, &channelclient.ArticleToCollectionRequest{UserId: user_id, Aid: req.ArticleId})
	if err != nil {
		return nil, err
	}
	key := fmt.Sprintf(globalkey.ArticleStatus, utils.Int64ToString(article.Aid))
	field := globalkey.ArticleCollectionNum
	l.svcCtx.RedisClient.Hincrby(key, field, 1)
	return &types.ArticleToCollectionResponse{
		ArticleId: article.Aid,
	}, nil
}
