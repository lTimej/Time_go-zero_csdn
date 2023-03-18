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

type ArticleToDisCollectionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewArticleToDisCollectionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ArticleToDisCollectionLogic {
	return &ArticleToDisCollectionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ArticleToDisCollectionLogic) ArticleToDisCollection(req *types.ArticleToDisCollectionRequest) (resp *types.ArticleToDisCollectionResponse, err error) {
	// todo: add your logic here and delete this line
	user_id := ctxdata.GetUidFromCtx(l.ctx)
	_, err = l.svcCtx.ChannelRpc.ArticleToDisCollection(l.ctx, &channelclient.ArticleToDisCollectionRequest{UserId: user_id, Aid: req.ArticleId})
	if err != nil {
		return nil, err
	}
	key := fmt.Sprintf(globalkey.ArticleStatus, utils.Int64ToString(req.ArticleId))
	field := globalkey.ArticleCollectionNum
	l.svcCtx.RedisClient.Hincrby(key, field, -1)
	return &types.ArticleToDisCollectionResponse{Message: "success"}, nil
}
