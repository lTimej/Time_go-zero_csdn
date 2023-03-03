package logic

import (
	"context"
	"liujun/Time_go-zero_csdn/csdn/user/cmd/rpc/userclient"

	"liujun/Time_go-zero_csdn/csdn/user/cmd/api/internal/svc"
	"liujun/Time_go-zero_csdn/csdn/user/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PhoneLoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPhoneLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PhoneLoginLogic {
	return &PhoneLoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PhoneLoginLogic) PhoneLogin(req *types.PhoneLoginRequest) (resp *types.PhoneLoginResponse, err error) {
	// todo: add your logic here and delete this line
	res, err := l.svcCtx.UserRpc.UserLogin(l.ctx, &userclient.LoginRequest{AuthType: "phone", Account: req.Phone, Password: req.Code})
	if err != nil {
		return nil, err
	}
	resp = new(types.PhoneLoginResponse)
	resp.Token = res.Token
	resp.RefreshToken = res.RefreshToken
	return
}
