package http

import (
	"context"
	"fmt"

	"github.com/ac-kurniawan/mychat/core"
	_ "github.com/ac-kurniawan/mychat/docs"
	"github.com/ac-kurniawan/mychat/library"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

//	@title			My Chat API
//	@version		1.0
//	@description	Service to manage chat by room
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	Ardhi
//	@contact.url	http://www.swagger.io/support
//	@contact.email	ac.kurniawan99@gmail.com

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

// @host		localhost:8080
type HttpServer struct {
	Port    string
	Service core.IMychatService
	Trace   library.AppTrace
	e       *echo.Echo
	handler HttpHandler
}

func (h *HttpServer) init() {
	h.e = echo.New()
	h.e.HideBanner = true
	h.handler = HttpHandler{
		Service: h.Service,
		Trace:   h.Trace,
	}
}

func (h *HttpServer) registerEndpoint() {
	h.e.GET("/swagger/*", echoSwagger.WrapHandler)

	h.e.GET("/room_chat", h.handler.GetRoomChatBySessionId)
	h.e.GET("/room_chats", h.handler.GetRoomChatsByParticipantId)
}

func (h *HttpServer) Start() {
	h.init()
	h.registerEndpoint()
	go func() {
		if err := h.e.Start(":" + h.Port); err != nil {
			fmt.Printf("[HTTP Server] - %s", err.Error())
		}
	}()
}

func (h *HttpServer) Stop(ctx context.Context) {
	if err := h.e.Shutdown(ctx); err != nil {
		fmt.Printf("[HTTP Server] - %s", err.Error())
	}
}

func NewHttpServer(module HttpServer) HttpServer {
	return module
}
