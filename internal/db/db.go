// Package db provides database handling for the app
package db

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
)

type Config struct {
	User     string
	Password string
	Host     string
	Port     string
	DBName   string
}

func NewConnection(ctx context.Context, cfg Config) (*sql.DB, error) {
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
