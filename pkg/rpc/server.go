package rpc

import (
	"github.com/vmkteam/rpcgen/v2"
	middleware "github.com/vmkteam/zenrpc-middleware"
	"github.com/vmkteam/zenrpc/v2"
	"log"
	"net/http"
	"news/pkg/db"
	"os"
)

func NewServer(repo db.NewsRepo) (zenrpc.Server, *rpcgen.RPCGen) {
	allowDebug := func(param string) middleware.AllowDebugFunc {
		return func(req *http.Request) bool {
			return req.FormValue(param) == "true"
		}
	}

	isDevel := true
	eLog := log.New(os.Stderr, "E", log.LstdFlags|log.Lshortfile)
	dLog := log.New(os.Stdout, "D", log.LstdFlags|log.Lshortfile)

	rpc := zenrpc.NewServer(zenrpc.Options{
		ExposeSMD: true,
		AllowCORS: true,
	})
	rpc.RegisterAll(map[string]zenrpc.Invoker{
		"project":  NewNewsService(repo),
		"category": NewCategoryService(repo),
	})
	rpc.Use(
		middleware.WithDevel(isDevel),
		middleware.WithHeaders(),
		middleware.WithAPILogger(dLog.Printf, middleware.DefaultServerName),
		middleware.WithSentry(middleware.DefaultServerName),
		middleware.WithNoCancelContext(),
		middleware.WithMetrics(middleware.DefaultServerName),
		middleware.WithTiming(isDevel, allowDebug("d")),
		//middleware.WithSQLLogger(&dbc, isDevel, allowDebug("d"), allowDebug("s")),
		middleware.WithErrorLogger(eLog.Printf, middleware.DefaultServerName),
	)

	gen := rpcgen.FromSMD(rpc.SMD())

	return rpc, gen
}
