package main

import (
	"crypto/tls"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/erik-jenkins/goblog/api/cmd/web/config"
)

func main() {
	addr := flag.String("addr", ":4000", "HTTP network address")
	dbHost := flag.String("dbhost", "mysql", "Hostname for MySQL server")
	dbPort := flag.String("dbport", "3306", "Port for MySQL server")
	dbUser := flag.String("dbuser", "blogger", "Username for MySQL user")
	dbPass := flag.String("dbpass", "password", "Password for MySQL user")
	dbRetries := flag.Int("dbretries", 10, "Number of retry attempts to connect to the DB")
	flag.Parse()

	// env holds app dependencies
	env := &config.Env{
		InfoLog:  log.New(os.Stdout, "INFO:\t", log.LstdFlags),
		ErrorLog: log.New(os.Stderr, "ERROR:\t", log.LstdFlags|log.Lshortfile),
	}
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/goblog?parseTime=true", *dbUser, *dbPass, *dbHost, *dbPort)
	err := env.OpenDB(dsn, *dbRetries)
	if err != nil {
		env.ErrorLog.Fatal(err)
	}

	// tlsConfig contains the non-default TLS settings that we want to use
	tlsConfig := &tls.Config{
		PreferServerCipherSuites: true,
		CurvePreferences:         []tls.CurveID{tls.X25519, tls.CurveP256},
	}

	// configure server
	srv := &http.Server{
		Addr:         *addr,
		ErrorLog:     env.ErrorLog,
		Handler:      routes(env),
		TLSConfig:    tlsConfig,
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	env.InfoLog.Printf("Starting server on %s\n", *addr)
	err = srv.ListenAndServeTLS("./tls/cert.pem", "./tls/key.pem")
	env.ErrorLog.Fatal(err)
}
