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
	fmt.Println("===========", l.ctx.Value(ctxdata.CtxKeyJwtUserId), "==========")
	user_id := ctxdata.GetUidFromCtx(l.ctx)
	res, err := l.svcCtx.UserRpc.UserCurrInfo(l.ctx, &userclient.UserCurrInfoRequest{UserId: user_id})
	if err != nil {
		return nil, err
	}
	fmt.Println(res, "&&&&&&**********((((((((")
	resp = new(types.UserCurrInfoResponse)
	resp.UserName = res.UserName
	resp.HeadPhoto = res.HeadPhoto
	resp.Introduce = res.Introduce
	resp.Career = res.Career
	resp.CodeYear = res.CodeYear
	return resp, nil
}
