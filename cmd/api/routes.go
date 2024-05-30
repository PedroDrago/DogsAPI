package main

import (
	"net/http"
)

func (app *application) route() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /v1/status", app.statusHandler)
	mux.HandleFunc("GET /v1/users", app.indexUserHandler)
	mux.HandleFunc("POST /v1/users", app.createUserHandler)
	mux.HandleFunc("GET /v1/users/{id}", app.viewUserHandler)
	mux.HandleFunc("PUT /v1/users/{id}", app.updateUserHandler)
	mux.HandleFunc("DELETE /v1/users/{id}", app.deleteUserHandler)

	return mux
}
