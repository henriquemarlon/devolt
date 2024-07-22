package main

import (
	"context"
	"github.com/rollmelette/rollmelette"
	"log/slog"
)

func main() {
	//////////////////////// Setup Application //////////////////////////
	app := SetupApplicationPersistent()

	///////////////////////// Rollmelette //////////////////////////
	ctx := context.Background()
	opts := rollmelette.NewRunOpts()
	err := rollmelette.Run(ctx, opts, app)
	if err != nil {
		slog.Error("application error", "error", err)
	}
}
