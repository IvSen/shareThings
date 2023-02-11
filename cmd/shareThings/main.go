package main

import (
	"context"

	"github.com/IvSen/shareThings/internal/app"

	"github.com/IvSen/shareThings/pkg/logging"

	"github.com/IvSen/shareThings/pkg/config"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	logging.Init(ctx)
	logger := logging.GetLogger()

	logger.Println("logger initialized")

	logger.Info("config initializing")
	cfg := config.GetConfig()

	a, err := app.NewApp(ctx, cfg, logger)
	if err != nil {
		logger.Fatal(err)
	}

	logger.Info("Running Application")
	err = a.Run(ctx)
	if err != nil {
		logger.Fatal(err)
		return
	}
}
