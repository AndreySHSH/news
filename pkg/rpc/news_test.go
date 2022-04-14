package rpc

import (
	"context"
	"fmt"
	"github.com/go-pg/pg/v10"
	"github.com/ivahaev/go-logger"
	"github.com/stretchr/testify/assert"
	"news/pkg/db"
	"testing"
)

func envPreparation() (*NewsService, error) {
	pgDataOnConnect := pg.Options{
		User:     "postgres",
		Password: "postgres",
		Database: "project",
		Addr:     fmt.Sprintf(`%s:%s`, "127.0.0.1", "5432"),
	}

	dbc := pg.Connect(&pgDataOnConnect)
	if err := dbc.Ping(context.Background()); err != nil {
		logger.Errorf(`fatal error connect DB, error: %s`, err.Error())
		return nil, err
	}

	dbLayer := db.New(dbc)
	repo := db.NewNewsRepo(dbLayer)

	return NewNewsService(repo), nil
}

func TestGetByID(t *testing.T) {
	tests := []News{
		0: {
			ID:      2,
			Title:   "еще одна новость",
			Content: "и тут что то",
			Tags: []Tag{
				0: {
					ID:    1,
					Title: "all",
				},
			},
			CategoryID: 1,
			Category: Category{
				ID:    1,
				Title: "all category",
			},
			CreatedAt: "12 Apr 22 16:34 MSK",
		},
		1: {
			ID:      1,
			Title:   "какая то новость",
			Content: "что то",
			Tags: []Tag{
				0: {
					ID:    1,
					Title: "all",
				},
				1: {
					ID:    2,
					Title: "new_tag",
				},
			},
			CategoryID: 1,
			Category: Category{
				ID:    1,
				Title: "all category",
			},
			CreatedAt: "12 Apr 22 14:44 MSK",
		},
	}

	mews, err := envPreparation()
	if err != nil {
		t.Error(err)
	}

	for _, tc := range tests {
		n, err := mews.GetByID(context.Background(), tc.ID)
		if err != nil {
			t.Error(err)
		}
		assert.Equal(t, tc, *n)
	}
}
