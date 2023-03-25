package logic

import (
	"context"
	"liujun/Time_go-zero_csdn/common/ctxdata"
	"liujun/Time_go-zero_csdn/csdn/user/cmd/rpc/userclient"

	"liujun/Time_go-zero_csdn/csdn/user/cmd/api/internal/svc"
	"liujun/Time_go-zero_csdn/csdn/user/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserInfoEditLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserInfoEditLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInfoEditLogic {
	return &UserInfoEditLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserInfoEditLogic) UserInfoEdit(req *types.UserInfoEditRequest) (resp *types.UserInfoEditResponse, err error) {
	// todo: add your logic here and delete this line
	user_id := ctxdata.GetUidFromCtx(l.ctx)
	info, err := l.svcCtx.UserRpc.UserInfoEdit(l.ctx, &userclient.UserInfoEditRequest{
		HeadPhoto: req.HeadPhoto,
		OldPwd:    req.OldPwd,
		NewPwd:    req.NewPwd,
		UserName:  req.UserName,
		Gender:    req.Gender,
		Introduce: req.Introduce,
		Tag:       req.Tag,
		AuthName:  req.AuthName,
		Birthday:  req.Birthday,
		Areas:     req.Areas,
		UserId:    user_id,
	})
	if err != nil {
		return nil, err
	}
	return &types.UserInfoEditResponse{
		HeadPhoto: info.HeadPhoto,
		Pwd:       info.Pwd,
		UserName:  info.UserName,
		Gender:    info.Gender,
		Introduce: info.Introduce,
		Tag:       info.Tag,
		AuthName:  info.AuthName,
		Birthday:  info.Birthday,
		Areas:     info.Areas,
	}, nil
}
