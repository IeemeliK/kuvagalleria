package main

import (
	"log"
	"net/http"

	m "github.com/IeemeliK/kuvagalleria/internal/middleware"
	"github.com/IeemeliK/kuvagalleria/internal/routes"
)

func main() {
	mux := http.NewServeMux()

	// Serve static files
	fileServer := http.FileServer(http.Dir("./static"))
	mux.Handle("GET /static/", http.StripPrefix("/static/", fileServer))

	// Home route
	mux.HandleFunc("GET /", routes.HomeHandler)

	log.Println("Server starting on :8080...")
	err := http.ListenAndServe(":8080", m.LoggingMiddleware(mux))
	if err != nil {
		log.Fatal(err)
	}
}
