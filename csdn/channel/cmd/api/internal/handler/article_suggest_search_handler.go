package handler

import (
	"net/http"

	"liujun/Time_go-zero_csdn/common/httpResp"
	"liujun/Time_go-zero_csdn/csdn/channel/cmd/api/internal/logic"
	"liujun/Time_go-zero_csdn/csdn/channel/cmd/api/internal/svc"
	"liujun/Time_go-zero_csdn/csdn/channel/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func ArticleSuggestSearchHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ArticleSuggestSearchRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}
		req.Keyword = r.URL.Query().Get("keyword")
		l := logic.NewArticleSuggestSearchLogic(r.Context(), svcCtx)
		resp, err := l.ArticleSuggestSearch(&req)
		httpResp.HttpResp(w, r, resp, err)
	}
}
