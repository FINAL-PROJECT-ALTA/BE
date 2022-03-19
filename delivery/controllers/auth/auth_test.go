package auth

import (
	"HealthFit/delivery/controllers/common"
	"HealthFit/entities"
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	// "github.com/aws/aws-sdk-go/aws/session"

	"github.com/labstack/echo/v4"

	"github.com/go-playground/assert"
	"github.com/go-playground/validator"
	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
	// "github.com/labstack/gommon/log"
)

type CustomValidator struct {
	validator *validator.Validate
}

var jwtTokenUser = ""
var jwtTokenAdmin = ""

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

//////

func TestLogin(t *testing.T) {
	t.Run(
		"1. Success Login User Test", func(t *testing.T) {
			e := echo.New()
			e.Validator = &CustomValidator{validator: validator.New()}

			requestBody, _ := json.Marshal(
				LoginReqFormat{
					Email:    "testuser@gmail.com",
					Password: "testuser",
				},
			)

			req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
			res := httptest.NewRecorder()
			req.Header.Set("Content-Type", "application/json")
			context := e.NewContext(req, res)
			context.SetPath("/users/login")

			authController := New(&MockAuthRepository{})
			authController.Login()(context)

			var response common.Response

			json.Unmarshal([]byte(res.Body.Bytes()), &response)
			data := (response.Data).(map[string]interface{})
			log.Info(data)
			log.Info(response)
			jwtTokenUser = data["token"].(string)

			assert.Equal(t, "USERS - berhasil masuk, mendapatkan token baru", response.Message)
		},
	)
	t.Run(
		"1. failed Login User Test", func(t *testing.T) {
			e := echo.New()
			e.Validator = &CustomValidator{validator: validator.New()}

			requestBody, _ := json.Marshal(
				LoginReqFormat{
					Email:    "tetuser@gmail.com",
					Password: "testuser",
				},
			)

			req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
			res := httptest.NewRecorder()
			req.Header.Set("Content-Type", "application/json")
			context := e.NewContext(req, res)
			context.SetPath("/users/login")

			authController := New(&MockAuthFailRepository{})
			authController.Login()(context)

			var response common.Response

			json.Unmarshal([]byte(res.Body.Bytes()), &response)

			assert.Equal(t, float64(http.StatusInternalServerError), response.Code)
		},
	)
	t.Run(
		"1. failed Login email not found User Test", func(t *testing.T) {
			e := echo.New()
			e.Validator = &CustomValidator{validator: validator.New()}

			requestBody, _ := json.Marshal(
				LoginReqFormat{
					Email:    "tetuser@gmail.com",
					Password: "testuser",
				},
			)

			req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
			res := httptest.NewRecorder()
			req.Header.Set("Content-Type", "application/json")
			context := e.NewContext(req, res)
			context.SetPath("/users/login")

			authController := New(&MockAuthNotFoundRepository{})
			authController.Login()(context)

			var response common.Response

			json.Unmarshal([]byte(res.Body.Bytes()), &response)

			assert.Equal(t, float64(http.StatusUnauthorized), response.Code)
		},
	)
	t.Run(
		"1. failed Login Incorrect Password User Test", func(t *testing.T) {
			e := echo.New()
			e.Validator = &CustomValidator{validator: validator.New()}

			requestBody, _ := json.Marshal(
				LoginReqFormat{
					Email:    "tetuser@gmail.com",
					Password: "testuser",
				},
			)

			req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
			res := httptest.NewRecorder()
			req.Header.Set("Content-Type", "application/json")
			context := e.NewContext(req, res)
			context.SetPath("/users/login")

			authController := New(&MockAuthIncorrectPasswordRepository{})
			authController.Login()(context)

			var response common.Response

			json.Unmarshal([]byte(res.Body.Bytes()), &response)

			assert.Equal(t, float64(http.StatusUnauthorized), response.Code)
		},
	)

	t.Run(
		"1. failed bind Login User Test", func(t *testing.T) {
			e := echo.New()
			e.Validator = &CustomValidator{validator: validator.New()}

			requestBody, _ := json.Marshal(
				LoginReqFormat{
					Email:    "tetuser",
					Password: "testuser",
				},
			)

			req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
			res := httptest.NewRecorder()
			req.Header.Set("Content-Type", "application/json")
			context := e.NewContext(req, res)
			context.SetPath("/users/login")

			authController := New(&MockAuthIncorrectPasswordRepository{})
			authController.Login()(context)

			var response common.Response

			json.Unmarshal([]byte(res.Body.Bytes()), &response)

			assert.Equal(t, float64(http.StatusBadRequest), response.Code)
		},
	)
	t.Run(
		"1. failed token generate Login User Test", func(t *testing.T) {
			e := echo.New()
			e.Validator = &CustomValidator{validator: validator.New()}

			requestBody, _ := json.Marshal(
				LoginReqFormat{
					Email:    "tetuser@gmail.com",
					Password: "testuser",
				},
			)

			req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
			res := httptest.NewRecorder()
			req.Header.Set("Content-Type", "application/json")
			context := e.NewContext(req, res)
			context.SetPath("/users/login")

			authController := New(&MockAuthFailTokenRepository{})
			authController.Login()(context)

			var response common.Response

			json.Unmarshal([]byte(res.Body.Bytes()), &response)

			assert.Equal(t, float64(http.StatusNotAcceptable), response.Code)
		},
	)
}
func TestAdminLogin(t *testing.T) {
	t.Run(
		"1. Success Login Admin Test", func(t *testing.T) {
			e := echo.New()
			e.Validator = &CustomValidator{validator: validator.New()}

			requestBody, _ := json.Marshal(
				LoginReqFormat{
					Email:    "testadmin@gmail.com",
					Password: "testadmin",
				},
			)

			req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
			res := httptest.NewRecorder()
			req.Header.Set("Content-Type", "application/json")
			context := e.NewContext(req, res)
			context.SetPath("/admin/login")

			authController := New(&MockAuthRepository{})
			authController.AdminLogin()(context)

			var response common.Response

			json.Unmarshal([]byte(res.Body.Bytes()), &response)
			data := (response.Data).(map[string]interface{})
			// log.Info(data)
			// log.Info(response)
			jwtTokenAdmin = data["token"].(string)
			log.Info(data)
			log.Info(response)
			assert.Equal(t, "ADMIN - berhasil masuk, mendapatkan token baru", response.Message)
		},
	)
}

