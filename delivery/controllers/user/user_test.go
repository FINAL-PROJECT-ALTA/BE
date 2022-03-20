package user

import (
	"HealthFit/delivery/controllers/auth"
	"HealthFit/delivery/controllers/common"
	"HealthFit/delivery/middlewares"
	"mime/multipart"

	// "HealthFit/delivery/middlewares"
	"HealthFit/entities"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/stretchr/testify/assert"

	"gorm.io/gorm"
)

type CustomValidator struct {
	validator *validator.Validate
}

var jwtTokenUser = ""
var jwtTokenAdmin = ""

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

type MockAuthLib struct{}

func (m *MockAuthLib) Login(email, password string) (entities.User, error) {
	if email == "testuser@gmail.com" {
		return entities.User{Model: gorm.Model{ID: 1}, Email: "testuser@gmail.com", Name: "testuser", Password: "testuser", Roles: false}, nil

	} else if email == "testadmin@gmail.com" {
		return entities.User{Model: gorm.Model{ID: 2}, Email: "testadmin@gmail.com", Name: "testadmin", Password: "testadmin", Roles: true}, nil
	}
	return entities.User{}, errors.New("")

}

type MockUserLib struct{}

func (m *MockUserLib) Register(newUser entities.User) (entities.User, error) {
	return entities.User{
		User_uid:      "testuser",
		Name:          "testuser",
		Email:         "testuser@mail.com",
		Password:      "testuser",
		Gender:        "male",
		Roles:         false,
		Image:         "",
		Goal:          []entities.Goal{},
		History:       []entities.User_history{},
		Goal_active:   false,
		Goal_exspired: false,
	}, nil

}

func (m *MockUserLib) GetById(userId string) (entities.User, error) {
	return entities.User{}, nil
}

func (m *MockUserLib) Update(userId string, newUser entities.User) (entities.User, error) {
	return entities.User{}, nil
}

func (m *MockUserLib) Delete(Userid string) error {
	return nil
}

type MockFalseLib struct{}

func (m *MockFalseLib) Register(newUser entities.User) (entities.User, error) {
	return entities.User{}, errors.New("invalid input")

}

func (m *MockFalseLib) GetById(userId string) (entities.User, error) {
	return entities.User{}, errors.New("False Object")
}

func (m *MockFalseLib) Update(userId string, newUser entities.User) (entities.User, error) {
	return entities.User{}, errors.New("False Object")
}

func (m *MockFalseLib) Delete(Userid string) error {
	return errors.New("False Object")
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

			authController := auth.New(&MockAuthLib{})
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

			authController := auth.New(&MockAuthLib{})
			authController.AdminLogin()(context)

			var response common.Response

			json.Unmarshal([]byte(res.Body.Bytes()), &response)
			data := (response.Data).(map[string]interface{})
			jwtTokenAdmin = data["token"].(string)
			log.Info(data)
			log.Info(response)
			assert.Equal(t, "ADMIN - berhasil masuk, mendapatkan token baru", response.Message)
		},
	)
}
func TestCreate(t *testing.T) {

	t.Run("success to create", func(t *testing.T) {
		e := echo.New()
		e.Validator = &CustomValidator{validator: validator.New()}

		reqBody, _ := json.Marshal(CreateUserRequestFormat{
			Name:     "testuser",
			Email:    "testuser@mail.com",
			Password: "testuser",
			Gender:   "male",
			Image:    "",
		})

		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(reqBody))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtTokenUser))

		context := e.NewContext(req, res)
		context.SetPath("/users")

		userController := New(&MockUserLib{}, &session.Session{})
		userController.Register()(context)

		response := common.Response{}

		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		assert.Equal(t, float64(201), response.Code)
		assert.Equal(t, "Success Create User", response.Message)

	})

	t.Run("failed to create", func(t *testing.T) {
		e := echo.New()
		e.Validator = &CustomValidator{validator: validator.New()}

		reqBody, _ := json.Marshal(CreateUserRequestFormat{
			Name:     "testuser",
			Password: "testuser",
			Gender:   "male",
			Image:    "",
		})

		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(reqBody))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtTokenUser))

		context := e.NewContext(req, res)
		context.SetPath("/users")

		userController := New(&MockFalseLib{}, &session.Session{})
		userController.Register()(context)

		response := common.Response{}

		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		assert.Equal(t, float64(400), response.Code)
		assert.Equal(t, "There is some problem from input", response.Message)

	})

	t.Run("failed to create", func(t *testing.T) {
		e := echo.New()
		e.Validator = &CustomValidator{validator: validator.New()}

		reqBody, _ := json.Marshal(CreateUserRequestFormat{
			Name:     "testusers",
			Email:    "test@gmail.com",
			Password: "testusers",
			Gender:   "male",
			Image:    "",
		})

		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(reqBody))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		// req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtTokenUser))

		context := e.NewContext(req, res)
		context.SetPath("/users")

		userController := New(&MockFalseLib{}, &session.Session{})
		userController.Register()(context)

		response := common.Response{}

		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		assert.Equal(t, float64(409), response.Code)

	})

	t.Run("failed to create", func(t *testing.T) {
		e := echo.New()
		e.Validator = &CustomValidator{validator: validator.New()}

		reqBody := new(bytes.Buffer)

		reqBodyDetail := multipart.NewWriter(reqBody)
		reqBodyDetail.WriteField("name", "testuser")
		reqBodyDetail.WriteField("email", "testuser@mail.com")
		reqBodyDetail.WriteField("password", "testuser")
		reqBodyDetail.WriteField("gender", "male")

		link, err := reqBodyDetail.CreateFormFile("", "test.jpg")
		if err != nil {
			log.Warn(err)
		}

		link.Write([]byte("sample"))
		reqBodyDetail.Close()

		req := httptest.NewRequest(http.MethodPost, "/", reqBody)
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		// req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtTokenUser))

		context := e.NewContext(req, res)
		context.SetPath("/users")

		userController := New(&MockFalseLib{}, &session.Session{})
		userController.Register()(context)

		response := common.Response{}

		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		assert.Equal(t, float64(400), response.Code)

	})

}

