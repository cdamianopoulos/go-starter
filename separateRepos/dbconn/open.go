package dbconn

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
	"time"

	"github.com/sirupsen/logrus"
)

// Open returns an instance of sql.DB to use for interacting with the database.
// If maxConnLifetime <= 0, connections are not closed due to a connection's age.
// If maxOpenConns <= 0, then there is no limit on the number of open connections. The default is 0 (unlimited).
// If maxIdleConns <= 0, then no idle connections are retained. The default is 2.
func Open(c driver.Connector, maxConnLifetime time.Duration, maxOpenConns, maxIdleConns int) (db *sql.DB, err error) {
	db = sql.OpenDB(c)
	db.SetConnMaxLifetime(maxConnLifetime)
	db.SetMaxOpenConns(maxOpenConns)
	db.SetMaxIdleConns(maxIdleConns)

	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("error pinging database: %w", err)
	}

	logrus.Info("Successfully established connection to database")
	return
}
