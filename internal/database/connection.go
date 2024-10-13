package database

import (
	"context"
	"fmt"

	"github.com/techforge-lat/bastion/internal/config"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Port interface{}

type Adapter struct {
	*pgxpool.Pool
}

func New(conf config.Root) (*Adapter, error) {
	config, err := pgxpool.ParseConfig(fmt.Sprintf("%s://%s:%s@%s:%d/%s?sslmode=%s",
		conf.Database.Driver,
		conf.Database.User,
		conf.Database.Password,
		conf.Database.Host,
		conf.Database.Port,
		conf.Database.Name,
		conf.Database.SSLMode,
	))
	if err != nil {
		return nil, fmt.Errorf("unable to parse config connection: %w", err)
	}

	dbPool, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		return nil, fmt.Errorf("unable to create connection pool: %w", err)
	}

	if err := dbPool.Ping(context.Background()); err != nil {
		return nil, fmt.Errorf("unable to ping database: %w", err)
	}

	return &Adapter{dbPool}, nil
}
