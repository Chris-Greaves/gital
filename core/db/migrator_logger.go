package db

import (
	"fmt"
	"log/slog"
)

type MigrateLogger struct {
	l       *slog.Logger
	verbose bool
}

func NewMigrateLogger(logger *slog.Logger, verbose bool) MigrateLogger {
	return MigrateLogger{
		l:       logger,
		verbose: verbose,
	}
}

func (ml MigrateLogger) Printf(format string, v ...interface{}) {
	ml.l.Info(fmt.Sprintf(format, v...))
}

func (ml MigrateLogger) Verbose() bool {
	return ml.verbose
}
