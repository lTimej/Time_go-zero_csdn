package logic

import (
	"context"
	"liujun/Time_go-zero_csdn/common/ctxdata"
	"liujun/Time_go-zero_csdn/csdn/user/cmd/rpc/userclient"

	"github.com/jinzhu/copier"

	"liujun/Time_go-zero_csdn/csdn/user/cmd/api/internal/svc"
	"liujun/Time_go-zero_csdn/csdn/user/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserFocusListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserFocusListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserFocusListLogic {
	return &UserFocusListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserFocusListLogic) UserFocusList(req *types.UserFocusListRequest) (resp *types.UserFocusListResponse, err error) {
	// todo: add your logic here and delete this line
	user_id := ctxdata.GetUidFromCtx(l.ctx)
	focus, err := l.svcCtx.UserRpc.UserFocusList(l.ctx, &userclient.UserFocusListRequest{UserId: user_id, Page: req.Page, PageNum: req.PageNum})
	if err != nil {
		return nil, err
	}
	resp = new(types.UserFocusListResponse)
	_ = copier.Copy(&resp, focus)
	return
}
