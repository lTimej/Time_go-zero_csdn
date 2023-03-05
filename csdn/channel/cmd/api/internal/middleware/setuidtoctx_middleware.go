package middleware

import (
	"liujun/Time_go-zero_csdn/common/httpResp"
	"liujun/Time_go-zero_csdn/common/xerr"
	"net/http"
	"strings"
)

type SetUidToCtxMiddleware struct {
}

func NewSetUidToCtxMiddleware() *SetUidToCtxMiddleware {
	return &SetUidToCtxMiddleware{}
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
			httpResp.HttpResp(w, r, nil, xerr.NewErrCodeMsg(xerr.OTHER_ERROR, "未登录"))
			return
		}
		r.Header.Set("Authorization", token)
		// Passthrough to next handler if need
		next(w, r)
	}
}
