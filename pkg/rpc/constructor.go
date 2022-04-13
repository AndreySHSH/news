package rpc

import (
	"github.com/go-pg/pg/v10"
	"github.com/vmkteam/rpcgen/v2"
	"github.com/vmkteam/zenrpc-middleware"
	"github.com/vmkteam/zenrpc/v2"
	"log"
	"net/http"
	"news/pkg/db"
	"os"
	"time"
)

func NewNewsRPC(dbc pg.DB) (zenrpc.Server, *rpcgen.RPCGen) {
	dbLayer := db.New(&dbc)
	repository := db.NewNewsRepo(dbLayer)

	news := &NewsService{
		Repo: repository,
	}

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
	rpc.Register("news", news)
	rpc.Use(
		middleware.WithDevel(isDevel),
		middleware.WithHeaders(),
		middleware.WithAPILogger(dLog.Printf, middleware.DefaultServerName),
		middleware.WithSentry(middleware.DefaultServerName),
		middleware.WithNoCancelContext(),
		middleware.WithMetrics(middleware.DefaultServerName),
		middleware.WithTiming(isDevel, allowDebug("d")),
		middleware.WithSQLLogger(&dbc, isDevel, allowDebug("d"), allowDebug("s")),
		middleware.WithErrorLogger(eLog.Printf, middleware.DefaultServerName),
	)

	gen := rpcgen.FromSMD(rpc.SMD())

	return rpc, gen
}

func newNewsList(in []db.News) []News {
	var news []News

	for _, v := range in {
		news = append(news, *newNews(&v))
	}

	return news
}

func newNews(in *db.News) *News {
	if in == nil {
		return nil
	}

	return &News{
		ID:         in.ID,
		Title:      in.Title,
		Content:    in.Content,
		TagIDs:     in.TagIDs,
		CategoryID: in.CategoryID,
		CreatedAt:  in.CreatedAt.Format(time.RFC822),
		Category: Category{
			ID:    in.Category.ID,
			Title: in.Category.Title,
		},
	}
}

func (in *NewsSearch) ToDB() *db.NewsSearch {
	if in == nil {
		return nil
	}

	return &db.NewsSearch{
		CategoryID: in.CategoryID,
		TagID:      in.TagID,
	}
}
