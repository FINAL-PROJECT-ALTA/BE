package goal

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
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/go-playground/assert"
	"github.com/go-playground/validator"
	"github.com/labstack/gommon/log"
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

			jwtTokenAdmin = data["token"].(string)
			log.Info(data)
			log.Info(response)
			assert.Equal(t, "ADMIN - berhasil masuk, mendapatkan token baru", response.Message)
		},
	)
}
func TestCreate(t *testing.T) {

	t.Run("success", func(t *testing.T) {
		e := echo.New()
		e.Validator = &CustomValidator{validator: validator.New()}

		reqBody, _ := json.Marshal(CreateGoalRequest{
			Height:        150,
			Weight:        55,
			Age:           24,
			Daily_active:  "not active",
			Weight_target: 2,
			Range_time:    30,
			Target:        "lose weight",
		})

		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(reqBody))
		res := httptest.NewRecorder()

		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtTokenUser))

		req.Header.Set("Content-Type", "application/json")

		context := e.NewContext(req, res)
		context.SetPath("/users/goals")

		goalController := New(&MockGoalRepository{})

		// goalController.Create()(context)

		err := middleware.JWT([]byte(configs.JWT_SECRET))(goalController.Create())(context)

		if err != nil {
			fmt.Println(err)
			return
		}

		var resp common.Response

		json.Unmarshal([]byte(res.Body.Bytes()), &resp)
		assert.Equal(t, float64(http.StatusCreated), resp.Code)
		assert.Equal(t, "Success create goal", resp.Message)

	})
	t.Run("Failed create", func(t *testing.T) {
		e := echo.New()
		e.Validator = &CustomValidator{validator: validator.New()}

		reqBody, _ := json.Marshal(map[string]interface{}{
			"height":        150,
			"weight":        55,
			"age":           24,
			"daily_active":  "not active",
			"weight_target": 2,
			"range_time":    30,
			"target":        "lose weight",
		})

		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(reqBody))
		res := httptest.NewRecorder()

		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtTokenUser))

		req.Header.Set("Content-Type", "application/json")

		context := e.NewContext(req, res)
		context.SetPath("/users/goals")

		goalController := New(&MockFailedGoalRepository{})

		// goalController.Create()(context)

		err := middleware.JWT([]byte(configs.JWT_SECRET))(goalController.Create())(context)

		if err != nil {
			fmt.Println(err)
			return
		}

		var resp common.Response

		json.Unmarshal([]byte(res.Body.Bytes()), &resp)
		assert.Equal(t, float64(http.StatusInternalServerError), resp.Code)
		assert.Equal(t, "There is some error on server", resp.Message)

	})
	t.Run("Failed bind create", func(t *testing.T) {
		e := echo.New()
		e.Validator = &CustomValidator{validator: validator.New()}

		reqBody, _ := json.Marshal(map[string]interface{}{
			"height":        "",
			"weight":        55,
			"age":           24,
			"daily_active":  "not active",
			"weight_target": 2,
			"range_time":    30,
			"target":        "lose weight",
		})

		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(reqBody))
		res := httptest.NewRecorder()

		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtTokenUser))

		req.Header.Set("Content-Type", "application/json")

		context := e.NewContext(req, res)
		context.SetPath("/users/goals")

		goalController := New(&MockFailedGoalRepository{})

		// goalController.Create()(context)

		err := middleware.JWT([]byte(configs.JWT_SECRET))(goalController.Create())(context)

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
			"height":        150,
			"weight":        55,
			"age":           24,
			"daily_active":  "not active",
			"weight_target": 2,
			"range_time":    30,
			"target":        "lose weight",
		})

		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(reqBody))
		res := httptest.NewRecorder()

		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtTokenAdmin))

		req.Header.Set("Content-Type", "application/json")

		context := e.NewContext(req, res)
		context.SetPath("/users/goals")

		goalController := New(&MockFailedGoalRepository{})

		err := middleware.JWT([]byte(configs.JWT_SECRET))(goalController.Create())(context)

		if err != nil {
			fmt.Println(err)
			return
		}

		var resp common.Response

		json.Unmarshal([]byte(res.Body.Bytes()), &resp)
		assert.Equal(t, float64(http.StatusBadRequest), resp.Code)
		assert.Equal(t, "access denied ", resp.Message)

	})
	t.Run("Failed impossible create", func(t *testing.T) {
		e := echo.New()
		e.Validator = &CustomValidator{validator: validator.New()}

		reqBody, _ := json.Marshal(map[string]interface{}{
			"height":        150,
			"weight":        55,
			"age":           24,
			"daily_active":  "not active",
			"weight_target": 10,
			"range_time":    30,
			"target":        "lose weight",
		})

		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(reqBody))
		res := httptest.NewRecorder()

		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtTokenUser))

		req.Header.Set("Content-Type", "application/json")

		context := e.NewContext(req, res)
		context.SetPath("/users/goals")

		goalController := New(&MockImpossibleGoalRepository{})

		err := middleware.JWT([]byte(configs.JWT_SECRET))(goalController.Create())(context)

		if err != nil {
			fmt.Println(err)
			return
		}

		var resp common.Response

		json.Unmarshal([]byte(res.Body.Bytes()), &resp)
		assert.Equal(t, float64(http.StatusBadRequest), resp.Code)
		assert.Equal(t, "impossible", resp.Message)

	})
}

