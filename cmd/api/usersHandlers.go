package main

import (
	"errors"
	"fmt"
	"net/http"
	_ "strconv"

	"github.com/PedroDrago/DogsAPI/internal/models"
	"github.com/PedroDrago/DogsAPI/internal/validator"
)

var ErrEditConflict = errors.New("edit conflict")

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
	var input struct {
		Name        string `json:"name"`
		Username    string `json:"username"`
		Email       string `json:"email"`
		BirthYear   int32  `json:"birth_year"`
		Address     string `json:"address"`
		PhoneNumber string `json:"phone_number"`
		Password    string `json:"password"`
	}
	err := readJSON(req, &input)
	if err != nil {
		app.responseBadRequest(writer, err)
		return
	}
	usr := models.User{
		Name:              input.Name,
		Username:          input.Username,
		Email:             input.Email,
		BirthYear:         input.BirthYear,
		Address:           input.Address,
		PhoneNumber:       input.PhoneNumber,
		PassowrdPlainText: input.Password,
	}
	usr.PasswordHash, err = HashPassword(usr.PassowrdPlainText)
	if err != nil {
		app.responseInternalServerError(writer, err)
	}
	v := validator.New()
	if !usr.Validate(v) {
		app.validationErrorResponse(writer, v.Errors)
		return
	}
	err = app.models.Users.Insert(&usr)
	if err != nil {
		switch {
		case err.Error() == `pq: duplicate key value violates unique constraint "users_email_key"`:
			app.responseBadRequest(writer, err)
		case err.Error() == `pq: duplicate key value violates unique constraint "users_user_key"`:
			app.responseBadRequest(writer, err)
		default:
			app.responseInternalServerError(writer, err)
		}
		return
	}
	err = writeJSON(writer, http.StatusCreated, Envelope{"user": usr}, nil)
	if err != nil {
		app.responseInternalServerError(writer, err)
	}
}

func (app *application) viewUserHandler(writer http.ResponseWriter, req *http.Request) {
	id, err := getIDParam(req)
	if err != nil {
		app.responseBadRequest(writer, err)
		return
	}
	usr, err := app.models.Users.Get(id)
	if err != nil {
		switch {
		case errors.Is(err, models.ErrRecordNotFound):
			app.responseBadRequest(writer, err)
		default:
			app.responseInternalServerError(writer, err)
		}
		return
	}
	err = app.models.Users.GetUserDogs(usr)
	if err != nil && err != models.ErrRecordNotFound {
		app.responseInternalServerError(writer, err)
		return
	}
	err = writeJSON(writer, http.StatusOK, Envelope{"user": usr}, nil)
	if err != nil {
		app.responseInternalServerError(writer, err)
	}
}

func (app *application) updateUserHandler(writer http.ResponseWriter, req *http.Request) {
	id, err := getIDParam(req)
	if err != nil {
		app.responseBadRequest(writer, err)
		return
	}

	usr, err := app.models.Users.Get(id)
	if err != nil {
		switch {
		case errors.Is(err, models.ErrRecordNotFound):
			app.responseNotFound(writer)
		default:
			app.responseInternalServerError(writer, err)
		}
		return
	}
	var input struct {
		Name        *string `json:"name"`
		Username    *string `json:"username"`
		Email       *string `json:"email"`
		BirthYear   *int32  `json:"birth_year"`
		Address     *string `json:"address"`
		PhoneNumber *string `json:"phone_number"`
		Password    *string `json:"password"`
	}
	err = readJSON(req, &input)
	if err != nil {
		app.responseInternalServerError(writer, err)
		return
	}
	if input.Name != nil {
		usr.Name = *input.Name
	}
	if input.Username != nil {
		usr.Username = *input.Username
	}
	if input.Email != nil {
		usr.Email = *input.Email
	}
	if input.BirthYear != nil {
		usr.BirthYear = *input.BirthYear
	}
	if input.Address != nil {
		usr.Address = *input.Address
	}
	if input.PhoneNumber != nil {
		usr.PhoneNumber = *input.PhoneNumber
	}
	if input.Password != nil {
		usr.PassowrdPlainText = *input.Password
		usr.PasswordHash, err = HashPassword(usr.PassowrdPlainText)
		if err != nil {
			app.responseInternalServerError(writer, err)
			return
		}
	}
	v := validator.New()
	if !usr.Validate(v) {
		app.validationErrorResponse(writer, v.Errors)
		return
	}
	err = app.models.Users.Update(usr)
	if err != nil {
		switch {
		case errors.Is(err, models.ErrRecordNotFound):
			app.responseNotFound(writer)
		case errors.Is(err, models.ErrDuplicateEmail):
			v.Errors["email"] = models.ErrDuplicateEmail.Error()
			app.validationErrorResponse(writer, v.Errors)
		case errors.Is(err, models.ErrDuplicateUsername):
			v.Errors["username"] = models.ErrDuplicateUsername.Error()
			app.validationErrorResponse(writer, v.Errors)
		default:
			app.responseInternalServerError(writer, err)
		}
		return
	}
	writer.Header().Set("Location", fmt.Sprintf("/v1/users/%d", usr.ID))
	err = writeJSON(writer, http.StatusOK, Envelope{"user": usr}, nil)
	if err != nil {
		app.responseInternalServerError(writer, err)
	}
}

func (app *application) deleteUserHandler(writer http.ResponseWriter, req *http.Request) {
	id, err := getIDParam(req)
	if err != nil {
		app.responseBadRequest(writer, err)
		return
	}
	err = app.models.Users.Delete(id)
	if err != nil {
		switch {
		case errors.Is(err, models.ErrRecordNotFound):
			app.responseNotFound(writer)
		default:
			app.responseInternalServerError(writer, err)
		}
		return
	}
	writer.WriteHeader(http.StatusOK)
}
