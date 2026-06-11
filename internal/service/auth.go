package service

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"golang.org/x/crypto/bcrypt"

	"github.com/IeemeliK/kuvagalleria/internal/repository"
)

var ErrInvalidCredentials = errors.New("invalid username or password")

type AuthService struct {
	DB *sql.DB
}

func NewAuthService(db *sql.DB) *AuthService {
	return &AuthService{DB: db}
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
