// Code generated by mfd-generator; DO NOT EDIT.

//nolint
//lint:file-ignore U1000 ignore unused code, it's generated
package db

import (
	"time"
)

var Columns = struct {
	Category struct {
		ID, Title, StatusID string
	}
	News struct {
		ID, Title, Content, TagIDs, CategoryID, CreatedAt, StatusID string

		Category string
	}
	Tag struct {
		ID, Title, StatusID string
	}
}{
	Category: struct {
		ID, Title, StatusID string
	}{
		ID:       "categoryId",
		Title:    "title",
		StatusID: "statusId",
	},
	News: struct {
		ID, Title, Content, TagIDs, CategoryID, CreatedAt, StatusID string

		Category string
	}{
		ID:         "newsId",
		Title:      "title",
		Content:    "content",
		TagIDs:     "tagIds",
		CategoryID: "categoryId",
		CreatedAt:  "createdAt",
		StatusID:   "statusId",

		Category: "Category",
	},
	Tag: struct {
		ID, Title, StatusID string
	}{
		ID:       "tagId",
		Title:    "title",
		StatusID: "statusId",
	},
}

var Tables = struct {
	Category struct {
		Name, Alias string
	}
	News struct {
		Name, Alias string
	}
	Tag struct {
		Name, Alias string
	}
}{
	Category: struct {
		Name, Alias string
	}{
		Name:  "categories",
		Alias: "t",
	},
	News: struct {
		Name, Alias string
	}{
		Name:  "news",
		Alias: "t",
	},
	Tag: struct {
		Name, Alias string
	}{
		Name:  "tags",
		Alias: "t",
	},
}

type Category struct {
	tableName struct{} `pg:"categories,alias:t,discard_unknown_columns"`

	ID       int64  `pg:"categoryId,pk"`
	Title    string `pg:"title,use_zero"`
	StatusID int    `pg:"statusId,use_zero"`
}

type News struct {
	tableName struct{} `pg:"news,alias:t,discard_unknown_columns"`

	ID         int64     `pg:"newsId,pk"`
	Title      string    `pg:"title,use_zero"`
	Content    string    `pg:"content,use_zero"`
	TagIDs     []int64   `pg:"tagIds,array"`
	CategoryID int64     `pg:"categoryId,use_zero"`
	CreatedAt  time.Time `pg:"createdAt,use_zero"`
	StatusID   int       `pg:"statusId,use_zero"`

	Category *Category `pg:"fk:categoryId,rel:has-one"`
}

type Tag struct {
	tableName struct{} `pg:"tags,alias:t,discard_unknown_columns"`

	ID       int64  `pg:"tagId,pk"`
	Title    string `pg:"title,use_zero"`
	StatusID int    `pg:"statusId,use_zero"`
}
