package user

// import (
// 	"HealthFit/delivery/controllers/auth"
// 	"HealthFit/delivery/middlewares"
// 	"HealthFit/entities"
// 	"bytes"
// 	"encoding/json"
// 	"errors"
// 	"fmt"
// 	"net/http"
// 	"net/http/httptest"
// 	"testing"

// 	"github.com/labstack/echo/v4"
// 	"github.com/stretchr/testify/assert"

// 	"gorm.io/gorm"
// )

// type MockAuthLib struct{}

// func (m *MockAuthLib) Login(email, password string) (entities.User, error) {
// 	if email != "test@admin.com" && password != "test" {
// 		return entities.User{}, errors.New("record not found")
// 	}
// 	return entities.User{Model: gorm.Model{ID: 1}, Email: email, Password: password}, nil
// }

// type MockAuthAdminLib struct{}

// func (m *MockAuthAdminLib) Login(email, password string) (entities.User, error) {
// 	if email != "admin" && password != "admin" {
// 		return entities.User{}, errors.New("record not found")
// 	}
// 	return entities.User{Model: gorm.Model{ID: 1}, Email: email, Password: password}, nil
// }

// type MockUserLib struct{}

// func (m *MockUserLib) Register(newUser entities.User) (entities.User, error) {
// 	if newUser.Name != "test" && newUser.Email != "test" && newUser.Password != "test" {
// 		return entities.User{}, errors.New("record not found")
// 	}
// 	return entities.User{}, nil

// }

// func (m *MockUserLib) GetById(userId int) (entities.User, error) {
// 	return entities.User{}, nil
// }

// func (m *MockUserLib) Update(userId int, newUser entities.User) (entities.User, error) {
// 	return entities.User{}, nil
// }

// func (m *MockUserLib) Delete(Userid int) error {
// 	return nil
// }

// func (m *MockUserLib) GetAll() ([]entities.User, error) {
// 	return []entities.User{}, nil
// }

// type MockFalseLib struct{}

// func (m *MockFalseLib) Register(newUser entities.User) (entities.User, error) {
// 	if newUser.Name != "test" && newUser.Email != "test" && newUser.Password != "test" {
// 		return entities.User{}, errors.New("record not found")
// 	}
// 	return entities.User{}, errors.New("invalid input")

// }

// func (m *MockFalseLib) GetById(userId int) (entities.User, error) {
// 	return entities.User{}, errors.New("False Object")
// }

// func (m *MockFalseLib) Update(userId int, newUser entities.User) (entities.User, error) {
// 	return entities.User{}, errors.New("False Object")
// }

// func (m *MockFalseLib) Delete(Userid int) error {
// 	return errors.New("False Object")
// }

// func (m *MockFalseLib) GetAll() ([]entities.User, error) {
// 	return []entities.User{}, errors.New("False Object")
// }

// func TestCreate(t *testing.T) {
// 	var jwtToken string

// 	t.Run("Success Login", func(t *testing.T) {
// 		e := echo.New()

// 		reqBody, _ := json.Marshal(map[string]string{
// 			"email":    "test",
// 			"password": "test",
// 		})

// 		req := httptest.NewRequest(http.MethodPost, "/users/login", bytes.NewBuffer(reqBody))
// 		res := httptest.NewRecorder()
// 		req.Header.Set("Content-Type", "application/json")
// 		context := e.NewContext(req, res)

// 		authController := auth.New(&MockAuthLib{})
// 		authController.Login()(context)

// 		response := auth.LoginRespFormat{}
// 		json.Unmarshal([]byte(res.Body.Bytes()), &response)

// 		jwtToken = response.Data["token"].(string)

// 		assert.Equal(t, response.Message, "berhasil masuk, mendapatkan token baru")
// 		assert.NotNil(t, response.Data["token"])
// 	})

// 	t.Run("Failed to Create", func(t *testing.T) {
// 		e := echo.New()

// 		reqBody, _ := json.Marshal(CreateUserRequestFormat{
// 			Name:     "",
// 			Email:    "",
// 			Password: "",
// 		})

