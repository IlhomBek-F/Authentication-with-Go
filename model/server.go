package model

import "database/sql"

type Server struct {
	Port int
	Db   *sql.DB
}
