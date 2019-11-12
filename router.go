package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"go-auth-api/handler"
	"net/http"
)

func newRouter() *echo.Echo {
	e := echo.New()

	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// index
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Go auth API")
	})

	// user
	e.POST("/sign-up", handler.SignUp)
	e.POST("/login", handler.Login)

	api := e.Group("/api")
	// authentication is required from here
	api.Use(middleware.JWTWithConfig(handler.Config))
	api.GET("/user", handler.GetUserInfo)
	api.GET("/todo", handler.GetUserTodos)
	api.POST("/todo", handler.CreateTodo)
	api.DELETE("/todo/:id", handler.DeleteTodo)

	return e
}
