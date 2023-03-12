package logic

import (
	"context"
	"fmt"
	"liujun/Time_go-zero_csdn/common/ctxdata"
	"liujun/Time_go-zero_csdn/csdn/channel/cmd/rpc/channelclient"
	"strconv"

	"liujun/Time_go-zero_csdn/csdn/channel/cmd/api/internal/svc"
	"liujun/Time_go-zero_csdn/csdn/channel/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ArticleReadLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewArticleReadLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ArticleReadLogic {
	return &ArticleReadLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ArticleReadLogic) ArticleRead(req *types.ArticleReadRequest) (resp *types.ArticleReadResponse, err error) {
	// todo: add your logic here and delete this line
	uid := strconv.FormatInt(ctxdata.GetUidFromCtx(l.ctx), 10)
	res, err := l.svcCtx.ChannelRpc.ArticleRead(l.ctx, &channelclient.ArticleReadRequest{ArticleId: req.ArticleId, UserId: uid})
	if err != nil {
		fmt.Println(err, "哈哈哈哈哈")
		return nil, err
	}
	return &types.ArticleReadResponse{
		Message: "ok",
		Aid:     res.Aid,
	}, nil
}
