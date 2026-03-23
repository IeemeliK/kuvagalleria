// Package routes route handlers
package routes

import (
	"log"
	"net/http"

	"github.com/IeemeliK/kuvagalleria/internal/middleware"
	"github.com/IeemeliK/kuvagalleria/internal/templates"
)

func HomeHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		username := r.Context().Value(middleware.UsernameKey).(string)

		data := PageData{
			Username: username,
			LoggedIn: true,
		}

		if err := templates.Render(w, "index.html", "", data); err != nil {
			log.Printf("template render error: %v", err)
			http.Error(w, "internal server error", http.StatusInternalServerError)
		}
	}
}
