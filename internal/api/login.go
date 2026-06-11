package api

import (
	"errors"
	"log"
	"net/http"

	"github.com/gorilla/sessions"

	"github.com/IeemeliK/kuvagalleria/internal/service"
	"github.com/IeemeliK/kuvagalleria/internal/templates"
)

const (
	InvalidCredentialsError = "Väärä käyttäjänimi tai salasana"
	MissingCredentialsError = "Käyttäjänimi ja salasana vaaditaan"
)

func LoginHandler(store *sessions.CookieStore, auth *service.AuthService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			handleLoginGet(w, r, store)
		case http.MethodPost:
			handleLoginPost(w, r, store, auth)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	}
}

func handleLoginGet(w http.ResponseWriter, r *http.Request, store *sessions.CookieStore) {
	session, err := store.Get(r, "session-name")
	if err != nil {
		log.Printf("Session error: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	if session.Values["user_id"] != nil {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	renderLogin(w, r, "")
}

func handleLoginPost(w http.ResponseWriter, r *http.Request, store *sessions.CookieStore, auth *service.AuthService) {
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

	userID, err := auth.Authenticate(r.Context(), username, password)
	if err != nil {
		if errors.Is(err, service.ErrInvalidCredentials) {
			renderLogin(w, r, InvalidCredentialsError)
			return
		}
		log.Printf("Authentication error: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	if err := saveUserSession(w, r, store, userID); err != nil {
		log.Printf("Session error: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("HX-Redirect", "/")
	w.WriteHeader(http.StatusOK)
}

func saveUserSession(w http.ResponseWriter, r *http.Request, store *sessions.CookieStore, userID string) error {
	session, err := store.Get(r, "session-name")
	if err != nil {
		return err
	}

	const sessionMaxAgeSeconds = 30 * 24 * 60 * 60
	if session.Options == nil {
		session.Options = &sessions.Options{}
	}

	session.Options.MaxAge = sessionMaxAgeSeconds
	session.Values["user_id"] = userID

	return session.Save(r, w)
}

func renderLogin(w http.ResponseWriter, r *http.Request, errorMsg string) {
	data := PageData{
		LoggedIn: false,
		Error:    errorMsg,
	}

	layout := ""
	if r.Header.Get("HX-Request") == "true" {
		layout = "login_form"
	}

	if err := templates.Render(w, "login.html", layout, data); err != nil {
		log.Printf("Template render error: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
}
