package handler

import (
	"liujun/Time_go-zero_csdn/common/httpResp"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"liujun/Time_go-zero_csdn/csdn/channel/cmd/api/internal/logic"
	"liujun/Time_go-zero_csdn/csdn/channel/cmd/api/internal/svc"
	"liujun/Time_go-zero_csdn/csdn/channel/cmd/api/internal/types"
)

func ArticleStatusHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ArticleStatusRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}
		aid := r.URL.Query().Get("aid")
		uid := r.URL.Query().Get("uid")
		req.ArticleId = aid
		req.UserId = uid
		l := logic.NewArticleStatusLogic(r.Context(), svcCtx)
		resp, err := l.ArticleStatus(&req)
		httpResp.HttpResp(w, r, resp, err)
	}
}
