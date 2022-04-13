package main

import (
	"context"
	"fmt"
	"github.com/go-pg/pg/v10"
	"news/pkg/app"
	"news/pkg/db"
	"news/pkg/rpc"
	"os"
	"time"

	"github.com/ivahaev/go-logger"
	"github.com/joho/godotenv"
)

const (
	shutDownDuration = 5 * time.Second
)

func main() {
	err := logger.SetLevel("debug")
	if err != nil {
		panic(fmt.Sprintf(`failed init logs, error: %s`, err.Error()))
	}

	err = godotenv.Load(".env")
	if err != nil {
		panic(fmt.Sprintf(`fatal error loading .env file, %s`, err.Error()))
	}

	pgDataOnConnect := pg.Options{
		User:     os.Getenv("POSTGRES_USER"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
		Database: os.Getenv("POSTGRES_DBNAME"),
		Addr:     fmt.Sprintf(`%s:%s`, os.Getenv("POSTGRES_HOST"), os.Getenv("POSTGRES_PORT")),
	}

	ctxApp, cancel := context.WithCancel(context.Background())
	defer cancel()

	pgConnect := app.NewDB(pgDataOnConnect)
	dbLayer := db.New(pgConnect)
	newRepo := db.NewNewsRepo(dbLayer)

	newsRPC, gen := rpc.NewNewsRPC(newRepo)
	go app.NewHTTP(newsRPC, *gen)

	<-app.GracefulShutdown()
	_, forceCancel := context.WithTimeout(ctxApp, shutDownDuration)

	logger.Notice("Graceful Shutdown")
	defer forceCancel()
}
