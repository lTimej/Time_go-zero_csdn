package logic

import (
	"context"
	"liujun/Time_go-zero_csdn/common/xerr"
	"liujun/Time_go-zero_csdn/csdn/user/cmd/rpc/userclient"

	"liujun/Time_go-zero_csdn/csdn/user/cmd/api/internal/svc"
	"liujun/Time_go-zero_csdn/csdn/user/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SendSmsCodeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSendSmsCodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SendSmsCodeLogic {
	return &SendSmsCodeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SendSmsCodeLogic) SendSmsCode(req *types.SendSmsCodeRequest, phone string) (resp *types.SendSmsCodeReponse, err error) {
	// todo: add your logic here and delete this line
	key := "sms:code:" + phone
	ok, _ := l.svcCtx.RedisClient.Exists(key)
	if ok {
		return nil, xerr.NewErrMsg("不要频繁发送")
	}
	_, err = l.svcCtx.UserRpc.SendSmsCode(l.ctx, &userclient.SmsRequest{Phone: phone})
	if err != nil {
		return nil, err
	}
	return
}
