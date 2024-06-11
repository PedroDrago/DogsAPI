package main

import (
	"fmt"
	"net/http"
	"runtime/debug"
)

func (app *application) internalServerErrorResponse(writer http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	app.errorLog.Output(2, trace)
	app.errorResponse(writer, http.StatusInternalServerError, "Internal server error")
}

func (app *application) errorResponse(writer http.ResponseWriter, status int, message any) {
	env := Envelope{"error": message}
	err := writeJSON(writer, status, env, nil)
	if err != nil {
		app.errorLog.Println(err)
		writer.WriteHeader(http.StatusInternalServerError)
	}
}

func (app *application) badRequestResponse(writer http.ResponseWriter, err error) {
	app.errorResponse(writer, http.StatusBadRequest, err.Error())
}

func (app *application) validationErrorResponse(writer http.ResponseWriter, errors map[string]string) {
	app.errorResponse(writer, http.StatusUnprocessableEntity, errors)
}

func (app *application) notFoundResponse(writer http.ResponseWriter) {
	app.errorResponse(writer, http.StatusNotFound, "Resource could not be found")
}

func (app *application) editConflictResponse(writer http.ResponseWriter) {
	app.errorResponse(writer, http.StatusConflict, "edit conflict")
}
