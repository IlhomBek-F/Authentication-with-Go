package repositories

import (
	"auth/model"
	"database/sql"
	"fmt"
)

func GetByEmail(db *sql.DB, email string) (model.User, error) {

	row := db.QueryRow("SELECT email FROM todos WHERE email = $1", email)

	var user model.User

	err := row.Scan(&user.Email, &user.Password)
	fmt.Println(err.Error())
	return user, err
}
