// Package routes route handlers
package routes

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	// Parse the templates
	files := []string{
		filepath.Join("views", "base.html"),
		filepath.Join("views", "index.html"),
	}

	tmpl, err := template.ParseFiles(files...)
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
