package handler

import (
	"fmt"
	"net/http"

	"liujun/Time_go-zero_csdn/csdn/im/cmd/api/internal/logic"
	"liujun/Time_go-zero_csdn/csdn/im/cmd/api/internal/svc"
	"liujun/Time_go-zero_csdn/csdn/im/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func UserChatHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserChatRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}
		path := r.URL.Query().Get("token")
		fmt.Println(path, "11111")
		l := logic.NewUserChatLogic(r.Context(), svcCtx)
		l.UserChat(w, r)
	}
}
