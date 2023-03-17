package logic

import (
	"context"
	"liujun/Time_go-zero_csdn/common/ctxdata"
	"liujun/Time_go-zero_csdn/csdn/user/cmd/rpc/userclient"

	"liujun/Time_go-zero_csdn/csdn/user/cmd/api/internal/svc"
	"liujun/Time_go-zero_csdn/csdn/user/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FocusUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFocusUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FocusUserLogic {
	return &FocusUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FocusUserLogic) FocusUser(req *types.FocusUserRequest) (resp *types.FocusUserResponse, err error) {
	// todo: add your logic here and delete this line
	user_id := ctxdata.GetUidFromCtx(l.ctx)
	_, err = l.svcCtx.UserRpc.FocueUser(l.ctx, &userclient.FocusUserRequest{UserId: user_id, TargetId: req.TargetUserId})
	if err != nil {
		return nil, err
	}
	return &types.FocusUserResponse{
		TargetUserId: req.TargetUserId,
	}, nil
}
