package handler

import (
	"github.com/zeromicro/go-zero/rest/httpx"
	"liujun/Time_go-zero_csdn/common/httpResp"
	"liujun/Time_go-zero_csdn/csdn/user/cmd/api/internal/logic"
	"liujun/Time_go-zero_csdn/csdn/user/cmd/api/internal/svc"
	"liujun/Time_go-zero_csdn/csdn/user/cmd/api/internal/types"
	"net/http"
)

func IsFocusUserHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.IsFocusUserRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}
		req.TargetUserId = r.URL.Query().Get("target")
		l := logic.NewIsFocusUserLogic(r.Context(), svcCtx)
		resp, err := l.IsFocusUser(&req)
		httpResp.HttpResp(w, r, resp, err)
	}
}
