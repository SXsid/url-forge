package db

import (
	"database/sql"
	"io/fs"

	"github.com/pressly/goose/v3"
)

func RunMigraton(dsn string, migrationFS fs.FS) error {
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return err
	}
	goose.SetBaseFS(migrationFS)
	if err := goose.SetDialect("postgres"); err != nil {
		return err
	}
	if err := goose.Up(db, "."); err != nil {
		return err
	}
	return err
}
