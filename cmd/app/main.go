package main

import (
	"context"

	"github.com/amagkn/translabor/config"
	"github.com/amagkn/translabor/internal/app"
	"github.com/amagkn/translabor/pkg/logger"
	"github.com/amagkn/translabor/pkg/validation"
)

func main() {
	ctx := context.Background()

	c, err := config.New()
	if err != nil {
		logger.Fatal(err, "config.New")
	}

	logger.Init(c.Logger)
	validation.Init()

	err = app.Run(ctx, c)
	if err != nil {
		logger.Fatal(err, "app.Run")
	}

	logger.Info("App stopped!")
}
