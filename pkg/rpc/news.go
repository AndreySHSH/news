//go:generate zenrpc
package rpc

import (
	"context"
	"errors"
	"github.com/vmkteam/zenrpc/v2"
	"news/pkg/db"
)

func (n *NewsService) Get(ctx context.Context, filter *NewsSearch) ([]News, error) {
	nl, err := n.Repository.NewsByFilters(ctx, filter.ToDB(), db.PagerNoLimit, n.Repository.FullNews())
	if err != nil {
		return nil, zenrpc.NewError(500, err)
	}

	return newNewsList(nl), nil
}

func (n *NewsService) GetByID(ctx context.Context, id int64) (*News, error) {
	news, err := n.Repository.NewsByID(ctx, id, n.Repository.FullNews())
	if err != nil {
		return nil, zenrpc.NewError(500, err)
	}

	if news == nil {
		return nil, zenrpc.NewError(404, errors.New("news is not found"))
	}

	return newNews(news), nil
}

func (n *NewsService) Count(ctx context.Context, filter *NewsSearch) (int, error) {
	count, err := n.Repository.CountNews(ctx, filter.ToDB(), n.Repository.FullNews())
	if err != nil {
		return 0, zenrpc.NewError(500, err)
	}

	return count, nil
}
