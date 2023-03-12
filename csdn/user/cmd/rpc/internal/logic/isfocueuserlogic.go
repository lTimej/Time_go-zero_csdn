package logic

import (
	"context"
	"fmt"

	"liujun/Time_go-zero_csdn/csdn/user/cmd/rpc/internal/svc"
	"liujun/Time_go-zero_csdn/csdn/user/cmd/rpc/types/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type IsFocueUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewIsFocueUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *IsFocueUserLogic {
	return &IsFocueUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *IsFocueUserLogic) IsFocueUser(in *user.IsFocusUserRequest) (*user.IsFocusUserResponse, error) {
	// todo: add your logic here and delete this line
	fmt.Println(in.UserId, in.TargetId, "&&&&&&&&&&&&&&")
	res, err := l.svcCtx.UserRelationModel.FindByUserIdTargetUserId(l.ctx, in.UserId, in.TargetId)
	if err != nil {
		return nil, err
	}
	var is_focus_user bool
	if res != nil && res.Relation != 0 {
		is_focus_user = true
	}
	return &user.IsFocusUserResponse{
		IsFocusUser: is_focus_user,
	}, nil
}
