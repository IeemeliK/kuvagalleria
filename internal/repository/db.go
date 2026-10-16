// Package repository provides database handling for the app
package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/IeemeliK/kuvagalleria/internal/config"
	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/jackc/pgx/v5/stdlib"
)

func NewConnectionPool(ctx context.Context, cfg *config.DatabaseConfig) (*pgxpool.Pool, error) {
	dbURL := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s",
		cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DBName,
	)

	dbPool, err := pgxpool.New(ctx, dbURL)
	if err != nil {
		return nil, fmt.Errorf("open db connection pool: %w", err)
	}

	pingCtx, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()
	pingErr := dbPool.Ping(pingCtx)
	if pingErr != nil {
		return nil, fmt.Errorf("getting connection from pool: %w", pingErr)
	}

	return dbPool, nil
}
