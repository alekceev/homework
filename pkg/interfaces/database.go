package interfaces

import "database/sql"

type DB interface {
	Open(host string) error
	Close() error
	Raw() *sql.DB
}
