package main

import (
	"log"
	"net/http"
	"os"

	"github.com/IeemeliK/kuvagalleria/internal"
	"github.com/IeemeliK/kuvagalleria/internal/db"
	"github.com/IeemeliK/kuvagalleria/internal/routes"
	"github.com/gorilla/sessions"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading env file(s)")
	}

	database := db.CreateConnection()
	defer func() {
		if err := database.Close(); err != nil {
			log.Printf("error closing database: %v", err)
		}
	}()

	store := sessions.NewCookieStore([]byte(os.Getenv("COOKIESTORE_SECRET")))
	mux := http.NewServeMux()

	mux.Handle("GET /static/", http.FileServerFS(internal.Static))
	mux.HandleFunc("/", routes.HomeHandler(store))
	mux.HandleFunc("/login", routes.LoginHandler(store, database))
	mux.HandleFunc("/logout", routes.LogoutHandler(store))

	log.Println("Server starting on :8080...")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal("http.ListenAndServe:", err)
	}
}
