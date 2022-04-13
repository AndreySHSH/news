//go:generate zenrpc
package rpc

import (
	"context"
	"github.com/vmkteam/zenrpc/v2"
	"news/pkg/db"
)

func (n *CategoryService) Get(ctx context.Context, filter *CategorySearch) ([]Category, error) {
	cl, err := n.Repository.CategoriesByFilters(ctx, filter.ToDB(), db.PagerNoLimit)
	if err != nil {
		return nil, zenrpc.NewError(500, err)
	}

	return newCategoryList(cl), nil
}
