package main

import (
	"context"
	"github.com/devolthq/devolt/internal/infra/cartesi/router"
	"github.com/rollmelette/rollmelette"
	"log/slog"
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
