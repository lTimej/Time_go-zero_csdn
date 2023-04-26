package middleware

import (
	"context"
	"fmt"
	"liujun/Time_go-zero_csdn/common/ctxdata"
	"liujun/Time_go-zero_csdn/common/httpResp"
	"liujun/Time_go-zero_csdn/common/xerr"
	"liujun/Time_go-zero_csdn/csdn/im/cmd/api/internal/config"
	"net/http"
	"strings"
)

type SetUidToCtxMiddleware struct {
	Config config.Config
}

func NewSetUidToCtxMiddleware(c config.Config) *SetUidToCtxMiddleware {
	return &SetUidToCtxMiddleware{Config: c}
}

func (m *SetUidToCtxMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO generate middleware implement function, delete after code implementation
		if r.URL.Path == "/v1/im/user/chat" {
			fmt.Println(r.Header.Get("Authorization"), "33333333333333")
			r.Header.Set("Authorization", r.URL.Query().Get("token"))
			// next(w, r)
			// return
		}
		header := r.Header.Get("Authorization")
		if header == "" {
			if r.URL.Path == "/v1/article/status" {
				next(w, r)
				return
			} else if r.URL.Path == "/v1/im/user/chat" {
				next(w, r)
				return
			} else {
				httpResp.HttpResp(w, r, nil, xerr.NewErrCodeMsg(xerr.SERVER_COMMON_ERROR, "未登录"))
				return
			}
		}
		if !strings.HasPrefix(header, "Bearer") {
			return
		}
		token := strings.Split(header, " ")[1]
		r.Header.Set("Authorization", token)
		claim, err := ctxdata.ParseToken(token, m.Config.JwtAuth.AccessSecret)
		if err != nil {
			httpResp.HttpResp(w, r, nil, xerr.NewErrCodeMsg(xerr.SERVER_COMMON_ERROR, "token认证失败"))
			return
		}
		ctx := context.WithValue(context.Background(), ctxdata.CtxKeyJwtUserId, claim.UserId)
		r = r.WithContext(ctx)
		next(w, r)
	}
}
