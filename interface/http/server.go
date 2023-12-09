package http

import (
	"context"
	"fmt"

	"github.com/ac-kurniawan/mychat/core"
	"github.com/ac-kurniawan/mychat/library"
	"github.com/labstack/echo/v4"
)

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
	h.e.GET("/room_chat", h.handler.GetRoomChatBySessionId)
}

func (h *HttpServer) Start() {
	h.init()
	h.registerEndpoint()
	go func() {
		if err := h.e.Start(":1323"); err != nil {
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
