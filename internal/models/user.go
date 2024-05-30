package models

type User struct {
	ID        int64
	Name      string
	Username  string
	Email     string
	BirthYear int32
	Address   string
	Phone     string
	Admin     bool
	Dogs      []Dog
}