// 		req := httptest.NewRequest(http.MethodPost, "/users/register", bytes.NewBuffer(reqBody))
// 		res := httptest.NewRecorder()
// 		req.Header.Set("Content-Type", "application/json")
// 		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))

// 		context := e.NewContext(req, res)
// 		context.SetPath("/users")

// 		userController := New(&MockFalseLib{})
// 		userController.Register()(context)

// 		response := InsertUserResponseFormat{}

// 		json.Unmarshal([]byte(res.Body.Bytes()), &response)

// 		assert.Equal(t, 400, response.Code)
// 		assert.Equal(t, "There is some problem from input", response.Message)

// 	})

// 	t.Run("Failed to Access", func(t *testing.T) {
// 		e := echo.New()

// 		reqBody, _ := json.Marshal(map[string]string{
// 			"email":    "anonim",
// 			"password": "anonim",
// 		})

// 		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(reqBody))
// 		res := httptest.NewRecorder()

// 		req.Header.Set("Content-Type", "application/json")
// 		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))

// 		context := e.NewContext(req, res)
// 		context.SetPath("/users")

// 		userController := New(&MockFalseLib{})
// 		userController.Register()(context)

// 		response := InsertUserResponseFormat{}

// 		json.Unmarshal([]byte(res.Body.Bytes()), &response)

// 		assert.Equal(t, 500, response.Code)
// 		assert.Equal(t, "There is some error on server", response.Message)

// 	})

// 	t.Run("Failed to Access", func(t *testing.T) {
// 		e := echo.New()

// 		reqBody, _ := json.Marshal(map[string]string{
// 			"email":    "test",
// 			"password": "test",
// 			"name":     "test",
// 		})

// 		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(reqBody))
// 		res := httptest.NewRecorder()

// 		req.Header.Set("Content-Type", "application/json")
// 		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))

// 		context := e.NewContext(req, res)
// 		context.SetPath("/users")

// 		userController := New(&MockUserLib{})
// 		userController.Register()(context)

// 		response := InsertUserResponseFormat{}

// 		json.Unmarshal([]byte(res.Body.Bytes()), &response)

// 		assert.Equal(t, 201, response.Code)
// 		assert.Equal(t, "Success Create User", response.Message)

// 	})
// }

// func TestGetById(t *testing.T) {
// 	var jwtToken string

// 	t.Run("Success Login", func(t *testing.T) {
// 		e := echo.New()

// 		reqBody, _ := json.Marshal(map[string]string{
// 			"email":    "test",
// 			"password": "test",
// 		})

// 		req := httptest.NewRequest(http.MethodPost, "/users/login", bytes.NewBuffer(reqBody))
// 		res := httptest.NewRecorder()
// 		req.Header.Set("Content-Type", "application/json")
// 		context := e.NewContext(req, res)

// 		authController := auth.New(&MockAuthLib{})
// 		authController.Login()(context)

// 		response := auth.LoginRespFormat{}
// 		json.Unmarshal([]byte(res.Body.Bytes()), &response)

// 		jwtToken = response.Data["token"].(string)

// 		assert.Equal(t, response.Message, "berhasil masuk, mendapatkan token baru")
// 		assert.NotNil(t, response.Data["token"])
// 	})

// 	t.Run("Fail to Get By Id", func(t *testing.T) {

// 		e := echo.New()

// 		req := httptest.NewRequest(http.MethodGet, "/", bytes.NewBuffer(nil))
// 		res := httptest.NewRecorder()

// 		req.Header.Set("Content-Type", "application/json")
// 		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))

// 		context := e.NewContext(req, res)
// 		context.SetPath("/users/:id")

// 		userController := New(&MockFalseLib{})
// 		if err := middlewares.JwtMiddleware()(userController.GetById())(context); err != nil {
// 			return
// 		}

// 		response := GetUsersResponseFormat{}

// 		json.Unmarshal([]byte(res.Body.Bytes()), &response)
// 		assert.Equal(t, 500, response.Code)
// 		assert.Equal(t, "Not Found", response.Message)
// 	})

// 	t.Run("Success Get By Id", func(t *testing.T) {

// 		e := echo.New()
// 		// userid := int(middlewares.ExtractTokenId(c))

