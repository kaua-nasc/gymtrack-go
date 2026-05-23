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

	// Connection Pool Settings
	db.SetMaxOpenConns(5)
	db.SetMaxIdleConns(2)
	db.SetConnMaxLifetime(0) // Connections are reused forever if not closed by the DB

	return db, nil
}
