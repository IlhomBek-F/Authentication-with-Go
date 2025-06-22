package repositories

import (
	"auth/model"
	"database/sql"
)

func GetByEmail(db *sql.DB, email string) (model.User, error) {

	row := db.QueryRow("SELECT * FROM users WHERE email = $1", email)

	var user model.User

	err := row.Scan(&user.Email, &user.Password, &user.Id)
	return user, err
}

func CreateUser(db *sql.DB, newUser model.User) (sql.Result, error) {
	result, err := db.Exec("INSERT INTO users (email, password) VALUES ($1, $2)", newUser.Email, newUser.Password)

	return result, err
}
