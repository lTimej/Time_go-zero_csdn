package logic

import (
	"context"
	"fmt"

	"liujun/Time_go-zero_csdn/csdn/user/cmd/api/internal/svc"
	"liujun/Time_go-zero_csdn/csdn/user/cmd/api/internal/types"
	"liujun/Time_go-zero_csdn/csdn/user/cmd/rpc/userclient"

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
	// copier.Copy(&resp.UserAddress, data.UserAddress)
	resp.UserAddress = &types.UserAddress{
		AddressId:  data.UserAddress.AddressId,
		Receiver:   data.UserAddress.Receiver,
		Mobile:     data.UserAddress.Mobile,
		ProvinceId: data.UserAddress.ProvinceId,
		CityId:     data.UserAddress.CityId,
		DistrictId: data.UserAddress.DistrictId,
		Province:   data.UserAddress.Province,
		City:       data.UserAddress.City,
		District:   data.UserAddress.District,
		Place:      data.UserAddress.Place,
		IsDefault:  data.UserAddress.IsDefault,
		Email:      data.UserAddress.Email,
	}
	return
}
