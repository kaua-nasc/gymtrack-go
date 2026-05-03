package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

// NewConnection opens a new database connection pool.
// The DSN should be in the format: "postgres://user:pass@host:port/dbname?sslmode=disable"
func NewConnection(dsn string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
