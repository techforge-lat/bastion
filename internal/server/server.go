package server

import (
	"context"
	"fmt"
	"net"
	"os"
	"os/signal"

	"github.com/labstack/echo/v4"
	"github.com/techforge-lat/bastion/internal/config"
	"github.com/techforge-lat/bastion/internal/database"
	"github.com/techforge-lat/bastion/internal/di"
	"github.com/techforge-lat/bastion/internal/logger"
)

const rateLimiterMemoryStore = 20

type Adapter struct {
	Config      config.Root
	Server      *echo.Echo
	DB          *database.Adapter
	Logger      *logger.Adapter
	DIContainer *di.Container
}

func New(ctx context.Context, conf config.Root) (*Adapter, error) {
	logger := logger.New()

	db, err := database.New(conf)
	if err != nil {
		logger.ErrorContext(ctx, "unable to create database connection", "error", err)
		return nil, err
	}

	return &Adapter{
		Config:      conf,
		Server:      newEcho(conf, nil),
		DB:          db,
		Logger:      logger,
		DIContainer: di.NewContainer(),
	}, nil
}

func (a Adapter) Execute() error {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	a.Server.Server.BaseContext = func(listener net.Listener) context.Context {
		return ctx
	}
	srvErr := make(chan error, 1)

	// Start server
	go func() {
		srvErr <- a.Server.Start(fmt.Sprintf(":%d", a.Config.ServerPort))
	}()

	// Wait for interruption.
	select {
	case err := <-srvErr:
		// Error when starting HTTP server.
		return err
	case <-ctx.Done():
		// Wait for first CTRL+C.
		// Stop receiving signal notifications as soon as possible.
		stop()
	}

	if err := a.Server.Shutdown(ctx); err != nil {
		a.Logger.ErrorContext(ctx, "unable to shutdown server", "error", err)
		return err
	}

	return nil
}
