package logger

import (
	"log/slog"
	"os"
)

// InitLogger initializes the global structured logger based on the environment.
func InitLogger(env, serviceName string) {
	var handler slog.Handler

	if env == "production" {
		// Production: Structured JSON format for machine parsing
		handler = slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
			Level: slog.LevelInfo,
		})
	} else {
		// Development: Pretty-printed text format for human readability
		handler = slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
			Level: slog.LevelDebug,
		})
	}

	// Create logger with default attributes like service name
	logger := slog.New(handler).With(
		slog.String("service", serviceName),
		slog.String("env", env),
	)

	// Set as the default global logger
	slog.SetDefault(logger)
}
