package main

import (
	"context"
	"log"
	"net/http"

	"github.com/gorilla/sessions"
	"github.com/joho/godotenv"

	"github.com/IeemeliK/kuvagalleria/internal/api"
	"github.com/IeemeliK/kuvagalleria/internal/config"
	"github.com/IeemeliK/kuvagalleria/internal/middleware"
	"github.com/IeemeliK/kuvagalleria/internal/repository"
	"github.com/IeemeliK/kuvagalleria/internal/service"
	"github.com/IeemeliK/kuvagalleria/internal/templates"
	"github.com/IeemeliK/kuvagalleria/web"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading env file(s)")
	}

	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("load config: %v", err)
	}

	database, err := repository.NewConnection(context.Background(), repository.Config{
		User:     cfg.Database.User,
		Password: cfg.Database.Password,
		Host:     cfg.Database.Host,
		Port:     cfg.Database.Port,
		DBName:   cfg.Database.DBName,
	})
	if err != nil {
		log.Fatalf("connect to database: %v", err)
	}
	defer func() {
		if cerr := database.Close(); cerr != nil {
			log.Printf("error closing database: %v", cerr)
		}
	}()

	if err := templates.Init(web.Templates()); err != nil {
		log.Fatalf("init templates: %v", err)
	}

	store := sessions.NewCookieStore([]byte(cfg.Session.Secret))
	mux := http.NewServeMux()

	auth := &middleware.Authenticator{Store: store, DB: database}
	handler := auth.Middleware(mux)

	authSvc := service.NewAuthService(database)

	mux.Handle("GET /static/", http.FileServerFS(web.Static))
	mux.HandleFunc("/", api.HomeHandler())
	mux.HandleFunc("/login", api.LoginHandler(store, authSvc))
	mux.HandleFunc("/logout", api.LogoutHandler(store))

	log.Println("Server starting on :8080...")
	if err := http.ListenAndServe(":8080", handler); err != nil {
		log.Fatal("http.ListenAndServe:", err)
	}
}
