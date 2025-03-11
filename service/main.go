package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"strings"
	"syscall"

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

	// Create a context that can be canceled
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel() // Ensure cancel is called at the end to clean up

	// Channel to listen for when shutting down
	done := make(chan bool)

	// Channel to listen for system signals (e.g., Ctrl+C)
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	// Run Scheduler
	scheduler := CreateScheduler(config, ctx)
	go func() {
		scheduler.Run(done)
		done <- true
	}()

	// Wait for an interrupt signal to initiate graceful shutdown
	<-sigChan

	// Handle shutdown signal (Ctrl+C or SIGTERM)
	slog.Info("Received shutdown signal, attempting to shut down gracefully.")

	// Cancel the context to notify all goroutines to stop
	cancel()

	// Wait for Scheduler to finish what it's doing
	<-done

	// Final cleanup before exiting
	slog.Info("Application shutdown complete.")

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
