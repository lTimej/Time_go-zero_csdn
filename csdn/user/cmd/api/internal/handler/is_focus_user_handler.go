package handler

import (
	"fmt"
	"liujun/Time_go-zero_csdn/common/httpResp"
	"net/http"
	"strconv"

	"github.com/zeromicro/go-zero/rest/httpx"
	"liujun/Time_go-zero_csdn/csdn/user/cmd/api/internal/logic"
	"liujun/Time_go-zero_csdn/csdn/user/cmd/api/internal/svc"
	"liujun/Time_go-zero_csdn/csdn/user/cmd/api/internal/types"
)

func IsFocusUserHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.IsFocusUserRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		target_user_id, _ := strconv.Atoi(r.URL.Query().Get("target"))
		fmt.Println(target_user_id, "===================")
		req.TargetUserId = int64(target_user_id)
		l := logic.NewIsFocusUserLogic(r.Context(), svcCtx)
		resp, err := l.IsFocusUser(&req)
		httpResp.HttpResp(w, r, resp, err)
	}
}
