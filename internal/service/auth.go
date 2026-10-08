package service

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"net/http"

	"github.com/gorilla/sessions"
	"golang.org/x/crypto/bcrypt"

	"github.com/IeemeliK/kuvagalleria/internal/repository"
)

var ErrInvalidCredentials = errors.New("invalid username or password")

type AuthService struct {
	DB    *sql.DB
	Store *sessions.CookieStore
}

func NewAuthService(db *sql.DB, store *sessions.CookieStore) *AuthService {
	return &AuthService{DB: db, Store: store}
}

func (s *AuthService) Authenticate(ctx context.Context, username, password string) (string, error) {
	hashedPassword, userID, err := repository.GetUserByUsername(ctx, s.DB, username)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", ErrInvalidCredentials
		}
		return "", fmt.Errorf("authenticate: %w", err)
	}

	if err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password)); err != nil {
		return "", ErrInvalidCredentials
	}

	return userID, nil
}

func (s *AuthService) SaveSession(w http.ResponseWriter, r *http.Request, userID string) error {
	session, err := s.Store.Get(r, "session-name")
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

func (s *AuthService) IsAuthenticated(r *http.Request) (string, bool) {
	session, err := s.Store.Get(r, "session-name")
	if err != nil {
		return "", false
	}

	userID, ok := session.Values["user_id"].(string)
	return userID, ok
}

func (s *AuthService) ClearSession(w http.ResponseWriter, r *http.Request) error {
	session, err := s.Store.Get(r, "session-name")
	if err != nil {
		return err
	}

	session.Options.MaxAge = -1
	return session.Save(r, w)
}

func (s *AuthService) GetCurrentUser(r *http.Request) (userID, username string, ok bool) {
	userID, ok = s.IsAuthenticated(r)
	if !ok {
		return "", "", false
	}

	err := s.DB.QueryRowContext(
		r.Context(),
		"SELECT username FROM users WHERE user_id = $1",
		userID,
	).Scan(&username)
	if err != nil {
		return "", "", false
	}

	return userID, username, true
}
