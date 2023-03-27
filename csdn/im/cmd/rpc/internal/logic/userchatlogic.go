package logic

import (
	"context"

	"liujun/Time_go-zero_csdn/csdn/im/cmd/rpc/internal/svc"
	"liujun/Time_go-zero_csdn/csdn/im/cmd/rpc/types/im"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserChatLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserChatLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserChatLogic {
	return &UserChatLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserChatLogic) UserChat(in *im.UserChatRequest) (*im.UserChatResponse, error) {
	// todo: add your logic here and delete this line

	return &im.UserChatResponse{}, nil
}
