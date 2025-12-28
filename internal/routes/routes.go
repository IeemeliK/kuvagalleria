// Package routes route handlers
package routes

import (
	"html/template"
	"log"
	"net/http"

	"github.com/IeemeliK/kuvagalleria/internal"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFS(internal.Templates, "base.tmpl", "index.tmpl"))

	if err := tmpl.Execute(w, nil); err != nil {
		log.Printf("template execute error: %v", err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
}
