package db

import (
	"database/sql"
	"embed"
	"log/slog"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/sqlite"
	"github.com/golang-migrate/migrate/v4/source/iofs"
)

//go:embed migrations/*.sql
var migrations embed.FS

func applyMigrations(db *sql.DB) error {
	// Use Embedded FS for migration SQL files
	sourceDriver, err := iofs.New(migrations, "migrations")
	if err != nil {
		return err
	}

	// Use the existing database instance for the migrations
	databaseDriver, err := sqlite.WithInstance(db, &sqlite.Config{})
	if err != nil {
		return err
	}

	// Create instance of Migrate using the custom drivers
	m, err := migrate.NewWithInstance("iofs", sourceDriver, "sqlite", databaseDriver)
	if err != nil {
		return err
	}

	// Set logger
	m.Log = NewMigrateLogger(slog.Default().WithGroup("migrator"), false)

	// Run the Migrate Instance and apply all the migrations
	err = m.Up()
	return err
}
