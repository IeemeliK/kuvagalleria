package handlers

import (
	"net/http"

	"github.com/IeemeliK/kuvagalleria/components"
)

type HomeHandler struct{}

func NewHomeHandler() *HomeHandler {
	return &HomeHandler{}
}

func (h *HomeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.Get(w, r)
}

func (h *HomeHandler) Get(w http.ResponseWriter, r *http.Request) {
	h.View(w, r)
}

func (h *HomeHandler) View(w http.ResponseWriter, r *http.Request) {
	components.Page(components.Index()).Render(r.Context(), w)
}