// 		req := httptest.NewRequest(http.MethodGet, "/", bytes.NewBuffer(nil))
// 		res := httptest.NewRecorder()

// 		req.Header.Set("Content-Type", "application/json")
// 		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))

// 		context := e.NewContext(req, res)
// 		context.SetPath("/users/:id")

// 		userController := New(&MockUserLib{})
// 		if err := middlewares.JwtMiddleware()(userController.GetById())(context); err != nil {
// 			return
// 		}

// 		response := GetUsersResponseFormat{}

// 		json.Unmarshal([]byte(res.Body.Bytes()), &response)
// 		assert.Equal(t, 200, response.Code)
// 		assert.Equal(t, "Success Get User", response.Message)
// 	})
// }

// func TestUpdate(t *testing.T) {
// 	var jwtToken string

// 	t.Run("Success Login", func(t *testing.T) {
// 		e := echo.New()

// 		reqBody, _ := json.Marshal(map[string]string{
// 			"email":    "test",
// 			"password": "test",
// 		})

// 		req := httptest.NewRequest(http.MethodPost, "/users/login", bytes.NewBuffer(reqBody))
// 		res := httptest.NewRecorder()
// 		req.Header.Set("Content-Type", "application/json")
// 		context := e.NewContext(req, res)

// 		authController := auth.New(&MockAuthLib{})
// 		authController.Login()(context)

// 		response := auth.LoginRespFormat{}
// 		json.Unmarshal([]byte(res.Body.Bytes()), &response)

// 		jwtToken = response.Data["token"].(string)

// 		assert.Equal(t, response.Message, "berhasil masuk, mendapatkan token baru")
// 		assert.NotNil(t, response.Data["token"])
// 	})

// 	t.Run("Error input Update", func(t *testing.T) {
// 		e := echo.New()

// 		reqBody, _ := json.Marshal(map[string]string{})
// 		req := httptest.NewRequest(http.MethodPut, "/", bytes.NewBuffer(reqBody))
// 		res := httptest.NewRecorder()

// 		req.Header.Set("Content-Type", "application/json")
// 		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", nil))

// 		context := e.NewContext(req, res)
// 		context.SetPath("/users/:id")

// 		userController := New(&MockFalseLib{})
// 		if err := middlewares.JwtMiddleware()(userController.Update())(context); err != nil {
// 			return
// 		}

// 		response := UpdateResponseFormat{}

// 		json.Unmarshal([]byte(res.Body.Bytes()), &response)
// 		assert.Equal(t, 400, response.Code)
// 		assert.Equal(t, "There is some problem from input", response.Message)
// 	})

// 	t.Run("Error Access Database", func(t *testing.T) {
// 		e := echo.New()

// 		reqBody, _ := json.Marshal(map[string]string{
// 			"name":     "test",
// 			"email":    "test",
// 			"password": "test",
// 		})
// 		req := httptest.NewRequest(http.MethodPut, "/", bytes.NewBuffer(reqBody))
// 		res := httptest.NewRecorder()

// 		req.Header.Set("Content-Type", "application/json")
// 		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))

// 		context := e.NewContext(req, res)
// 		context.SetPath("/users/:id")

// 		userController := New(&MockFalseLib{})
// 		if err := middlewares.JwtMiddleware()(userController.Update())(context); err != nil {
// 			return
// 		}

// 		response := UpdateResponseFormat{}

// 		json.Unmarshal([]byte(res.Body.Bytes()), &response)
// 		assert.Equal(t, 500, response.Code)
// 		assert.Equal(t, "There is some error on server", response.Message)
// 	})

// 	t.Run("Error Access Database", func(t *testing.T) {
// 		e := echo.New()

// 		reqBody, _ := json.Marshal(map[string]string{
// 			"name":     "test",
// 			"email":    "test",
// 			"password": "test",
// 		})
// 		req := httptest.NewRequest(http.MethodPut, "/", bytes.NewBuffer(reqBody))
// 		res := httptest.NewRecorder()

// 		req.Header.Set("Content-Type", "application/json")
// 		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))

// 		context := e.NewContext(req, res)
// 		context.SetPath("/users/:id")

