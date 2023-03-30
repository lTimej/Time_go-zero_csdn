package handler

import (
	"liujun/Time_go-zero_csdn/common/httpResp"
	"liujun/Time_go-zero_csdn/common/utils"
	"net/http"

	"liujun/Time_go-zero_csdn/csdn/im/cmd/api/internal/logic"
	"liujun/Time_go-zero_csdn/csdn/im/cmd/api/internal/svc"
	"liujun/Time_go-zero_csdn/csdn/im/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func UserChatRecordHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserChatRecordRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}
		target_user_id := r.URL.Query().Get("target_user_id")
		page := utils.StringToInt64(r.URL.Query().Get("page"))
		page_num := utils.StringToInt64(r.URL.Query().Get("page_num"))
		req.TargetUserId = target_user_id
		req.Page = page
		req.PageNum = page_num
		l := logic.NewUserChatRecordLogic(r.Context(), svcCtx)
		resp, err := l.UserChatRecord(&req)
		httpResp.HttpResp(w, r, resp, err)
	}
}