func TestGetById(t *testing.T) {

	t.Run("success", func(t *testing.T) {
		e := echo.New()
		e.Validator = &CustomValidator{validator: validator.New()}

		req := httptest.NewRequest(http.MethodGet, "/", nil)
		res := httptest.NewRecorder()
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtTokenUser))

		req.Header.Set("Content-Type", "application/json")

		context := e.NewContext(req, res)
		context.SetPath("/users/goals")
		context.SetParamNames("goal_uid")
		context.SetParamValues("xyz")

		goalController := New(&MockGoalRepository{})

		err := middleware.JWT([]byte(configs.JWT_SECRET))(goalController.GetById())(context)

		if err != nil {
			fmt.Println(err)
			return
		}

		var response common.Response

		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		assert.Equal(t, float64(http.StatusOK), response.Code)
		assert.Equal(t, "Success get goal", response.Message)

	})
	t.Run("Failed get goal ById", func(t *testing.T) {
		e := echo.New()
		e.Validator = &CustomValidator{validator: validator.New()}

		req := httptest.NewRequest(http.MethodGet, "/", nil)
		res := httptest.NewRecorder()
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtTokenUser))

		req.Header.Set("Content-Type", "application/json")

		context := e.NewContext(req, res)
		context.SetPath("/users/goals")
		context.SetParamNames("goal_uid")
		context.SetParamValues("xyz")

		goalController := New(&MockFailedGoalRepository{})

		err := middleware.JWT([]byte(configs.JWT_SECRET))(goalController.GetById())(context)

		if err != nil {
			fmt.Println(err)
			return
		}

		var response common.Response

		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		assert.Equal(t, float64(http.StatusInternalServerError), response.Code)
		assert.Equal(t, "There is some error on server", response.Message)

	})
	t.Run("Failed not found get goal ById", func(t *testing.T) {
		e := echo.New()
		e.Validator = &CustomValidator{validator: validator.New()}

		req := httptest.NewRequest(http.MethodGet, "/", nil)
		res := httptest.NewRecorder()
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtTokenUser))

		req.Header.Set("Content-Type", "application/json")

		context := e.NewContext(req, res)
		context.SetPath("/users/goals")
		context.SetParamNames("goal_uid")
		context.SetParamValues("xyz")

		goalController := New(&MockNotFoundGoalRepository{})

		err := middleware.JWT([]byte(configs.JWT_SECRET))(goalController.GetById())(context)

		if err != nil {
			fmt.Println(err)
			return
		}

		var response common.Response

		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		assert.Equal(t, float64(http.StatusNotFound), response.Code)
		assert.Equal(t, "Goal is not found", response.Message)

	})
	t.Run("Failed diff time get goal ById", func(t *testing.T) {
		e := echo.New()
		e.Validator = &CustomValidator{validator: validator.New()}

		req := httptest.NewRequest(http.MethodGet, "/", nil)
		res := httptest.NewRecorder()
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtTokenUser))

		req.Header.Set("Content-Type", "application/json")

		context := e.NewContext(req, res)
		context.SetPath("/users/goals")
		context.SetParamNames("goal_uid")
		context.SetParamValues("xyz")

		goalController := New(&MockDiffTimeGoalRepository{})

		err := middleware.JWT([]byte(configs.JWT_SECRET))(goalController.GetById())(context)

		if err != nil {
			fmt.Println(err)
			return
		}

		var response common.Response

		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		assert.Equal(t, float64(http.StatusOK), response.Code)
		assert.Equal(t, "Success get goal", response.Message)

	})
	t.Run("Failed access get goal ById", func(t *testing.T) {
		e := echo.New()
		e.Validator = &CustomValidator{validator: validator.New()}

		req := httptest.NewRequest(http.MethodGet, "/", nil)
		res := httptest.NewRecorder()
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtTokenAdmin))

		req.Header.Set("Content-Type", "application/json")

		context := e.NewContext(req, res)
		context.SetPath("/users/goals")
		context.SetParamNames("food_uid")
		context.SetParamValues("xyz")

		goalController := New(&MockFailedGoalRepository{})

		err := middleware.JWT([]byte(configs.JWT_SECRET))(goalController.GetById())(context)

		if err != nil {
			fmt.Println(err)
			return
		}

		var response common.Response

		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		assert.Equal(t, float64(http.StatusBadRequest), response.Code)
		assert.Equal(t, "access denied ", response.Message)

	})

}

