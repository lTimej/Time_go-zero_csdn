package logic

import (
	"context"
	"liujun/Time_go-zero_csdn/common/xerr"
	"liujun/Time_go-zero_csdn/csdn/order/cmd/rpc/internal/svc"
	"liujun/Time_go-zero_csdn/csdn/order/cmd/rpc/types/order"
	"liujun/Time_go-zero_csdn/csdn/order/model"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type OrderUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewOrderUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OrderUpdateLogic {
	return &OrderUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *OrderUpdateLogic) OrderUpdate(in *order.OrderUpdateRequest) (*order.OrderUpdateResponse, error) {
	// todo: add your logic here and delete this line
	order_desc, err := l.svcCtx.OrderModel.FindOneBySn(l.ctx, in.Sn)
	if err != nil && err != model.ErrNotFound {
		return nil, errors.Wrapf(xerr.ErrDBError, "err:%v", err)
	}
	if order_desc == nil {
		return nil, errors.Wrapf(xerr.NewErrMsg("订单不存在"), "order no exists  in : %+v", in)
	}

	if order_desc.PayStatus == in.PayStatus {
		return nil, nil
	}
	// 2、Verify order status
	if err := l.verifyOrderTradeState(in.PayStatus, order_desc.PayStatus); err != nil {
		return nil, errors.WithMessagef(err, " , in : %+v", in)
	}
	// 3、Pre-update status judgment.
	order_desc.PayStatus = in.PayStatus
	if err := l.svcCtx.OrderModel.Update(l.ctx, order_desc); err != nil {
		return nil, errors.Wrapf(xerr.NewErrMsg("Failed to update homestay order status"), "Failed to update homestay order status db UpdateWithVersion err:%v , in : %v", err, in)
	}

	//4、notify user
	// if in.PayStatus == model.OrderTradeStateWaitUse {
	// 	payload, err := json.Marshal(jobtype.PaySuccessNotifyUserPayload{Order: homestayOrder})
	// 	if err != nil {
	// 		logx.WithContext(l.ctx).Errorf("pay success notify user task json Marshal fail, err :%+v , sn : %s",err,homestayOrder.Sn)
	// 	}else{
	// 		_, err := l.svcCtx.AsynqClient.Enqueue(asynq.NewTask(jobtype.MsgPaySuccessNotifyUser, payload))
	// 		if err != nil {
	// 			logx.WithContext(l.ctx).Errorf("pay success notify user  insert queue fail err :%+v , sn : %s",err,homestayOrder.Sn)
	// 		}
	// 	}
	// }
	return &order.OrderUpdateResponse{Sn: in.Sn}, nil
}

func (l *OrderUpdateLogic) verifyOrderTradeState(newTradeState, oldTradeState int64) error {
	if newTradeState == model.OrderTradeStateWaitPay {
		return errors.Wrapf(xerr.NewErrMsg("Changing this status is not supported"),
			"Changing this status is not supported newTradeState: %d, oldTradeState: %d",
			newTradeState,
			oldTradeState)
	}

	if newTradeState == model.OrderTradeStateCancel {
		if oldTradeState != model.OrderTradeStateWaitPay {
			return errors.Wrapf(xerr.NewErrMsg("只有待支付的订单才能被取消"),
				"Only orders pending payment can be cancelled newTradeState: %d, oldTradeState: %d",
				newTradeState,
				oldTradeState)
		}

	} else if newTradeState == model.OrderTradeStateWaitUse {
		if oldTradeState != model.OrderTradeStateWaitPay {
			return errors.Wrapf(xerr.NewErrMsg("Only orders pending payment can change this status"),
				"Only orders pending payment can change this status newTradeState: %d, oldTradeState: %d",
				newTradeState,
				oldTradeState)
		}
	} else if newTradeState == model.OrderTradeStateUsed {
		if oldTradeState != model.OrderTradeStateWaitUse {
			return errors.Wrapf(xerr.NewErrMsg("Only unused orders can be changed to this status"),
				"Only unused orders can be changed to this status newTradeState: %d, oldTradeState: %d",
				newTradeState,
				oldTradeState)
		}
	} else if newTradeState == model.OrderTradeStateExpire {
		if oldTradeState != model.OrderTradeStateWaitUse {
			return errors.Wrapf(xerr.NewErrMsg("Only unused orders can be changed to this status"),
				"Only unused orders can be changed to this status newTradeState: %d, oldTradeState: %d",
				newTradeState,
				oldTradeState)
		}
	}

	return nil
}
