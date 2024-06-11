package main

import (
	"net/http"
)

func (app *application) route() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /v1/status", app.statusHandler)
	mux.HandleFunc("POST /v1/users", app.createUserHandler)
	mux.HandleFunc("GET /v1/users/{id}", app.viewUserHandler)
	mux.HandleFunc("GET /v1/users/{id}/dogs", app.listUserDogsHandler)
	mux.HandleFunc("PUT /v1/users/{id}", app.updateUserHandler)
	mux.HandleFunc("DELETE /v1/users/{id}", app.deleteUserHandler)
	// mux.HandleFunc("DELETE /v1/users/{id}/dogs", app.deleteUserHandler)
	mux.HandleFunc("POST /v1/dogs", app.createDogHandler)
	mux.HandleFunc("GET /v1/dogs/{id}", app.viewDogHandler)
	// mux.HandleFunc("PUT /v1/dogs/{id}", app.updateDogHandler)
	// mux.HandleFunc("DELETE /v1/dogs/{id}", app.deleteDogHandler)

	return mux
}
