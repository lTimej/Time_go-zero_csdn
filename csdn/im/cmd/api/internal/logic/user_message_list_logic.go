package logic

import (
	"context"
	"liujun/Time_go-zero_csdn/common/ctxdata"
	"liujun/Time_go-zero_csdn/csdn/im/cmd/rpc/imclient"

	"github.com/jinzhu/copier"

	"liujun/Time_go-zero_csdn/csdn/im/cmd/api/internal/svc"
	"liujun/Time_go-zero_csdn/csdn/im/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserMessageListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserMessageListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserMessageListLogic {
	return &UserMessageListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserMessageListLogic) UserMessageList(req *types.UserMessageListRequest) (resp *types.UserMessageListResponse, err error) {
	// todo: add your logic here and delete this line
	user_id := ctxdata.GetUidFromCtx(l.ctx)
	res, err := l.svcCtx.ImRpc.UserMessageList(l.ctx, &imclient.UserMessageListRequest{UserId: user_id})
	if err != nil {
		return nil, err
	}
	resp = new(types.UserMessageListResponse)
	copier.Copy(resp, res)
	return
}
