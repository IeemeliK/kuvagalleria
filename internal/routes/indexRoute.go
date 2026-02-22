// Package routes route handlers
package routes

import (
	"html/template"
	"log"
	"net/http"

	"github.com/IeemeliK/kuvagalleria/internal"
	"github.com/gorilla/sessions"
)

type PageData struct {
	LoggedIn bool
	Error    string
}

func HomeHandler(store *sessions.CookieStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		session, err := store.Get(r, "session-name")
		if err != nil {
			log.Printf("Session error: %v", err)
			http.Error(w, "internal server error", http.StatusInternalServerError)
			return
		}

		if session.Values["user_id"] == nil {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}

		tmpl := template.Must(template.ParseFS(internal.Templates, "base.html", "index.html"))

		if err := tmpl.Execute(w, nil); err != nil {
			log.Printf("template execute error: %v", err)
			http.Error(w, "internal server error", http.StatusInternalServerError)
			return
		}
	}
}
