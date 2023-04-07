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

func (s *ProductServer) ProductSpuList(ctx context.Context, in *product.ProductSpuListRequest) (*product.ProductSpuListResponse, error) {
	l := logic.NewProductSpuListLogic(ctx, s.svcCtx)
	return l.ProductSpuList(in)
}

func (s *ProductServer) ProductCategory(ctx context.Context, in *product.ProductCategoryRequest) (*product.ProductCategoryResponse, error) {
	l := logic.NewProductCategoryLogic(ctx, s.svcCtx)
	return l.ProductCategory(in)
}

func (s *ProductServer) ProductDesc(ctx context.Context, in *product.ProductDescRequest) (*product.ProductDescResponse, error) {
	l := logic.NewProductDescLogic(ctx, s.svcCtx)
	return l.ProductDesc(in)
}

func (s *ProductServer) AddCart(ctx context.Context, in *product.AddCartRequest) (*product.AddCartResponse, error) {
	l := logic.NewAddCartLogic(ctx, s.svcCtx)
	return l.AddCart(in)
}

func (s *ProductServer) GetCart(ctx context.Context, in *product.GetCartRequest) (*product.GetCartResponse, error) {
	l := logic.NewGetCartLogic(ctx, s.svcCtx)
	return l.GetCart(in)
}
