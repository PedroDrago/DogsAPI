package main

import (
	"net/http"
)

func (app *application) route() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/status", app.statusHandler)
	return mux
}
