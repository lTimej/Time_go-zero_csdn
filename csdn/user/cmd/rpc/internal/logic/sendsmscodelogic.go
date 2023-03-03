package logic

import (
	"context"
	"liujun/Time_go-zero_csdn/common/utils"
	"liujun/Time_go-zero_csdn/common/xerr"

	"liujun/Time_go-zero_csdn/csdn/user/cmd/rpc/internal/svc"
	"liujun/Time_go-zero_csdn/csdn/user/cmd/rpc/types/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type SendSmsCodeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSendSmsCodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SendSmsCodeLogic {
	return &SendSmsCodeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SendSmsCodeLogic) SendSmsCode(in *user.SmsRequest) (*user.SmsResponse, error) {
	// todo: add your logic here and delete this line
	code := utils.GetRandNum(6)
	// err := sms.SendSmsCode(code, in.Phone)
	// if err != nil {
	// 	return nil, err
	// }
	key := "sms:code:" + in.Phone
	err := l.svcCtx.RedisClient.Setex(key, code, 60)
	if err != nil {
		return nil, xerr.NewErrMsg("验证码存储错误")
	}
	return &user.SmsResponse{}, nil
}
