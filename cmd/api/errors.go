package main

import (
	"encoding/json"
	"net/http"
)

func (app *application) responseInternalServerError(writer http.ResponseWriter, err error) {
	app.errorLog.Println(err)
	http.Error(writer, "Internal Server Error", http.StatusInternalServerError)
}

func (app *application) responseBadRequest(writer http.ResponseWriter, errMsg string) {
	js, err := json.Marshal(map[string]string{
		"error": errMsg,
	})
	if err != nil {
		app.responseInternalServerError(writer, err)
		return
	}
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusBadRequest)
	writer.Write(js)
}
