package handler

import (
	"net/http"

	"liujun/Time_go-zero_csdn/common/httpResp"
	"liujun/Time_go-zero_csdn/csdn/user/cmd/api/internal/logic"
	"liujun/Time_go-zero_csdn/csdn/user/cmd/api/internal/svc"
	"liujun/Time_go-zero_csdn/csdn/user/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func UserPasswordLoginHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserPasswordLoginRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}
		l := logic.NewUserPasswordLoginLogic(r.Context(), svcCtx)
		resp, err := l.UserPasswordLogin(&req)
		httpResp.HttpResp(w, r, resp, err)
	}
}