func TestGetByID(t *testing.T) {

	t.Run("failed to get", func(t *testing.T) {
		e := echo.New()
		e.Validator = &CustomValidator{validator: validator.New()}

		req := httptest.NewRequest(http.MethodGet, "/", nil)
		res := httptest.NewRecorder()
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtTokenUser))
		req.Header.Set("Content-Type", "application/json")

		context := e.NewContext(req, res)
		context.SetPath("/users")

		userController := New(&MockFalseLib{}, &session.Session{})

		if err := middlewares.JwtMiddleware()(userController.GetById())(context); err != nil {
			return
		}
		response := common.Response{}
		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		assert.Equal(t, float64(404), response.Code)

	})

	t.Run("success to get", func(t *testing.T) {
		e := echo.New()
		e.Validator = &CustomValidator{validator: validator.New()}

		req := httptest.NewRequest(http.MethodGet, "/", nil)
		res := httptest.NewRecorder()
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtTokenUser))
		req.Header.Set("Content-Type", "application/json")

		context := e.NewContext(req, res)
		context.SetPath("/users")

		userController := New(&MockUserLib{}, &session.Session{})

		if err := middlewares.JwtMiddleware()(userController.GetById())(context); err != nil {
			return
		}
		response := common.Response{}
		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		assert.Equal(t, float64(200), response.Code)
		assert.Equal(t, "Success get user", response.Message)

	})

}

func TestUpdate(t *testing.T) {

	t.Run("failed to update", func(t *testing.T) {
		e := echo.New()
		e.Validator = &CustomValidator{validator: validator.New()}

		reqBody, _ := json.Marshal(map[string]interface{}{
			"name": 123,
		})

		req := httptest.NewRequest(http.MethodPut, "/", bytes.NewBuffer(reqBody))
		res := httptest.NewRecorder()
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtTokenUser))
		req.Header.Set("Content-Type", "application/json")

		context := e.NewContext(req, res)
		context.SetPath("/users")

		userController := New(&MockFalseLib{}, &session.Session{})

		if err := middlewares.JwtMiddleware()(userController.Update())(context); err != nil {
			return
		}
		response := common.Response{}
		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		assert.Equal(t, float64(500), response.Code)

	})

	t.Run("failed to update", func(t *testing.T) {
		e := echo.New()
		e.Validator = &CustomValidator{validator: validator.New()}

		reqBody, _ := json.Marshal(CreateUserRequestFormat{})

		req := httptest.NewRequest(http.MethodPut, "/", bytes.NewBuffer(reqBody))
		res := httptest.NewRecorder()
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtTokenUser))
		req.Header.Set("Content-Type", "application/json")

		context := e.NewContext(req, res)
		context.SetPath("/users")

		userController := New(&MockFalseLib{}, &session.Session{})

		if err := middlewares.JwtMiddleware()(userController.Update())(context); err != nil {
			return
		}
		response := common.Response{}
		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		assert.Equal(t, float64(500), response.Code)

	})

	t.Run("success to update", func(t *testing.T) {
		e := echo.New()
		e.Validator = &CustomValidator{validator: validator.New()}

		reqBody, _ := json.Marshal(CreateUserRequestFormat{
			Name:     "testusers",
			Email:    "test@gmail.com",
			Password: "testusers",
			Gender:   "male",
			Image:    "",
		})

		req := httptest.NewRequest(http.MethodPut, "/", bytes.NewBuffer(reqBody))
		res := httptest.NewRecorder()
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtTokenUser))
		req.Header.Set("Content-Type", "application/json")

		context := e.NewContext(req, res)
		context.SetPath("/users")

		userController := New(&MockUserLib{}, &session.Session{})

		if err := middlewares.JwtMiddleware()(userController.Update())(context); err != nil {
			return
		}
		response := common.Response{}
		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		assert.Equal(t, float64(200), response.Code)
		assert.Equal(t, "Success Update User", response.Message)

	})

}

func TestDelete(t *testing.T) {

	t.Run("failed to get", func(t *testing.T) {
		e := echo.New()
		e.Validator = &CustomValidator{validator: validator.New()}

		req := httptest.NewRequest(http.MethodDelete, "/", nil)
		res := httptest.NewRecorder()
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtTokenUser))
		req.Header.Set("Content-Type", "application/json")

		context := e.NewContext(req, res)
		context.SetPath("/users")

		userController := New(&MockFalseLib{}, &session.Session{})

		if err := middlewares.JwtMiddleware()(userController.Delete())(context); err != nil {
			return
		}
		response := common.Response{}
		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		assert.Equal(t, float64(500), response.Code)

	})

	t.Run("success to get", func(t *testing.T) {
		e := echo.New()
		e.Validator = &CustomValidator{validator: validator.New()}

		req := httptest.NewRequest(http.MethodDelete, "/", nil)
		res := httptest.NewRecorder()
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtTokenUser))
		req.Header.Set("Content-Type", "application/json")

		context := e.NewContext(req, res)
		context.SetPath("/users")

		userController := New(&MockUserLib{}, &session.Session{})

		if err := middlewares.JwtMiddleware()(userController.Delete())(context); err != nil {
			return
		}
		response := common.Response{}
		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		assert.Equal(t, float64(200), response.Code)
		assert.Equal(t, "Success Delete User", response.Message)

	})

}
