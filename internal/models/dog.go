package models

import "time"

type Dog struct {
	ID           int64
	Name         string
	BirthYear    int32
	Breed        string
	Sex          string
	SpecialNeeds []string
	Neutered     bool
	CreatedAt    time.Time
	UpdatedAt    time.Time
	Tutor        User
}
