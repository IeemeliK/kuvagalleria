package routes

import (
	"database/sql"
	"html/template"
	"log"
	"net/http"

	"github.com/IeemeliK/kuvagalleria/internal"
	"github.com/gorilla/sessions"
	"golang.org/x/crypto/bcrypt"
)

func LoginHandler(w http.ResponseWriter, r *http.Request, store *sessions.CookieStore, database *sql.DB) {
	if r.Method != http.MethodPost && r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	session, err := store.Get(r, "session-name")
	if err != nil {
		log.Printf("Session error: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	if r.Method == http.MethodGet {
		if session.Values["user_id"] != nil {
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}

		tmpl := template.Must(template.ParseFS(internal.Templates, "base.html", "login.html"))
		if err := tmpl.Execute(w, nil); err != nil {
			log.Printf("template execute error: %v", err)
			http.Error(w, "internal server error", http.StatusInternalServerError)
		}
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

	var hashedPassword, userID string
	err = database.QueryRow("SELECT password_hash, user_id FROM users WHERE username = $1", username).Scan(&hashedPassword, &userID)
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
	session, err = store.Get(r, "session-name")
	if err != nil {
		log.Printf("Session error: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	const sessionMaxAgeSeconds = 30 * 24 * 60 * 60
	if session.Options == nil {
		session.Options = &sessions.Options{}
	}
	session.Options.MaxAge = sessionMaxAgeSeconds
	session.Values["user_id"] = userID
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
