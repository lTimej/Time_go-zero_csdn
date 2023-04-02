package logic

import (
	"context"
	"fmt"

	"liujun/Time_go-zero_csdn/common/ctxdata"
	"liujun/Time_go-zero_csdn/common/globalkey"
	"liujun/Time_go-zero_csdn/csdn/im/cmd/api/internal/svc"
	"liujun/Time_go-zero_csdn/csdn/im/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserChatCountLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserChatCountLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserChatCountLogic {
	return &UserChatCountLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserChatCountLogic) UserChatCount(req *types.UserChatCountRequest) (resp *types.UserChatCountResponse, err error) {
	// todo: add your logic here and delete this line
	user_id := ctxdata.GetUidFromCtx(l.ctx)
	user_chat_count_key := fmt.Sprintf(globalkey.UserChatCount, user_id)
	msg_count := l.svcCtx.RedisIm.ZCard(l.ctx, user_chat_count_key).Val()
	return &types.UserChatCountResponse{
		Count: msg_count,
	}, nil
}
