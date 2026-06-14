package main

import (
	"context"
	"fmt"
	"log"

	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"

	"github.com/IeemeliK/kuvagalleria/internal/config"
	"github.com/IeemeliK/kuvagalleria/internal/repository"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading env file(s)")
	}

	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("load config: %v", err)
	}

	db, err := repository.NewConnection(context.Background(), repository.Config{
		User:     cfg.Database.User,
		Password: cfg.Database.Password,
		Host:     cfg.Database.Host,
		Port:     cfg.Database.Port,
		DBName:   cfg.Database.DBName,
	})
	if err != nil {
		log.Fatalf("connect to database: %v", err)
	}
	defer db.Close()

	hash, err := bcrypt.GenerateFromPassword([]byte("admin"), bcrypt.DefaultCost)
	if err != nil {
		log.Fatalf("hash password: %v", err)
	}

	_, err = db.ExecContext(context.Background(),
		`INSERT INTO users (username, password_hash)
		 VALUES ($1, $2)
		 ON CONFLICT (username) DO UPDATE SET password_hash = $2`,
		"admin", string(hash),
	)
	if err != nil {
		log.Fatalf("seed user: %v", err)
	}

	fmt.Println("Seeded admin user (username: admin, password: admin)")
}
