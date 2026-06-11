package api

import (
	"log"
	"net/http"

	"github.com/IeemeliK/kuvagalleria/internal/middleware"
	"github.com/IeemeliK/kuvagalleria/internal/templates"
)

func HomeHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		username, ok := middleware.UsernameFromContext(r.Context())
		if !ok {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

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
