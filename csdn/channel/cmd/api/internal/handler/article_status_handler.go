package handler

import (
	"fmt"
	"liujun/Time_go-zero_csdn/common/httpResp"
	"net/http"
	"strconv"

	"liujun/Time_go-zero_csdn/csdn/channel/cmd/api/internal/logic"
	"liujun/Time_go-zero_csdn/csdn/channel/cmd/api/internal/svc"
	"liujun/Time_go-zero_csdn/csdn/channel/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func ArticleStatusHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ArticleStatusRequest
		if err := httpx.Parse(r, &req); err != nil {
			fmt.Println(err, "呃呃呃呃呃呃呃呃呃呃呃呃呃呃呃呃呃呃")
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}
		aid, _ := strconv.ParseInt(r.URL.Query().Get("aid"), 10, 64)
		uid := r.URL.Query().Get("uid")
		req.ArticleId = aid
		req.UserId = uid
		l := logic.NewArticleStatusLogic(r.Context(), svcCtx)
		resp, err := l.ArticleStatus(&req)
		httpResp.HttpResp(w, r, resp, err)
	}
}
