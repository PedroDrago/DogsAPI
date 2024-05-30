package models

import (
	"database/sql"
	"time"

	_ "github.com/lib/pq"
)

type UserModel struct {
	DB *sql.DB
}

type User struct {
	ID           int64     `json:"id"`
	Name         string    `json:"name"`
	Username     string    `json:"username"`
	Email        string    `json:"email"`
	BirthYear    int32     `json:"birth_year,omitempty"`
	Address      string    `json:"address,omitempty"`
	PhoneNumber  string    `json:"phone_number,omitempty"`
	Admin        bool      `json:"-"`
	CreatedAt    time.Time `json:"-"`
	UpdatedAt    time.Time `json:"-"`
	PasswordHash string    `json:"-"`
	Dogs         []Dog     `json:"dogs,omitempty"`
}

func (model UserModel) Insert(user *User) error {
	query := `
	INSERT INTO users(name, username, email, birth_year, address, phone_number, password_hash) 
	VALUES ($1, $2, $3, $4, $5, $6, $7) 
	RETURNING id, created_at, updated_at;
	` // TODO: I guess I'll need to Join here to get the Dogs array
	args := []any{user.Name, user.Username, user.Email, user.BirthYear, user.Address, user.PhoneNumber, user.PasswordHash}
	err := model.DB.QueryRow(query, args...).Scan(&user.ID, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return err
	}
	return nil
}

func (model UserModel) Get(id int64) (User, error) {
	usr := User{}
	query := `
	SELECT id, name, username, email, birth_year, address, phone_number, admin, created_at, updated_at, password_hash
	FROM users
	WHERE id = $1;
	`
	args := []any{&usr.ID, &usr.Name, &usr.Username, &usr.Email, &usr.BirthYear, &usr.Address, &usr.PhoneNumber, &usr.Admin, &usr.CreatedAt, &usr.UpdatedAt, &usr.PasswordHash}
	err := model.DB.QueryRow(query, id).Scan(args...)
	if err != nil {
		return usr, err
	}
	return usr, nil
}

func (model UserModel) Update(user *User) error {
	return nil
}

func (model UserModel) Delete(id int64) error {
	// query := `
	// DELETE FROM users
	// WHERE id = $1;
	// `
	return nil
}
