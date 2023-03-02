package logic

import (
	"context"
	"liujun/Time_go-zero_csdn/csdn/user/cmd/rpc/userclient"

	"liujun/Time_go-zero_csdn/csdn/user/cmd/api/internal/svc"
	"liujun/Time_go-zero_csdn/csdn/user/cmd/api/internal/types"
	"liujun/Time_go-zero_csdn/csdn/user/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserPasswordLoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserPasswordLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserPasswordLoginLogic {
	return &UserPasswordLoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserPasswordLoginLogic) UserPasswordLogin(req *types.UserPasswordLoginRequest) (resp *types.UserPasswordLoginResponse, err error) {
	// todo: add your logic here and delete this line
	res, err := l.svcCtx.UserRpc.UserLogin(l.ctx, &userclient.LoginRequest{
		AuthType: model.UserAuthTypeUsername,
		Account:  req.UserName,
		Password: req.Password,
	})
	if err != nil {
		return nil, err
	}

	resp = new(types.UserPasswordLoginResponse)
	resp.Token = res.Token
	resp.RefreshToken = res.RefreshToken
	return
}
