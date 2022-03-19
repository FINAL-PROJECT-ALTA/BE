package userhistory

import (
	"HealthFit/configs"
	"HealthFit/delivery/controllers/auth"
	"HealthFit/delivery/controllers/common"
	"HealthFit/entities"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	// "github.com/aws/aws-sdk-go/aws/session"

	"github.com/labstack/echo/v4"

	"github.com/labstack/echo/v4/middleware"

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
func TestInsert(t *testing.T) {

	t.Run("success", func(t *testing.T) {
		e := echo.New()
		e.Validator = &CustomValidator{validator: validator.New()}

		reqBody, _ := json.Marshal(CreateUserHistoryRequestFormat{
			Menu_uid: "xyzmenu",
			Goal_uid: "xyzgoal",
		})

		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(reqBody))
		res := httptest.NewRecorder()

		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtTokenUser))

		req.Header.Set("Content-Type", "application/json")

		context := e.NewContext(req, res)
		context.SetPath("/userhistories")

		userhistoryController := New(&MockUserHistoryRepository{})

		err := middleware.JWT([]byte(configs.JWT_SECRET))(userhistoryController.Insert())(context)

		if err != nil {
			fmt.Println(err)
			return
		}

		var resp common.Response

		json.Unmarshal([]byte(res.Body.Bytes()), &resp)
		assert.Equal(t, float64(http.StatusCreated), resp.Code)
		assert.Equal(t, "Success Create User History", resp.Message)

	})
	t.Run("Failed create", func(t *testing.T) {
		e := echo.New()
		e.Validator = &CustomValidator{validator: validator.New()}

		reqBody, _ := json.Marshal(map[string]interface{}{
			"menu_uid": "xyzmenu",
			"goal_uid": "xyzgoal",
		})

		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(reqBody))
		res := httptest.NewRecorder()

		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtTokenUser))

		req.Header.Set("Content-Type", "application/json")

		context := e.NewContext(req, res)
		context.SetPath("/userhistories")

		userhistoryController := New(&MockFailedUserHistoryRepository{})

		err := middleware.JWT([]byte(configs.JWT_SECRET))(userhistoryController.Insert())(context)

		if err != nil {
			fmt.Println(err)
			return
		}

		var resp common.Response

		json.Unmarshal([]byte(res.Body.Bytes()), &resp)
		assert.Equal(t, float64(http.StatusInternalServerError), resp.Code)
		assert.Equal(t, "Internal Server Error", resp.Message)

	})
	t.Run("Failed bind create", func(t *testing.T) {
		e := echo.New()
		e.Validator = &CustomValidator{validator: validator.New()}

		reqBody, _ := json.Marshal(map[string]interface{}{
			"menu_uid": "",
			"goal_uid": "xyzgoal",
		})

		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(reqBody))
		res := httptest.NewRecorder()

		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtTokenUser))

		req.Header.Set("Content-Type", "application/json")

		context := e.NewContext(req, res)
		context.SetPath("/userhistories")

		userhistoryController := New(&MockFailedUserHistoryRepository{})

		err := middleware.JWT([]byte(configs.JWT_SECRET))(userhistoryController.Insert())(context)

		if err != nil {
			fmt.Println(err)
			return
		}

		var resp common.Response

		json.Unmarshal([]byte(res.Body.Bytes()), &resp)
		assert.Equal(t, float64(http.StatusBadRequest), resp.Code)
		assert.Equal(t, "There is some problem from input", resp.Message)

	})
	t.Run("Failed access create", func(t *testing.T) {
		e := echo.New()
		e.Validator = &CustomValidator{validator: validator.New()}

		reqBody, _ := json.Marshal(map[string]interface{}{
			"menu_uid": "xyzmenu",
			"goal_uid": "xyzgoal",
		})

		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(reqBody))
		res := httptest.NewRecorder()

		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtTokenAdmin))

		req.Header.Set("Content-Type", "application/json")

		context := e.NewContext(req, res)
		context.SetPath("/userhistories")

		userhistoryController := New(&MockFailedUserHistoryRepository{})

		err := middleware.JWT([]byte(configs.JWT_SECRET))(userhistoryController.Insert())(context)

		if err != nil {
			fmt.Println(err)
			return
		}

		var resp common.Response

		json.Unmarshal([]byte(res.Body.Bytes()), &resp)
		assert.Equal(t, float64(http.StatusUnauthorized), resp.Code)
		assert.Equal(t, "access denied", resp.Message)

	})
}

