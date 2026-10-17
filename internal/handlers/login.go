package handlers

import (
	"net/http"

	"github.com/IeemeliK/kuvagalleria/components"
)

const (
	InvalidCredentialsError = "Väärä käyttäjänimi tai salasana"
	MissingCredentialsError = "Käyttäjänimi ja salasana vaaditaan"
)

type LoginHandler struct{}

func NewLoginHandler() *LoginHandler {
	return &LoginHandler{}
}

func (h *LoginHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.Get(w, r)
	case http.MethodPost:
		h.Post(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (h *LoginHandler) Get(w http.ResponseWriter, r *http.Request) {
	h.View(w, r, viewProps{""})
}

func (h *LoginHandler) Post(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	username := r.FormValue("username")
	password := r.FormValue("password")

	if username == "" || password == "" {
		h.View(w, r, viewProps{MissingCredentialsError})
		return
	}

	if username == "tester" && password == "tester" {
		w.Header().Set("HX-Redirect", "/")
		w.WriteHeader(http.StatusOK)
	} else {
		h.View(w, r, viewProps{InvalidCredentialsError})
	}
}

type viewProps struct {
	errorMsg string
}

func (h *LoginHandler) View(w http.ResponseWriter, r *http.Request, props viewProps) {
	components.Page(components.Login(props.errorMsg)).Render(r.Context(), w)
}
