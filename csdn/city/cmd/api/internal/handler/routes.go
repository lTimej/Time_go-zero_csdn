// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	"liujun/Time_go-zero_csdn/csdn/city/cmd/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/v1/city/china/map",
				Handler: ChinaMapListHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/v1/city/city",
				Handler: CityListHandler(serverCtx),
			},
		},
	)
}