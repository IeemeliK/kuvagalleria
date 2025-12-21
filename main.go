package main

import (
	"log"
	"net/http"

	"github.com/IeemeliK/kuvagalleria/internal"
	"github.com/IeemeliK/kuvagalleria/internal/routes"
)

func main() {
	mux := http.NewServeMux()

	fileServer := http.FileServer(http.FS(internal.Static))
	mux.Handle("GET /static/", http.StripPrefix("/static/", fileServer))

	mux.HandleFunc("GET /", routes.HomeHandler)

	log.Println("Server starting on :8080...")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal("http.ListenAndServe:", err)
	}
}
