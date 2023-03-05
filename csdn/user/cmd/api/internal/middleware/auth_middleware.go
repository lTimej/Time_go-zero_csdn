package middleware

import (
	"liujun/Time_go-zero_csdn/common/httpResp"
	"liujun/Time_go-zero_csdn/common/xerr"
	"liujun/Time_go-zero_csdn/csdn/user/cmd/api/internal/config"
	"net/http"
	"strings"
)

type AuthMiddleWare struct {
	config config.Config
}

func NewAuthMiddleWare() *AuthMiddleWare {
	return &AuthMiddleWare{}
}

func (am *AuthMiddleWare) Handle(next http.HandlerFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		header := r.Header.Get("Authorization")
		if !strings.HasPrefix(header, "Bearer") {
			return
		}
		token := strings.Split(header, " ")[1]
		if token == "" {
			httpResp.HttpResp(w, r, nil, xerr.NewErrCodeMsg(xerr.OTHER_ERROR, "未登录"))
			return
		}
		r.Header.Set("Authorization", token)
		next(w, r)
	}
}
