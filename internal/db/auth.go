package db

import (
	"database/sql"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

var ErrInvalidCredentials = errors.New("invalid username or password")

func AuthenticateUser(database *sql.DB, username, password string) (string, error) {
	var hashedPassword, userID string

	err := database.QueryRow("SELECT password_hash, user_id FROM users WHERE username = $1", username).Scan(&hashedPassword, &userID)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", ErrInvalidCredentials
		}
		return "", err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password)); err != nil {
		return "", ErrInvalidCredentials
	}

	return userID, nil
}
