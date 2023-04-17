package logic

import (
	"context"
	"fmt"
	"liujun/Time_go-zero_csdn/common/ctxdata"
	"liujun/Time_go-zero_csdn/csdn/user/cmd/rpc/userclient"

	"github.com/jinzhu/copier"

	"liujun/Time_go-zero_csdn/csdn/user/cmd/api/internal/svc"
	"liujun/Time_go-zero_csdn/csdn/user/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserAddressLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserAddressLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserAddressLogic {
	return &GetUserAddressLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserAddressLogic) GetUserAddress(req *types.GetUserAddressRequest) (resp *types.GetUserAddressResponse, err error) {
	// todo: add your logic here and delete this line
	user_id := ctxdata.GetUidFromCtx(l.ctx)
	fmt.Println(user_id, "###########333333333")
	data, err := l.svcCtx.UserRpc.UserGetAddress(l.ctx, &userclient.GetUserAddressRequest{UserId: user_id})
	if err != nil {
		fmt.Println(err, "111111111222222222")
		return nil, err
	}
	resp = new(types.GetUserAddressResponse)
	copier.Copy(&resp.UserAddress, data.UserAddress)
	return
}