func TestUpdate(t *testing.T) {

	t.Run("success", func(t *testing.T) {
		e := echo.New()
		e.Validator = &CustomValidator{validator: validator.New()}

		reqBody, _ := json.Marshal(UpdateGoalRequest{
			Height:        150,
			Weight:        55,
			Age:           24,
			Daily_active:  "not active",
			Weight_target: 2,
			Range_time:    30,
			Target:        "lose weight",
		})

		req := httptest.NewRequest(http.MethodPut, "/", bytes.NewBuffer(reqBody))
		res := httptest.NewRecorder()

		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtTokenUser))

		req.Header.Set("Content-Type", "application/json")

		context := e.NewContext(req, res)
		context.SetPath("/users/goals")
		context.SetParamNames("goal_uid")
		context.SetParamValues("xyz")

		goalController := New(&MockGoalRepository{})

		// goalController.Create()(context)

		err := middleware.JWT([]byte(configs.JWT_SECRET))(goalController.Update())(context)

		if err != nil {
			fmt.Println(err)
			return
		}

		var resp common.Response

		json.Unmarshal([]byte(res.Body.Bytes()), &resp)
		assert.Equal(t, float64(http.StatusOK), resp.Code)
		assert.Equal(t, "Success update goal", resp.Message)

	})
	t.Run("failed update food", func(t *testing.T) {
		e := echo.New()
		e.Validator = &CustomValidator{validator: validator.New()}

		reqBody, _ := json.Marshal(UpdateGoalRequest{
			Height:        150,
			Weight:        55,
			Age:           24,
			Daily_active:  "not active",
			Weight_target: 2,
			Range_time:    30,
			Target:        "lose weight",
		})

		req := httptest.NewRequest(http.MethodPut, "/", bytes.NewBuffer(reqBody))
		res := httptest.NewRecorder()

		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtTokenUser))

		req.Header.Set("Content-Type", "application/json")

		context := e.NewContext(req, res)
		context.SetPath("/users/goals")
		context.SetParamNames("goal_uid")
		context.SetParamValues("xyz")

		goalController := New(&MockFailedGoalRepository{})

		// goalController.Create()(context)

		err := middleware.JWT([]byte(configs.JWT_SECRET))(goalController.Update())(context)

		if err != nil {
			fmt.Println(err)
			return
		}

		var resp common.Response

		json.Unmarshal([]byte(res.Body.Bytes()), &resp)
		assert.Equal(t, float64(http.StatusInternalServerError), resp.Code)
		assert.Equal(t, "There is some error on server", resp.Message)

	})
	t.Run("failed bind update goal", func(t *testing.T) {
		e := echo.New()
		e.Validator = &CustomValidator{validator: validator.New()}

		reqBody, _ := json.Marshal(map[string]interface{}{
			"height":        150,
			"weight":        55,
			"age":           24,
			"daily_active":  "not active@",
			"weight_target": 2,
			"range_time":    30,
			"target":        "lose weight",
		})

		req := httptest.NewRequest(http.MethodPut, "/", bytes.NewBuffer(reqBody))
		res := httptest.NewRecorder()

		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtTokenUser))

		req.Header.Set("Content-Type", "application/json")

		context := e.NewContext(req, res)
		context.SetPath("/users/goals")
		context.SetParamNames("goal_uid")
		context.SetParamValues("xyz")

		goalController := New(&MockGoalRepository{})

		// goalController.Create()(context)

		err := middleware.JWT([]byte(configs.JWT_SECRET))(goalController.Update())(context)

		if err != nil {
			fmt.Println(err)
			return
		}

		var resp common.Response

		json.Unmarshal([]byte(res.Body.Bytes()), &resp)
		assert.Equal(t, float64(http.StatusBadRequest), resp.Code)
		assert.Equal(t, "There is some problem from input", resp.Message)

	})
	t.Run("Failed access update", func(t *testing.T) {
		e := echo.New()
		e.Validator = &CustomValidator{validator: validator.New()}

		reqBody, _ := json.Marshal(map[string]interface{}{
			"name":          "makanan",
			"calories":      100,
			"energy":        200,
			"carbohidrate":  300,
			"protein":       400,
			"food_category": "snack",
			"unit":          "ons",
			"unit_value":    1,
		})

		req := httptest.NewRequest(http.MethodPut, "/", bytes.NewBuffer(reqBody))
		res := httptest.NewRecorder()

		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtTokenAdmin))

		req.Header.Set("Content-Type", "application/json")

		context := e.NewContext(req, res)
		context.SetPath("/users/goals")
		context.SetParamNames("goal_uid")
		context.SetParamValues("xyz")

		goalController := New(&MockFailedGoalRepository{})

		err := middleware.JWT([]byte(configs.JWT_SECRET))(goalController.Update())(context)

		if err != nil {
			fmt.Println(err)
			return
		}

		var resp common.Response

		json.Unmarshal([]byte(res.Body.Bytes()), &resp)
		assert.Equal(t, float64(http.StatusBadRequest), resp.Code)
		assert.Equal(t, "access denied ", resp.Message)

	})
}

