package sqlite

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func SqliteInit() (db *sql.DB, err error) {
	db, err = sql.Open("sqlite3", "file:/tmp/catalog.db?_mutex=full&_cslike=false")
	if err != nil {
		return
	}
	_, err = db.Exec(init_query())
	return
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
