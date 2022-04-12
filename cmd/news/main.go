package main

import (
	"context"
	"fmt"
	"news/pkg/app/loaders"
	"news/pkg/db"
	"news/pkg/rpc"
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

	ctxApp, cancel := context.WithCancel(context.Background())
	defer cancel()

	dataBase := db.New(loaders.NewDB())
	newRepo := db.NewNewsRepo(dataBase)

	go rpc.NewNewsRPC(newRepo)

	<-loaders.GracefulShutdown()
	_, forceCancel := context.WithTimeout(ctxApp, shutDownDuration)

	logger.Notice("Graceful Shutdown")
	defer forceCancel()
}
