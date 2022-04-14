//go:generate zenrpc
package rpc

import (
	"context"
	"errors"
	"github.com/vmkteam/zenrpc/v2"
	"net/http"
	"news/pkg/db"
	"news/pkg/project"
)

type NewsService struct {
	zenrpc.Service
	repo db.NewsRepo
	nm   project.NewsManger
}

func NewNewsService(repo db.NewsRepo) *NewsService {
	return &NewsService{repo: repo, nm: project.NewNewsManager(repo)}
}

func (n *NewsService) Get(ctx context.Context, filter *NewsSearch) ([]News, error) {
	nl, err := n.nm.Get(ctx, filter.ToDB())
	if err != nil {
		return nil, zenrpc.NewError(http.StatusInternalServerError, err)
	}

	return newNewsList(nl), nil
}

func (n *NewsService) GetByID(ctx context.Context, id int64) (*News, error) {
	ne, err := n.nm.GetByID(ctx, id)
	if err != nil {
		return nil, zenrpc.NewError(http.StatusInternalServerError, err)
	} else if ne == nil {
		return nil, zenrpc.NewError(http.StatusNotFound, errors.New("project is not found"))
	}

	return newNews(ne), nil
}

func (n *NewsService) Count(ctx context.Context, filter *NewsSearch) (int, error) {
	count, err := n.nm.Count(ctx, filter.ToDB())
	if err != nil {
		return 0, zenrpc.NewError(http.StatusInternalServerError, err)
	}

	return count, nil
}
