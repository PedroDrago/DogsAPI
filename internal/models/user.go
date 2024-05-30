package models

import (
	"database/sql"
	"time"
)

type UserModel struct {
	DB *sql.DB
}

type User struct {
	ID        int64
	Name      string
	Username  string
	Email     string
	BirthYear int32
	Address   string
	Phone     string
	Admin     bool
	CreatedAt time.Time
	UpdatedAt time.Time
	Dogs      []Dog
}

func (model UserModel) Insert(user *User) error {
	query := `
	INSERT INTO users(name, username, email, birth_year, address, phone) 
	VALUES ($1, $2, $3, $4, $5, $6) 
	RETURNING id, created_at, updated_at;
	`
	args := []any{user.Name, user.Username, user.Email, user.BirthYear, user.Address, user.Phone}
	err := model.DB.QueryRow(query, args...).Scan(&user.ID, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return err
	}
	return nil
}

func (model UserModel) Get(id int64) (User, error) {
	usr := User{}
	return usr, nil
}

func (model UserModel) Update(user *User) error {
	return nil
}

func (model UserModel) Delete(id int64) error {
	return nil
}
