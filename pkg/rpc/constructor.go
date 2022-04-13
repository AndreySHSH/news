package rpc

import (
	"context"
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
		Repository: repository,
	}

	category := &CategoryService{
		Repository: repository,
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
	rpc.Register("category", category)
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

func newCategoryList(in []db.Category) []Category {
	var news []Category

	for _, v := range in {
		news = append(news, *newCategory(&v))
	}

	return news
}

func newCategory(in *db.Category) *Category {
	if in == nil {
		return nil
	}

	return &Category{
		ID:    in.ID,
		Title: in.Title,
	}
}

func (n *NewsService) getTags(ctx context.Context) error {
	tl, err := n.Repository.TagsByFilters(ctx, nil, db.PagerNoLimit)
	if err != nil {
		return err
	}

	for _, v := range tl {
		n.tags.Store(v.ID, Tag{
			ID:    v.ID,
			Title: v.Title,
		})
	}
	return nil
}

func (n *NewsService) newNewsList(in []db.News) []News {
	var news []News

	for _, v := range in {
		news = append(news, *n.newNews(&v))
	}

	return news
}

func (n *NewsService) newNews(in *db.News) *News {
	if in == nil {
		return nil
	}

	var tags []Tag

	for _, v := range in.TagIDs {

		if value, is := n.tags.Load(v); is != false {
			tags = append(tags, value.(Tag))
		}

	}

	return &News{
		ID:         in.ID,
		Title:      in.Title,
		Content:    in.Content,
		Tags:       tags,
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

func (in *CategorySearch) ToDB() *db.CategorySearch {
	if in == nil {
		return nil
	}

	return &db.CategorySearch{
		ID: in.CategoryID,
	}
}
