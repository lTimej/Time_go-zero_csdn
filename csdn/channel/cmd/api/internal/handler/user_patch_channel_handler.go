package handler

import (
	"liujun/Time_go-zero_csdn/common/httpResp"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"liujun/Time_go-zero_csdn/csdn/channel/cmd/api/internal/logic"
	"liujun/Time_go-zero_csdn/csdn/channel/cmd/api/internal/svc"
	"liujun/Time_go-zero_csdn/csdn/channel/cmd/api/internal/types"
)

func UserPatchChannelHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserPatchChannelRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewUserPatchChannelLogic(r.Context(), svcCtx)
		resp, err := l.UserPatchChannel(&req)
		httpResp.HttpResp(w, r, resp, err)
	}
}
