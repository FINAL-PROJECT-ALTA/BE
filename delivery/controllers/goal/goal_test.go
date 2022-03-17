package goal

import (
	"HealthFit/delivery/controllers/auth"
	"HealthFit/delivery/controllers/common"
	"HealthFit/entities"
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

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
type MockAuthRepository struct{}

func (m *MockAuthRepository) Login(email, password string) (entities.User, error) {
	if email == "testuser@gmail.com" {
		return entities.User{Model: gorm.Model{ID: 1}, Email: "testuser@gmail.com", Name: "testuser", Password: "testuser", Roles: false}, nil

	} else if email == "testadmin@gmail.com" {
		return entities.User{Model: gorm.Model{ID: 2}, Email: "testadmin@gmail.com", Name: "testadmin", Password: "testadmin", Roles: true}, nil
	}
	return entities.User{}, errors.New("")

}

func TestLogin(t *testing.T) {
	t.Run(
		"1. Success Login User Test", func(t *testing.T) {
			e := echo.New()
			e.Validator = &CustomValidator{validator: validator.New()}

			requestBody, _ := json.Marshal(
				auth.LoginReqFormat{
					Email:    "testuser@gmail.com",
					Password: "testuser",
				},
			)

			req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
			res := httptest.NewRecorder()
			req.Header.Set("Content-Type", "application/json")
			context := e.NewContext(req, res)
			context.SetPath("/users/login")

			authController := auth.New(&MockAuthRepository{})
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
		"1. Success Login Admin Test", func(t *testing.T) {
			e := echo.New()
			e.Validator = &CustomValidator{validator: validator.New()}

			requestBody, _ := json.Marshal(
				auth.LoginReqFormat{
					Email:    "testadmin@gmail.com",
					Password: "testadmin",
				},
			)

			req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
			res := httptest.NewRecorder()
			req.Header.Set("Content-Type", "application/json")
			context := e.NewContext(req, res)
			context.SetPath("/admin/login")

			authController := auth.New(&MockAuthRepository{})
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

type MockGoalRepository struct{}

func (m *MockGoalRepository) Create(food entities.Goal) (entities.Goal, error) {

	return entities.Goal{
		User_uid:      "xyz",
		Height:        150,
		Weight:        55,
		Age:           24,
		Daily_active:  "not active",
		Weight_target: 2,
		Range_time:    30,
		Target:        "lose weight",
		Status:        "active",
	}, nil
}
func (m *MockGoalRepository) GetById(goal_uid string, user_uid string) (entities.Goal, error) {
	return entities.Goal{
		User_uid:      "xyz",
		Height:        150,
		Weight:        55,
		Age:           24,
		Daily_active:  "not active",
		Weight_target: 2,
		Range_time:    30,
		Target:        "lose weight",
		Status:        "active",
	}, nil

}

func (m *MockGoalRepository) Update(goal_uid string, newSood entities.Goal) (entities.Goal, error) {

	return entities.Goal{
		User_uid:      "xyz",
		Height:        150,
		Weight:        55,
		Age:           24,
		Daily_active:  "not active",
		Weight_target: 2,
		Range_time:    30,
		Target:        "lose weight",
		Status:        "active",
	}, nil
}
func (m *MockGoalRepository) Delete(goal_uid string, user_uid string) error {

	return nil
}
func (m *MockGoalRepository) GetAll(goal_uid string) ([]entities.Goal, error) {

	return []entities.Goal{{
		User_uid:      "xyz",
		Height:        150,
		Weight:        55,
		Age:           24,
		Daily_active:  "not active",
		Weight_target: 2,
		Range_time:    30,
		Target:        "lose weight",
		Status:        "active",
	}}, nil
}
func (m *MockGoalRepository) CancelGoal(user_uid string) (entities.Goal, error) {

	return entities.Goal{}, nil
}
