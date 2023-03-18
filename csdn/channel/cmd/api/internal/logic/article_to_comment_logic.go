package logic

import (
	"context"
	"liujun/Time_go-zero_csdn/common/ctxdata"
	"liujun/Time_go-zero_csdn/csdn/channel/cmd/rpc/channelclient"

	"liujun/Time_go-zero_csdn/csdn/channel/cmd/api/internal/svc"
	"liujun/Time_go-zero_csdn/csdn/channel/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ArticleToCommentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewArticleToCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ArticleToCommentLogic {
	return &ArticleToCommentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ArticleToCommentLogic) ArticleToComment(req *types.ArticleToCommentRequest) (resp *types.ArticleToCommentResponse, err error) {
	// todo: add your logic here and delete this line
	user_id := ctxdata.GetUidFromCtx(l.ctx)
	comment, err := l.svcCtx.ChannelRpc.ArticleToComment(l.ctx, &channelclient.ArticleToCommnetRequest{UserId: user_id, ArticleId: req.ArticleId, CommentId: req.CommentId, Content: req.Content})
	if err != nil {
		return nil, err
	}
	return &types.ArticleToCommentResponse{
		ArticleId:       comment.ArticleId,
		CommentId:       comment.CommentId,
		CommentParentId: comment.CommentParentId,
	}, nil
}
