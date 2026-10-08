package api

import (
	"log/slog"
	"net/http"

	"github.com/IeemeliK/kuvagalleria/internal/middleware"
	"github.com/IeemeliK/kuvagalleria/internal/templates"
)

type HomePageData struct {
	HeaderData HeaderData
}

func HomeHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		username, ok := middleware.UsernameFromContext(r.Context())
		if !ok {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		data := HomePageData{
			HeaderData: HeaderData{
				Username: username,
				LoggedIn: true,
			},
		}

		if err := templates.Render(w, "index.html", "", data); err != nil {
			slog.Error("rendering home template", "error", err)
			http.Error(w, "internal server error", http.StatusInternalServerError)
		}
	}
}
