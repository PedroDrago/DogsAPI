package models

import (
	"database/sql"
	"time"
)

type DogModel struct {
	DB *sql.DB
}

type Dog struct {
	ID           int64     `json:"id"`
	Name         string    `json:"name"`
	BirthYear    int32     `json:"birth_year,omitempty"`
	Breed        string    `json:"breed,omitempty"`
	Sex          string    `json:"sex,omitempty"`
	SpecialNeeds []string  `json:"special_needs,omitempty"`
	Neutered     bool      `json:"neutered,omitempty"`
	CreatedAt    time.Time `json:"-"`
	UpdatedAt    time.Time `json:"-"`
	TutorID      int64     `json:"user_id"`
	Version      int32     `json:"version"`
}

func (model DogModel) Insert(dog *Dog, userID int64) error {
	return nil
}

func (model DogModel) Get(id int64) (Dog, error) {
	dog := Dog{}
	return dog, nil
}

func (model DogModel) Update(user *User) error {
	return nil
}

func (model DogModel) Delete(id int64) error {
	return nil
}

func (model DogModel) Index(id int64) ([]Dog, error) {
	dogs := []Dog{}
	return dogs, nil
}
