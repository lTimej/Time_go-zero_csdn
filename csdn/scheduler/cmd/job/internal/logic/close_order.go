package logic

import (
	"context"
	"encoding/json"
	"liujun/Time_go-zero_csdn/common/xerr"
	"liujun/Time_go-zero_csdn/csdn/order/cmd/rpc/types/order"
	"liujun/Time_go-zero_csdn/csdn/order/model"
	"liujun/Time_go-zero_csdn/csdn/scheduler/cmd/job/internal/svc"
	"liujun/Time_go-zero_csdn/csdn/scheduler/cmd/job/jobtype"

	"github.com/hibiken/asynq"
	"github.com/pkg/errors"
)

var ErrCloseOrderFal = xerr.NewErrMsg("关闭订单失败")

type CloseProductOrderHandler struct {
	svcCtx *svc.ServiceContext
}

func NewCloseProductOrderHandler(svcCtx *svc.ServiceContext) *CloseProductOrderHandler {
	return &CloseProductOrderHandler{
		svcCtx: svcCtx,
	}
}

// every one minute exec : if return err != nil , asynq will retry
func (l *CloseProductOrderHandler) ProcessTask(ctx context.Context, t *asynq.Task) error {
	var p jobtype.DeferCloseProductOrderPayload
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		return errors.Wrapf(ErrCloseOrderFal, "closeHomestayOrderStateMqHandler payload err:%v, payLoad:%+v", err, t.Payload())
	}

	resp, err := l.svcCtx.OrderRpc.OrderDesc(ctx, &order.OrderDescRequest{
		Sn: p.Sn,
	})
	if err != nil || resp.OrderDescInfo == nil {
		return errors.Wrapf(ErrCloseOrderFal, "closeHomestayOrderStateMqHandler  get order fail or order no exists err:%v, sn:%s ,HomestayOrder : %+v", err, p.Sn, resp.OrderDescInfo)
	}

	if resp.OrderDescInfo.PayStatus == model.OrderTradeStateWaitPay {
		_, err := l.svcCtx.OrderRpc.OrderUpdate(ctx, &order.OrderUpdateRequest{
			Sn:        p.Sn,
			PayStatus: model.OrderTradeStateCancel,
		})
		if err != nil {
			return errors.Wrapf(ErrCloseOrderFal, "CloseHomestayOrderHandler close order fail  err:%v, sn:%s ", err, p.Sn)
		}
	}
	return nil
}
