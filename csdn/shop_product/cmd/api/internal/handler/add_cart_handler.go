package handler

import (
	"fmt"
	"liujun/Time_go-zero_csdn/common/httpResp"
	"net/http"

	"liujun/Time_go-zero_csdn/csdn/shop_product/cmd/api/internal/logic"
	"liujun/Time_go-zero_csdn/csdn/shop_product/cmd/api/internal/svc"
	"liujun/Time_go-zero_csdn/csdn/shop_product/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func AddCartHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AddCartRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}
		fmt.Println(r, "==============")
		l := logic.NewAddCartLogic(r.Context(), svcCtx)
		resp, err := l.AddCart(&req, w, r)
		httpResp.HttpResp(w, r, resp, err)
	}
}
