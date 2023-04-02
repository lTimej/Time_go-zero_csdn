package logic

import (
	"context"

	"liujun/Time_go-zero_csdn/csdn/im/cmd/api/internal/svc"
	"liujun/Time_go-zero_csdn/csdn/im/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserChatUploadPictureLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserChatUploadPictureLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserChatUploadPictureLogic {
	return &UserChatUploadPictureLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserChatUploadPictureLogic) UserChatUploadPicture(req *types.UserChatUploadPictureRequest) (resp *types.UserChatUploadPictureResponse, err error) {
	// todo: add your logic here and delete this line
	resp = new(types.UserChatUploadPictureResponse)
	resp.Picture = req.Picture
	return
}
