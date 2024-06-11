package main

import (
	"crypto/sha256"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"golang.org/x/crypto/bcrypt"
)

type Envelope map[string]any

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

func getIDParam(req *http.Request) (int64, error) {
	id := req.PathValue("id")
	return strconv.ParseInt(id, 10, 64)
}

func HashPassword(password string) ([]byte, error) {
	quickHash := sha256.Sum256([]byte(password))
	slowHash, err := bcrypt.GenerateFromPassword(quickHash[:], 12)
	if err != nil {
		return nil, err
	}
	return slowHash, nil
}
