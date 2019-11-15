package handler

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
	"go-auth-api/model"
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
	userJson := `{"name": "test", "password": "test1234"}`

	e := NewRouter()
	req := httptest.NewRequest(http.MethodPost, "/sign-up", strings.NewReader(userJson))
	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)

	testDB, err := gorm.Open("sqlite3", "/tmp/test.db")
	if err != nil {
		t.Error()
	}
	defer testDB.Close()
	testDB.AutoMigrate(&model.User{})
	testDB.AutoMigrate(&model.Todo{})

	h := &Handler{testDB}

	if assert.NoError(t, h.SignUp(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
	}

	var user model.User
	testDB.First(&user)
	assert.Equal(t, "test", user.Name)
}
