package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/twirapp/language-detector/internal/http"
	"github.com/twirapp/language-detector/internal/lingua"
)

func main() {
	appCtx, appCtxCancel := context.WithCancel(context.Background())

	port := os.Getenv("PORT")
	if port == "" {
		port = "3012"
	}

	appEnv := os.Getenv("APP_ENV")
	if appEnv == "" {
		appEnv = "development"
	}

	lingua := lingua.New(appEnv)
	http.New(appCtx, port, lingua)

	exitSignal := make(chan os.Signal, 1)
	signal.Notify(exitSignal, syscall.SIGINT, syscall.SIGTERM)

	<-exitSignal
	appCtxCancel()
}
