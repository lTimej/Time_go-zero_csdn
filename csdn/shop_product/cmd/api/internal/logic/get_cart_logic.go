package logic

import (
	"context"
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
	sepc_structs := []*productclient.SpecStruct{}
	// if user_id == "0" {
	carts := make(map[string]int64)
	rc, err := r.Cookie(name)
	if err != nil {
		fmt.Println(rc, "7777777777777777", err)
	}
	if rc != nil {
		c := rc.Value
		cks := strings.Split(c, "-")

		for _, ck := range cks {
			sku_str := strings.Split(ck, ":")
			carts[sku_str[0]] = utils.StringToInt64(sku_str[1])
			sepc_structs = append(sepc_structs, &productclient.SpecStruct{
				Name:  sku_str[0],
				Count: sku_str[1],
			})
		}
	}
	cart_info, err := l.svcCtx.ProductRpc.GetCart(l.ctx, &productclient.GetCartRequest{UserId: user_id, SpecStructs: sepc_structs})
	if err != nil {
		fmt.Println(err, "444444444444")
		return nil, err
	}
	copier.Copy(resp, cart_info)
	return resp, nil
}
