package middleware

import (
	"context"
	"testing"
)

func TestUserIDFromContext(t *testing.T) {
	tests := []struct {
		name    string
		ctx     context.Context
		wantID  string
		wantOK  bool
	}{
		{
			name:   "empty context",
			ctx:    context.Background(),
			wantOK: false,
		},
		{
			name:    "with valid user_id",
			ctx:     context.WithValue(context.Background(), UserIDKey, "user-abc-123"),
			wantID:  "user-abc-123",
			wantOK:  true,
		},
		{
			name:   "with wrong type",
			ctx:    context.WithValue(context.Background(), UserIDKey, 42),
			wantOK: false,
		},
		{
			name:   "with empty string",
			ctx:    context.WithValue(context.Background(), UserIDKey, ""),
			wantID: "",
			wantOK: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			id, ok := UserIDFromContext(tt.ctx)
			if ok != tt.wantOK {
				t.Errorf("UserIDFromContext() ok = %v, want %v", ok, tt.wantOK)
			}
			if id != tt.wantID {
				t.Errorf("UserIDFromContext() id = %q, want %q", id, tt.wantID)
			}
		})
	}
}

func TestUsernameFromContext(t *testing.T) {
	tests := []struct {
		name       string
		ctx        context.Context
		wantUser   string
		wantOK     bool
	}{
		{
			name:     "empty context",
			ctx:      context.Background(),
			wantOK:   false,
		},
		{
			name:     "with valid username",
			ctx:      context.WithValue(context.Background(), UsernameKey, "alice"),
			wantUser: "alice",
			wantOK:   true,
		},
		{
			name:     "with wrong type",
			ctx:      context.WithValue(context.Background(), UsernameKey, true),
			wantOK:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			user, ok := UsernameFromContext(tt.ctx)
			if ok != tt.wantOK {
				t.Errorf("UsernameFromContext() ok = %v, want %v", ok, tt.wantOK)
			}
			if user != tt.wantUser {
				t.Errorf("UsernameFromContext() user = %q, want %q", user, tt.wantUser)
			}
		})
	}
}
