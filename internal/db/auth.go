package db

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

var ErrInvalidCredentials = errors.New("invalid username or password")

func AuthenticateUser(ctx context.Context, dbc *sql.DB, username, password string) (string, error) {
	var hashedPassword, userID string

	err := dbc.QueryRowContext(ctx, "SELECT password_hash, user_id FROM users WHERE username = $1", username).Scan(&hashedPassword, &userID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", ErrInvalidCredentials
		}
		return "", fmt.Errorf("query user: %w", err)
	}

	if err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password)); err != nil {
		return "", ErrInvalidCredentials
	}

	return userID, nil
}
