package app

import (
	"github.com/go-pg/pg/v10"
)

func NewDB(in pg.Options) *pg.DB {
	return pg.Connect(&in)
}
