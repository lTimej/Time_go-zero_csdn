package logic

import (
	"context"
	"fmt"

	"liujun/Time_go-zero_csdn/csdn/user/cmd/rpc/internal/svc"
	"liujun/Time_go-zero_csdn/csdn/user/cmd/rpc/types/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAddressLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetAddressLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAddressLogic {
	return &GetAddressLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetAddressLogic) GetAddress(in *user.GetAddressRequest) (*user.GetAddressResponse, error) {
	// todo: add your logic here and delete this line
	address_builder := l.svcCtx.UserAddressModel.AddressBuilder()
	data, err := l.svcCtx.UserAddressModel.FindOneByAddressId(l.ctx, address_builder, in.AddressId)
	if err != nil {
		return nil, err
	}
	fmt.Println(data, "8888888888888")
	resp := new(user.GetAddressResponse)
	resp.UserAddress = &user.UserAddress{
		AddressId:  data.AddressId,
		Receiver:   data.Receiver,
		Mobile:     data.Mobile,
		ProvinceId: data.ProvinceId,
		CityId:     data.CityId,
		DistrictId: data.DistrictId,
		Province:   data.Province,
		City:       data.City,
		District:   data.District,
		Place:      data.Place,
		IsDefault:  data.IsDefault,
		Email:      data.Email,
	}
	// copier.Copy(&resp.UserAddress, data)
	fmt.Println(resp.UserAddress, "777777777777")
	return resp, nil
}
