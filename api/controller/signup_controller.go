package controller

import (
	"auth/model"
	"auth/repositories"
	"database/sql"
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

	_, err = repositories.GetByEmail(s.Db, newUser.Email)

	if err == nil {
		return c.JSON(http.StatusInternalServerError, model.ErrorResponse{Status: http.StatusConflict, Message: "User already exist with this email address"})
	}

	encryptedPassword, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.DefaultCost)
	newUser.Password = string(encryptedPassword)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, model.ErrorResponse{Status: http.StatusInternalServerError, Message: err.Error()})
	}

	_, err = repositories.CreateUser(s.Db, newUser)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, model.ErrorResponse{Status: http.StatusInternalServerError, Message: err.Error()})
	}

	resp := model.SuccessResponse{
		Status:  http.StatusOK,
		Message: "Success",
	}

	return c.JSON(http.StatusCreated, resp)
}
