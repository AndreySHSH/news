package project

import (
	"context"
	"news/pkg/db"
)

type NewsManger struct {
	repo db.NewsRepo
}

func NewNewsManager(repository db.NewsRepo) NewsManger {
	return NewsManger{
		repo: repository,
	}
}

func (m *NewsManger) Get(ctx context.Context, filter *db.NewsSearch) ([]News, error) {
	dbNewsList, err := m.repo.NewsByFilters(ctx, filter, db.PagerNoLimit, m.repo.FullNews())
	if err != nil {
		return nil, err
	}

	nl := newNewsList(dbNewsList)

	if tags, err := m.getTags(ctx, nl); err != nil {
		return nil, err
	} else {
		nl.SetTags(tags)
	}

	return nl, nil
}

func (m *NewsManger) GetByID(ctx context.Context, id int64) (*News, error) {
	dbNews, err := m.repo.NewsByID(ctx, id, m.repo.FullNews())
	if err != nil {
		return nil, err
	} else if dbNews == nil {
		return nil, nil
	}

	n := newNews(dbNews)
	tags, err := m.getTags(ctx, NewsList{*n})
	if err != nil {
		return nil, err
	}

	n.setTags(tags)

	return n, nil
}

func (m *NewsManger) Count(ctx context.Context, filter *db.NewsSearch) (int, error) {
	count, err := m.repo.CountNews(ctx, filter, m.repo.FullNews())
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (m *NewsManger) getTags(ctx context.Context, nl NewsList) (map[int64]db.Tag, error) {
	var tags = make(map[int64]db.Tag)
	var tagsIDs []int64

	for k := range nl.TagIDs() {
		tagsIDs = append(tagsIDs, k)
	}

	tl, err := m.repo.TagsByFilters(ctx, &db.TagSearch{IDs: tagsIDs}, db.PagerNoLimit)
	if err != nil {
		return nil, err
	}

	for _, v := range tl {
		tags[v.ID] = db.Tag{
			ID:    v.ID,
			Title: v.Title,
		}
	}

	return tags, nil
}
