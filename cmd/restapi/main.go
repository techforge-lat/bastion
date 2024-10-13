package main

import (
	"context"

	"github.com/techforge-lat/bastion/internal/config"
	"github.com/techforge-lat/bastion/internal/server"
)

func main() {
	conf := config.Load()

	ctx := context.Background()

	server, err := server.New(ctx, conf)
	if err != nil {
		server.Logger.ErrorContext(ctx, "unable to create server", "error", err)
		return
	}

	if err := server.Execute(); err != nil {
		server.Logger.ErrorContext(ctx, "unable to execute server", "error", err)
	}
}
