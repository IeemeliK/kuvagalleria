package repository

import (
	"context"
	"database/sql"
	"fmt"
)

func GetUserByUsername(ctx context.Context, dbc *sql.DB, username string) (hashedPassword, userID string, err error) {
	err = dbc.QueryRowContext(ctx, "SELECT password_hash, user_id FROM users WHERE username = $1", username).Scan(&hashedPassword, &userID)
	if err != nil {
		return "", "", fmt.Errorf("query user by username: %w", err)
	}
	return hashedPassword, userID, nil
}
