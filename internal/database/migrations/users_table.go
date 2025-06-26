package migrations

import (
	"database/sql"
)

func InitMigrations(db *sql.DB) error {
	userTable := `CREATE TABLE IF NOT EXISTS public.users  (
	           id serial4 PRIMARY KEY NOT NULL,
	           email varchar NOT NULL,
	           password varchar NOT NULL,
			   created_at timestamp DEFAULT current_timestamp,
			   updated_at timestamp DEFAULT current_timestamp,
			   deleted_at timestamp DEFAULT current_timestamp
	           );`

	_, err := db.Exec(userTable)

	return err
}
