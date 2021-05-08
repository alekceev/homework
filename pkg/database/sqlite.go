package database

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

type DB struct {
	db *sqlx.DB
}

func Connect(host string) (*DB, error) {
	db := &DB{}
	err := db.Open(host)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func (d *DB) Open(host string) error {
	sqlite, err := sqlx.Open("sqlite3", host)
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

func (d *DB) Raw() *sqlx.DB {
	return d.db
}

func init_query() string {
	return `
CREATE TABLE IF NOT EXISTS items (
	id          INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
	name        VARCHAR(255) COLLATE NOCASE NOT NULL,
	description TEXT COLLATE NOCASE NOT NULL,
	number      VARCHAR(255) COLLATE NOCASE NOT NULL,
	category    VARCHAR(255) COLLATE NOCASE NOT NULL,
	price       NUMERIC CHECK (price > 0) NOT NULL,
	sale_price  NUMERIC CHECK (price > 0) NOT NULL,
	amount      NUMERIC NOT NULL
);
`
}
