package db

import (
	"database/sql"

	"github.com/Chris-Greaves/gital/core"
	_ "github.com/glebarez/go-sqlite"
)

type Database struct {
	config *core.Config
	db     *sql.DB
}

func CreateAndOpenTheDatabase(config *core.Config) (Database, error) {
	db_path := config.GetString("database_location")
	db, err := sql.Open("sqlite", db_path+"/gital.db")
	if err != nil {
		return Database{}, err
	}

	// Apply any Schema changes

	return Database{
		config: config,
		db:     db,
	}, nil
}
