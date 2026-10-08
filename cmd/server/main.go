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
	"github.com/IeemeliK/kuvagalleria/internal/templates"
	"github.com/IeemeliK/kuvagalleria/web"
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

	config.InitLogger(cfg)

	database, err := repository.NewConnection(context.Background(), repository.Config{
		User:     cfg.Database.User,
		Password: cfg.Database.Password,
		Host:     cfg.Database.Host,
		Port:     cfg.Database.Port,
		DBName:   cfg.Database.DBName,
	})
	if err != nil {
		slog.Error("connecting to database", "error", err)
		os.Exit(1)
	}
	defer func() {
		if cerr := database.Close(); cerr != nil {
			slog.Warn("closing database", "error", cerr)
		}
	}()

	if err := templates.Init(web.Templates()); err != nil {
		slog.Error("initializing templates", "error", err)
		os.Exit(1)
	}

	store := sessions.NewCookieStore([]byte(cfg.Session.Secret))

	authSvc := service.NewAuthService(database, store)
	authMdw := &middleware.Authenticator{Store: store, DB: database}

	handler := router.New(authSvc, authMdw)

	slog.Info("server starting", "addr", ":8080")
	if err := http.ListenAndServe(":8080", handler); err != nil {
		slog.Error("http listen and serve", "error", err)
		os.Exit(1)
	}
}
