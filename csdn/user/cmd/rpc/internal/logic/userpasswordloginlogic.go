package logic

import (
	"context"
	"errors"
	"fmt"
	"liujun/Time_go-zero_csdn/csdn/user/cmd/rpc/internal/svc"
	"liujun/Time_go-zero_csdn/csdn/user/cmd/rpc/types/user"
	"liujun/Time_go-zero_csdn/csdn/user/model"
	"liujun/Time_go-zero_csdn/csdn/user/utils"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserPasswordLoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserPasswordLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserPasswordLoginLogic {
	return &UserPasswordLoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserPasswordLoginLogic) UserPasswordLogin(in *user.UserPasswordRequest) (*user.UserPasswordResponse, error) {
	// todo: add your logic here and delete this line
	username := in.Username
	password := in.Password
	fmt.Println(username, password)
	user_basic := new(model.UserBasic)
	resp := new(user.UserPasswordResponse)
	l.svcCtx.UserMysql.Where("mobile = ? AND password = ?", username, password).First(user_basic)
	if user_basic.UserName == "" {
		return resp, errors.New("用户名错误")
	}
	token, err := utils.GenToken(user_basic.UserId, time.Hour*2)
	if err != nil {
		return resp, err
	}
	refresh_token, err := utils.GenToken(user_basic.UserId, time.Hour*7)
	if err != nil {
		return resp, err
	}
	resp.Token = token
	resp.RefreshToken = refresh_token
	fmt.Println(resp)
	return resp, nil
}
