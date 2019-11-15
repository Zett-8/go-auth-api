package handler

import (
	"github.com/labstack/echo"
	"go-auth-api/model"
	"net/http"
	"strconv"
)

func GetUserTodos(c echo.Context) error {
	userId := retrieveUserIdFromToken(c)

	if user, _ := model.GetUser(&model.User{ID: userId}); user.ID == 0 {
		return echo.ErrNotFound
	}

	todos, _ := model.GetTodos(&model.Todo{UserID: uint(userId)})

	return c.JSON(http.StatusOK, todos)
}

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
	if user, _ := model.GetUser(&model.User{ID: id}); user.ID == 0 {
		return echo.ErrNotFound
	}

	todo.UserID = uint(id)
	model.CreateTodo(todo)

	return c.JSON(http.StatusCreated, todo)
}

func GetUserInfo(c echo.Context) error {
	userId := retrieveUserIdFromToken(c)

	user, _ := model.GetUser(&model.User{ID: userId})

	return c.JSON(http.StatusOK, user)
}

func PutTodo(c echo.Context) error {
	var newTodo *model.Todo
	err := c.Bind(&newTodo)
	if err != nil {
		return err
	}

	userId := retrieveUserIdFromToken(c)
	if user, _ := model.GetUser(&model.User{ID: userId}); user.ID == 0 {
		return echo.ErrNotFound
	}

	todoID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.ErrNotFound
	}

	todos, _ := model.GetTodos(&model.Todo{ID: todoID, UserID: uint(userId)})
	if len(todos) == 0 {
		return echo.ErrNotFound
	}
	todo := todos[0]
	todo.Name = newTodo.Name
	todo.Done = newTodo.Done

	if err := model.UpdateTodo(&todo); err != nil {
		return echo.ErrNotFound
	}

	return c.JSON(http.StatusOK, todo)
}

func DeleteTodo(c echo.Context) error {
	userId := retrieveUserIdFromToken(c)
	if user, _ := model.GetUser(&model.User{ID: userId}); user.ID == 0 {
		return echo.ErrNotFound
	}

	todoID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.ErrNotFound
	}

	if err := model.DeleteTodo(&model.Todo{ID: todoID, UserID: uint(userId)}); err != nil {
		return echo.ErrNotFound
	}

	return c.NoContent(http.StatusNoContent)
}
