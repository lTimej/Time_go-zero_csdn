package logic

import (
	"context"
	"liujun/Time_go-zero_csdn/csdn/shop_product/cmd/rpc/productclient"

	"github.com/jinzhu/copier"

	"liujun/Time_go-zero_csdn/csdn/shop_product/cmd/api/internal/svc"
	"liujun/Time_go-zero_csdn/csdn/shop_product/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ProductCategoryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewProductCategoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProductCategoryLogic {
	return &ProductCategoryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ProductCategoryLogic) ProductCategory(req *types.ProductCategoryRequest) (resp *types.ProductCategoryResponse, err error) {
	// todo: add your logic here and delete this line
	cats, err := l.svcCtx.ProductRpc.ProductCategory(l.ctx, &productclient.ProductCategoryRequest{})
	if err != nil {
		return
	}

	resp = new(types.ProductCategoryResponse)
	copier.Copy(resp, cats)
	return
}
