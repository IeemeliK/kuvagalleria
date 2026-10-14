// Package repository provides database handling for the app
package repository

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/IeemeliK/kuvagalleria/internal/config"
	_ "github.com/jackc/pgx/v5/stdlib"
)

func NewConnection(ctx context.Context, cfg *config.DatabaseConfig) (*pgxpool.Pool, error) {
	dbURL := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s",
		cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DBName,
	)

	db, err := sql.Open("pgx", dbURL)
	if err != nil {
		return nil, fmt.Errorf("open database: %w", err)
	}

	pingCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	if err := db.PingContext(pingCtx); err != nil {
		db.Close()
		return nil, fmt.Errorf("ping database: %w", err)
	}

	return db, nil
}
