package handler

import (
	"fmt"
	"github.com/zeromicro/go-zero/rest/httpx"
	"liujun/Time_go-zero_csdn/common/httpResp"
	"liujun/Time_go-zero_csdn/csdn/channel/cmd/api/internal/logic"
	"liujun/Time_go-zero_csdn/csdn/channel/cmd/api/internal/svc"
	"liujun/Time_go-zero_csdn/csdn/channel/cmd/api/internal/types"
	"net/http"
	"strconv"
)

func AllArticleHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	fmt.Println(111112)
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("进来了.....................")
		var req types.AllArticleRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}
		channel_id, _ := strconv.Atoi(r.URL.Path[len("/v1/article/articles/"):])
		page, _ := strconv.Atoi(r.URL.Query().Get("page"))
		page_num, _ := strconv.Atoi(r.URL.Query().Get("page_num"))
		l := logic.NewAllArticleLogic(r.Context(), svcCtx)
		resp, err := l.AllArticle(&req, int64(channel_id), int32(page), int32(page_num))
		httpResp.HttpResp(w, r, resp, err)
	}
}
