package main

import (
	"context"
	"log/slog"
	"os"

	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"

	"github.com/IeemeliK/kuvagalleria/internal/config"
	"github.com/IeemeliK/kuvagalleria/internal/repository"
)

func main() {
	if err := godotenv.Load(); err != nil {
		slog.Error("loading env file", "error", err)
		os.Exit(1)
	}

	cfg, err := config.Load()
	if err != nil {
		slog.Error("loading config", "error", err)
		os.Exit(1)
	}

	db, err := repository.NewConnection(context.Background(), repository.Config{
		User:     cfg.Database.User,
		Password: cfg.Database.Password,
		Host:     "localhost",
		Port:     cfg.Database.Port,
		DBName:   cfg.Database.DBName,
	})
	if err != nil {
		slog.Error("connecting to database", "error", err)
		os.Exit(1)
	}
	defer db.Close()

	hash, err := bcrypt.GenerateFromPassword([]byte("admin"), bcrypt.DefaultCost)
	if err != nil {
		slog.Error("hashing password", "error", err)
		os.Exit(1)
	}

	_, err = db.ExecContext(
		context.Background(),
		`INSERT INTO users (username, password_hash)
		 VALUES ($1, $2)
		 ON CONFLICT (username) DO UPDATE SET password_hash = $2`,
		"admin", string(hash),
	)
	if err != nil {
		slog.Error("seeding user", "error", err)
		os.Exit(1)
	}

	slog.Info("seeded admin user")
}
