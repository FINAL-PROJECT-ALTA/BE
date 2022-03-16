package admin

import (
	"HealthFit/delivery/controllers/common"
	"HealthFit/entities"
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"

	// "github.com/labstack/gommon/log"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

type CustomValidator struct {
	validator *validator.Validate
}

var jwtToken = ""

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

//////
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

func TestRegister(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		e := echo.New()
		e.Validator = &CustomValidator{validator: validator.New()}

		reqBody, _ := json.Marshal(CreateAdminRequestFormat{
			Name:     "test",
			Email:    "test@gmail.com",
			Password: "test",
			Gender:   "Male",
		})

		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(reqBody))
		res := httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")

		context := e.NewContext(req, res)
		context.SetPath("/admin/register")

		AdminController := New(&MockAdminRepository{})

		AdminController.Register()(context)

		var resp common.Response

		json.Unmarshal([]byte(res.Body.Bytes()), &resp)
		assert.Equal(t, float64(http.StatusCreated), resp.Code)
		assert.Equal(t, "Success create admin", resp.Message)

	})
}

type MockAdminRepository struct{}

func (m *MockAdminRepository) Register(user entities.User) (entities.User, error) {

	return entities.User{Name: "test", Email: "test@mail.com", Password: "test", Gender: "Male", Roles: true, Image: ""}, nil
}

type MockFailedAdminRepository struct{}

func (m *MockFailedAdminRepository) Register(user entities.User) (entities.User, error) {

	return entities.User{}, errors.New("There is some problem from server")
}
