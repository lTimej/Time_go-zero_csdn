package logic

import (
	"context"
	"fmt"

	// "database/sql"
	"liujun/Time_go-zero_csdn/csdn/shop_product/cmd/rpc/internal/svc"
	"liujun/Time_go-zero_csdn/csdn/shop_product/cmd/rpc/types/product"

	"github.com/Masterminds/squirrel"
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
	cate, err := l.svcCtx.ProductCategoryModel.FindOne(l.ctx, in.CategoryId)
	if err != nil {
		fmt.Println(err, "3333333")
		return nil, err
	}
	var builder squirrel.SelectBuilder
	if cate.ParentId.Int64 == 0 {
		builder = l.svcCtx.ProductSpuModel.Builder().Where("tb_spu.category1_id = ?", in.CategoryId)
	} else {
		builder = l.svcCtx.ProductSpuModel.Builder().Where("tb_spu.category2_id = ?", in.CategoryId)
	}
	spus, err := l.svcCtx.ProductSpuModel.FindAllByCategoryId(l.ctx, builder)
	if err != nil {
		fmt.Println(err, "22222222222222")
		return nil, err
	}
	resp := new(product.ProductSpuListResponse)
	for _, spu := range spus {
		resp.ProductSpus = append(resp.ProductSpus, &product.ProductSpuList{
			Name:         spu.Name,
			DefaultImage: "http://172.20.16.20:9000/" + spu.DefaultImage,
			Sales:        spu.Sales,
			Cfavs:        spu.Cfavs,
			SpuId:        spu.Id,
			Price:        spu.Price,
			NowPrice:     spu.NowPrice,
		})
	}
	return resp, nil
}
