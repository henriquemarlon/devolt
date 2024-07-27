package main

import (
	"context"
	"log/slog"
	"github.com/devolthq/devolt/internal/infra/cartesi/router"
	"github.com/rollmelette/rollmelette"
)

func main() {
	//////////////////////// Setup Application //////////////////////////
	app := router.NewApp()

	///////////////////////// Rollmelette //////////////////////////
	ctx := context.Background()
	opts := rollmelette.NewRunOpts()
	err := rollmelette.Run(ctx, opts, app)
	if err != nil {
		slog.Error("application error", "error", err)
	}
}