func TestGetByUid(t *testing.T) {

	t.Run("success", func(t *testing.T) {
		e := echo.New()
		e.Validator = &CustomValidator{validator: validator.New()}

		req := httptest.NewRequest(http.MethodGet, "/", nil)
		res := httptest.NewRecorder()
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtTokenUser))

		req.Header.Set("Content-Type", "application/json")

		context := e.NewContext(req, res)
		context.SetPath("/userhistories")
		context.SetParamNames("user_history_uid")
		context.SetParamValues("xyz")

		userhistoryController := New(&MockUserHistoryRepository{})

		err := middleware.JWT([]byte(configs.JWT_SECRET))(userhistoryController.GetByUid())(context)

		if err != nil {
			fmt.Println(err)
			return
		}

		var response common.Response

		json.Unmarshal([]byte(res.Body.Bytes()), &response)
		// data := (response.Data).(map[string]interface{})
		// log.Info(data)
		// log.Info(response)
		assert.Equal(t, float64(http.StatusOK), response.Code)
		assert.Equal(t, "Success get user", response.Message)

	})
	t.Run("Failed get user history uid", func(t *testing.T) {
		e := echo.New()
		e.Validator = &CustomValidator{validator: validator.New()}

		req := httptest.NewRequest(http.MethodGet, "/", nil)
		res := httptest.NewRecorder()
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtTokenUser))

		req.Header.Set("Content-Type", "application/json")

		context := e.NewContext(req, res)
		context.SetPath("/userhistories")
		context.SetParamNames("user_history_uid")
		context.SetParamValues("xyz")

		userhistoryController := New(&MockFailedUserHistoryRepository{})

		err := middleware.JWT([]byte(configs.JWT_SECRET))(userhistoryController.GetByUid())(context)

		if err != nil {
			fmt.Println(err)
			return
		}

		var response common.Response

		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		assert.Equal(t, float64(http.StatusInternalServerError), response.Code)
		assert.Equal(t, "Internal Server Error", response.Message)

	})
	t.Run("Failed access get by userhistory uid", func(t *testing.T) {
		e := echo.New()
		e.Validator = &CustomValidator{validator: validator.New()}

		reqBody, _ := json.Marshal(map[string]interface{}{
			"menu_uid": "xyzmenu",
			"goal_uid": "xyzgoal",
		})

		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(reqBody))
		res := httptest.NewRecorder()

		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtTokenAdmin))

		req.Header.Set("Content-Type", "application/json")

		context := e.NewContext(req, res)
		context.SetPath("/userhistories")

		userhistoryController := New(&MockFailedUserHistoryRepository{})

		err := middleware.JWT([]byte(configs.JWT_SECRET))(userhistoryController.GetByUid())(context)

		if err != nil {
			fmt.Println(err)
			return
		}

		var resp common.Response

		json.Unmarshal([]byte(res.Body.Bytes()), &resp)
		assert.Equal(t, float64(http.StatusUnauthorized), resp.Code)
		assert.Equal(t, "access denied", resp.Message)

	})
}
func TestGeAll(t *testing.T) {

	t.Run("success", func(t *testing.T) {
		e := echo.New()
		e.Validator = &CustomValidator{validator: validator.New()}

		req := httptest.NewRequest(http.MethodGet, "/", nil)
		res := httptest.NewRecorder()
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtTokenUser))

		req.Header.Set("Content-Type", "application/json")

		context := e.NewContext(req, res)
		context.SetPath("/userhistories")

		userhistoryController := New(&MockUserHistoryRepository{})

		err := middleware.JWT([]byte(configs.JWT_SECRET))(userhistoryController.GetAll())(context)

		if err != nil {
			fmt.Println(err)
			return
		}

		var response common.Response

		json.Unmarshal([]byte(res.Body.Bytes()), &response)
		// data := (response.Data).(map[string]interface{})
		// log.Info(data)
		// log.Info(response)
		assert.Equal(t, float64(http.StatusOK), response.Code)
		assert.Equal(t, "Success get user histories", response.Message)

	})
	t.Run("Failed get all history", func(t *testing.T) {
		e := echo.New()
		e.Validator = &CustomValidator{validator: validator.New()}

		req := httptest.NewRequest(http.MethodGet, "/", nil)
		res := httptest.NewRecorder()
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtTokenUser))

		req.Header.Set("Content-Type", "application/json")

		context := e.NewContext(req, res)
		context.SetPath("/userhistories")

		userhistoryController := New(&MockFailedUserHistoryRepository{})

		err := middleware.JWT([]byte(configs.JWT_SECRET))(userhistoryController.GetAll())(context)

		if err != nil {
			fmt.Println(err)
			return
		}

		var response common.Response

		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		assert.Equal(t, float64(http.StatusInternalServerError), response.Code)
		assert.Equal(t, "There is some error on server", response.Message)

	})
	t.Run("Failed access get by userhistory uid", func(t *testing.T) {
		e := echo.New()
		e.Validator = &CustomValidator{validator: validator.New()}

		reqBody, _ := json.Marshal(map[string]interface{}{
			"menu_uid": "xyzmenu",
			"goal_uid": "xyzgoal",
		})

		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(reqBody))
		res := httptest.NewRecorder()

		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtTokenAdmin))

		req.Header.Set("Content-Type", "application/json")

		context := e.NewContext(req, res)
		context.SetPath("/userhistories")

		userhistoryController := New(&MockFailedUserHistoryRepository{})

		err := middleware.JWT([]byte(configs.JWT_SECRET))(userhistoryController.GetAll())(context)

		if err != nil {
			fmt.Println(err)
			return
		}

		var resp common.Response

		json.Unmarshal([]byte(res.Body.Bytes()), &resp)
		assert.Equal(t, float64(http.StatusUnauthorized), resp.Code)
		assert.Equal(t, "access denied", resp.Message)

	})

}

