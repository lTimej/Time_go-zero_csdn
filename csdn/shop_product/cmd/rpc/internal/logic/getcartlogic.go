package logic

import (
	"context"
	"encoding/json"
	"fmt"
	"liujun/Time_go-zero_csdn/common/utils"

	"liujun/Time_go-zero_csdn/common/globalkey"
	"liujun/Time_go-zero_csdn/csdn/shop_product/cmd/rpc/internal/svc"
	"liujun/Time_go-zero_csdn/csdn/shop_product/cmd/rpc/types/product"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCartLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetCartLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCartLogic {
	return &GetCartLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetCartLogic) GetCart(in *product.GetCartRequest) (*product.GetCartResponse, error) {
	// todo: add your logic here and delete this line
	key := fmt.Sprintf(globalkey.UserCartList, in.UserId)
	carts, err := l.svcCtx.RedisClient.Hgetall(key)
	if err != nil {
		fmt.Println(err, "1111111111")
		return nil, err
	}
	resp := new(product.GetCartResponse)
	for key, val := range carts {
		cart_obj := product.Carts{}
		var sku_ids []int64
		json.Unmarshal([]byte(key), &sku_ids)
		for _, sku_id := range sku_ids {
			builderByskuId := l.svcCtx.ProductSkuModel.BuilderBySkuId()
			cart_info, err := l.svcCtx.ProductSkuModel.FindOneSkuInfoBySkuId(l.ctx, builderByskuId, sku_id)
			if err != nil {
				fmt.Println(err, "2222222222")
				return nil, err
			}
			cart_obj.DefaultImage = cart_info.DefaultImage
			cart_obj.Price = cart_info.Price
			cart_obj.Title = cart_info.Title
			cart_obj.SpecLabel = append(cart_obj.SpecLabel, &product.SpecLabel{
				Name:  cart_info.Name,
				Label: cart_info.Label,
			})
			cart_obj.Count = utils.StringToInt64(val)
		}
		resp.Carts = append(resp.Carts, &cart_obj)
	}
	return resp, nil
}
