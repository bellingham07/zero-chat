// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	chat "zero-chat/chat/api/internal/handler/chat"
	"zero-chat/chat/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.Authority},
			[]rest.Route{
				{
					Method:  http.MethodGet,
					Path:    "/chat",
					Handler: chat.ChatHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/send",
					Handler: chat.SendMsgHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/all",
					Handler: chat.GetAllChatHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/detail",
					Handler: chat.GetChatDetailHandler(serverCtx),
				},
			}...,
		),
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
		rest.WithPrefix("/chat"),
	)
}
