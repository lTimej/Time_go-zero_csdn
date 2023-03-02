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

func NewAuthMiddleWare(c config.Config) *AuthMiddleWare {
	return &AuthMiddleWare{
		config: c,
	}
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
		// user_id, err := utils.VerifyToken(token, am.config.JwtAuth.AccessSecret)
		// if err != nil {
		// 	httpResp.HttpResp(w, r, nil, xerr.NewErrCodeMsg(xerr.OTHER_ERROR, "非法token"))
		// 	return
		// }
		// _ = context.WithValue(r.Context(), ctxdata.CtxKeyJwtUserId, user_id)
		next(w, r)
	}
}
