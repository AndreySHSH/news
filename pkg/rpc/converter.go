package rpc

import (
	"news/pkg/db"
	"news/pkg/project"
	"time"
)

func newCategoryList(in []db.Category) []Category {
	var out []Category
	for _, v := range in {
		out = append(out, *newCategory(&v))
	}

	return out
}

func newCategory(in *db.Category) *Category {
	if in == nil {
		return nil
	}

	return &Category{
		ID:    in.ID,
		Title: in.Title,
	}
}

func newNewsList(in []project.News) []News {
	var out []News
	for _, v := range in {
		out = append(out, *newNews(&v))
	}

	return out
}

func newNews(in *project.News) *News {
	if in == nil {
		return nil
	}

	var tags []Tag
	for _, v := range in.Tags {
		tags = append(tags, Tag{
			ID:    v.ID,
			Title: v.Title,
		})
	}

	return &News{
		ID:         in.ID,
		Title:      in.Title,
		Content:    in.Content,
		Tags:       tags,
		CategoryID: in.CategoryID,
		CreatedAt:  in.CreatedAt.Format(time.RFC822),
		Category:   *newCategory(in.Category),
	}
}

func (in *CategorySearch) ToDB() *db.CategorySearch {
	if in == nil {
		return nil
	}

	return &db.CategorySearch{
		ID: in.CategoryID,
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
