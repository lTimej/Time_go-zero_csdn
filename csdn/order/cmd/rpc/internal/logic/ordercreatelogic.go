package logic

import (
	"context"
	"errors"
	"fmt"
	"liujun/Time_go-zero_csdn/common/order_var"
	"liujun/Time_go-zero_csdn/common/utils"
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
	sn := utils.GetSN("LJ")
	orders := model.Order{
		UserId:     in.UserId,
		AddressId:  in.AddressId,
		TotalCount: in.TotalCount,
		TotalPrice: in.TotalPrice,
		Freight:    order_var.Freight,
		Version:    order_var.Version,
		Sn:         sn,
	}
	_, err := l.svcCtx.OrderModel.Insert(l.ctx, &orders)
	if err != nil {
		fmt.Println(err, "11111111111111")
		return nil, err
	}
	o, err := l.svcCtx.OrderModel.FindOneBySn(l.ctx, sn)
	if err != nil {
		fmt.Println(err, "222222222222")
		return nil, err
	}
	if o == nil {
		fmt.Println(err, "333333333333333")
		return nil, errors.New("订单不存在")
	}
	order_id := o.Id
	fmt.Println(order_id, in.Sku, "55555555555555555")
	for _, s := range in.Sku {
		user_order := model.UserOrder{
			OrderId:     order_id,
			SkuId:       s.SkuId,
			SpecId:      s.SpecId,
			Specs:       s.Specs,
			Count:       s.Count,
			Comment:     "",
			Score:       0,
			IsAnonymous: 0,
			IsCommented: 0,
		}
		_, err = l.svcCtx.OrderUserModel.Insert(l.ctx, &user_order)
		if err != nil {
			fmt.Println(err, "444444444444444444")
			return nil, err
		}
	}

	return &order.OrderCreateResponse{
		Sn: sn,
	}, nil
}
