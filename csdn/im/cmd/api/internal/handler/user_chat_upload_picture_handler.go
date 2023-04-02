package handler

import (
	"liujun/Time_go-zero_csdn/common/httpResp"
	"liujun/Time_go-zero_csdn/common/minIO"
	"liujun/Time_go-zero_csdn/csdn/im/cmd/api/internal/logic"
	"liujun/Time_go-zero_csdn/csdn/im/cmd/api/internal/svc"
	"liujun/Time_go-zero_csdn/csdn/im/cmd/api/internal/types"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func UserChatUploadPictureHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserChatUploadPictureRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}
		file_name, err := minIO.MinIOUpload(r, "picture")
		if err != nil {
			httpResp.HttpResp(w, r, nil, err)
			return
		}
		req.Picture = file_name
		l := logic.NewUserChatUploadPictureLogic(r.Context(), svcCtx)
		resp, err := l.UserChatUploadPicture(&req)
		httpResp.HttpResp(w, r, resp, err)
	}
}
