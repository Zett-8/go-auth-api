package handler

import (
	"github.com/labstack/echo"
	"net/http"
)

func GetTodos(c echo.Context) error {
	return c.String(http.StatusOK, "yasss")
}