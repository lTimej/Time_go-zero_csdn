package logic

import (
	"context"

	"liujun/Time_go-zero_csdn/csdn/user/cmd/rpc/internal/svc"
	"liujun/Time_go-zero_csdn/csdn/user/cmd/rpc/types/user"
	"liujun/Time_go-zero_csdn/csdn/user/model"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type UserUpdateAddressLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserUpdateAddressLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserUpdateAddressLogic {
	return &UserUpdateAddressLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserUpdateAddressLogic) UserUpdateAddress(in *user.UpdateUserAddressRequest) (*user.UpdateUserAddressResponse, error) {
	// todo: add your logic here and delete this line
	address := &model.Address{
		Id:         in.AddressId,
		Receiver:   in.Receiver,
		Mobile:     in.Mobile,
		ProvinceId: in.ProvinceId,
		CityId:     in.CityId,
		DistrictId: in.DistrictId,
		Place:      in.Place,
	}
	addr, err := l.svcCtx.UserAddressModel.FindOne(l.ctx, in.AddressId)
	if err != nil {
		return nil, err
	}
	if addr == nil {
		return nil, errors.New("地址不存在")
	}
	err = l.svcCtx.UserAddressModel.Update(l.ctx, address)
	if err != nil {
		return nil, err
	}
	return &user.UpdateUserAddressResponse{}, nil
}
