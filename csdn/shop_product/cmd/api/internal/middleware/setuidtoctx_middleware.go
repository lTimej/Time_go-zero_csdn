package middleware

import (
	"context"
	"liujun/Time_go-zero_csdn/common/ctxdata"
	"liujun/Time_go-zero_csdn/common/httpResp"
	"liujun/Time_go-zero_csdn/common/xerr"
	"liujun/Time_go-zero_csdn/csdn/shop_product/cmd/api/internal/config"
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
		header := r.Header.Get("Authorization")
		if !strings.HasPrefix(header, "Bearer") {
			return
		}
		token := strings.Split(header, " ")[1]
		if token == "" {
			if r.URL.Path == "/v1/article/status" {
				next(w, r)
				return
			} else {
				httpResp.HttpResp(w, r, nil, xerr.NewErrCodeMsg(xerr.OTHER_ERROR, "未登录"))
				return
			}
		}
		r.Header.Set("Authorization", token)
		claim, err := ctxdata.ParseToken(token, m.Config.JwtAuth.AccessSecret)
		if err != nil {
			httpResp.HttpResp(w, r, nil, xerr.NewErrCodeMsg(xerr.OTHER_ERROR, "token认证失败"))
			return
		}
		ctx := context.WithValue(context.Background(), ctxdata.CtxKeyJwtUserId, claim.UserId)
		r = r.WithContext(ctx)
		next(w, r)
	}
}
