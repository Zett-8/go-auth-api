package handler

import (
	"github.com/labstack/echo"
	"go-auth-api/model"
	"net/http"
)

func CreateTodo(c echo.Context) error {
	todo := new(model.Todo)
	if err := c.Bind(todo); err != nil {
		return err
	}

	if todo.Name == "" {
		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: "invalid form",
		}
	}

	id := retrieveUserIdFromToken(c)
	if user := model.GetUser(model.User{ID: id}); user.ID == 0 {
		return echo.ErrNotFound
	}

	todo.UserID = uint(id)
	model.CreateTodo(todo)

	return c.JSON(http.StatusCreated, todo)
}
