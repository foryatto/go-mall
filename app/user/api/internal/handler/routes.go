// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	user "mall/app/user/api/internal/handler/user"
	"mall/app/user/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/login",
				Handler: user.LoginHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/register",
				Handler: user.RegisterHandler(serverCtx),
			},
		},
		rest.WithPrefix("/v1/user"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/info",
				Handler: user.InfoHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/addAddress",
				Handler: user.AddAddressHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/editAddress",
				Handler: user.EditAddressHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/delAddress",
				Handler: user.DelAddressHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/getAddressList",
				Handler: user.UserAddressListHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/addCollection",
				Handler: user.UserCollectionAddHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/delCollection",
				Handler: user.UserCollectionDelHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/getCollectionList",
				Handler: user.UserCollectionListHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.JwtAuth.AccessSecret),
		rest.WithPrefix("/v1/user"),
	)
}
