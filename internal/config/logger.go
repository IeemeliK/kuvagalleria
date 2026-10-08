package config

import (
	"log/slog"
	"os"
)

func InitLogger(cfg Config) {
	level := new(slog.LevelVar)
	level.Set(slog.LevelInfo)

	var l slog.Level
	if err := l.UnmarshalText([]byte(cfg.LogLevel)); err != nil {
		slog.Warn("invalid log level, defaulting to info", "value", cfg.LogLevel)
	} else {
		level.Set(l)
	}

	slog.SetDefault(slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{
		Level: level,
	})))
}
