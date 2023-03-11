// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	"liujun/Time_go-zero_csdn/csdn/channel/cmd/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/v1/channel/articles/channel",
				Handler: AllChannelHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/v1/channel/default/channel",
				Handler: DefaultChannelHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/v1/article/articles/:channel_id",
				Handler: AllArticleHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/v1/article/status",
				Handler: ArticleStatusHandler(serverCtx),
			},
		},
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.SetUidToCtxMiddleware},
			[]rest.Route{
				{
					Method:  http.MethodGet,
					Path:    "/v1/channel/user/channel",
					Handler: UserChannelHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/v1/channel/user/channel",
					Handler: UserAddChannelHandler(serverCtx),
				},
				{
					Method:  http.MethodPatch,
					Path:    "/v1/channel/user/channel",
					Handler: UserPatchChannelHandler(serverCtx),
				},
			}...,
		),
		rest.WithJwt(serverCtx.Config.JwtAuth.AccessSecret),
	)
}
