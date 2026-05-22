package db

import (
	"database/sql"
	"database/sql/driver"
	"net"
	"time"

	"github.com/lib/pq"
)

type ipv4Dialer struct{}

func (d *ipv4Dialer) Dial(network, address string) (net.Conn, error) {
	return net.Dial("tcp4", address)
}

func (d *ipv4Dialer) DialTimeout(network, address string, timeout time.Duration) (net.Conn, error) {
	return net.DialTimeout("tcp4", address, timeout)
}

type IPv4Driver struct{}

func (d *IPv4Driver) Open(name string) (driver.Conn, error) {
	return pq.DialOpen(&ipv4Dialer{}, name)
}

func init() {
	sql.Register("postgres-ipv4", &IPv4Driver{})
}

// NewConnection opens a new database connection pool.
// The DSN should be in the format: "postgres://user:pass@host:port/dbname?sslmode=disable"
func NewConnection(dsn string) (*sql.DB, error) {
	db, err := sql.Open("postgres-ipv4", dsn)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	// Connection Pool Settings
	db.SetMaxOpenConns(15)
	db.SetMaxIdleConns(15)
	db.SetConnMaxLifetime(0) // Connections are reused forever if not closed by the DB

	return db, nil
}
