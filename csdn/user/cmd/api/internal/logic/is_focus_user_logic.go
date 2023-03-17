package logic

import (
	"context"
	"liujun/Time_go-zero_csdn/common/ctxdata"
	"liujun/Time_go-zero_csdn/csdn/user/cmd/rpc/userclient"

	"liujun/Time_go-zero_csdn/csdn/user/cmd/api/internal/svc"
	"liujun/Time_go-zero_csdn/csdn/user/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type IsFocusUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewIsFocusUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *IsFocusUserLogic {
	return &IsFocusUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *IsFocusUserLogic) IsFocusUser(req *types.IsFocusUserRequest) (resp *types.IsFocusUserResponse, err error) {
	// todo: add your logic here and delete this line
	user_id := ctxdata.GetUidFromCtx(l.ctx)
	res, err := l.svcCtx.UserRpc.IsFocueUser(l.ctx, &userclient.IsFocusUserRequest{UserId: user_id, TargetId: req.TargetUserId})
	if err != nil {
		return nil, err
	}
	return &types.IsFocusUserResponse{IsFocusUser: res.IsFocusUser}, nil
}
