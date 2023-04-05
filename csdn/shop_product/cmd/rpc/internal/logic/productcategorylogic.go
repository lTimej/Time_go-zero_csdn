package logic

import (
	"context"
	"encoding/json"
	"fmt"
	"liujun/Time_go-zero_csdn/common/globalkey"

	"liujun/Time_go-zero_csdn/csdn/shop_product/cmd/rpc/internal/svc"
	"liujun/Time_go-zero_csdn/csdn/shop_product/cmd/rpc/types/product"

	"github.com/zeromicro/go-zero/core/logx"
)

type ProductCategoryLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewProductCategoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProductCategoryLogic {
	return &ProductCategoryLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ProductCategoryLogic) ProductCategory(in *product.ProductCategoryRequest) (*product.ProductCategoryResponse, error) {
	// todo: add your logic here and delete this line
	data := new(product.ProductCategoryResponse)
	category_key := globalkey.ProductCategory
	ok, _ := l.svcCtx.RedisClient.Exists(category_key)
	if ok {
		data_redis, _ := l.svcCtx.RedisClient.Get(category_key)
		json.Unmarshal([]byte(data_redis), data)
	} else {
		builder := l.svcCtx.ProductCategoryModel.Builder()
		parents, err := l.svcCtx.ProductCategoryModel.FindAllByParent(l.ctx, builder, 0)
		if err != nil {
			return nil, err
		}

		for _, parent := range parents {
			category := &product.CategoryList{}
			category.Id = parent.Id
			category.Name = parent.Name
			category.SubCategorys = []*product.SubCategory{}
			cats, err := l.svcCtx.ProductCategoryModel.FindAllByParent(l.ctx, builder, parent.Id)
			if err != nil {
				return nil, err
			}
			for _, cat := range cats {
				sku_obj, err := l.svcCtx.ProductSkuModel.FindOneByCategoryId(l.ctx, cat.Id)
				var default_image string
				if err != nil {
					fmt.Println(err, "888888888888888")
					return nil, err
				}
				if sku_obj == nil {
					default_image = ""
				} else {
					default_image = "http://172.20.16.20:9000/" + sku_obj.DefaultImage
				}
				fmt.Println(default_image, "999999999999999")
				sub_category := &product.SubCategory{}
				sub_category.Id = cat.Id
				sub_category.Name = cat.Name
				sub_category.DefaultImage = default_image
				category.SubCategorys = append(category.SubCategorys, sub_category)
			}
			data.Categorys = append(data.Categorys, category)
		}
		data_redis, _ := json.Marshal(data)
		l.svcCtx.RedisClient.Set(category_key, string(data_redis))
	}
	return data, nil
}
