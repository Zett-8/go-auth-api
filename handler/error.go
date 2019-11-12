package handler

import (
	"github.com/labstack/echo"
	"net/http"
)

func ErrorHandler(err error, c echo.Context) {
	c.String(http.StatusInternalServerError, "original error message")
}
