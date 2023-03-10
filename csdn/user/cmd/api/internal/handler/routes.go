// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	"liujun/Time_go-zero_csdn/csdn/user/cmd/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/v1/user/login",
				Handler: UserPasswordLoginHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/v1/user/login/auth",
				Handler: PhoneLoginHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/v1/user/login/smscode/:phone",
				Handler: SendSmsCodeHandler(serverCtx),
			},
		},
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.SetUidToCtxMiddleware},
			[]rest.Route{
				{
					Method:  http.MethodGet,
					Path:    "/v1/user/curr/user",
					Handler: UserCurrInfoHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/v1/user/isfocus",
					Handler: IsFocusUserHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/v1/user/focus",
					Handler: FocusUserHandler(serverCtx),
				},
				{
					Method:  http.MethodDelete,
					Path:    "/v1/user/focus",
					Handler: CancelFocusUserHandler(serverCtx),
				},
			}...,
		),
		rest.WithJwt(serverCtx.Config.JwtAuth.AccessSecret),
	)
}
