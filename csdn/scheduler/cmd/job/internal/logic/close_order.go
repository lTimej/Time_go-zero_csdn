package logic

import (
	"context"
	"encoding/json"
	"liujun/Time_go-zero_csdn/csdn/scheduler/cmd/job/internal/svc"
	"liujun/Time_go-zero_csdn/csdn/scheduler/cmd/job/jobtype"

	"github.com/hibiken/asynq"
	"github.com/pkg/errors"
)

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

	resp, err := l.svcCtx.OrderRpc.HomestayOrderDetail(ctx, &order.HomestayOrderDetailReq{
		Sn: p.Sn,
	})
	if err != nil || resp.HomestayOrder == nil {
		return errors.Wrapf(ErrCloseOrderFal, "closeHomestayOrderStateMqHandler  get order fail or order no exists err:%v, sn:%s ,HomestayOrder : %+v", err, p.Sn, resp.HomestayOrder)
	}

	if resp.HomestayOrder.TradeState == model.HomestayOrderTradeStateWaitPay {
		_, err := l.svcCtx.OrderRpc.UpdateHomestayOrderTradeState(ctx, &order.UpdateHomestayOrderTradeStateReq{
			Sn:         p.Sn,
			TradeState: model.HomestayOrderTradeStateCancel,
		})
		if err != nil {
			return errors.Wrapf(ErrCloseOrderFal, "CloseHomestayOrderHandler close order fail  err:%v, sn:%s ", err, p.Sn)
		}
	}
	return nil
}
