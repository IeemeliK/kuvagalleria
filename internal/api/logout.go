package api

import (
	"log"
	"net/http"

	"github.com/gorilla/sessions"
)

func LogoutHandler(store *sessions.CookieStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		session, err := store.Get(r, "session-name")
		if err != nil {
			log.Printf("Session error: %v", err)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		session.Options.MaxAge = -1
		if err = session.Save(r, w); err != nil {
			log.Printf("Session error: %v", err)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		w.Header().Set("HX-Redirect", "/login")
		w.WriteHeader(http.StatusSeeOther)
	}
}
