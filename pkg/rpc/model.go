package rpc

import (
	"github.com/vmkteam/zenrpc/v2"
	"news/pkg/db"
	"sync"
)

type NewsService struct {
	zenrpc.Service
	Repository db.NewsRepo
	tags       sync.Map
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

type Tag struct {
	ID    int64  `json:"id"`
	Title string `json:"title"`
}

type Category struct {
	ID    int64  `json:"id"`
	Title string `json:"title"`
}

type News struct {
	ID         int64    `json:"id"`
	Title      string   `json:"title"`
	Content    string   `json:"content"`
	Tags       []Tag    `json:"tags"`
	CategoryID int64    `json:"categoryID"`
	CreatedAt  string   `json:"createdAt"`
	Category   Category `json:"category"`
}
