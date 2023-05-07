package logic

import (
	"context"

	"liujun/Time_go-zero_csdn/csdn/user/cmd/rpc/internal/svc"
	"liujun/Time_go-zero_csdn/csdn/user/cmd/rpc/types/user"

	"github.com/jinzhu/copier"
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
	resp := new(user.GetAddressResponse)
	copier.Copy(&resp.UserAddress, data)
	return resp, nil
}
