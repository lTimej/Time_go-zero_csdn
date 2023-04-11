package logic

import (
	"context"
	"encoding/json"
	"fmt"
	"liujun/Time_go-zero_csdn/common/ctxdata"
	"liujun/Time_go-zero_csdn/common/globalkey"
	"liujun/Time_go-zero_csdn/common/utils"
	"liujun/Time_go-zero_csdn/csdn/shop_product/cmd/rpc/productclient"
	"net/http"
	"strings"

	"github.com/jinzhu/copier"

	"liujun/Time_go-zero_csdn/csdn/shop_product/cmd/api/internal/svc"
	"liujun/Time_go-zero_csdn/csdn/shop_product/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCartLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetCartLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCartLogic {
	return &GetCartLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetCartLogic) GetCart(req *types.GetCartRequest, r *http.Request) (resp *types.GetCartResponse, err error) {
	// todo: add your logic here and delete this line
	resp = new(types.GetCartResponse)
	user_id := ctxdata.GetUidFromCtx(l.ctx)
	name := globalkey.AnonymityUserCartList
	if user_id == "0" {
		carts := make(map[string]int64)
		rc, err := r.Cookie(name)
		if err != nil {
			fmt.Println(rc, "7777777777777777", err)
			return nil, err
		}
		if rc != nil {
			c := rc.Value
			cks := strings.Split(c, "-")
			for _, ck := range cks {
				sku_str := strings.Split(ck, ":")
				carts[sku_str[0]] = utils.StringToInt64(sku_str[1])
			}
		} else {
			return nil, nil
		}
		fmt.Println(carts, "666666666666666")
		for key, val := range carts {
			cart_obj := types.Carts{}
			builderByskuId := l.svcCtx.ProductSkuModel.BuilderBySkuId()
			fmt.Println(key, "******************777777777777777")
			var sku_ids []int64
			json.Unmarshal([]byte(key), &sku_ids)
			for _, sku_id := range sku_ids {
				cart_info, err := l.svcCtx.ProductSkuModel.FindOneSkuInfoBySkuId(l.ctx, builderByskuId, sku_id)
				if err != nil {
					fmt.Println(err, "2222222222")
					return nil, err
				}
				cart_obj.DefaultImage = cart_info.DefaultImage
				cart_obj.Price = cart_info.Price
				cart_obj.Title = cart_info.Title
				cart_obj.SpecLabel = append(cart_obj.SpecLabel, types.SpecLabel{
					Name:  cart_info.Name,
					Label: cart_info.Label,
				})
				cart_obj.Count = val
			}
			resp.Carts = append(resp.Carts, cart_obj)
		}
	} else {
		cart_info, err := l.svcCtx.ProductRpc.GetCart(l.ctx, &productclient.GetCartRequest{UserId: user_id})
		if err != nil {
			fmt.Println(err, "444444444444")
			return nil, err
		}
		copier.Copy(resp, cart_info)
	}
	return resp, nil
}
