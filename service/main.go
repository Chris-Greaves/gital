package main

import (
	"log/slog"
	"os"
	"strings"

	"github.com/Chris-Greaves/gital/core"
)

func main() {
	// Setup logging
	logLevel := &slog.LevelVar{} // INFO by Default
	opts := &slog.HandlerOptions{
		Level: logLevel,
	}

	slog.SetDefault(slog.New(slog.NewJSONHandler(os.Stdout, opts))) // Set as default Logger

	slog.Info("Starting Gital Service")
	config, err := core.LoadConfig()
	if err != nil {
		slog.Error("Error loading config.", slog.Any("error", err))
		panic("Config was not loaded correctly")
	}

	// Set Log Level now we have access to the config
	SetLogLevel(config, logLevel)

	if !config.IsSet("directories") {
		slog.Warn("No directories to scan, stopping!")
		return
	}
}

func SetLogLevel(config *core.Config, logLevel *slog.LevelVar) {
	switch strings.ToUpper(config.GetString("MinLogLevel")) {
	case "DEBUG":
		logLevel.Set(slog.LevelDebug)
	case "INFO":
		logLevel.Set(slog.LevelInfo)
	case "WARNING", "WARN":
		logLevel.Set(slog.LevelWarn)
	case "ERROR", "ERR":
		logLevel.Set(slog.LevelError)
	default:
		logLevel.Set(slog.LevelInfo)
	}
}
