package app

import (
	"github.com/ivahaev/go-logger"
	"github.com/vmkteam/rpcgen/v2"
	"github.com/vmkteam/zenrpc/v2"
	"net/http"
	"os"
)

func NewHTTP(rpc zenrpc.Server, gen rpcgen.RPCGen) {
	http.Handle("/v1/rpc/", rpc)
	http.Handle("/v1/rpc/doc/", http.HandlerFunc(zenrpc.SMDBoxHandler))
	http.Handle("/v1/rpc/news/client.go", http.HandlerFunc(rpcgen.Handler(gen.GoClient())))

	logger.Noticef("starting server on %s", os.Getenv("HTTP_ADDR"))

	panic(http.ListenAndServe(os.Getenv("HTTP_ADDR"), nil))
}
