package logic

import (
	"context"

	"liujun/Time_go-zero_csdn/common/ctxdata"
	"liujun/Time_go-zero_csdn/csdn/user/cmd/api/internal/svc"
	"liujun/Time_go-zero_csdn/csdn/user/cmd/api/internal/types"
	"liujun/Time_go-zero_csdn/csdn/user/cmd/rpc/userclient"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type UserFansListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserFansListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserFansListLogic {
	return &UserFansListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserFansListLogic) UserFansList(req *types.UserFansListRequest) (resp *types.UserFansListResponse, err error) {
	// todo: add your logic here and delete this line
	user_id := ctxdata.GetUidFromCtx(l.ctx)
	fans, err := l.svcCtx.UserRpc.UserFansList(l.ctx, &userclient.UserFansListRequest{UserId: user_id, Page: req.Page, PageNum: req.PageNum})
	if err != nil {
		return nil, err
	}
	resp = new(types.UserFansListResponse)
	_ = copier.Copy(&resp, fans)
	return
}