func TestDelete(t *testing.T) {

	t.Run("success", func(t *testing.T) {
		e := echo.New()
		e.Validator = &CustomValidator{validator: validator.New()}

		req := httptest.NewRequest(http.MethodDelete, "/", nil)
		res := httptest.NewRecorder()

		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtTokenUser))

		req.Header.Set("Content-Type", "application/json")

		context := e.NewContext(req, res)
		context.SetPath("/users/goals")
		context.SetParamNames("goal_uid")
		context.SetParamValues("xyz")

		goalController := New(&MockGoalRepository{})

		// goalController.Create()(context)

		err := middleware.JWT([]byte(configs.JWT_SECRET))(goalController.Delete())(context)

		if err != nil {
			fmt.Println(err)
			return
		}

		var resp common.Response

		json.Unmarshal([]byte(res.Body.Bytes()), &resp)
		assert.Equal(t, float64(http.StatusOK), resp.Code)
		assert.Equal(t, "Success delete goal", resp.Message)

	})
	t.Run("failed delete goal", func(t *testing.T) {
		e := echo.New()
		e.Validator = &CustomValidator{validator: validator.New()}

		req := httptest.NewRequest(http.MethodDelete, "/", nil)
		res := httptest.NewRecorder()

		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtTokenUser))

		req.Header.Set("Content-Type", "application/json")

		context := e.NewContext(req, res)
		context.SetPath("/users/goals")
		context.SetParamNames("goal_uid")
		context.SetParamValues("xyz")

		goalController := New(&MockFailedGoalRepository{})

		// goalController.Create()(context)

		err := middleware.JWT([]byte(configs.JWT_SECRET))(goalController.Delete())(context)

		if err != nil {
			fmt.Println(err)
			return
		}

		var resp common.Response

		json.Unmarshal([]byte(res.Body.Bytes()), &resp)
		assert.Equal(t, float64(http.StatusInternalServerError), resp.Code)
		assert.Equal(t, "There is some error on server", resp.Message)

	})
	t.Run("Failed access delete", func(t *testing.T) {
		e := echo.New()
		e.Validator = &CustomValidator{validator: validator.New()}
		req := httptest.NewRequest(http.MethodDelete, "/", nil)
		res := httptest.NewRecorder()

		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtTokenAdmin))

		req.Header.Set("Content-Type", "application/json")

		context := e.NewContext(req, res)
		context.SetPath("/users/goals")
		context.SetParamNames("goal_uid")
		context.SetParamValues("xyz")

		goalController := New(&MockFailedGoalRepository{})

		err := middleware.JWT([]byte(configs.JWT_SECRET))(goalController.Delete())(context)

		if err != nil {
			fmt.Println(err)
			return
		}

		var resp common.Response

		json.Unmarshal([]byte(res.Body.Bytes()), &resp)
		assert.Equal(t, float64(http.StatusBadRequest), resp.Code)
		assert.Equal(t, "access denied ", resp.Message)

	})

}

