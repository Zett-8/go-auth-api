package handler

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"go-auth-api/model"
	"net/http"
)

type Handler struct {
	DB *gorm.DB
}

func NewRouter() *echo.Echo {
	e := echo.New()
	e.HTTPErrorHandler = ErrorHandler

	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	DB, err := gorm.Open("postgres", model.DBConfig)
	if err != nil {
		panic(err)
	}
	h := &Handler{DB}

	// index
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Go auth API")
	})

	// user
	e.POST("/sign-up", h.SignUp)
	e.POST("/login", h.Login)

	api := e.Group("/api")
	// authentication is required from here
	api.Use(middleware.JWTWithConfig(Config))
	api.GET("/user", h.GetUserInfo)
	api.GET("/todo", h.GetUserTodos)
	api.POST("/todo", h.CreateTodo)
	api.PUT("/todo/:id", h.PutTodo)
	api.DELETE("/todo/:id", h.DeleteTodo)

	return e
}
