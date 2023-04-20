package logic

import (
	"context"
	"liujun/Time_go-zero_csdn/common/ctxdata"
	"liujun/Time_go-zero_csdn/csdn/order/cmd/rpc/orderclient"

	"github.com/jinzhu/copier"

	"liujun/Time_go-zero_csdn/csdn/order/cmd/api/internal/svc"
	"liujun/Time_go-zero_csdn/csdn/order/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type OrderCreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewOrderCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OrderCreateLogic {
	return &OrderCreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *OrderCreateLogic) OrderCreate(req *types.OrderCreateRequest) (resp *types.OrderCreateResponse, err error) {
	// todo: add your logic here and delete this line
	user_id := ctxdata.GetUidFromCtx(l.ctx)
	var sku []*orderclient.Sku
	copier.Copy(&sku, req.Sku)
	data, err := l.svcCtx.OrderRpc.OrderCreate(l.ctx, &orderclient.OrderCreateRequest{
		UserId:     user_id,
		AddressId:  req.AddressId,
		TotalCount: req.TotalCount,
		TotalPrice: req.TotalPrice,
		Freight:    req.Freight,
		Version:    req.Version,
		Sn:         req.Sn,
		Sku:        sku,
	})
	if err != nil {
		return nil, err
	}
	resp = new(types.OrderCreateResponse)
	resp.Sn = data.Sn
	return
}
