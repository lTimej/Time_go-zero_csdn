package logic

import (
	"context"
	"liujun/Time_go-zero_csdn/csdn/channel/cmd/rpc/channelclient"

	"liujun/Time_go-zero_csdn/csdn/channel/cmd/api/internal/svc"
	"liujun/Time_go-zero_csdn/csdn/channel/cmd/api/internal/types"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type ArticleCommentListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewArticleCommentListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ArticleCommentListLogic {
	return &ArticleCommentListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ArticleCommentListLogic) ArticleCommentList(req *types.ArticleCommentListRequest) (resp *types.ArticleCommentListResponse, err error) {
	// todo: add your logic here and delete this line
	data, err := l.svcCtx.ChannelRpc.ArticleCommentList(l.ctx, &channelclient.ArticleCommentListRequest{
		Type:      req.Ty,
		ArticleId: req.ArticleId,
		Offset:    req.Offset,
		Limit:     req.Limit,
	})

	if err != nil {
		return nil, err
	}
	resp = new(types.ArticleCommentListResponse)
	_ = copier.Copy(&resp, data)
	return
}
