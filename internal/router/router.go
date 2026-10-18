package router

import (
	"net/http"

	"github.com/IeemeliK/kuvagalleria/internal/handlers"
	"github.com/IeemeliK/kuvagalleria/internal/middleware"
	"github.com/IeemeliK/kuvagalleria/internal/service"
)

func New(authSvc *service.AuthService, authMdw *middleware.Authenticator) http.Handler {
	mux := http.NewServeMux()

	mux.Handle("GET /static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./public/static"))))

	// mux.Handle("POST /logout", handlers.LogoutHandler(authSvc))
	mux.Handle("/login", handlers.NewLoginHandler())

	protected := http.NewServeMux()
	protected.Handle("GET /", handlers.NewHomeHandler())
	mux.Handle("/", protected)

	return middleware.LoggingMiddleware(mux)
}
