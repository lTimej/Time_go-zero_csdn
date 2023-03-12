package logic

import (
	"context"
	"liujun/Time_go-zero_csdn/common/ctxdata"
	"liujun/Time_go-zero_csdn/csdn/user/cmd/rpc/userclient"
	"strconv"

	"liujun/Time_go-zero_csdn/csdn/user/cmd/api/internal/svc"
	"liujun/Time_go-zero_csdn/csdn/user/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CancelFocusUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCancelFocusUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CancelFocusUserLogic {
	return &CancelFocusUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CancelFocusUserLogic) CancelFocusUser(req *types.CancelFocusUserRequest) (resp *types.CancelFocusUserResponse, err error) {
	// todo: add your logic here and delete this line
	user_id := strconv.FormatInt(ctxdata.GetUidFromCtx(l.ctx), 10)
	_, err = l.svcCtx.UserRpc.CancelFocueUser(l.ctx, &userclient.CancelFocusUserRequest{UserId: user_id, TargetId: req.TargetUserId})
	if err != nil {
		return nil, err
	}
	return &types.CancelFocusUserResponse{Message: "ok"}, nil
}
