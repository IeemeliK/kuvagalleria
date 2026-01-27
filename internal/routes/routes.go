// Package routes route handlers
package routes

import (
	"database/sql"
	"html/template"
	"log"
	"net/http"

	"github.com/IeemeliK/kuvagalleria/internal"
	"github.com/IeemeliK/kuvagalleria/internal/db"
	"github.com/gorilla/sessions"
	"golang.org/x/crypto/bcrypt"
)

type PageData struct {
	LoggedIn bool
	Error    string
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	data := PageData{LoggedIn: false}
	tmpl := template.Must(template.ParseFS(internal.Templates, "*.html"))

	if err := tmpl.Execute(w, data); err != nil {
		log.Printf("template execute error: %v", err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
}

func LoginHandler(w http.ResponseWriter, r *http.Request, store *sessions.CookieStore) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	if err := r.ParseForm(); err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	username := r.FormValue("username")
	password := r.FormValue("password")

	if username == "" || password == "" {
		renderLoginError(w, "Username and password are required")
		return
	}

	database := db.CreateConnection()
	defer database.Close()

	var hashedPassword string
	err := database.QueryRow("SELECT password_hash FROM users WHERE username = $1", username).Scan(&hashedPassword)
	if err != nil {
		if err == sql.ErrNoRows {
			renderLoginError(w, "Invalid username or password")
			return
		}
		log.Printf("Database error: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password)); err != nil {
		renderLoginError(w, "Invalid username or password")
		return
	}

	// Successful login
	session, err := store.Get(r, "session-name")
	if err != nil {
		log.Printf("Session error: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	session.Values["user_id"] = username // Or actual ID if you have one
	session.Save(r, w)

	// HTMX redirect
	w.Header().Set("HX-Redirect", "/")
	w.WriteHeader(http.StatusOK)
}

func renderLoginError(w http.ResponseWriter, errorMsg string) {
	data := PageData{
		LoggedIn: false,
		Error:    errorMsg,
	}
	tmpl := template.Must(template.ParseFS(internal.Templates, "login.html"))
	if err := tmpl.Execute(w, data); err != nil {
		log.Printf("Template error: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
}
