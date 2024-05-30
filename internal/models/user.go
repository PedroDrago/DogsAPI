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
	query := `
	UPDATE USERS
	SET name = $1,
	username = $2,
	email = $3,
	birth_year = $4,
	address = $5,
	phone_number = $6,
	updated_at = CURRENT_TIMESTAMP,
	password_hash = $7
	WHERE id = $8
	returning id, updated_at;
	`
	args := []any{user.Name, user.Username, user.Email, user.BirthYear, user.Address, user.PhoneNumber, user.PasswordHash, user.ID}
	err := model.DB.QueryRow(query, args...).Scan(&user.ID, &user.UpdatedAt)
	if err != nil {
		return err
	}
	return nil
}

func (model UserModel) Delete(id int64) error {
	query := `
	DELETE FROM users
	WHERE id = $1;
	`
	_, err := model.DB.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}

func (model UserModel) Index() ([]User, error) {
	var users []User
	var usr User
	query := `
	SELECT id, name, username, email, birth_year, address, phone_number, created_at, updated_at, password_hash 
	FROM USERS;
	`
	rows, err := model.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	args := []any{&usr.ID, &usr.Name, &usr.Username, &usr.Email, &usr.BirthYear, &usr.Address, &usr.PhoneNumber, &usr.CreatedAt, &usr.UpdatedAt, &usr.PasswordHash}
	for rows.Next() {
		err = rows.Scan(args...)
		users = append(users, usr)
		if err != nil {
			return nil, err
		}
	}
	return users, nil
}
