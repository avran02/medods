package logger

import (
	"log/slog"
	"os"

	"github.com/avran02/medods/config"
)

func Setup(config config.ServerConfig) {
	var ll slog.Leveler
	var isDefaultLogLevel bool
	switch config.LogLevel {
	case "DEBUG":
		ll = slog.LevelDebug
	case "INFO":
		ll = slog.LevelInfo
	case "WARN":
		ll = slog.LevelWarn
	case "ERROR":
		ll = slog.LevelError
	default:
		isDefaultLogLevel = true
		ll = slog.LevelInfo
	}

	log := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: ll,
	}))

	slog.SetDefault(log)

	if isDefaultLogLevel {
		slog.Warn("Logger using default value", "log level", ll)
		return
	}
	slog.Info("Logger", "log level", ll)
}
