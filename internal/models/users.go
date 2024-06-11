package models

import (
	"context"
	"database/sql"
	"errors"
	"regexp"
	"time"

	"github.com/PedroDrago/DogsAPI/internal/validator"
	"github.com/lib/pq"
	_ "github.com/lib/pq"
)

var (
	ErrDuplicateEmail    = errors.New("duplicate email")
	ErrDuplicateUsername = errors.New("duplicate username")
	ErrRecordNotFound    = errors.New("record not found")
	EmailRX              = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
)

type UserModel struct {
	DB *sql.DB
}

type PubUser struct {
	User
	Password string `json:"password"`
}

type User struct {
	ID                int64     `json:"id"`
	Name              string    `json:"name"`
	Username          string    `json:"username"`
	Email             string    `json:"email"`
	BirthYear         int32     `json:"birth_year,omitempty"`
	Address           string    `json:"address,omitempty"`
	PhoneNumber       string    `json:"phone_number,omitempty"`
	Admin             bool      `json:"-"`
	CreatedAt         time.Time `json:"-"`
	UpdatedAt         time.Time `json:"-"`
	PasswordHash      []byte    `json:"-"`
	PassowrdPlainText string    `json:"-"`
	Dogs              []Dog     `json:"dogs,omitempty"`
	Version           int32     `json:"version"`
}

func (model UserModel) Insert(user *User) error {
	query := `
    INSERT INTO users (name, username, email, birth_year, address, phone_number, password_hash)
    VALUES ($1, $2, $3, $4, $5, $6, $7)
    RETURNING id, version, created_at, updated_at
    `

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	args := []any{user.Name, user.Username, user.Email, user.BirthYear, user.Address, user.PhoneNumber, user.PasswordHash}
	err := model.DB.QueryRowContext(ctx, query, args...).Scan(
		&user.ID,
		&user.Version,
		&user.CreatedAt,
		&user.UpdatedAt)
	if err != nil {
		switch {
		case err.Error() == `pq: duplicate key value violates unique constraint "users_email_key"`:
			return ErrDuplicateEmail
		case err.Error() == `pq: duplicate key value violates unique constraint "users_username_key"`:
			return ErrDuplicateUsername
		default:
			return err
		}
	}
	return nil
}

func (model UserModel) GetUserDogs(usr *User) error {
	query := `
    SELECT dogs.id, dogs.name, dogs.birth_year, dogs.breed, dogs.sex, dogs.special_needs, dogs.neutered, dogs.created_at, dogs.updated_at, dogs.version
    FROM dogs
    INNER JOIN users
    ON dogs.user_id = users.id
    WHERE users.id = $1
    `
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	rows, err := model.DB.QueryContext(ctx, query, usr.ID)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return ErrRecordNotFound
		default:
			return err
		}
	}
	defer rows.Close()
	for rows.Next() {
		var dog Dog
		err = rows.Scan(
			&dog.ID,
			&dog.Name,
			&dog.BirthYear,
			&dog.Breed,
			&dog.Sex,
			pq.Array(dog.SpecialNeeds),
			&dog.Neutered,
			&dog.CreatedAt,
			&dog.UpdatedAt,
			&dog.Version,
		)
		if err != nil {
			return err
		}
		usr.Dogs = append(usr.Dogs, dog)
	}
	return nil
}

func (model UserModel) Get(id int64) (*User, error) {
	query := `
    SELECT id, name, username, email, birth_year, address, phone_number, admin, created_at, updated_at, password_hash, version
    FROM users
    WHERE id = $1
    `

	user := &User{}
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	args := []any{
		&user.ID,
		&user.Name,
		&user.Username,
		&user.Email,
		&user.BirthYear,
		&user.Address,
		&user.PhoneNumber,
		&user.Admin,
		&user.CreatedAt,
		&user.UpdatedAt,
		&user.PasswordHash,
		&user.Version,
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
	return user, nil
}

func (model UserModel) Update(user *User) error {
	query := `
    UPDATE users
    SET name = $1,
    username = $2,
    email = $3,
    birth_year = $4,
    address = $5,
    phone_number = $6,
    password_hash = $7,
    version = version + 1
    WHERE id = $8 AND version = $9
    returning VERSION
    `

	args := []any{
		user.Name,
		user.Username,
		user.Email,
		user.BirthYear,
		user.Address,
		user.PhoneNumber,
		user.PasswordHash,
		user.ID,
		user.Version,
	}
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := model.DB.QueryRowContext(ctx, query, args...).Scan(&user.Version)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return ErrRecordNotFound
		case err.Error() == `pq: duplicate key value violates unique constraint "users_email_key"`:
			return ErrDuplicateEmail
		case err.Error() == `pq: duplicate key value violates unique constraint "users_username_key"`:
			return ErrDuplicateUsername
		}
		return err
	}
	return nil
}

func (model UserModel) Delete(id int64) error {
	query := `
    DELETE FROM USERS
    WHERE id = $1
    `
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	res, err := model.DB.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}
	rows, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rows == 0 {
		return ErrRecordNotFound
	}
	return nil
}

func (model UserModel) Index() ([]User, error) {
	return nil, nil
}

func (usr *User) Validate(v *validator.Validator) bool {
	nameLen := len(usr.Name)
	v.Check(usr.Name != "", "name", "must be provided")
	v.Check(nameLen >= 4, "name", "must be at least 4 bytes long")
	v.Check(nameLen <= 50, "name", "must at most 50 bytes long")
	l := len(usr.Username)
	v.Check(usr.Username != "", "name", "must be provided")
	v.Check(l >= 5, "usrname", "must be at least 5 bytes long")
	v.Check(l <= 50, "usrname", "must at most 20 bytes long")
	v.Check(usr.Email != "", "email", "must be provided")
	v.Check(v.Match(usr.Email, EmailRX), "email", "must be a valid email address")
	v.Check(usr.BirthYear > 1900, "birth_year", "must be greater than than 1900")
	v.Check(usr.BirthYear < int32(time.Now().Year()), "birth_year", "must be not in the future")
	v.Check(usr.Address != "", "address", "must be provided")
	l = len(usr.Address) // In a real app should be validated as a real address
	v.Check(l >= 5, "address", "must be at least 5 bytes")
	v.Check(l <= 100, "address", "must be at most 100 bytes")
	v.Check(usr.PhoneNumber != "", "phone_number", "must be provided")
	l = len(usr.PhoneNumber) // In a real app should be validated as a real phone number with country code, DDD etc.
	v.Check(l >= 8, "phone_number", "must be at least 8 bytes")
	v.Check(l <= 20, "phone_number", "must be at most 20 bytes")
	v.Check(usr.PassowrdPlainText != "", "password", "must be provided")
	l = len(usr.PassowrdPlainText)
	v.Check(l >= 16, "password", "must be at least 16 bytes long")
	v.Check(l <= 256, "password", "must be at most 256 bytes long")
	return v.Valid()
}

func (model UserModel) UserExists(id int64) (bool, error) {
	query := `
    SELECT COUNT(*)
    FROM USERS
    WHERE id = $1
    `
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var count int
	err := model.DB.QueryRowContext(ctx, query, id).Scan(&count)
	if err != nil {
		return false, err
	}
	if count == 0 {
		return false, nil
	}
	return true, nil
}
