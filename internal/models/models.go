package models

import "database/sql"

type Models struct {
	Users UserModel
	Dogs  DogModel
}

func NewModels(db *sql.DB) Models {
	models := Models{
		Users: UserModel{
			DB: db,
		},
		Dogs: DogModel{
			DB: db,
		},
	}
	return models
}
