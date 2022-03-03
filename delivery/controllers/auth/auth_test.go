package auth

import (
	"HealthFit/entities"
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

type MockAuthLib struct{}

func (m *MockAuthLib) Login(email, password string) (entities.User, error) {
	if email != "test" && password != "test" {
		return entities.User{}, errors.New("record not found")
	}
	return entities.User{Model: gorm.Model{ID: 1}, Email: email, Password: password}, nil
}

type MockAuthLibToken struct{}

func (m *MockAuthLibToken) Login(email, password string) (entities.User, error) {
	return entities.User{}, nil
}

func TestLogin(t *testing.T) {
	t.Run("error in input file", func(t *testing.T) {
		e := echo.New()

		reqBody, _ := json.Marshal(map[string]string{
			"email": "anonim@123",
		})

		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(reqBody))
		res := httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")

		context := e.NewContext(req, res)
		context.SetPath("/login")

		authCont := New(&MockAuthLib{})
		authCont.Login()(context)

		resp := LoginRespFormat{}

		json.Unmarshal([]byte(res.Body.Bytes()), &resp)
		assert.Equal(t, 400, resp.Code)
		assert.Equal(t, "error in input file", resp.Message)

	})

	t.Run("error in call database", func(t *testing.T) {
		e := echo.New()

		reqBody, _ := json.Marshal(map[string]string{
			"email":    "anonim",
			"password": "anonim",
		})
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(reqBody))
		res := httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")
		context := e.NewContext(req, res)
		context.SetPath("/login")
		authController := New(&MockAuthLib{})
		authController.Login()(context)
		response := LoginRespFormat{}
		log.Info(res.Body)
		json.Unmarshal([]byte(res.Body.Bytes()), &response)
		assert.Equal(t, 500, response.Code)
		assert.Equal(t, "error in call database", response.Message)
	})

	t.Run("fail in process token", func(t *testing.T) {
		e := echo.New()
		reqBody, _ := json.Marshal(map[string]string{
			"email":    "anonim@123",
			"password": "anonim123",
		})
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(reqBody))
		res := httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")
		context := e.NewContext(req, res)
		context.SetPath("/login")
		authController := New(&MockAuthLibToken{})
		authController.Login()(context)
		response := LoginRespFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		assert.Equal(t, 406, response.Code)
		assert.Equal(t, "error in process token", response.Message)
	})

	t.Run("success login", func(t *testing.T) {
		e := echo.New()
		reqBody, _ := json.Marshal(map[string]string{
			"email":    "test",
			"password": "test",
		})
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(reqBody))
		res := httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")
		context := e.NewContext(req, res)
		context.SetPath("/login")
		authController := New(&MockAuthLib{})
		authController.Login()(context)
		response := LoginRespFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		assert.Equal(t, 200, response.Code)
	})
}
