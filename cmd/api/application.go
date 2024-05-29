package main

import (
	"database/sql"
	"flag"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/lib/pq"
)

func openDB() (*sql.DB, error) {
	db, err := sql.Open("postgres", os.Getenv("DOGSAPI_DB_DSN"))
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}

type application struct {
	db       *sql.DB
	srv      http.Server
	infoLog  *log.Logger
	errorLog *log.Logger
	warnLog  *log.Logger
}

func newApplication() *application {
	var err error
	app := application{
		infoLog:  log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime),
		warnLog:  log.New(os.Stdout, "WARN\t", log.Ldate|log.Ltime),
		errorLog: log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile),
	}
	app.db, err = openDB()
	if err != nil {
		app.errorLog.Fatal(err)
	}
	app.infoLog.Println("Connection with Postgres established")
	app.srv = http.Server{
		Addr:         ":" + *flag.String("port", "4000", "Http server Port"),
		Handler:      http.NewServeMux(),
		ReadTimeout:  time.Second * 5, // TODO: Adjust these fields
		IdleTimeout:  time.Second * 5, // TODO: Adjust these fields
		WriteTimeout: time.Second * 5, // TODO: Adjust these fields
	}
	return &app
}
