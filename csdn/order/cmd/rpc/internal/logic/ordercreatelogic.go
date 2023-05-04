package logic

import (
	"context"
	"encoding/json"
	"liujun/Time_go-zero_csdn/common/order_var"
	"liujun/Time_go-zero_csdn/common/utils"
	"liujun/Time_go-zero_csdn/common/xerr"
	"liujun/Time_go-zero_csdn/csdn/order/cmd/rpc/internal/svc"
	"liujun/Time_go-zero_csdn/csdn/order/cmd/rpc/types/order"
	"liujun/Time_go-zero_csdn/csdn/order/model"
	"time"

	"liujun/Time_go-zero_csdn/csdn/scheduler/cmd/job/jobtype"

	"github.com/hibiken/asynq"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

const CloseOrderTimeMinutes = 30 //defer close order time
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
		return nil, errors.Wrapf(xerr.ErrDBError, "err:%v", err)
	}
	o, err := l.svcCtx.OrderModel.FindOneBySn(l.ctx, sn)
	if err != nil && err != model.ErrNotFound {
		return nil, errors.Wrapf(xerr.ErrDBError, "err:%v", err)
	}
	if o == nil {
		return nil, errors.Wrapf(xerr.NewErrMsg("订单不存在"), "err:%v", err)
	}
	order_id := o.Id
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
			return nil, errors.Wrapf(xerr.ErrDBError, "err:%v", err)
		}
	}

	//2、Delayed closing of order tasks.
	payload, err := json.Marshal(jobtype.DeferCloseProductOrderPayload{Sn: sn})
	if err != nil {
		logx.WithContext(l.ctx).Errorf("create defer close order task json Marshal fail err :%+v , sn : %s", err, sn)
	} else {
		_, err = l.svcCtx.AsynqClient.Enqueue(asynq.NewTask(jobtype.DeferCloseProductOrder, payload), asynq.ProcessIn(CloseOrderTimeMinutes*time.Minute))
		if err != nil {
			logx.WithContext(l.ctx).Errorf("create defer close order task insert queue fail err :%+v , sn : %s", err, sn)
		}
	}
	return &order.OrderCreateResponse{
		Sn: sn,
	}, nil
}
