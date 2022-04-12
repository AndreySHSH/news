package loaders

import (
	"os"
	"os/signal"
	"syscall"
)

func GracefulShutdown() chan os.Signal {
	done := make(chan os.Signal, 1)
	signal.Notify(done, syscall.SIGINT, syscall.SIGTERM)

	return done
}
