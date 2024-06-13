package main

import (
	"net/http"
)

func (app *application) route() *http.ServeMux {
	router := http.NewServeMux()
	router.HandleFunc("GET /v1/status", app.statusHandler)
	router.HandleFunc("POST /v1/users", app.createUserHandler)
	router.HandleFunc("GET /v1/users/{id}", app.viewUserHandler)
	router.HandleFunc("GET /v1/users/{id}/dogs", app.listUserDogsHandler)
	router.HandleFunc("PUT /v1/users/{id}", app.updateUserHandler)
	router.HandleFunc("DELETE /v1/users/{id}", app.deleteUserHandler)
	// mux.HandleFunc("DELETE /v1/users/{id}/dogs", app.deleteUserHandler)
	router.HandleFunc("POST /v1/dogs", app.createDogHandler)
	router.HandleFunc("GET /v1/dogs/{id}", app.viewDogHandler)
	// mux.HandleFunc("PUT /v1/dogs/{id}", app.updateDogHandler)
	// mux.HandleFunc("DELETE /v1/dogs/{id}", app.deleteDogHandler)

	return router
}
