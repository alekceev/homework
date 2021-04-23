package database

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type DB struct {
	db *sql.DB
}

func (d *DB) Open() error {
	sqlite, err := sql.Open("sqlite3", "file:/tmp/catalog.db?_mutex=full&_cslike=false")
	if err != nil {
		return err
	}
	_, err = sqlite.Exec(init_query())

	d.db = sqlite
	return nil
}

func (d *DB) Close() error {
	return d.db.Close()
}

func (d *DB) Dbh() *sql.DB {
	return d.db
}

func init_query() string {
	return `
CREATE TABLE IF NOT EXISTS items (
	id          INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
	name        VARCHAR(255) COLLATE NOCASE NOT NULL,
	description TEXT COLLATE NOCASE NOT NULL,
	price       NUMERIC CHECK (price > 0) NOT NULL
);
`
}
