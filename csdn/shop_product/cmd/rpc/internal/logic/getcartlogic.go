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
	key := ""
	if in.UserId == "0" { //未登录
		key = globalkey.AnonymityUserCartList
	} else {
		key = fmt.Sprintf(globalkey.UserCartList, in.UserId)
	}
	carts, err := l.svcCtx.RedisClient.Hgetall(key)
	if err != nil {
		return nil, err
	}
	resp := new(product.GetCartResponse)
	for key, val := range carts {
		fmt.Println(key, val, "hahhah")
		var spus []int64
		json.Unmarshal([]byte(key), &spus)
		for _, spu := range spus {
			builderByskuId := l.svcCtx.ProductSkuModel.BuilderBySkuId()
			cart_info, err := l.svcCtx.ProductSkuModel.FindOneSkuInfoBySkuId(l.ctx, builderByskuId, spu)
			if err != nil {
				return nil, err
			}
			resp.DefaultImage = cart_info.DefaultImage
			resp.Price = cart_info.Price
			resp.Title = cart_info.Title
			resp.SpecLabel = append(resp.SpecLabel, &product.SpecLabel{
				Name:  cart_info.Name,
				Label: cart_info.Label,
			})
		}
		resp.Count = utils.StringToInt64(val)
	}
	return resp, nil
}