// 		userController := New(&MockUserLib{})
// 		if err := middlewares.JwtMiddleware()(userController.Update())(context); err != nil {
// 			return
// 		}

// 		response := UpdateResponseFormat{}

// 		json.Unmarshal([]byte(res.Body.Bytes()), &response)
// 		assert.Equal(t, 200, response.Code)
// 		assert.Equal(t, "Success Update User", response.Message)
// 	})
// }

// func TestDeleteByID(t *testing.T) {
// 	var jwtToken string

// 	t.Run("Success Login", func(t *testing.T) {
// 		e := echo.New()
// 		reqBody, _ := json.Marshal(map[string]string{
// 			"email":    "test",
// 			"password": "test",
// 		})

// 		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(reqBody))
// 		res := httptest.NewRecorder()
// 		req.Header.Set("Content-Type", "application/json")
// 		context := e.NewContext(req, res)
// 		context.SetPath("/login")

// 		authController := auth.New(&MockAuthLib{})
// 		authController.Login()(context)

// 		response := auth.LoginRespFormat{}
// 		json.Unmarshal([]byte(res.Body.Bytes()), &response)

// 		jwtToken = response.Data["token"].(string)

// 		assert.Equal(t, response.Message, "berhasil masuk, mendapatkan token baru")
// 		assert.NotNil(t, response.Data["token"])
// 	})

// 	t.Run("Fail to Delete", func(t *testing.T) {
// 		e := echo.New()
// 		reqBody, _ := json.Marshal(map[string]string{})
// 		req := httptest.NewRequest(http.MethodDelete, "/", bytes.NewBuffer(reqBody))
// 		res := httptest.NewRecorder()

// 		req.Header.Set("Content-Type", "application/json")
// 		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", nil))

// 		context := e.NewContext(req, res)
// 		context.SetPath("/users/:id")

// 		userController := New(&MockFalseLib{})
// 		if err := middlewares.JwtMiddleware()(userController.Delete())(context); err != nil {
// 			return
// 		}

// 		response := DeleteUserResponseFormat{}

// 		json.Unmarshal([]byte(res.Body.Bytes()), &response)
// 		assert.Equal(t, 400, response.Code)
// 		assert.Equal(t, "error in request Delete", response.Message)

// 	})

// 	t.Run("Fail to access Delete", func(t *testing.T) {
// 		e := echo.New()
// 		reqBody, _ := json.Marshal(map[string]string{
// 			"name":     "test",
// 			"email":    "test",
// 			"password": "test",
// 		})
// 		req := httptest.NewRequest(http.MethodDelete, "/", bytes.NewBuffer(reqBody))
// 		res := httptest.NewRecorder()

// 		req.Header.Set("Content-Type", "application/json")
// 		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))

// 		context := e.NewContext(req, res)
// 		context.SetPath("/users/:id")

// 		userController := New(&MockFalseLib{})
// 		if err := middlewares.JwtMiddleware()(userController.Delete())(context); err != nil {
// 			return
// 		}

// 		response := DeleteUserResponseFormat{}

// 		json.Unmarshal([]byte(res.Body.Bytes()), &response)
// 		assert.Equal(t, 500, response.Code)
// 		assert.Equal(t, "There is some error on server", response.Message)

// 	})

// 	t.Run("Success Delete", func(t *testing.T) {
// 		e := echo.New()
// 		reqBody, _ := json.Marshal(map[string]string{
// 			"name":     "test",
// 			"email":    "test",
// 			"password": "test",
// 		})
// 		req := httptest.NewRequest(http.MethodDelete, "/", bytes.NewBuffer(reqBody))
// 		res := httptest.NewRecorder()

// 		req.Header.Set("Content-Type", "application/json")
// 		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))

// 		context := e.NewContext(req, res)
// 		context.SetPath("/users/:id")

// 		userController := New(&MockUserLib{})
// 		if err := middlewares.JwtMiddleware()(userController.Delete())(context); err != nil {
// 			return
// 		}

// 		response := DeleteUserResponseFormat{}

// 		json.Unmarshal([]byte(res.Body.Bytes()), &response)
// 		assert.Equal(t, 200, response.Code)
// 		assert.Equal(t, "Success Delete User", response.Message)

// 	})

// }