func TestGetAll(t *testing.T) {

	t.Run("success", func(t *testing.T) {
		e := echo.New()
		e.Validator = &CustomValidator{validator: validator.New()}

		req := httptest.NewRequest(http.MethodGet, "/", nil)
		res := httptest.NewRecorder()

		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtTokenUser))

		req.Header.Set("Content-Type", "application/json")

		context := e.NewContext(req, res)
		context.SetPath("/users/goals")

		goalController := New(&MockGoalRepository{})
		err := middleware.JWT([]byte(configs.JWT_SECRET))(goalController.GetAll())(context)

		if err != nil {
			fmt.Println(err)
			return
		}

		var response common.Response

		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		assert.Equal(t, float64(http.StatusOK), response.Code)
		assert.Equal(t, "Success get all goal", response.Message)

	})
	t.Run("Failed diff time get all goal", func(t *testing.T) {
		e := echo.New()
		e.Validator = &CustomValidator{validator: validator.New()}

		req := httptest.NewRequest(http.MethodGet, "/", nil)
		res := httptest.NewRecorder()
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtTokenUser))

		req.Header.Set("Content-Type", "application/json")

		context := e.NewContext(req, res)
		context.SetPath("/users/goals")

		goalController := New(&MockDiffTimeGoalRepository{})

		err := middleware.JWT([]byte(configs.JWT_SECRET))(goalController.GetAll())(context)

		if err != nil {
			fmt.Println(err)
			return
		}

		var response common.Response

		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		assert.Equal(t, float64(http.StatusOK), response.Code)
		assert.Equal(t, "Success get all goal", response.Message)

	})
	t.Run("Failed get all goal", func(t *testing.T) {
		e := echo.New()
		e.Validator = &CustomValidator{validator: validator.New()}

		req := httptest.NewRequest(http.MethodGet, "/", nil)
		res := httptest.NewRecorder()

		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtTokenUser))

		req.Header.Set("Content-Type", "application/json")

		context := e.NewContext(req, res)
		context.SetPath("/users/goals")

		goalController := New(&MockFailedGoalRepository{})

		err := middleware.JWT([]byte(configs.JWT_SECRET))(goalController.GetAll())(context)

		if err != nil {
			fmt.Println(err)
			return
		}

		var response common.Response

		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		assert.Equal(t, float64(http.StatusInternalServerError), response.Code)
		assert.Equal(t, "There is some error on server", response.Message)

	})
	t.Run("Failed access delete", func(t *testing.T) {
		e := echo.New()
		e.Validator = &CustomValidator{validator: validator.New()}
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		res := httptest.NewRecorder()

		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtTokenAdmin))

		req.Header.Set("Content-Type", "application/json")

		context := e.NewContext(req, res)
		context.SetPath("/users/goals")

		goalController := New(&MockFailedGoalRepository{})

		err := middleware.JWT([]byte(configs.JWT_SECRET))(goalController.GetAll())(context)

		if err != nil {
			fmt.Println(err)
			return
		}

		var resp common.Response

		json.Unmarshal([]byte(res.Body.Bytes()), &resp)
		assert.Equal(t, float64(http.StatusBadRequest), resp.Code)
		assert.Equal(t, "access denied ", resp.Message)

	})
}

