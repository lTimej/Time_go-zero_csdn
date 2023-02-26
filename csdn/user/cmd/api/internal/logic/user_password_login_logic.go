package logic

import (
	"context"
	"fmt"
	"liujun/Time_go-zero_csdn/csdn/user/cmd/rpc/userclient"

	"liujun/Time_go-zero_csdn/csdn/user/cmd/api/internal/svc"
	"liujun/Time_go-zero_csdn/csdn/user/cmd/api/internal/types"

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
	res, err := l.svcCtx.UserRpc.UserPasswordLogin(l.ctx, &userclient.UserPasswordRequest{
		Username: req.UserName,
		Password: req.Password,
	})
	fmt.Println("10000000000000000000000000000000000000000000")
	if err != nil {
		fmt.Println("22222222222222222222=======22222222222222222222", err, 33333)
		return nil, err
	}
	fmt.Println("嘻嘻嘻嘻嘻嘻嘻嘻嘻", resp)
	resp = new(types.UserPasswordLoginResponse)
	resp.Token = res.Token
	resp.RefreshToken = res.RefreshToken
	fmt.Println("哈哈哈哈哈哈哈哈啊哈哈哈哈", resp)
	return
}
