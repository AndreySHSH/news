//go:generate zenrpc
package rpc

import (
	"context"
	"errors"
	"github.com/vmkteam/zenrpc/v2"
	"news/pkg/db"
)

func (n *NewsService) Get(ctx context.Context, filter *NewsSearch) ([]News, error) {
	nn, err := n.Repo.NewsByFilters(ctx, filter.ToDB(), db.PagerNoLimit, n.Repo.FullNews())
	if err != nil {
		return nil, zenrpc.NewError(500, err)
	}

	return newNewsList(nn), nil
}

func (n *NewsService) GetByID(ctx context.Context, id int64) (*News, error) {
	news, err := n.Repo.NewsByID(ctx, id, n.Repo.FullNews())
	if err != nil {
		return nil, zenrpc.NewError(500, err)
	}

	if news == nil {
		return nil, zenrpc.NewError(404, errors.New("news is not found"))
	}

	return newNews(news), nil
}

func (n *NewsService) Count(ctx context.Context, filter *NewsSearch) (int, error) {
	count, err := n.Repo.CountNews(ctx, filter.ToDB(), n.Repo.FullNews())
	if err != nil {
		return 0, zenrpc.NewError(500, err)
	}

	return count, nil
}
