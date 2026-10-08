package api

import (
	"errors"
	"log/slog"
	"net/http"

	"github.com/IeemeliK/kuvagalleria/internal/service"
	"github.com/IeemeliK/kuvagalleria/internal/templates"
)

const (
	InvalidCredentialsError = "Väärä käyttäjänimi tai salasana"
	MissingCredentialsError = "Käyttäjänimi ja salasana vaaditaan"
)

type loginPageData struct {
	Error      string
	HeaderData HeaderData
}

func LoginHandler(authSvc *service.AuthService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			handleLoginGet(w, r, authSvc)
		case http.MethodPost:
			handleLoginPost(w, r, authSvc)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	}
}

func handleLoginGet(w http.ResponseWriter, r *http.Request, authSvc *service.AuthService) {
	if _, ok := authSvc.IsAuthenticated(r); ok {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	renderLogin(w, r, "")
}

func handleLoginPost(w http.ResponseWriter, r *http.Request, authSvc *service.AuthService) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	username := r.FormValue("username")
	password := r.FormValue("password")

	if username == "" || password == "" {
		renderLogin(w, r, MissingCredentialsError)
		return
	}

	userID, err := authSvc.Authenticate(r.Context(), username, password)
	if err != nil {
		if errors.Is(err, service.ErrInvalidCredentials) {
			renderLogin(w, r, InvalidCredentialsError)
			return
		}
		slog.Error("authentication error", "error", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	if err := authSvc.SaveSession(w, r, userID); err != nil {
		slog.Error("saving session", "error", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("HX-Redirect", "/")
	w.WriteHeader(http.StatusOK)
}

func renderLogin(w http.ResponseWriter, r *http.Request, errorMsg string) {
	data := loginPageData{
		Error: errorMsg,
		HeaderData: HeaderData{
			LoggedIn: false,
		},
	}

	layout := ""
	if r.Header.Get("HX-Request") == "true" {
		layout = "login_form"
	}

	if err := templates.Render(w, "login.html", layout, data); err != nil {
		slog.Error("rendering login template", "error", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
}
