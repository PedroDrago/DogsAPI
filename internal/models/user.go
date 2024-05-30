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
	` // TODO: I guess I'll need to Join here to get the Dogs array
	args := []any{user.Name, user.Username, user.Email, user.BirthYear, user.Address, user.Phone}
	err := model.DB.QueryRow(query, args...).Scan(&user.ID, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return err
	}
	return nil
}

func (model UserModel) Get(id int64) (User, error) {
	usr := User{}
	query := `
	SELECT id, name, username, email, birth_year, address, phone, admin, created_at, updated_at 
	FROM users
	WHERE id = $1;
	`
	args := []any{&usr.ID, &usr.Name, &usr.Username, &usr.Email, &usr.BirthYear, &usr.Address, &usr.Phone, &usr.Admin, &usr.CreatedAt, &usr.UpdatedAt}
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
	return nil
}
