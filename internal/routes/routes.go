// Package routes route handlers
package routes

import (
	t "html/template"
	"log"
	"net/http"

	"github.com/IeemeliK/kuvagalleria/internal"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := t.ParseFS(internal.Templates, "index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println("Template parsing error:", err)
		return
	}

	// Execute the template (wrapper is base.html which invokes "content" from index.html)
	err = tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println("Template execution error:", err)
	}
}
