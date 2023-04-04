package handler

import (
	"liujun/Time_go-zero_csdn/common/httpResp"
	"liujun/Time_go-zero_csdn/common/utils"
	"net/http"

	"liujun/Time_go-zero_csdn/csdn/shop_product/cmd/api/internal/logic"
	"liujun/Time_go-zero_csdn/csdn/shop_product/cmd/api/internal/svc"
	"liujun/Time_go-zero_csdn/csdn/shop_product/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func ProducSputListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ProductSpuListRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}
		req.CategoryId = utils.StringToInt64(r.URL.Query().Get("category_id"))
		l := logic.NewProducSputListLogic(r.Context(), svcCtx)
		resp, err := l.ProducSputList(&req)
		httpResp.HttpResp(w, r, resp, err)
	}
}
