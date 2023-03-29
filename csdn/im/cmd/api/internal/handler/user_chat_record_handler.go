package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"liujun/Time_go-zero_csdn/csdn/im/cmd/api/internal/logic"
	"liujun/Time_go-zero_csdn/csdn/im/cmd/api/internal/svc"
	"liujun/Time_go-zero_csdn/csdn/im/cmd/api/internal/types"
)

func UserChatRecordHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserChatRecordRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewUserChatRecordLogic(r.Context(), svcCtx)
		resp, err := l.UserChatRecord(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
