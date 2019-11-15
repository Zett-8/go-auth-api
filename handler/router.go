package handler

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"net/http"
)

func NewRouter() *echo.Echo {
	e := echo.New()
	e.HTTPErrorHandler = ErrorHandler

	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// index
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Go auth API")
	})

	// user
	e.POST("/sign-up", SignUp)
	e.POST("/login", Login)

	api := e.Group("/api")
	// authentication is required from here
	api.Use(middleware.JWTWithConfig(Config))
	api.GET("/user", GetUserInfo)
	api.GET("/todo", GetUserTodos)
	api.POST("/todo", CreateTodo)
	api.PUT("/todo/:id", PutTodo)
	api.DELETE("/todo/:id", DeleteTodo)

	return e
}
