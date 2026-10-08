package api

import (
	"log/slog"
	"net/http"

	"github.com/IeemeliK/kuvagalleria/internal/service"
)

func LogoutHandler(authSvc *service.AuthService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := authSvc.ClearSession(w, r); err != nil {
			slog.Error("clearing session", "error", err)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		w.Header().Set("HX-Redirect", "/login")
		w.WriteHeader(http.StatusOK)
	}
}
