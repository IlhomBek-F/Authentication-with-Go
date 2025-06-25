package controller

import (
	"auth/model"
	"auth/repositories"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (c *Controller) GetUsers(e echo.Context) error {

	users, err := repositories.GetUsers(c.Db)

	if err != nil {
		return e.JSON(http.StatusInternalServerError, model.ErrorResponse{Status: http.StatusInternalServerError, Message: "Internal server error"})
	}

	resp := model.SuccessResponse{
		Status:  http.StatusOK,
		Message: "success",
		Data:    users,
	}

	return e.JSON(http.StatusOK, resp)
}
