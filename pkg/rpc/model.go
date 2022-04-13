package rpc

import (
	"github.com/vmkteam/zenrpc/v2"
	"news/pkg/db"
)

type NewsService struct {
	zenrpc.Service
	Repository db.NewsRepo
}

type CategoryService struct {
	zenrpc.Service
	Repository db.NewsRepo
}

type NewsSearch struct {
	CategoryID *int64 `json:"categoryID"`
	TagID      *int64 `json:"tagID"`
}

type CategorySearch struct {
	CategoryID *int64 `json:"id"`
}

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
