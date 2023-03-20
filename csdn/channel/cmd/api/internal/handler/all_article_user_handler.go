package handler

import (
	"liujun/Time_go-zero_csdn/common/httpResp"
	"net/http"
	"strconv"

	"liujun/Time_go-zero_csdn/csdn/channel/cmd/api/internal/logic"
	"liujun/Time_go-zero_csdn/csdn/channel/cmd/api/internal/svc"
	"liujun/Time_go-zero_csdn/csdn/channel/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func AllArticleUserHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AllArticleUserRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}
		page, _ := strconv.Atoi(r.URL.Query().Get("page"))
		page_num, _ := strconv.Atoi(r.URL.Query().Get("page_num"))
		req.Page = int32(page)
		req.PageNum = int32(page_num)
		l := logic.NewAllArticleUserLogic(r.Context(), svcCtx)
		resp, err := l.AllArticleUser(&req)
		httpResp.HttpResp(w, r, resp, err)
	}
}
