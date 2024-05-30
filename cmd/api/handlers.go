package main

import (
	"net/http"
)

func (app *application) statusHandler(writer http.ResponseWriter, req *http.Request) {
	status := map[string]any{
		"status": "available",
		"system_info": map[string]string{
			"version":     app.cfg.ver,
			"environment": app.cfg.env,
		},
	}
	err := writeJSON(writer, http.StatusOK, status, nil)
	if err != nil {
		app.errorLog.Println(err)
		writer.WriteHeader(http.StatusInternalServerError)
	}
}

func (app *application) createUserHandler(writer http.ResponseWriter, req *http.Request) {
}

func (app *application) viewUserHandler(writer http.ResponseWriter, req *http.Request) {
}

func (app *application) updateUserHandler(writer http.ResponseWriter, req *http.Request) {
}

func (app *application) deleteUserHandler(writer http.ResponseWriter, req *http.Request) {
}
