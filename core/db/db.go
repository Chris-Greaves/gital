package db

import (
	"context"
	"database/sql"
	"errors"
	"os"
	"time"

	"github.com/Chris-Greaves/gital/core"
	"github.com/golang-migrate/migrate/v4"
	_ "modernc.org/sqlite"
)

const upsertRepoScript string = "INSERT INTO Repositories(name,path,current_branch,last_updated) VALUES(?,?,?,?) ON CONFLICT(path) DO UPDATE SET name=excluded.name,current_branch=excluded.current_branch,last_updated=excluded.last_updated;"

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
	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return Database{}, err
	}

	return Database{
		config: config,
		db:     db,
	}, nil
}

func (d Database) Close() {
	d.db.Close()
}

func (d Database) UpsertRepo(ctx context.Context, name, path, branch string) error {
	_, err := d.db.ExecContext(ctx, upsertRepoScript, name, path, branch, time.Now().Unix())

	return err
}
