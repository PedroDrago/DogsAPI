package models

type Dog struct {
	Name         string
	Age          int32
	Breed        string
	Sex          string
	Tutor        User
	Neutered     bool
	SpecialNeeds []string
}
