package loaders

import (
	"fmt"
	"github.com/go-pg/pg/v10"
	"os"
)

func NewDB() *pg.DB {
	db := pg.Connect(&pg.Options{
		User:     os.Getenv("POSTGRES_USER"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
		Database: os.Getenv("POSTGRES_DBNAME"),
		Addr:     fmt.Sprintf(`%s:%s`, os.Getenv("POSTGRES_HOST"), os.Getenv("POSTGRES_PORT")),
	})

	return db
}
