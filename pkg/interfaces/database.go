package interfaces

import "database/sql"

type DB interface {
	Open() error
	Close() error
	Dbh() *sql.DB
}