type MockUserHistoryRepository struct{}

func (m *MockUserHistoryRepository) Insert(user_history entities.User_history) (entities.User_history, error) {

	return entities.User_history{
		Menu_uid:         "xyzmenu",
		User_uid:         "xyzuser",
		User_history_uid: "xyzuserhistory",
		Goal_uid:         "xyzgoal",
	}, nil
}
func (m *MockUserHistoryRepository) GetById(user_uid, user_history_uid string) (entities.User_history, error) {
	return entities.User_history{
		Menu_uid: "xyzmenu",
		Menu: entities.Menu{
			User_uid:      "xyz",
			Menu_category: "food",
			Created_by:    "admin",
			Detail_menu: []entities.Detail_menu{
				{Food: entities.Food{Food_uid: "JSKJDSK"}},
				{Food: entities.Food{Food_uid: "sfsaafsd"}},
			},
		},
		User_uid:         "xyzuser",
		User_history_uid: "xyzuserhistory",
		Goal_uid:         "xyzgoal",
	}, nil

}

func (m *MockUserHistoryRepository) GetAll(user_uid string) ([]entities.User_history, error) {

	return []entities.User_history{{
		Menu_uid: "xyzmenu",
		Menu: entities.Menu{
			User_uid:      "xyz",
			Menu_category: "food",
			Created_by:    "admin",
			Detail_menu: []entities.Detail_menu{
				{Food: entities.Food{Food_uid: "JSKJDSK"}},
				{Food: entities.Food{Food_uid: "sfsaafsd"}},
			},
		},
		User_uid:         "xyzuser",
		User_history_uid: "xyzuserhistory",
		Goal_uid:         "xyzgoal",
	}}, nil
}

type MockFailedUserHistoryRepository struct{}

func (m *MockFailedUserHistoryRepository) Insert(user_history entities.User_history) (entities.User_history, error) {

	return entities.User_history{}, errors.New("There is some error on server")
}

func (m *MockFailedUserHistoryRepository) GetById(user_uid, user_history_uid string) (entities.User_history, error) {

	return entities.User_history{}, errors.New("There is some error on server")
}

func (m *MockFailedUserHistoryRepository) GetAll(user_uid string) ([]entities.User_history, error) {

	return []entities.User_history{}, errors.New("There is some error on server")
}
