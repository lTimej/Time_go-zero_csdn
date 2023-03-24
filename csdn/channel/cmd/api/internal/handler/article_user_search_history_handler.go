package handler

import (
	"liujun/Time_go-zero_csdn/common/httpResp"
	"net/http"

	"liujun/Time_go-zero_csdn/csdn/channel/cmd/api/internal/logic"
	"liujun/Time_go-zero_csdn/csdn/channel/cmd/api/internal/svc"
	"liujun/Time_go-zero_csdn/csdn/channel/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func ArticleUserSearchHistoryHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ArticleUserSearchHistoryRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewArticleUserSearchHistoryLogic(r.Context(), svcCtx)
		resp, err := l.ArticleUserSearchHistory(&req)
		httpResp.HttpResp(w, r, resp, err)
	}
}
