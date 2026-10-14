package router

import (
	"net/http"

	"github.com/IeemeliK/kuvagalleria/internal/handlers"
	"github.com/IeemeliK/kuvagalleria/internal/middleware"
	"github.com/IeemeliK/kuvagalleria/internal/service"
	"github.com/IeemeliK/kuvagalleria/web"
)

func New(authSvc *service.AuthService, authMdw *middleware.Authenticator) http.Handler {
	mux := http.NewServeMux()

	mux.Handle("GET /static/", http.StripPrefix("/static/", http.FileServerFS(web.StaticFS())))

	mux.HandleFunc("POST /logout", handlers.LogoutHandler(authSvc))
	mux.HandleFunc("/login", handlers.LoginHandler(authSvc))

	protected := http.NewServeMux()
	protected.HandleFunc("GET /", handlers.HomeHandler())
	mux.Handle("/", authMdw.Middleware(protected))

	return middleware.LoggingMiddleware(mux)
}
