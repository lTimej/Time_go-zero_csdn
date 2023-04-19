package handler

import (
	"liujun/Time_go-zero_csdn/common/httpResp"
	"liujun/Time_go-zero_csdn/common/utils"
	"net/http"

	"liujun/Time_go-zero_csdn/csdn/city/cmd/api/internal/logic"
	"liujun/Time_go-zero_csdn/csdn/city/cmd/api/internal/svc"
	"liujun/Time_go-zero_csdn/csdn/city/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func CityListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CityRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}
		req.Pid = utils.StringToInt64(r.URL.Query().Get("pid"))
		l := logic.NewCityListLogic(r.Context(), svcCtx)
		resp, err := l.CityList(&req)
		httpResp.HttpResp(w, r, resp, err)
	}
}
