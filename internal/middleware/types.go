package middleware

import (
	"context"
	"database/sql"

	"github.com/gorilla/sessions"
)

type contextKey string

const (
	UserIDKey   contextKey = "user_id"
	UsernameKey contextKey = "username"
)

func UserIDFromContext(ctx context.Context) (string, bool) {
	v, ok := ctx.Value(UserIDKey).(string)
	return v, ok
}

func UsernameFromContext(ctx context.Context) (string, bool) {
	v, ok := ctx.Value(UsernameKey).(string)
	return v, ok
}

type Authenticator struct {
	Store *sessions.CookieStore
	DB    *sql.DB
}
