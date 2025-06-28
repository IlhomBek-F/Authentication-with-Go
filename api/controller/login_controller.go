package controller

import (
	"auth/internal/tokenutil"
	"auth/model"
	"auth/repositories"
	"net/http"
	"os"
	"strconv"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

func (s *Controller) Login(e echo.Context) error {
	var userCredential = new(model.User)

	err := e.Bind(&userCredential)

	if err != nil {
		return e.JSON(http.StatusInternalServerError, model.ErrorResponse{Status: http.StatusInternalServerError, Message: "Internal server error"})
	}

	if validationError := e.Validate(userCredential); validationError != nil {
		return e.JSON(http.StatusUnprocessableEntity, model.ErrorResponse{Status: http.StatusUnprocessableEntity, Message: validationError.Error()})
	}

	user, err := repositories.GetByEmail(s.Db, userCredential.Email)

	if err != nil {
		return e.JSON(http.StatusNotFound, model.ErrorResponse{Status: http.StatusFound, Message: "User does not exist with this email"})
	}

	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userCredential.Password)) != nil {
		return e.JSON(http.StatusUnauthorized, model.ErrorResponse{Status: http.StatusUnauthorized, Message: "Invalid credentials"})
	}

	jwtPrivateKey := os.Getenv("ACCESS_TOKEN_SECRET")
	accessTokenExpiryHour, _ := strconv.Atoi(os.Getenv("ACCESS_TOKEN_EXPIRY_HOUR"))
	accessToken, err := tokenutil.GenerateAccessToken(user, jwtPrivateKey, accessTokenExpiryHour)

	if err != nil {
		return e.JSON(http.StatusInternalServerError, model.ErrorResponse{Status: http.StatusInternalServerError, Message: "Internal server error"})
	}

	resp := model.SuccessResponse{
		Status:  http.StatusOK,
		Message: "Success",
		Data: map[string]any{
			"id":          user.Id,
			"email":       user.Email,
			"accessToken": accessToken,
		},
	}

	return e.JSON(http.StatusOK, resp)
}
