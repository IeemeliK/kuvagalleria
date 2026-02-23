package middleware

import (
	"database/sql"

	"github.com/gorilla/sessions"
)

type contextKey string

const (
	UserIDKey   contextKey = "user_id"
	UsernameKey contextKey = "username"
)

type Authenticator struct {
	Store *sessions.CookieStore
	DB    *sql.DB
}
