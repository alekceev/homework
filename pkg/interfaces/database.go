package interfaces

import "github.com/jmoiron/sqlx"

type DB interface {
	Open(host string) error
	Close() error
	Raw() *sqlx.DB
}
