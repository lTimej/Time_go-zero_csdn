// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	"liujun/Time_go-zero_csdn/csdn/shop_product/cmd/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/v1/shop/product/list",
				Handler: ProductListHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/v1/shop/product/category",
				Handler: ProductCategoryHandler(serverCtx),
			},
		},
	)
}
