package config

import (
	"fmt"
	"os"
)

type Config struct {
	Database DatabaseConfig
	Session  SessionConfig
}

type DatabaseConfig struct {
	User     string
	Password string
	Host     string
	Port     string
	DBName   string
}

type SessionConfig struct {
	Secret string
}

func Load() (Config, error) {
	required := map[string]string{
		"POSTGRES_USER":       "",
		"POSTGRES_PASSWORD":   "",
		"HOST":                "",
		"PORT":                "",
		"POSTGRES_DB":         "",
		"COOKIESTORE_SECRET":  "",
	}

	for k := range required {
		v := os.Getenv(k)
		if v == "" {
			return Config{}, fmt.Errorf("required env var %q is not set or empty", k)
		}
		required[k] = v
	}

	return Config{
		Database: DatabaseConfig{
			User:     required["POSTGRES_USER"],
			Password: required["POSTGRES_PASSWORD"],
			Host:     required["HOST"],
			Port:     required["PORT"],
			DBName:   required["POSTGRES_DB"],
		},
		Session: SessionConfig{
			Secret: required["COOKIESTORE_SECRET"],
		},
	}, nil
}
