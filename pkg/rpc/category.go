//go:generate zenrpc
package rpc

import (
	"context"
	"github.com/vmkteam/zenrpc/v2"
	"net/http"
	"news/pkg/db"
)

type CategoryService struct {
	zenrpc.Service
	repo db.NewsRepo
}

func NewCategoryService(repo db.NewsRepo) *CategoryService {
	return &CategoryService{repo: repo}
}

func (n *CategoryService) Get(ctx context.Context, filter *CategorySearch) ([]Category, error) {
	cl, err := n.repo.CategoriesByFilters(ctx, filter.ToDB(), db.PagerNoLimit)
	if err != nil {
		return nil, zenrpc.NewError(http.StatusInternalServerError, err)
	}

	return newCategoryList(cl), nil
}
