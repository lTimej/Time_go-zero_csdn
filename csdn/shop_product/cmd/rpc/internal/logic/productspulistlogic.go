package logic

import (
	"context"

	"github.com/jinzhu/copier"

	"liujun/Time_go-zero_csdn/csdn/shop_product/cmd/rpc/internal/svc"
	"liujun/Time_go-zero_csdn/csdn/shop_product/cmd/rpc/types/product"

	"github.com/zeromicro/go-zero/core/logx"
)

type ProductSpuListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewProductSpuListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProductSpuListLogic {
	return &ProductSpuListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ProductSpuListLogic) ProductSpuList(in *product.ProductSpuListRequest) (*product.ProductSpuListResponse, error) {
	// todo: add your logic here and delete this line
	builder := l.svcCtx.ProductSpuModel.Builder().Where("tb_sku.category1_id = ?", in.CategoryId)
	spus, err := l.svcCtx.ProductSpuModel.FindAllByCategoryId(l.ctx, builder)
	if err != nil {
		return nil, err
	}
	resp := new(product.ProductSpuListResponse)
	copier.Copy(resp.ProductSpus, spus)
	return resp, nil
}
