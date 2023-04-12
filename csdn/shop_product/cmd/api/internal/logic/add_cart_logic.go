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
	"time"

	"liujun/Time_go-zero_csdn/csdn/shop_product/cmd/api/internal/svc"
	"liujun/Time_go-zero_csdn/csdn/shop_product/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddCartLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddCartLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddCartLogic {
	return &AddCartLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddCartLogic) AddCart(req *types.AddCartRequest, w http.ResponseWriter, r *http.Request) (resp *types.AddCartResponse, err error) {
	// todo: add your logic here and delete this line
	user_id := ctxdata.GetUidFromCtx(l.ctx)
	if user_id == "0" {
		name := globalkey.AnonymityUserCartList
		skus := make(map[string]int64)
		rc, _ := r.Cookie(name)
		if rc != nil {
			c := rc.Value
			cks := strings.Split(c, "-")
			for _, ck := range cks {
				sku_str := strings.Split(ck, ":")
				skus[sku_str[0]] = utils.StringToInt64(sku_str[1])
			}
		}
		sku_ids := req.SkuId
		data, _ := json.Marshal(sku_ids)
		sku := string(data)
		// sku := utils.Int64ToString(req.SkuId)
		if _, ok := skus[sku]; ok {
			skus[sku] += req.Count
		} else {
			skus[sku] = req.Count
		}
		var value string
		for key, val := range skus {
			value += key + ":" + utils.Int64ToString(val) + "-"
		}
		value = value[:len(value)-1]
		expire := time.Now().Add(time.Duration(l.svcCtx.Config.JwtAuth.AccessExpire) * time.Second)
		cookie := http.Cookie{Name: name, Value: value, Expires: expire}
		http.SetCookie(w, &cookie)
		return
	}
	fmt.Println(user_id, "99999999999999999")
	_, err = l.svcCtx.ProductRpc.AddCart(l.ctx, &productclient.AddCartRequest{UserId: user_id, SkuId: req.SkuId, Count: req.Count})
	if err != nil {
		return nil, err
	}
	return
}
