package logic

import (
	"context"
	"errors"

	"liujun/Time_go-zero_csdn/csdn/order/cmd/rpc/internal/svc"
	"liujun/Time_go-zero_csdn/csdn/order/cmd/rpc/types/order"
	"liujun/Time_go-zero_csdn/csdn/order/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type OrderCreateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewOrderCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OrderCreateLogic {
	return &OrderCreateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *OrderCreateLogic) OrderCreate(in *order.OrderCreateRequest) (*order.OrderCreateResponse, error) {
	// todo: add your logic here and delete this line
	orders := model.Order{
		UserId:     in.UserId,
		AddressId:  in.AddressId,
		TotalCount: in.TotalCount,
		TotalPrice: in.TotalPrice,
		Freight:    in.Freight,
		Version:    in.Version,
		Sn:         in.Sn,
	}
	_, err := l.svcCtx.OrderModel.Insert(l.ctx, &orders)
	if err != nil {
		return nil, err
	}
	o, err := l.svcCtx.OrderModel.FindOneBySn(l.ctx, in.Sn)
	if err != nil {
		return nil, err
	}
	if o == nil {
		return nil, errors.New("订单不存在")
	}
	order_id := o.Id
	for _, s := range in.Sku {
		user_order := model.UserOrder{
			OrderId:     order_id,
			SkuId:       s.SkuId,
			SpecId:      s.SpecId,
			Specs:       s.Specs,
			Comment:     "",
			Score:       0,
			IsAnonymous: 0,
			IsCommented: 0,
		}
		_, err = l.svcCtx.OrderUserModel.Insert(l.ctx, &user_order)
		if err != nil {
			return nil, err
		}
	}
	return &order.OrderCreateResponse{
		Sn: o.Sn,
	}, nil
}
