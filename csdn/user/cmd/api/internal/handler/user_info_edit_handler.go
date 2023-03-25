package handler

import (
	"liujun/Time_go-zero_csdn/common/httpResp"
	"liujun/Time_go-zero_csdn/common/minIO"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"liujun/Time_go-zero_csdn/csdn/user/cmd/api/internal/logic"
	"liujun/Time_go-zero_csdn/csdn/user/cmd/api/internal/svc"
	"liujun/Time_go-zero_csdn/csdn/user/cmd/api/internal/types"
)

func UserInfoEditHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserInfoEditRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}
		file_name, err := minIO.MinIOUpload(r, "head_photo")
		if err != nil {
			httpResp.HttpResp(w, r, nil, err)
			return
		}
		req.HeadPhoto = file_name
		l := logic.NewUserInfoEditLogic(r.Context(), svcCtx)
		resp, err := l.UserInfoEdit(&req)
		httpResp.HttpResp(w, r, resp, err)
	}
}
