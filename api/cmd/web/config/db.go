package config

import (
	"database/sql"
	"errors"
	"time"

	// MySQL DB driver
	_ "github.com/go-sql-driver/mysql"
)

// ErrCannotConnectDB is returned when a connection to the database cannot be made.
var ErrCannotConnectDB = errors.New("cannot connect to the given database")

// OpenDB opens a connection to the MySQL database indicated by the dsn. If
// the connection is unsuccessful, it will be retried a number of times
// given by dbRetries.
func (env *Env) OpenDB(dsn string, maxRetries int) error {
	waitPeriod := 1 * time.Second

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return err
	}

	for numTries := 1; numTries <= maxRetries; numTries++ {
		if err = db.Ping(); err == nil {
			env.DB = db
			env.InfoLog.Printf("Connected to database!")
			return nil
		}

		env.InfoLog.Printf("Failled to connect to the DB %d times. Waiting %v...", numTries, waitPeriod)
		time.Sleep(waitPeriod)
		waitPeriod *= 2
	}

	return ErrCannotConnectDB
}
