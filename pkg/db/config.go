package db

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"fmt"
	"time"

	"go-starter/separateRepos/dbconn"
)

type Config struct {
	DBName          string        `envconfig:"MYSQL_DATABASE" required:"true"`
	Username        string        `envconfig:"MYSQL_USERNAME" required:"true"`
	Password        string        `envconfig:"MYSQL_PASSWORD" required:"true"`
	HostURL         string        `envconfig:"MYSQL_HOST_URL" required:"true"`
	HostPort        uint16        `envconfig:"MYSQL_HOST_PORT" required:"true"`
	SSLCert         string        `envconfig:"RDS_SSL_CERT" required:"true"`
	MaxConnLifetime time.Duration `envconfig:"MAX_CONN_LIFETIME" default:"8h"`
	// Mysql keeps connections alive for 8 hours on server side by default. Setting this option will delete
	// connections from client's connection pool after 8 hours to avoid using already closed connections.
}

// OpenConnection returns an instance of sql.DB to use for interacting with the database.
func OpenConnection(config Config) (*sql.DB, error) {
	return dbconn.Open(config, config.MaxConnLifetime)
}

// Connect implements the driver.Connector interface.
func (c Config) Connect(ctx context.Context) (driver.Conn, error) {
	return nil, nil
}

// Driver implements the driver.Connector interface.
func (c Config) Driver() driver.Driver {
	return c
}

// Open implements the driver.Driver interface.
func (c Config) Open(_ string) (driver.Conn, error) {
	return nil, fmt.Errorf("open method is not supported on custom driver")
}
