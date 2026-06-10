package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/IeemeliK/kuvagalleria/internal"
	"github.com/IeemeliK/kuvagalleria/internal/db"
	"github.com/IeemeliK/kuvagalleria/internal/middleware"
	"github.com/IeemeliK/kuvagalleria/internal/routes"
	"github.com/IeemeliK/kuvagalleria/internal/templates"
	"github.com/gorilla/sessions"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading env file(s)")
	}

	if err := internal.InitTemplates(); err != nil {
		log.Fatalf("init templates: %v", err)
	}
	if err := templates.Init(); err != nil {
		log.Fatalf("init templates: %v", err)
	}

	cfg := db.Config{
		User:     os.Getenv("POSTGRES_USER"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
		Host:     os.Getenv("HOST"),
		Port:     os.Getenv("PORT"),
		DBName:   os.Getenv("POSTGRES_DB"),
	}
	database, err := db.NewConnection(context.Background(), cfg)
	if err != nil {
		log.Fatalf("connect to database: %v", err)
	}
	defer func() {
		if cerr := database.Close(); cerr != nil {
			log.Printf("error closing database: %v", cerr)
		}
	}()

	store := sessions.NewCookieStore([]byte(os.Getenv("COOKIESTORE_SECRET")))
	mux := http.NewServeMux()

	auth := &middleware.Authenticator{Store: store, DB: database}
	handler := auth.Middleware(mux)

	mux.Handle("GET /static/", http.FileServerFS(internal.Static))
	mux.HandleFunc("/", routes.HomeHandler())
	mux.HandleFunc("/login", routes.LoginHandler(store, database))
	mux.HandleFunc("/logout", routes.LogoutHandler(store))

	log.Println("Server starting on :8080...")
	if err := http.ListenAndServe(":8080", handler); err != nil {
		log.Fatal("http.ListenAndServe:", err)
	}
}
