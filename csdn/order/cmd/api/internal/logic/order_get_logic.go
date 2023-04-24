package logic

import (
	"context"
	"liujun/Time_go-zero_csdn/common/ctxdata"
	"liujun/Time_go-zero_csdn/common/xerr"
	"liujun/Time_go-zero_csdn/csdn/order/cmd/rpc/orderclient"

	"github.com/jinzhu/copier"
	"github.com/pkg/errors"

	"liujun/Time_go-zero_csdn/csdn/order/cmd/api/internal/svc"
	"liujun/Time_go-zero_csdn/csdn/order/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type OrderGetLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewOrderGetLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OrderGetLogic {
	return &OrderGetLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *OrderGetLogic) OrderGet(req *types.OrderGetRequest) (resp *types.OrderGetResponse, err error) {
	// todo: add your logic here and delete this line
	user_id := ctxdata.GetUidFromCtx(l.ctx)
	data, err := l.svcCtx.OrderRpc.OrderGet(l.ctx, &orderclient.OrderGetRequest{UserId: user_id, PayStatus: req.PayStatus})
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrMsg("获取订单列表失败"), "req: %+v,rpc err:%+v", req, err)
	}
	resp = new(types.OrderGetResponse)
	copier.Copy(&resp, data)
	return
}
