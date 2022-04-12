//go:generate zenrpc
package rpc

import (
	"context"
	"github.com/ivahaev/go-logger"
	"log"
	"net/http"
	"news/pkg/db"
	"os"
	"time"

	"github.com/vmkteam/zenrpc/v2"
)

type Category struct {
	ID    int64  `json:"id"`
	Title string `json:"title"`
}

type News struct {
	ID         int64    `json:"id"`
	Title      string   `json:"title"`
	Content    string   `json:"content"`
	TagIDs     []int64  `json:"tagIDs"`
	CategoryID int64    `json:"categoryID"`
	CreatedAt  string   `json:"createdAt"`
	Category   Category `json:"category"`
}

type NewsService struct {
	zenrpc.Service
	Repo db.NewsRepo
}

func NewNewsRPC(repo db.NewsRepo) {
	news := &NewsService{
		Repo: repo,
	}

	rpc := zenrpc.NewServer(zenrpc.Options{ExposeSMD: true})
	rpc.Register("news", news)
	rpc.Use(zenrpc.Logger(log.New(os.Stderr, "", log.LstdFlags)))

	//gen := rpcgen.FromSMD(rpc.SMD())

	http.Handle("/v1/rpc/", rpc)
	http.Handle("/v1/rpc/doc/", http.HandlerFunc(zenrpc.SMDBoxHandler))
	//http.Handle("/v1/rpc/news/api.go", http.HandlerFunc(rpcgen.Handler(gen.GoClient())))

	logger.Noticef("starting server on %s", os.Getenv("HTTP_ADDR"))
	logger.Crit(http.ListenAndServe(os.Getenv("HTTP_ADDR"), nil))
}

func (n *NewsService) Get(ctx context.Context) ([]News, error) {
	var news []News

	newsRepo, err := n.Repo.NewsByFilters(ctx, &db.NewsSearch{}, db.PagerNoLimit, n.Repo.FullNews())
	if err != nil {
		return nil, zenrpc.NewError(500, err)
	}

	for _, v := range newsRepo {
		news = append(news, News{
			ID:         v.ID,
			Title:      v.Title,
			Content:    v.Content,
			TagIDs:     v.TagIDs,
			CategoryID: v.CategoryID,
			CreatedAt:  v.CreatedAt.Format(time.RFC822),
			Category: Category{
				ID:    v.Category.ID,
				Title: v.Category.Title,
			},
		})
	}

	return news, nil
}

func (n *NewsService) GetByID(ctx context.Context, id int64) ([]News, error) {
	var news []News

	newsRepo, err := n.Repo.NewsByFilters(ctx, &db.NewsSearch{ID: &id}, db.PagerNoLimit, n.Repo.FullNews())
	if err != nil {
		return nil, zenrpc.NewError(500, err)
	}

	for _, v := range newsRepo {
		news = append(news, News{
			ID:         v.ID,
			Title:      v.Title,
			Content:    v.Content,
			TagIDs:     v.TagIDs,
			CategoryID: v.CategoryID,
			CreatedAt:  v.CreatedAt.Format(time.RFC822),
			Category: Category{
				ID:    v.Category.ID,
				Title: v.Category.Title,
			},
		})
	}

	return news, nil
}

func (n *NewsService) Count(ctx context.Context) (int, error) {
	newsRepo, err := n.Repo.CountNews(ctx, &db.NewsSearch{}, n.Repo.FullNews())
	if err != nil {
		return 0, zenrpc.NewError(500, err)
	}

	return newsRepo, nil
}
