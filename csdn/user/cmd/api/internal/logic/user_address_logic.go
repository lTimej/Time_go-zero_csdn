package logic

import (
	"context"
	"liujun/Time_go-zero_csdn/common/ctxdata"
	"liujun/Time_go-zero_csdn/csdn/user/cmd/rpc/userclient"

	"liujun/Time_go-zero_csdn/csdn/user/cmd/api/internal/svc"
	"liujun/Time_go-zero_csdn/csdn/user/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserAddressLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserAddressLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserAddressLogic {
	return &UserAddressLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserAddressLogic) UserAddress(req *types.UserAddressRequest) (resp *types.UserAddressResponse, err error) {
	// todo: add your logic here and delete this line
	user_id := ctxdata.GetUidFromCtx(l.ctx)
	_, err = l.svcCtx.UserRpc.UserAddress(l.ctx, &userclient.UserAddressRequest{
		UserId:     user_id,
		Receiver:   req.Receiver,
		Mobile:     req.Mobile,
		ProvinceId: req.ProvinceId,
		CityId:     req.CityId,
		DistrictId: req.DistrictId,
		Place:      req.Place,
	})
	if err != nil {
		return nil, err
	}
	return
}
