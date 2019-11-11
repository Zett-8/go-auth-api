package handler

import (
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
	"go-auth-api/model"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestConnection(t *testing.T) {
	e := echo.New()
	req, _ := http.NewRequest("GET", "/", nil)
	rc := httptest.NewRecorder()

	e.NewContext(req, rc)

	assert.Equal(t, rc.Code, http.StatusOK, "success to connect to server")
}

func TestSignUp(t *testing.T) {
	e := echo.New()

	user := model.User{Name: "admin", Password: "test1234"}

	req, _ := http.NewRequest("POST", "/sign-up", user)
	rc := httptest.NewRecorder()

	e.NewContext(req, rc)

	assert.Equal(t, rc.Code, http.StatusCreated)
}
