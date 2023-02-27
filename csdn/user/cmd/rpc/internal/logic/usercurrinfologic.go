package logic

import (
	"context"

	"liujun/Time_go-zero_csdn/csdn/user/cmd/rpc/internal/svc"
	"liujun/Time_go-zero_csdn/csdn/user/cmd/rpc/types/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserCurrInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserCurrInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserCurrInfoLogic {
	return &UserCurrInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserCurrInfoLogic) UserCurrInfo(in *user.UserCurrInfoRequest) (*user.UserCurrInfoResponse, error) {
	// todo: add your logic here and delete this line
	resp := new(user.UserCurrInfoResponse)
	l.svcCtx.UserMysql.Table("user_basic").Joins("left join user_profile on user_basic.user_id = user_profile.user_id").Select("user_basic.user_name,user_basic.profile_photo as head_photo,user_basic.introduce,user_basic.code_year,user_profile.career").Scan(resp)
	return resp, nil
}
