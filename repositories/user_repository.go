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

func GetUsers(db *sql.DB) ([]model.User, error) {
	rows, err := db.Query("SELECT email, id FROM users")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var user model.User
	users := []model.User{}

	for rows.Next() {
		if err := rows.Scan(&user.Email, &user.Id); err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}
