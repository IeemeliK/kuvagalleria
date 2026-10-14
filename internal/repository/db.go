// Package repository provides database handling for the app
package repository

import (
	"context"
	"fmt"

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
		return nil, fmt.Errorf("open database: %w", err)
	}

	return dbPool, nil
}
