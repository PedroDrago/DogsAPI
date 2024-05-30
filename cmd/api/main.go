package main

import (
	_ "github.com/lib/pq"
)

const version = "1.0.0"

func main() {
	app := newApplication()
	defer app.db.Close()

	app.infoLog.Printf("HTTP server listening on port %s", app.srv.Addr)
	err := app.srv.ListenAndServe()
	app.errorLog.Fatal(err)
}
