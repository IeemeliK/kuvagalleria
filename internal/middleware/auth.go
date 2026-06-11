// Package middleware
package middleware

import (
	"context"
	"database/sql"
	"errors"
	"log"
	"net/http"
	"strings"
)

func (a *Authenticator) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/login" || r.URL.Path == "/logout" || strings.HasPrefix(r.URL.Path, "/static/") {
			next.ServeHTTP(w, r)
			return
		}

		session, err := a.Store.Get(r, "session-name")
		if err != nil {
			log.Printf("Session decode error (treating as unauthenticated): %v", err)
			session.Values = make(map[any]any)
		}

		userID, ok := session.Values["user_id"].(string)
		if !ok || userID == "" {
			a.unauthorized(w, r)
			return
		}

		var username string
		err = a.DB.QueryRowContext(r.Context(),
			"SELECT username FROM users WHERE user_id = $1",
			userID,
		).Scan(&username)
		if errors.Is(err, sql.ErrNoRows) {
			a.unauthorized(w, r)
			return
		}
		if err != nil {
			log.Printf("Database error: %v", err)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
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
