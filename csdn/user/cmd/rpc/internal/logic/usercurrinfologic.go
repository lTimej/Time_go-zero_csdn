package logic

import (
	"context"
	"liujun/Time_go-zero_csdn/common/minIO"
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
	user_info, err := l.svcCtx.UserModel.FindOneJoinUserProfileByUserId(l.ctx, in.UserId)
	if err != nil {
		return nil, err
	}
	var head_photo string
	if user_info.HeadPhoto == "" {
		head_photo = minIO.DefaultHeadPhoto
	} else {
		head_photo = "http://172.20.16.20:9000/" + user_info.HeadPhoto
	}
	m := NewUserFocusListLogic(l.ctx, l.svcCtx)
	focus, err := m.get_focus_user_id(in.UserId)
	if err != nil {
		return nil, err
	}
	fans, err := m.get_fans_user_id(in.UserId)
	if err != nil {
		return nil, err
	}
	return &user.UserCurrInfoResponse{
		UserName:  user_info.UserName,
		HeadPhoto: head_photo,
		Introduce: user_info.Introduce,
		CodeYear:  user_info.CodeYear,
		Career:    user_info.Career,
		Focus:     int32(len(focus)),
		Fans:      int32(len(fans)),
		Visitor:   0,
	}, nil
}
