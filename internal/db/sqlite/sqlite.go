package sqlite

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type SqliteDB struct {
	conn *sql.DB
}

func New(dsn string) (*SqliteDB, error) {
	c, err := sql.Open("sqlite3", dsn)
	if err != nil {
		return nil, err
	}
	return &SqliteDB{conn: c}, nil
}

func (s *SqliteDB) Conn() *sql.DB {
	return s.conn
}
