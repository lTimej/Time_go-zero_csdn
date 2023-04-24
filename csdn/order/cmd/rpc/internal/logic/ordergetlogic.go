package logic

import (
	"context"
	"liujun/Time_go-zero_csdn/csdn/order/model"

	"liujun/Time_go-zero_csdn/common/xerr"
	"liujun/Time_go-zero_csdn/csdn/order/cmd/rpc/internal/svc"
	"liujun/Time_go-zero_csdn/csdn/order/cmd/rpc/types/order"

	"github.com/Masterminds/squirrel"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type OrderGetLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewOrderGetLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OrderGetLogic {
	return &OrderGetLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *OrderGetLogic) OrderGet(in *order.OrderGetRequest) (*order.OrderGetResponse, error) {
	// todo: add your logic here and delete this line
	resp := new(order.OrderGetResponse)
	var ubuilder squirrel.SelectBuilder
	if in.PayStatus == 0 {
		ubuilder = l.svcCtx.OrderModel.Builder().Where("user_id = ?", in.UserId)
	} else {
		ubuilder = l.svcCtx.OrderModel.Builder().Where("user_id = ? and pay_status = ?", in.UserId, in.PayStatus)
	}

	orders, err := l.svcCtx.OrderModel.FindAllByUserId(l.ctx, ubuilder)
	if err != nil && err != model.ErrNotFound {
		return nil, errors.Wrapf(xerr.ErrDBError, "err:%v", err)
	}
	for _, or := range orders {
		var order_info order.OrderInfo
		order_info.AddressId = or.AddressId
		order_info.TotalPrice = or.TotalPrice
		order_info.TotalCount = or.TotalCount
		order_info.Sn = or.Sn
		order_info.Freight = or.Freight
		order_info.PayStatus = or.PayStatus
		order_info.OrderSpec = []*order.OrderSpec{}
		obuilder := l.svcCtx.OrderUserModel.Builder().Where("order_id = ?", or.Id)
		spus, err := l.svcCtx.OrderUserModel.FindOneByOrderId(l.ctx, obuilder)
		if err != nil && err != model.ErrNotFound {
			return nil, errors.Wrapf(xerr.ErrDBError, "err:%v", err)
		}
		for _, spu := range spus {
			order_info.OrderSpec = append(order_info.OrderSpec, &order.OrderSpec{
				SkuId:        spu.SkuId,
				SpecId:       spu.SpecId,
				Specs:        spu.Specs,
				Count:        spu.Count,
				Price:        spu.Price,
				DefaultImage: spu.DefaultImage,
				Title:        spu.Title,
			})
		}
		resp.OrderInfo = append(resp.OrderInfo, &order_info)
	}
	return resp, nil
}
