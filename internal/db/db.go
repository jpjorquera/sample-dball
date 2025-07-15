package db

import "database/sql"

type DB interface {
	Conn() *sql.DB
}
