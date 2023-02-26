package handler

import (
	"fmt"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"liujun/Time_go-zero_csdn/csdn/user/cmd/api/internal/logic"
	"liujun/Time_go-zero_csdn/csdn/user/cmd/api/internal/svc"
	"liujun/Time_go-zero_csdn/csdn/user/cmd/api/internal/types"
)

func UserPasswordLoginHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserPasswordLoginRequest
		fmt.Println(r, 1111111)
		if err := httpx.Parse(r, &req); err != nil {
			fmt.Println(err, 222222)
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}
		fmt.Println("+++++++++++++++++++++++++++++++++++++++++")
		l := logic.NewUserPasswordLoginLogic(r.Context(), svcCtx)
		resp, err := l.UserPasswordLogin(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			fmt.Println("进来了。。。。。。。。。。。。。。")
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
