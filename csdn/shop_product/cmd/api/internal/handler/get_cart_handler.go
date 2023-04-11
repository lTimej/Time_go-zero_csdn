package handler

import (
	"liujun/Time_go-zero_csdn/common/httpResp"
	"liujun/Time_go-zero_csdn/csdn/shop_product/cmd/api/internal/logic"
	"liujun/Time_go-zero_csdn/csdn/shop_product/cmd/api/internal/svc"
	"liujun/Time_go-zero_csdn/csdn/shop_product/cmd/api/internal/types"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetCartHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetCartRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewGetCartLogic(r.Context(), svcCtx)
		resp, err := l.GetCart(&req, r)
		httpResp.HttpResp(w, r, resp, err)
	}
}
