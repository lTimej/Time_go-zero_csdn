package logic

import (
	"context"
	"liujun/Time_go-zero_csdn/common/xerr"
	"liujun/Time_go-zero_csdn/csdn/order/cmd/api/internal/svc"
	"liujun/Time_go-zero_csdn/csdn/order/cmd/api/internal/types"
	"liujun/Time_go-zero_csdn/csdn/order/cmd/rpc/orderclient"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type OrderUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewOrderUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OrderUpdateLogic {
	return &OrderUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *OrderUpdateLogic) OrderUpdate(req *types.OrderUpdateRequest) (resp *types.OrderUpdateResponse, err error) {
	// todo: add your logic here and delete this line
	data, err := l.svcCtx.OrderRpc.OrderUpdate(l.ctx, &orderclient.OrderUpdateRequest{
		Sn:        req.Sn,
		PayStatus: req.PayStatus,
	})
	if err != nil {
		return nil, errors.Wrapf(xerr.ErrDBError, "err:%v", err)
	}
	resp = new(types.OrderUpdateResponse)
	resp.Sn = data.Sn
	return
}
