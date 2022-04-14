package app

import (
	"github.com/ivahaev/go-logger"
	"github.com/labstack/echo/v4"
	"github.com/vmkteam/rpcgen/v2"
	middleware "github.com/vmkteam/zenrpc-middleware"
	"github.com/vmkteam/zenrpc/v2"
	"net/http"
)

func NewHTTP(rpc zenrpc.Server, gen rpcgen.RPCGen, addr string) error {
	e := echo.New()

	e.Any("/v1/project/", middleware.EchoHandler(rpc))
	e.Any("/v1/rpc/project/doc/", echo.WrapHandler(http.HandlerFunc(zenrpc.SMDBoxHandler)))
	e.Any("/v1/project/client.go", echo.WrapHandler(http.HandlerFunc(rpcgen.Handler(gen.GoClient()))))

	logger.Noticef("starting server on %s", addr)

	return e.Start(addr)
}
