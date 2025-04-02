package main

import (
	"context"
	"flag"
	"os"
	"os/signal"
	"syscall"

	"github.com/twirapp/language-detector/internal/http"
	"github.com/twirapp/language-detector/internal/predictor"
)

var modelPath string

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

	//lid.176.bin

	flag.StringVar(&modelPath, "modelpath", "", "Path to lang model")
	flag.Parse()

	if modelPath == "" {
		wd, err := os.Getwd()
		if err != nil {
			panic(err)
		}
		modelPath = wd + "/lid.176.bin"
	}

	pr, err := predictor.New(modelPath)
	if err != nil {
		panic(err)
	}
	http.New(appCtx, port, pr)

	exitSignal := make(chan os.Signal, 1)
	signal.Notify(exitSignal, syscall.SIGINT, syscall.SIGTERM)

	<-exitSignal
	appCtxCancel()
}
