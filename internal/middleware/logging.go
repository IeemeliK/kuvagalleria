package middleware

import (
	"log/slog"
	"net/http"
	"time"
)

func LoggingMiddleware(next http.Handler) http.Handler {
	f := func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		slog.Info("completed", "method", r.Method, "path", r.URL.Path, "duration", time.Since(start))
	}
	return http.HandlerFunc(f)
}