func TestCancelGoal(t *testing.T) {

	t.Run("success", func(t *testing.T) {
		e := echo.New()
		e.Validator = &CustomValidator{validator: validator.New()}

		req := httptest.NewRequest(http.MethodPut, "/", nil)
		res := httptest.NewRecorder()

		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtTokenUser))

		req.Header.Set("Content-Type", "application/json")

		context := e.NewContext(req, res)
		context.SetPath("/users/goals/cancel")

		goalController := New(&MockGoalRepository{})
		err := middleware.JWT([]byte(configs.JWT_SECRET))(goalController.CencelGoal())(context)

		if err != nil {
			fmt.Println(err)
			return
		}

		var response common.Response

		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		assert.Equal(t, float64(http.StatusOK), response.Code)
		assert.Equal(t, "Success cancel goal", response.Message)

	})
	t.Run("Failed cancel goal", func(t *testing.T) {
		e := echo.New()
		e.Validator = &CustomValidator{validator: validator.New()}

		req := httptest.NewRequest(http.MethodPut, "/", nil)
		res := httptest.NewRecorder()

		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtTokenUser))

		req.Header.Set("Content-Type", "application/json")

		context := e.NewContext(req, res)
		context.SetPath("/users/goals/cancel")

		goalController := New(&MockFailedGoalRepository{})

		err := middleware.JWT([]byte(configs.JWT_SECRET))(goalController.CencelGoal())(context)

		if err != nil {
			fmt.Println(err)
			return
		}

		var response common.Response

		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		assert.Equal(t, float64(http.StatusInternalServerError), response.Code)
		assert.Equal(t, "There is some error on server", response.Message)

	})
	t.Run("Failed access cencel goal", func(t *testing.T) {
		e := echo.New()
		e.Validator = &CustomValidator{validator: validator.New()}
		req := httptest.NewRequest(http.MethodPut, "/", nil)
		res := httptest.NewRecorder()

		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtTokenAdmin))

		req.Header.Set("Content-Type", "application/json")

		context := e.NewContext(req, res)
		context.SetPath("/users/goals/cancel")

		goalController := New(&MockFailedGoalRepository{})

		err := middleware.JWT([]byte(configs.JWT_SECRET))(goalController.CencelGoal())(context)

		if err != nil {
			fmt.Println(err)
			return
		}

		var resp common.Response

		json.Unmarshal([]byte(res.Body.Bytes()), &resp)
		assert.Equal(t, float64(http.StatusBadRequest), resp.Code)
		assert.Equal(t, "access denied ", resp.Message)

	})
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

type MockFailedGoalRepository struct{}

func (m *MockFailedGoalRepository) Create(food entities.Goal) (entities.Goal, error) {

	return entities.Goal{}, errors.New("There is some error on server")
}

func (m *MockFailedGoalRepository) GetById(goal_uid string, user_uid string) (entities.Goal, error) {

	return entities.Goal{}, errors.New("There is some error on server")
}

func (m *MockFailedGoalRepository) Update(goal_uid string, newSood entities.Goal) (entities.Goal, error) {

	return entities.Goal{}, errors.New("There is some error on server")
}
func (m *MockFailedGoalRepository) Delete(goal_uid string, user_uid string) error {

	return errors.New("")
}
func (m *MockFailedGoalRepository) GetAll(user_uid string) ([]entities.Goal, error) {

	return []entities.Goal{}, errors.New("There is some error on server")
}
func (m *MockFailedGoalRepository) CancelGoal(user_uid string) (entities.Goal, error) {

	return entities.Goal{}, errors.New("There is some error on server")
}

