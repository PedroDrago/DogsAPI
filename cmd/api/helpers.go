package main

import (
	"encoding/json"
	"errors"
	"net/http"
)

func writeJSON(writer http.ResponseWriter, status int, data any, headers http.Header) error {
	js, err := json.Marshal(data)
	if err != nil {
		return errors.New("[TMP] Error when Marshalling JSON")
	}
	for key, value := range headers {
		writer.Header()[key] = value
	}
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(status)
	writer.Write(js)
	return nil
}

func readJSON(req *http.Request, dst any) error {
	dec := json.NewDecoder(req.Body)
	err := dec.Decode(dst)
	if err != nil {
		return err
	}
	return nil
}
