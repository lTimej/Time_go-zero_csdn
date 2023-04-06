package logic

import (
	"context"
	"fmt"

	"liujun/Time_go-zero_csdn/csdn/shop_product/cmd/rpc/internal/svc"
	"liujun/Time_go-zero_csdn/csdn/shop_product/cmd/rpc/types/product"

	"github.com/zeromicro/go-zero/core/logx"
)

type ProductDescLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewProductDescLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProductDescLogic {
	return &ProductDescLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ProductDescLogic) ProductDesc(in *product.ProductDescRequest) (*product.ProductDescResponse, error) {
	// todo: add your logic here and delete this line
	resp := new(product.ProductDescResponse)
	builder_spu_specification := l.svcCtx.ProductSpuSpecificationModel.Builder().Where("spu_id = ?", in.SpuId)
	spu_specificetions, err := l.svcCtx.ProductSpuSpecificationModel.FindBySpuId(l.ctx, builder_spu_specification)
	if err != nil {
		fmt.Println(err, "11111111111111111")
		return nil, err
	}
	var label string
	for _, item := range spu_specificetions {
		label += item.Name
	}
	builder := l.svcCtx.ProductSkuModel.Builder().Where("spu_id = ?", in.SpuId)
	sku_base_infos, err := l.svcCtx.ProductSkuModel.FindAllSkuBasicInfoBySpuId(l.ctx, builder)
	if err != nil {
		fmt.Println(err, "2222222222")
		return nil, err
	}
	resp.Title = sku_base_infos[0].Name
	resp.Price = sku_base_infos[0].Price

	resp.NowPrice = sku_base_infos[0].NowPrice
	resp.SkuSpec = new(product.SkuSpec)
	resp.SkuSpec.Stock = sku_base_infos[0].Stock
	resp.SkuSpec.Label = label
	for _, item := range sku_base_infos {
		resp.SwiperImages = append(resp.SwiperImages, item.DefaultImage)
	}
	builder_spec := l.svcCtx.ProductSkuModel.BuilderSpec()
	sku_spec_infos, err := l.svcCtx.ProductSkuModel.FindAllSkuSpecBySpuId(l.ctx, builder_spec, in.SpuId)
	if err != nil {
		fmt.Println(err, "3333333333333")
		return nil, err
	}
	check_labl_name := make(map[string]int)
	for _, item := range sku_spec_infos {
		label_name := item.Label
		spec_list := new(product.SpecList)
		if _, ok := check_labl_name[label_name]; !ok {
			check_labl_name[label_name] += 1
			spec_list.LabelName = label_name
			spec_list.SpecId = item.SpecId
		} else {
			continue
		}
		for _, item := range sku_spec_infos {
			if item.Label == label_name {
				spec_list.Specs = append(spec_list.Specs, &product.Specs{
					SkuId:  item.SkuId,
					SkuImg: item.DefaultImage,
					Name:   item.Name,
				})
			}
		}
		resp.SkuSpec.SpecList = append(resp.SkuSpec.SpecList, spec_list)
	}
	resp.SpuDesc = new(product.SpuDesc)
	builder_spu_desc := l.svcCtx.ProductSpuDescModel.Builder().Where("spu_id = ?", in.SpuId)
	descs, err := l.svcCtx.ProductSpuDescModel.FindAllBySpuId(l.ctx, builder_spu_desc)
	if err != nil {
		fmt.Println(err, "555555555555")
		return nil, err
	}
	resp.SpuDesc.DescInfo = descs[0].DetailInfo
	for _, desc := range descs {
		resp.SpuDesc.DescImage = append(resp.SpuDesc.DescImage, desc.DescImage)
	}
	return resp, nil
}
