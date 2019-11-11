package handler

import (
	"github.com/labstack/echo"
	"go-auth-api/model"
	"net/http"
)

func GetUser(c echo.Context) error {
	user := model.GetUser(1)
	return c.JSON(http.StatusOK, user)
}
