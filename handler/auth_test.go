package handler

import (
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestConnection(t *testing.T) {
	e := echo.New()
	req, _ := http.NewRequest(http.MethodGet, "/", nil)
	rc := httptest.NewRecorder()

	e.NewContext(req, rc)

	assert.Equal(t, http.StatusOK, rc.Code, "success to connect to server")
}
