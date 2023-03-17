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

type ArticleStatusLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewArticleStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ArticleStatusLogic {
	return &ArticleStatusLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ArticleStatusLogic) ArticleStatus(req *types.ArticleStatusRequest) (resp *types.ArticleStatusResponse, err error) {
	// todo: add your logic here and delete this line
	user_id := ctxdata.GetUidFromCtx(l.ctx)
	fmt.Println(ctxdata.GetUidFromCtx(l.ctx), "^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^")
	article_status, err := l.svcCtx.ChannelRpc.ArticleStatus(l.ctx, &channelclient.ArticlestatusRequest{ArticleId: req.ArticleId, UserId: user_id, TargetId: req.UserId})
	if err != nil {
		return nil, err
	}
	return &types.ArticleStatusResponse{
		CollectionNum: article_status.CollectionNum,
		ReadNum:       article_status.ReadNum,
		LikeNum:       article_status.LikeNum,
		Aid:           article_status.Aid,
		Iscollection:  article_status.Iscollection,
		Islike:        article_status.Islike,
		Isfocus:       article_status.Isfocus,
	}, nil
}
