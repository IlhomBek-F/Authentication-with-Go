package model

import "database/sql"

type Server struct {
	port int
	db   *sql.DB
}
