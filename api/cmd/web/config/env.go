package config

import (
	"database/sql"
	"log"
)

// Env contains depedencies that are useful to the handlers, such as a DB
// connection and logging handles.
type Env struct {
	DB       *sql.DB
	InfoLog  *log.Logger
	ErrorLog *log.Logger
}
