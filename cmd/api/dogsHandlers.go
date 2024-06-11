package main

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/PedroDrago/DogsAPI/internal/models"
	"github.com/PedroDrago/DogsAPI/internal/validator"
)

func (app *application) viewDogHandler(writer http.ResponseWriter, req *http.Request) {
	id, err := getIDParam(req)
	if err != nil {
		app.responseBadRequest(writer, err)
		return
	}
	dog, err := app.models.Dogs.Get(id)
	if err != nil {
		switch {
		case errors.Is(err, models.ErrRecordNotFound):
			app.responseNotFound(writer)
		default:
			app.responseInternalServerError(writer, err)
		}
		return
	}
	err = writeJSON(writer, http.StatusOK, Envelope{"dog": dog}, nil)
	if err != nil {
		app.responseInternalServerError(writer, err)
	}
}

func (app *application) createDogHandler(writer http.ResponseWriter, req *http.Request) {
	var input struct {
		Name         string   `json:"name"`
		BirthYear    int32    `json:"birth_year"`
		Breed        string   `json:"breed"`
		Sex          string   `json:"sex"`
		SpecialNeeds []string `json:"special_needs"`
		Neutered     bool     `json:"neutered"`
		TutorID      int64    `json:"tutor_id"`
	}
	err := readJSON(req, &input)
	if err != nil {
		app.infoLog.Println(err)
		app.responseBadRequest(writer, err)
		return
	}
	dog := models.Dog{
		Name:         input.Name,
		BirthYear:    input.BirthYear,
		Breed:        input.Breed,
		Sex:          input.Sex,
		SpecialNeeds: input.SpecialNeeds,
		Neutered:     input.Neutered,
		TutorID:      input.TutorID,
	}
	v := validator.New()
	if !dog.Validate(v) {
		app.validationErrorResponse(writer, v.Errors)
		return
	}
	exists, err := app.models.Users.UserExists(dog.TutorID)
	if err != nil {
		app.responseInternalServerError(writer, err)
		return
	}
	if !exists {
		v.Errors["tutor_id"] = "must be an existent tutor"
		app.validationErrorResponse(writer, v.Errors)
		return

	}
	err = app.models.Dogs.Insert(&dog)
	if err != nil {
		app.responseInternalServerError(writer, err)
		return
	}
	writer.Header().Set("Location", fmt.Sprintf("/v1/dogs/%d", dog.ID))
	err = writeJSON(writer, http.StatusCreated, Envelope{"dog": dog}, nil)
	if err != nil {
		app.responseInternalServerError(writer, err)
	}
}
