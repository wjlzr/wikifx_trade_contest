// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"
	oauthuser "wikifx-trade-contest/api/internal/handler/oauth/oauth"
	sysuser "wikifx-trade-contest/api/internal/handler/sys/user"
	"wikifx-trade-contest/api/internal/svc"

	"github.com/tal-tech/go-zero/rest"
)

func RegisterHandlers(engine *rest.Server, serverCtx *svc.ServiceContext) {
	engine.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.Customintercept},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/api/v1/oauth/login",
					Handler: oauthuser.LoginHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/api/v1/oauth/send-code",
					Handler: oauthuser.SendCodeHandler(serverCtx),
				},
			}...,
		),
	)

	engine.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/api/v1/user/add",
				Handler: sysuser.UserAddHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
	)
}
