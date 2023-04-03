// Code generated by goctl. DO NOT EDIT.
// Source: product.proto

package server

import (
	"context"

	"liujun/Time_go-zero_csdn/csdn/shop_product/cmd/rpc/internal/logic"
	"liujun/Time_go-zero_csdn/csdn/shop_product/cmd/rpc/internal/svc"
	"liujun/Time_go-zero_csdn/csdn/shop_product/cmd/rpc/types/product"
)

type ProductServer struct {
	svcCtx *svc.ServiceContext
	product.UnimplementedProductServer
}

func NewProductServer(svcCtx *svc.ServiceContext) *ProductServer {
	return &ProductServer{
		svcCtx: svcCtx,
	}
}

func (s *ProductServer) ProductList(ctx context.Context, in *product.ProductListRequest) (*product.ProductListResponse, error) {
	l := logic.NewProductListLogic(ctx, s.svcCtx)
	return l.ProductList(in)
}
