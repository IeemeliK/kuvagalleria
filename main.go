package main

import (
	"log"
	"net/http"
	"os"

	"github.com/IeemeliK/kuvagalleria/internal"
	"github.com/IeemeliK/kuvagalleria/internal/routes"
	"github.com/gorilla/sessions"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
	// Initialize session store
	store := sessions.NewCookieStore([]byte(os.Getenv("COOKIESTORE_SECRET")))

	mux := http.NewServeMux()

	// Serve static files
	mux.Handle("GET /static/", http.FileServerFS(internal.Static))

	mux.HandleFunc("/", routes.HomeHandler)
	mux.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		routes.LoginHandler(w, r, store)
	})

	log.Println("Server starting on :8080...")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal("http.ListenAndServe:", err)
	}
}
