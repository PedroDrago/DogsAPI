package main

import (
	"net/http"
	"strconv"

	"github.com/PedroDrago/DogsAPI/internal/models"
	"golang.org/x/crypto/bcrypt"
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
	usr := models.PubUser{}
	err := readJSON(req, &usr)
	if err != nil {
		app.responseBadRequest(writer, err.Error())
		return
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(usr.Password), bcrypt.MinCost)
	if err != nil {
		app.responseInternalServerError(writer, err)
		return
	}
	usr.PasswordHash = string(hash)
	err = app.models.Users.Insert(&usr.User)
	if err != nil {
		app.responseInternalServerError(writer, err)
	}
}

func (app *application) viewUserHandler(writer http.ResponseWriter, req *http.Request) {
	id, err := strconv.ParseInt(req.PathValue("id"), 10, 64)
	if err != nil || id < 1 {
		app.responseBadRequest(writer, "Invalid User ID")
		return
	}
	usr, err := app.models.Users.Get(id)
	if err != nil {
		http.NotFound(writer, req)
		return
	}
	err = writeJSON(writer, http.StatusOK, usr, nil)
	if err != nil {
		http.Error(writer, "Internal Server Error", http.StatusInternalServerError)
	}
}

func (app *application) updateUserHandler(writer http.ResponseWriter, req *http.Request) {
	id, err := strconv.ParseInt(req.PathValue("id"), 10, 64)
	if err != nil || id < 1 {
		app.responseBadRequest(writer, "Invalid User ID")
		return
	}
	usr := models.User{}
	err = readJSON(req, &usr)
	if err != nil {
		app.responseBadRequest(writer, err.Error())
		return
	}
	usr.ID = id
	err = app.models.Users.Update(&usr)
	if err != nil {
		app.responseInternalServerError(writer, err)
		return
	}
	writer.WriteHeader(http.StatusOK)
}

func (app *application) deleteUserHandler(writer http.ResponseWriter, req *http.Request) {
	id, err := strconv.ParseInt(req.PathValue("id"), 10, 64)
	if err != nil || id < 1 {
		app.responseBadRequest(writer, "Invalid User ID")
		return
	}
	err = app.models.Users.Delete(id)
	if err != nil {
		app.responseInternalServerError(writer, err)
		return
	}
	writer.WriteHeader(http.StatusOK)
}

func (app *application) indexUserHandler(writer http.ResponseWriter, req *http.Request) {
	users, err := app.models.Users.Index()
	if err != nil {
		app.responseInternalServerError(writer, err)
		return
	}
	err = writeJSON(writer, http.StatusOK, users, nil)
	if err != nil {
		app.responseInternalServerError(writer, err)
	}
}

func (app *application) indexDogHandler(writer http.ResponseWriter, req *http.Request) {
}

func (app *application) createDogHandler(writer http.ResponseWriter, req *http.Request) {
}

func (app *application) viewDogHandler(writer http.ResponseWriter, req *http.Request) {
}

func (app *application) updateDogHandler(writer http.ResponseWriter, req *http.Request) {
}

func (app *application) deleteDogHandler(writer http.ResponseWriter, req *http.Request) {
}
