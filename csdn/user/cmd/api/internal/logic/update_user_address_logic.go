package logic

import (
	"context"
	"liujun/Time_go-zero_csdn/csdn/user/cmd/rpc/userclient"

	"liujun/Time_go-zero_csdn/csdn/user/cmd/api/internal/svc"
	"liujun/Time_go-zero_csdn/csdn/user/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateUserAddressLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateUserAddressLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserAddressLogic {
	return &UpdateUserAddressLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateUserAddressLogic) UpdateUserAddress(req *types.UpdateUserAddressRequest) (resp *types.UpdateUserAddressResponse, err error) {
	// todo: add your logic here and delete this line
	_, err = l.svcCtx.UserRpc.UserUpdateAddress(l.ctx, &userclient.UpdateUserAddressRequest{
		AddressId:  req.AddressId,
		Receiver:   req.Receiver,
		Mobile:     req.Mobile,
		ProvinceId: req.ProvinceId,
		CityId:     req.CityId,
		DistrictId: req.DistrictId,
		Place:      req.Place,
		Email:      req.Email,
		IsDefault:  req.IsDefault,
	})
	if err != nil {
		return nil, err
	}
	return
}
