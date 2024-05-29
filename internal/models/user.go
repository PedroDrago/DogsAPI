package models

type User struct {
	Name     string
	Username string
	Age      int32
	Address  string
	Phone    string
	Admin    bool
	Dogs     []Dog
}
