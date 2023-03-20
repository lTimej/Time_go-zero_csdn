package handler

import (
	"liujun/Time_go-zero_csdn/common/httpResp"
	"liujun/Time_go-zero_csdn/common/utils"
	"net/http"

	"liujun/Time_go-zero_csdn/csdn/channel/cmd/api/internal/logic"
	"liujun/Time_go-zero_csdn/csdn/channel/cmd/api/internal/svc"
	"liujun/Time_go-zero_csdn/csdn/channel/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func ArticleCommentListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ArticleCommentListRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}
		ty := r.URL.Query().Get("type")
		article_id := utils.StringToInt64(r.URL.Query().Get("article_id"))
		offset := utils.StringToInt64(r.URL.Query().Get("offset"))
		limit := utils.StringToInt64(r.URL.Query().Get("limit"))
		req.Ty = ty
		req.ArticleId = article_id
		req.Offset = offset
		req.Limit = limit
		l := logic.NewArticleCommentListLogic(r.Context(), svcCtx)
		resp, err := l.ArticleCommentList(&req)
		httpResp.HttpResp(w, r, resp, err)
	}
}
