package logic

import (
	"context"

	"liujun/Time_go-zero_csdn/csdn/user/cmd/rpc/internal/svc"
	"liujun/Time_go-zero_csdn/csdn/user/cmd/rpc/types/user"
	"liujun/Time_go-zero_csdn/csdn/user/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserAddressLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserAddressLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserAddressLogic {
	return &UserAddressLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserAddressLogic) UserAddress(in *user.UserAddressRequest) (*user.UserAddressResponse, error) {
	// todo: add your logic here and delete this line
	address := &model.Address{
		UserId:     in.UserId,
		Receiver:   in.Receiver,
		Mobile:     in.Mobile,
		ProvinceId: in.ProvinceId,
		CityId:     in.CityId,
		DistrictId: in.DistrictId,
		Place:      in.Place,
	}

	_, err := l.svcCtx.UserAddressModel.Insert(l.ctx, address)
	if err != nil {
		return nil, err
	}
	return &user.UserAddressResponse{}, nil
}
