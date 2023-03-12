package handler

import (
	"liujun/Time_go-zero_csdn/common/httpResp"
	"net/http"
	"strconv"

	"github.com/zeromicro/go-zero/rest/httpx"
	"liujun/Time_go-zero_csdn/csdn/channel/cmd/api/internal/logic"
	"liujun/Time_go-zero_csdn/csdn/channel/cmd/api/internal/svc"
	"liujun/Time_go-zero_csdn/csdn/channel/cmd/api/internal/types"
)

func ArticleLikeHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ArticleLikeRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}
		aid, _ := strconv.ParseInt(r.URL.Query().Get("aid"), 10, 64)
		req.ArticleId = aid
		l := logic.NewArticleLikeLogic(r.Context(), svcCtx)
		resp, err := l.ArticleLike(&req)
		httpResp.HttpResp(w, r, resp, err)
	}
}
