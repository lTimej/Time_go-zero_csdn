package handler

import (
	"liujun/Time_go-zero_csdn/common/httpResp"
	"liujun/Time_go-zero_csdn/common/utils"
	"net/http"

	"liujun/Time_go-zero_csdn/csdn/order/cmd/api/internal/logic"
	"liujun/Time_go-zero_csdn/csdn/order/cmd/api/internal/svc"
	"liujun/Time_go-zero_csdn/csdn/order/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func OrderGetHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.OrderGetRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}
		req.PayStatus = utils.StringToInt64(r.URL.Query().Get("pay_status"))
		l := logic.NewOrderGetLogic(r.Context(), svcCtx)
		resp, err := l.OrderGet(&req)
		httpResp.HttpResp(w, r, resp, err)
	}
}
