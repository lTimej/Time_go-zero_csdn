package handler

import (
	"fmt"
	"liujun/Time_go-zero_csdn/common/httpResp"
	"net/http"

	"liujun/Time_go-zero_csdn/csdn/order/cmd/api/internal/logic"
	"liujun/Time_go-zero_csdn/csdn/order/cmd/api/internal/svc"
	"liujun/Time_go-zero_csdn/csdn/order/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func OrderCreateHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.OrderCreateRequest
		if err := httpx.Parse(r, &req); err != nil {
			fmt.Println(err, "哈哈哈哈哈哈哈")
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewOrderCreateLogic(r.Context(), svcCtx)
		resp, err := l.OrderCreate(&req)
		httpResp.HttpResp(w, r, resp, err)
	}
}