type MockAuthRepository struct{}

func (m *MockAuthRepository) Login(email, password string) (entities.User, error) {

	return entities.User{Model: gorm.Model{ID: 1}, Email: "testuser@gmail.com", Name: "testuser", Password: "testuser", Roles: false}, nil

}
func (m *MockAuthRepository) LoginAdmin(email, password string) (entities.User, error) {

	return entities.User{Model: gorm.Model{ID: 2}, Email: "testadmin@gmail.com", Name: "testadmin", Password: "testadmin", Roles: true}, nil

}

type MockAuthFailRepository struct{}

func (m *MockAuthFailRepository) Login(email, password string) (entities.User, error) {

	return entities.User{}, errors.New("")

}
func (m *MockAuthFailRepository) LoginAdmin(email, password string) (entities.User, error) {

	return entities.User{}, errors.New("")

}

type MockAuthNotFoundRepository struct{}

func (m *MockAuthNotFoundRepository) Login(email, password string) (entities.User, error) {

	return entities.User{}, errors.New("email not found")

}
func (m *MockAuthNotFoundRepository) LoginAdmin(email, password string) (entities.User, error) {

	return entities.User{}, errors.New("email not found")

}

type MockAuthIncorrectPasswordRepository struct{}

func (m *MockAuthIncorrectPasswordRepository) Login(email, password string) (entities.User, error) {

	return entities.User{}, errors.New("incorrect password")

}
func (m *MockAuthIncorrectPasswordRepository) LoginAdmin(email, password string) (entities.User, error) {

	return entities.User{}, errors.New("incorrect password")

}

type MockAuthFailTokenRepository struct{}

func (m *MockAuthFailTokenRepository) Login(email, password string) (entities.User, error) {

	return entities.User{Model: gorm.Model{ID: 0}, Email: "testuser@gmail.com", Name: "testuser", Password: "testuser", Roles: false}, nil

}
func (m *MockAuthFailTokenRepository) LoginAdmin(email, password string) (entities.User, error) {

	return entities.User{Model: gorm.Model{ID: 0}, Email: "testadmin@gmail.com", Name: "testadmin", Password: "testadmin", Roles: false}, nil
}
