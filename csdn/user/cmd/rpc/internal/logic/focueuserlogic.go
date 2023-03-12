package logic

import (
	"context"

	"liujun/Time_go-zero_csdn/csdn/user/cmd/rpc/internal/svc"
	"liujun/Time_go-zero_csdn/csdn/user/cmd/rpc/types/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type FocueUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFocueUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FocueUserLogic {
	return &FocueUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FocueUserLogic) FocueUser(in *user.FocusUserRequest) (*user.FocusUserResponse, error) {
	// todo: add your logic here and delete this line

	return &user.FocusUserResponse{}, nil
}
