// Package middleware
package middleware

import (
	"context"
	"log"
	"net/http"
)

func (a *Authenticator) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/login" || r.URL.Path == "/logout" || r.URL.Path == "/static/" {
			next.ServeHTTP(w, r)
			return
		}

		session, err := a.Store.Get(r, "session-name")
		if err != nil {
			log.Printf("Session error: %v", err)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		userID, ok := session.Values["user_id"].(string)
		if !ok || userID == "" {
			a.unauthorized(w, r)
			return
		}

		var (
			exists   bool
			username string
		)
		err = a.DB.QueryRow(
			"SELECT EXISTS(SELECT 1 FROM users WHERE user_id = $1), username FROM users WHERE user_id = $1",
			userID,
		).Scan(&exists, &username)
		if err != nil {
			log.Printf("Database error: %v", err)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
		if !exists {
			a.unauthorized(w, r)
			return
		}

		ctx := context.WithValue(r.Context(), UserIDKey, userID)
		ctx = context.WithValue(ctx, UsernameKey, username)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (a *Authenticator) unauthorized(w http.ResponseWriter, r *http.Request) {
	if isAPIRequest(r) {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}

func isAPIRequest(r *http.Request) bool {
	return r.Header.Get("HX-Request") == "true" || r.Header.Get("Accept") == "application/json"
}
