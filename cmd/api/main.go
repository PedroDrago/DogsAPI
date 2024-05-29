package main

import (
	_ "github.com/lib/pq"
)

func main() {
	app := newApplication()

	app.infoLog.Printf("HTTP server listening on port %s", app.srv.Addr)
	err := app.srv.ListenAndServe()
	app.errorLog.Fatal(err)
}
