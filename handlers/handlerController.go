package handlers

import "database/sql"

type HandlerController struct {
	DB *sql.DB
}
