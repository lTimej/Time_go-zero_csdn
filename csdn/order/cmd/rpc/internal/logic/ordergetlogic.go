package logic

import (
	"context"
	"fmt"

	"github.com/jinzhu/copier"

	"liujun/Time_go-zero_csdn/csdn/order/cmd/rpc/internal/svc"
	"liujun/Time_go-zero_csdn/csdn/order/cmd/rpc/types/order"

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
	builder := l.svcCtx.OrderModel.Builder()
	data, err := l.svcCtx.OrderModel.FindAllByUserId(l.ctx, builder, in.UserId)
	if err != nil {
		return nil, err
	}
	fmt.Println(data, "111111111")
	resp := new(order.OrderGetResponse)
	copier.Copy(&resp.Orders, data)
	fmt.Println(resp, "22222222222")
	return resp, nil
}
