package main

import (
	"context"
	"github.com/rollmelette/rollmelette"
	"log/slog"
)

func main() {
	//////////////////////// Setup Application //////////////////////////
	app := NewApp()

	///////////////////////// Rollmelette //////////////////////////
	ctx := context.Background()
	opts := rollmelette.NewRunOpts()
	opts.RollupURL = "http://127.0.0.1:5004"
	err := rollmelette.Run(ctx, opts, app)
	if err != nil {
		slog.Error("application error", "error", err)
	}
}
