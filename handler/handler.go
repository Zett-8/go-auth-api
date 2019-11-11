package handler

import (
	"github.com/labstack/echo"
	"go-auth-api/model"
	"net/http"
)

func SignUp(c echo.Context) error {
	user := new(model.User)
	if err := c.Bind(user); err != nil {
		return err
	}

	if user.Name == "" || user.Password == "" {
		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: "invalid name or password",
		}
	}

	if u := model.GetUser(model.User{Name: user.Name}); u.Name == user.Name {
		return &echo.HTTPError{
			Code:    http.StatusConflict,
			Message: "name already exists",
		}
	}

	model.CreateUser(user)
	user.Password = "****"

	return c.JSON(http.StatusCreated, user)
}
