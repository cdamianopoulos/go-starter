package dbconn

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
	"time"

	"github.com/sirupsen/logrus"
)

// Open returns an instance of sql.DB to use for interacting with the database.
func Open(c driver.Connector, maxConnLifetime time.Duration) (db *sql.DB, err error) {
	db = sql.OpenDB(c)
	db.SetConnMaxLifetime(maxConnLifetime)

	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("error pinging database: %w", err)
	}

	logrus.Info("Successfully established connection to database")
	return
}
