package logic

import (
	"context"
	"liujun/Time_go-zero_csdn/common/xerr"
	"liujun/Time_go-zero_csdn/csdn/order/cmd/rpc/internal/svc"
	"liujun/Time_go-zero_csdn/csdn/order/cmd/rpc/types/order"
	"liujun/Time_go-zero_csdn/csdn/order/model"

	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type OrderDescLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewOrderDescLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OrderDescLogic {
	return &OrderDescLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *OrderDescLogic) OrderDesc(in *order.OrderDescRequest) (*order.OrderDescResponse, error) {
	// todo: add your logic here and delete this line
	order_desc, err := l.svcCtx.OrderModel.FindOneBySn(l.ctx, in.Sn)
	if err != nil && err != model.ErrNotFound {
		return nil, errors.Wrapf(xerr.ErrDBError, "err:%v", err)
	}
	resp := new(order.OrderDescResponse)
	if order_desc != nil {
		_ = copier.Copy(&resp, order_desc)
	}
	return resp, nil
}
