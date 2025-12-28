package main

import (
	"log"
	"net/http"

	"github.com/IeemeliK/kuvagalleria/internal"
	"github.com/IeemeliK/kuvagalleria/internal/routes"
)

func main() {
	mux := http.NewServeMux()

	mux.Handle("GET /static/", http.FileServerFS(internal.Static))

	mux.HandleFunc("GET /", routes.HomeHandler)

	log.Println("Server starting on :8080...")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal("http.ListenAndServe:", err)
	}
}
