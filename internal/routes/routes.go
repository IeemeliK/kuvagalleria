// Package routes route handlers
package routes

import (
	"html/template"
	"log"
	"net/http"

	"github.com/IeemeliK/kuvagalleria/internal"
)

type PageData struct {
	LoggedIn bool
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	data := PageData{LoggedIn: false}
	tmpl := template.Must(template.ParseFS(internal.Templates, "*.html"))

	if err := tmpl.Execute(w, data); err != nil {
		log.Printf("template execute error: %v", err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
}
