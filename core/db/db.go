package db

import (
	"database/sql"
	"os"

	"github.com/Chris-Greaves/gital/core"
	_ "modernc.org/sqlite"
)

type Database struct {
	config *core.Config
	db     *sql.DB
}

func CreateAndOpenTheDatabase(config *core.Config) (Database, error) {
	db_path := config.GetString(core.KeyDatabasePath)

	err := os.MkdirAll(db_path, os.ModeAppend)
	if err != nil {
		return Database{}, err
	}

	db, err := sql.Open("sqlite", db_path+string(os.PathSeparator)+"gital.db")
	if err != nil {
		return Database{}, err
	}

	// Apply any Schema changes
	err = applyMigrations(db)
	if err != nil {
		return Database{}, err
	}

	return Database{
		config: config,
		db:     db,
	}, nil
}
