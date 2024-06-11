package models

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/PedroDrago/DogsAPI/internal/validator"
	"github.com/lib/pq"
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

func (model DogModel) Insert(dog *Dog) error {
	query := `
    INSERT INTO dogs (name, birth_year, breed, sex, special_needs, neutered, user_id)
    VALUES ($1, $2, $3, $4, $5, $6, $7)
    RETURNING id, created_at, updated_at, version
    `

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	args := []any{dog.Name, dog.BirthYear, dog.Breed, dog.Sex, pq.Array(dog.SpecialNeeds), dog.Neutered, dog.TutorID}
	err := model.DB.QueryRowContext(ctx, query, args...).Scan(&dog.ID, &dog.CreatedAt, &dog.UpdatedAt, &dog.Version)
	if err != nil {
		return err
	}
	return nil
}

func (model DogModel) Get(id int64) (*Dog, error) {
	query := `
    SELECT id, name, birth_year, breed, sex, special_needs, neutered, created_at, updated_at, user_id, version
    FROM dogs
    WHERE id = $1
    `
	dog := &Dog{}
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	args := []any{
		&dog.ID,
		&dog.Name,
		&dog.BirthYear,
		&dog.Breed,
		&dog.Sex,
		&dog.SpecialNeeds,
		&dog.Neutered,
		&dog.CreatedAt,
		&dog.UpdatedAt,
		&dog.TutorID,
		&dog.Version,
	}
	err := model.DB.QueryRowContext(ctx, query, id).Scan(args...)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrRecordNotFound
		default:
			return nil, err
		}
	}
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

var SpecialNeedsList = []string{"blind", "deaf"} // TODO: add more

func (dog *Dog) Validate(v *validator.Validator) bool {
	v.Check(dog.Name != "", "name", "must be provided")
	l := len(dog.Name)
	v.Check(l <= 20, "name", "must be at most 20 bytes long")
	v.Check(dog.BirthYear >= 1980, "birth_year", "must be greater than 1980")
	v.Check(dog.BirthYear <= int32(time.Now().Year()), "birth_year", "must not be in the future")
	v.Check(dog.Sex != "", "sex", "must be provided")
	v.Check(validator.PermittedValue(dog.Sex, "male", "female"), "sex", "must be male or female")
	for _, val := range dog.SpecialNeeds {
		v.Check(validator.PermittedValue(val, SpecialNeedsList...), "special_needs", "must be a valid special need")
	}
	return v.Valid()
}
