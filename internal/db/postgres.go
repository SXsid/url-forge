package db

import (
	"context"
	"time"

	"github/SXsid/url-forge/internal/config"

	"github.com/jackc/pgx/v5/pgxpool"
)

func NewPool(ctx context.Context, dbConfig config.DBConfig) (*pgxpool.Pool, error) {
	cfg, err := pgxpool.ParseConfig(dbConfig.DSN)
	if err != nil {
		return nil, err
	}
	cfg.MaxConnIdleTime = time.Duration(dbConfig.MaxConnIdelTime) * time.Minute
	cfg.MaxConnLifetime = time.Duration(dbConfig.MaxConnLifeTime) * time.Minute
	cfg.MaxConns = int32(dbConfig.MaxConn)
	cfg.MinConns = int32(dbConfig.MinConn)
	pgx, err := pgxpool.NewWithConfig(ctx, cfg)
	if err != nil {
		return nil, err
	}
	return pgx, nil
}

func Test(ctx context.Context, pgxpool *pgxpool.Pool) error {
	timeCtx, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()
	if err := pgxpool.Ping(timeCtx); err != nil {
		return err
	}
	return nil
}