type MockImpossibleGoalRepository struct{}

func (m *MockImpossibleGoalRepository) Create(food entities.Goal) (entities.Goal, error) {

	return entities.Goal{}, errors.New("impossible")
}

func (m *MockImpossibleGoalRepository) GetById(goal_uid string, user_uid string) (entities.Goal, error) {

	return entities.Goal{}, errors.New("There is some error on server")
}

func (m *MockImpossibleGoalRepository) Update(goal_uid string, newSood entities.Goal) (entities.Goal, error) {

	return entities.Goal{}, errors.New("There is some error on server")
}
func (m *MockImpossibleGoalRepository) Delete(goal_uid string, user_uid string) error {

	return errors.New("")
}
func (m *MockImpossibleGoalRepository) GetAll(user_uid string) ([]entities.Goal, error) {

	return []entities.Goal{}, errors.New("There is some error on server")
}
func (m *MockImpossibleGoalRepository) CancelGoal(user_uid string) (entities.Goal, error) {

	return entities.Goal{}, errors.New("There is some error on server")
}

type MockNotFoundGoalRepository struct{}

func (m *MockNotFoundGoalRepository) Create(food entities.Goal) (entities.Goal, error) {

	return entities.Goal{}, errors.New("impossible")
}

func (m *MockNotFoundGoalRepository) GetById(goal_uid string, user_uid string) (entities.Goal, error) {

	return entities.Goal{}, errors.New("not found")
}

func (m *MockNotFoundGoalRepository) Update(goal_uid string, newSood entities.Goal) (entities.Goal, error) {

	return entities.Goal{}, errors.New("There is some error on server")
}
func (m *MockNotFoundGoalRepository) Delete(goal_uid string, user_uid string) error {

	return errors.New("")
}
func (m *MockNotFoundGoalRepository) GetAll(user_uid string) ([]entities.Goal, error) {

	return []entities.Goal{}, errors.New("There is some error on server")
}
func (m *MockNotFoundGoalRepository) CancelGoal(user_uid string) (entities.Goal, error) {

	return entities.Goal{}, errors.New("There is some error on server")
}

type MockDiffTimeGoalRepository struct{}

func (m *MockDiffTimeGoalRepository) Create(food entities.Goal) (entities.Goal, error) {

	return entities.Goal{}, errors.New("impossible")
}

func (m *MockDiffTimeGoalRepository) GetById(goal_uid string, user_uid string) (entities.Goal, error) {
	time := time.Date(2022, 03, 15, 10, 20, 0, 0, time.UTC)
	return entities.Goal{
		User_uid:      "xyz",
		Height:        150,
		Weight:        55,
		Age:           24,
		Daily_active:  "not active",
		Weight_target: 1,
		Range_time:    30,
		Target:        "lose weight",
		Status:        "active",
		CreatedAt:     time,
	}, nil
}

func (m *MockDiffTimeGoalRepository) Update(goal_uid string, newSood entities.Goal) (entities.Goal, error) {

	return entities.Goal{}, errors.New("There is some error on server")
}
func (m *MockDiffTimeGoalRepository) Delete(goal_uid string, user_uid string) error {

	return errors.New("")
}
func (m *MockDiffTimeGoalRepository) GetAll(user_uid string) ([]entities.Goal, error) {
	time := time.Date(2022, 03, 15, 10, 20, 0, 0, time.UTC)
	return []entities.Goal{{
		User_uid:      "xyz",
		Height:        150,
		Weight:        55,
		Age:           24,
		Daily_active:  "not active",
		Weight_target: 1,
		Range_time:    30,
		Target:        "lose weight",
		Status:        "active",
		CreatedAt:     time,
	}}, nil
}
func (m *MockDiffTimeGoalRepository) CancelGoal(user_uid string) (entities.Goal, error) {

	return entities.Goal{}, errors.New("There is some error on server")
}
