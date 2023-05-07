package logic

import (
	"context"
	"fmt"

	"liujun/Time_go-zero_csdn/csdn/user/cmd/api/internal/svc"
	"liujun/Time_go-zero_csdn/csdn/user/cmd/api/internal/types"
	"liujun/Time_go-zero_csdn/csdn/user/cmd/rpc/userclient"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetOrderAddressLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetOrderAddressLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetOrderAddressLogic {
	return &GetOrderAddressLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetOrderAddressLogic) GetOrderAddress(req *types.GetAddressRequest) (resp *types.GetAddressResponse, err error) {
	// todo: add your logic here and delete this line
	data, err := l.svcCtx.UserRpc.GetAddress(l.ctx, &userclient.GetAddressRequest{AddressId: req.AddressId})
	if err != nil {
		fmt.Println(err, "111111111222222222")
		return nil, err
	}
	resp = new(types.GetAddressResponse)
	copier.Copy(&resp.UserAddress, data.UserAddress)
	return
}
