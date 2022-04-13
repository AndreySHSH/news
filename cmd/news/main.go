package main

import (
	"context"
	"fmt"
	"github.com/go-pg/pg/v10"
	"news/pkg/app"
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
		panic(fmt.Sprintf(`fatal error loading .env file, error: %s`, err.Error()))
	}

	pgDataOnConnect := pg.Options{
		User:     os.Getenv("POSTGRES_USER"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
		Database: os.Getenv("POSTGRES_DBNAME"),
		Addr:     fmt.Sprintf(`%s:%s`, os.Getenv("POSTGRES_HOST"), os.Getenv("POSTGRES_PORT")),
	}

	httpAdder := os.Getenv("HTTP_ADDR")
	if httpAdder == "" {
		panic("failed to get http address from environment")
	}

	ctxApp, cancel := context.WithCancel(context.Background())
	defer cancel()

	db := pg.Connect(&pgDataOnConnect)
	if err := db.Ping(ctxApp); err != nil {
		panic(fmt.Sprintf(`fatal error connect DB, error: %s`, err.Error()))
	}

	defer func(db *pg.DB) {
		err := db.Close()
		if err != nil {
			logger.Crit(err)
		}
	}(db)

	newsRPC, gen := rpc.NewNewsRPC(*db)
	go func() {
		err := app.NewHTTP(newsRPC, *gen, httpAdder)
		if err != nil {
			panic(err)
		}
	}()

	<-app.GracefulShutdown()
	_, forceCancel := context.WithTimeout(ctxApp, shutDownDuration)
	defer forceCancel()

	logger.Notice("Graceful Shutdown")
}
