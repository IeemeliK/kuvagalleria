package config

import (
	"log/slog"
	"os"
)

func (c *Config) InitLogger() {
	level := new(slog.LevelVar)
	level.Set(slog.LevelInfo)

	var l slog.Level
	if err := l.UnmarshalText([]byte(c.LogLevel)); err != nil {
		slog.Warn("invalid log level, defaulting to info", "value", c.LogLevel)
	} else {
		level.Set(l)
	}

	slog.SetDefault(slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{
		Level: level,
	})))
}
