package middleware

import (
	"context"
	"liujun/Time_go-zero_csdn/common/ctxdata"
	"liujun/Time_go-zero_csdn/common/httpResp"
	"liujun/Time_go-zero_csdn/common/xerr"
	"liujun/Time_go-zero_csdn/csdn/user/cmd/api/internal/config"
	"net/http"
	"strings"
)

type AuthMiddleWare struct {
	Config config.Config
}

func NewAuthMiddleWare(c config.Config) *AuthMiddleWare {
	return &AuthMiddleWare{Config: c}
}

func (am *AuthMiddleWare) Handle(next http.HandlerFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		header := r.Header.Get("Authorization")
		if !strings.HasPrefix(header, "Bearer") {
			return
		}
		token := strings.Split(header, " ")[1]
		if token == "" {
			httpResp.HttpResp(w, r, nil, xerr.NewErrCodeMsg(xerr.SERVER_COMMON_ERROR, "未登录"))
			return
		}
		r.Header.Set("Authorization", token)
		claim, err := ctxdata.ParseToken(token, am.Config.JwtAuth.AccessSecret)
		if err != nil {
			httpResp.HttpResp(w, r, nil, xerr.NewErrCodeMsg(xerr.SERVER_COMMON_ERROR, "token认证失败"))
			return
		}
		ctx := context.WithValue(context.Background(), ctxdata.CtxKeyJwtUserId, claim.UserId)
		r = r.WithContext(ctx)
		next(w, r)
	}
}
