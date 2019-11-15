package handler

import (
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestConnect(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	_ = e.NewContext(req, rec)

	if assert.NoError(t, nil) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}

func TestSignUp(t *testing.T) {
	userJson := `{"name": "test4", "password": "test1234"}`

	e := NewRouter()
	req := httptest.NewRequest(http.MethodPost, "/sign-up", strings.NewReader(userJson))
	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)

	if assert.NoError(t, SignUp(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
	}
}
