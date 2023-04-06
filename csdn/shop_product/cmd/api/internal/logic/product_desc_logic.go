package logic

import (
	"context"
	"liujun/Time_go-zero_csdn/csdn/shop_product/cmd/rpc/productclient"

	"github.com/jinzhu/copier"

	"liujun/Time_go-zero_csdn/csdn/shop_product/cmd/api/internal/svc"
	"liujun/Time_go-zero_csdn/csdn/shop_product/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ProductDescLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewProductDescLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProductDescLogic {
	return &ProductDescLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ProductDescLogic) ProductDesc(req *types.ProductDescRequest) (resp *types.ProductDescResponse, err error) {
	// todo: add your logic here and delete this line
	sku_descs, err := l.svcCtx.ProductRpc.ProductDesc(l.ctx, &productclient.ProductDescRequest{SpuId: req.SpuId})
	if err != nil {
		return nil, err
	}
	resp = new(types.ProductDescResponse)
	copier.Copy(resp, sku_descs)
	return
}
