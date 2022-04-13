package rpc

import (
	"github.com/vmkteam/rpcgen/v2"
	"github.com/vmkteam/zenrpc/v2"
	"log"
	"news/pkg/db"
	"os"
	"time"
)

func NewNewsRPC(repo db.NewsRepo) (zenrpc.Server, *rpcgen.RPCGen) {
	news := &NewsService{
		Repo: repo,
	}

	rpc := zenrpc.NewServer(zenrpc.Options{ExposeSMD: true})
	rpc.Register("news", news)
	rpc.Use(zenrpc.Logger(log.New(os.Stderr, "", log.LstdFlags)))

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
