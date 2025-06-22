package controller

import (
	"auth/model"
	"auth/repositories"
	"database/sql"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

type Controller struct {
	Port int
	Db   *sql.DB
}

func (s *Controller) SignUp(c echo.Context) error {
	var newUser model.User

	err := c.Bind(&newUser)

	if err != nil {
		return c.JSON(http.StatusBadRequest, model.ErrorResponse{Status: http.StatusBadRequest, Message: err.Error()})
	}

	user, err := repositories.GetByEmail(s.Db, newUser.Email)

	if err != nil {
		return c.JSON(http.StatusConflict, model.ErrorResponse{Status: http.StatusConflict, Message: "User already exist with this email address"})
	}

	encryptedPassword, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.DefaultCost)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, model.ErrorResponse{Status: http.StatusInternalServerError, Message: err.Error()})
	}

	fmt.Println("You can save user with ", encryptedPassword, user)

	return nil
}
