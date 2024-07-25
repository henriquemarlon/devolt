package main

import (
	"context"
	"github.com/rollmelette/rollmelette"
	"github.com/devolthq/devolt/internal/infra/cartesi/router"
	"log/slog"
)

func main() {
	//////////////////////// Setup Application //////////////////////////
	app := router.SetupApplication()

	///////////////////////// Rollmelette //////////////////////////
	ctx := context.Background()
	opts := rollmelette.NewRunOpts()
	err := rollmelette.Run(ctx, opts, app)
	if err != nil {
		slog.Error("application error", "error", err)
	}
}
