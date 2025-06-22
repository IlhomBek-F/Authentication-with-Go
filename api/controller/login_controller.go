package controller

import (
	"auth/model"
	"auth/repositories"
	"net/http"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

func (s *Controller) Login(e echo.Context) error {
	var userCredential model.User

	err := e.Bind(&userCredential)

	if err != nil {
		return e.JSON(http.StatusInternalServerError, model.ErrorResponse{Status: http.StatusInternalServerError, Message: "Internal server error"})
	}

	user, err := repositories.GetByEmail(s.Db, userCredential.Email)

	if err != nil {
		return e.JSON(http.StatusNotFound, model.ErrorResponse{Status: http.StatusFound, Message: "User does not exist with this email"})
	}

	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userCredential.Password)) != nil {
		return e.JSON(http.StatusUnauthorized, model.ErrorResponse{Status: http.StatusUnauthorized, Message: "Invalid credentials"})
	}

	resp := model.SuccessResponse{
		Status:  http.StatusOK,
		Message: "Success",
		Data:    userCredential,
	}

	return e.JSON(http.StatusOK, resp)
}
