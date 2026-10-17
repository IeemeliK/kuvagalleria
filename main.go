package main

import (
	"context"
	"log/slog"
	"net/http"
	"os"

	"github.com/gorilla/sessions"
	"github.com/joho/godotenv"

	"github.com/IeemeliK/kuvagalleria/internal/config"
	"github.com/IeemeliK/kuvagalleria/internal/middleware"
	"github.com/IeemeliK/kuvagalleria/internal/repository"
	"github.com/IeemeliK/kuvagalleria/internal/router"
	"github.com/IeemeliK/kuvagalleria/internal/service"
)

func main() {
	if err := godotenv.Load(); err != nil {
		slog.Error("loading env file", "error", err)
		os.Exit(1)
	}

	cfg, err := config.New()
	if err != nil {
		slog.Error("loading config", "error", err)
		os.Exit(1)
	}

	cfg.InitLogger()

	dbpool, err := repository.NewConnectionPool(context.Background(), &cfg.Database)
	if err != nil {
		slog.Error("connecting to database", "error", err)
		os.Exit(1)
	}
	defer dbpool.Close()

	store := sessions.NewCookieStore([]byte(cfg.Session.Secret))

	authSvc := service.NewAuthService(dbpool, store)
	authMdw := &middleware.Authenticator{Store: store, DB: dbpool}

	handler := router.New(authSvc, authMdw)

	slog.Info("server starting", "addr", ":8080")
	if err := http.ListenAndServe(":8080", handler); err != nil {
		slog.Error("http listen and serve", "error", err)
		os.Exit(1)
	}
}
