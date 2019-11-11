package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"go-auth-api/handler"
	"net/http"
)

func newRouter() *echo.Echo {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	//index
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Go auth API")
	})

	//auth
	e.POST("/sign-up", handler.SignUp)
	e.POST("/login", handler.Login)

	//api := e.Group("/api")
	//api.GET("/user", handler.GetUser)

	return e
}
