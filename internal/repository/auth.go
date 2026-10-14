package repository

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

func GetUserByUsername(ctx context.Context, dbc *pgxpool.Pool, username string) (hashedPassword, userID string, err error) {
	err = dbc.QueryRow(ctx, "SELECT password_hash, user_id FROM users WHERE username = $1", username).Scan(&hashedPassword, &userID)
	if err != nil {
		return "", "", fmt.Errorf("query user by username: %w", err)
	}
	return hashedPassword, userID, nil
}
