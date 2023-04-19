package handler

import (
	"liujun/Time_go-zero_csdn/common/httpResp"
	"net/http"

	"liujun/Time_go-zero_csdn/csdn/city/cmd/api/internal/logic"
	"liujun/Time_go-zero_csdn/csdn/city/cmd/api/internal/svc"
	"liujun/Time_go-zero_csdn/csdn/city/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func ChinaMapListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ChinaMapRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewChinaMapListLogic(r.Context(), svcCtx)
		resp, err := l.ChinaMapList(&req)
		httpResp.HttpResp(w, r, resp, err)
	}
}
