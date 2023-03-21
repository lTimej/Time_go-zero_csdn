package logic

import (
	"context"
	"fmt"

	"liujun/Time_go-zero_csdn/common/ctxdata"
	"liujun/Time_go-zero_csdn/csdn/user/cmd/api/internal/svc"
	"liujun/Time_go-zero_csdn/csdn/user/cmd/api/internal/types"
	"liujun/Time_go-zero_csdn/csdn/user/cmd/rpc/userclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserCurrInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserCurrInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserCurrInfoLogic {
	return &UserCurrInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserCurrInfoLogic) UserCurrInfo(req *types.UserCurrInfoRequest) (resp *types.UserCurrInfoResponse, err error) {
	// todo: add your logic here and delete this line
	user_id := ctxdata.GetUidFromCtx(l.ctx)
	fmt.Println(user_id, "二二二热呃呃呃呃呃呃呃")
	res, err := l.svcCtx.UserRpc.UserCurrInfo(l.ctx, &userclient.UserCurrInfoRequest{UserId: user_id})
	if err != nil {
		fmt.Println(err, "潍坊市看到房价是否考虑")
		return nil, err
	}
	resp = new(types.UserCurrInfoResponse)
	resp.UserName = res.UserName
	resp.HeadPhoto = res.HeadPhoto
	resp.Introduce = res.Introduce
	resp.Career = res.Career
	resp.CodeYear = res.CodeYear
	resp.Focus = res.Focus
	resp.Fans = res.Fans
	return resp, nil
}
