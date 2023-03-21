package handler

import (
	"net/http"

	"liujun/Time_go-zero_csdn/common/httpResp"
	"liujun/Time_go-zero_csdn/csdn/user/cmd/api/internal/logic"
	"liujun/Time_go-zero_csdn/csdn/user/cmd/api/internal/svc"
	"liujun/Time_go-zero_csdn/csdn/user/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func UserFansListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserFansListRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewUserFansListLogic(r.Context(), svcCtx)
		resp, err := l.UserFansList(&req)
		httpResp.HttpResp(w, r, resp, err)
	}
}
