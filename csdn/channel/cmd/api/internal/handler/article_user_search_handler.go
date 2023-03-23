package handler

import (
	"liujun/Time_go-zero_csdn/common/httpResp"
	"liujun/Time_go-zero_csdn/common/utils"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"liujun/Time_go-zero_csdn/csdn/channel/cmd/api/internal/logic"
	"liujun/Time_go-zero_csdn/csdn/channel/cmd/api/internal/svc"
	"liujun/Time_go-zero_csdn/csdn/channel/cmd/api/internal/types"
)

func ArticleUserSearchHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ArticleUserSearchRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}
		req.Page = utils.StringToInt64(r.URL.Query().Get("page"))
		req.PageNum = utils.StringToInt64(r.URL.Query().Get("page_num"))
		req.Keyword = r.URL.Query().Get("keyword")
		l := logic.NewArticleUserSearchLogic(r.Context(), svcCtx)
		resp, err := l.ArticleUserSearch(&req)
		httpResp.HttpResp(w, r, resp, err)
	}
}
