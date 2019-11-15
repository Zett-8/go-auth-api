package handler

import (
	"github.com/labstack/echo"
	"go-auth-api/model"
	"net/http"
	"strconv"
)

func (h *Handler) GetUserTodos(c echo.Context) error {
	userId := retrieveUserIdFromToken(c)

	if user, _ := model.GetUser(&model.User{ID: userId}, h.DB); user.ID == 0 {
		return echo.ErrNotFound
	}

	todos, _ := model.GetTodos(&model.Todo{UserID: uint(userId)}, h.DB)

	return c.JSON(http.StatusOK, todos)
}

func (h *Handler) CreateTodo(c echo.Context) error {
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
	if user, _ := model.GetUser(&model.User{ID: id}, h.DB); user.ID == 0 {
		return echo.ErrNotFound
	}

	todo.UserID = uint(id)
	model.CreateTodo(todo, h.DB)

	return c.JSON(http.StatusCreated, todo)
}

func (h *Handler) GetUserInfo(c echo.Context) error {
	userId := retrieveUserIdFromToken(c)

	user, _ := model.GetUser(&model.User{ID: userId}, h.DB)

	return c.JSON(http.StatusOK, user)
}

func (h *Handler) PutTodo(c echo.Context) error {
	var newTodo *model.Todo
	err := c.Bind(&newTodo)
	if err != nil {
		return err
	}

	userId := retrieveUserIdFromToken(c)
	if user, _ := model.GetUser(&model.User{ID: userId}, h.DB); user.ID == 0 {
		return echo.ErrNotFound
	}

	todoID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.ErrNotFound
	}

	todos, _ := model.GetTodos(&model.Todo{ID: todoID, UserID: uint(userId)}, h.DB)
	if len(todos) == 0 {
		return echo.ErrNotFound
	}
	todo := todos[0]
	todo.Name = newTodo.Name
	todo.Done = newTodo.Done

	if err := model.UpdateTodo(&todo, h.DB); err != nil {
		return echo.ErrNotFound
	}

	return c.JSON(http.StatusOK, todo)
}

func (h *Handler) DeleteTodo(c echo.Context) error {
	userId := retrieveUserIdFromToken(c)
	if user, _ := model.GetUser(&model.User{ID: userId}, h.DB); user.ID == 0 {
		return echo.ErrNotFound
	}

	todoID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.ErrNotFound
	}

	if err := model.DeleteTodo(&model.Todo{ID: todoID, UserID: uint(userId)}, h.DB); err != nil {
		return echo.ErrNotFound
	}

	return c.NoContent(http.StatusNoContent)
}
