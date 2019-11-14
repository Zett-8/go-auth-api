package handler

import (
	//"fmt"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestSignUp(t *testing.T) {
	userJson := `{"name": "test", "password": "test1234"}`

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/sign-up", strings.NewReader(userJson))
	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)

	if assert.NoError(t, SignUp(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
	}
}
