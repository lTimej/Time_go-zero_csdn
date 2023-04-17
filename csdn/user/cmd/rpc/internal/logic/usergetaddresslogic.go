package logic

import (
	"context"
	"fmt"

	"github.com/jinzhu/copier"

	"liujun/Time_go-zero_csdn/csdn/user/cmd/rpc/internal/svc"
	"liujun/Time_go-zero_csdn/csdn/user/cmd/rpc/types/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserGetAddressLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserGetAddressLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserGetAddressLogic {
	return &UserGetAddressLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserGetAddressLogic) UserGetAddress(in *user.GetUserAddressRequest) (*user.GetUserAddressResponse, error) {
	// todo: add your logic here and delete this line
	address_builder := l.svcCtx.UserAddressModel.AddressBuilder()
	data, err := l.svcCtx.UserAddressModel.FindAllByUserId(l.ctx, address_builder, in.UserId)
	fmt.Println(data, "####333333333333", len(data))
	if err != nil {
		return nil, err
	}
	resp := new(user.GetUserAddressResponse)
	copier.Copy(&resp.UserAddress, data)
	fmt.Println(resp.UserAddress, "444444444444", len(resp.UserAddress))
	return resp, nil
}
