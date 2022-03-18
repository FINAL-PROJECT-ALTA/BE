package menu

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

func TestCreate(t *testing.T) {

	t.Run("success create menu by user", func(t *testing.T) {
		e := echo.New()
		e.Validator = &CustomValidator{validator: validator.New()}

		reqBody, _ := json.Marshal(map[string]interface{}{
			"menu_category": "food",
			"foods":         []entities.Food{{Food_uid: "xyz"}},
		})

		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(reqBody))
		// log.Info(req)

		res := httptest.NewRecorder()

		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtTokenUser))

		req.Header.Set("Content-Type", "application/json")

		context := e.NewContext(req, res)
		context.SetPath("/menus")

		goalController := New(&MockMenuRepository{})

		// goalController.Create()(context)

		err := middleware.JWT([]byte(configs.JWT_SECRET))(goalController.Create())(context)

		if err != nil {
			fmt.Println(err)
			return
		}

		var resp common.Response

		json.Unmarshal([]byte(res.Body.Bytes()), &resp)
		assert.Equal(t, float64(http.StatusCreated), resp.Code)
		assert.Equal(t, "Success create menu", resp.Message)

	})

	t.Run("success create menu by admin", func(t *testing.T) {
		e := echo.New()
		e.Validator = &CustomValidator{validator: validator.New()}

		reqBody, _ := json.Marshal(map[string]interface{}{
			"menu_category": "food",
			"foods":         []entities.Food{{Food_uid: "xyz"}},
		})

		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(reqBody))
		// log.Info(req)

		res := httptest.NewRecorder()

		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtTokenAdmin))

		req.Header.Set("Content-Type", "application/json")

		context := e.NewContext(req, res)
		context.SetPath("/menus")

		goalController := New(&MockMenuAdminRepository{})

		// goalController.Create()(context)

		err := middleware.JWT([]byte(configs.JWT_SECRET))(goalController.Create())(context)

		if err != nil {
			fmt.Println(err)
			return
		}

		var resp common.Response

		json.Unmarshal([]byte(res.Body.Bytes()), &resp)
		assert.Equal(t, float64(http.StatusCreated), resp.Code)
		assert.Equal(t, "Success create menu", resp.Message)

	})

	t.Run("Failed create by user", func(t *testing.T) {
		e := echo.New()
		e.Validator = &CustomValidator{validator: validator.New()}

		reqBody, _ := json.Marshal(map[string]interface{}{
			"menu_category": "food",
			"foods":         []entities.Food{},
		})

		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(reqBody))
		res := httptest.NewRecorder()

		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtTokenUser))

		req.Header.Set("Content-Type", "application/json")

		context := e.NewContext(req, res)
		context.SetPath("/menus")

		goalController := New(&MockFailedMenuRepository{})

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

	t.Run("Failed create by admin", func(t *testing.T) {
		e := echo.New()
		e.Validator = &CustomValidator{validator: validator.New()}

		reqBody, _ := json.Marshal(map[string]interface{}{
			"menu_category": "food",
			"foods":         []entities.Food{},
		})

		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(reqBody))
		res := httptest.NewRecorder()

		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtTokenAdmin))

		req.Header.Set("Content-Type", "application/json")

		context := e.NewContext(req, res)
		context.SetPath("/menus")

		goalController := New(&MockFailedMenuRepository{})

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

	t.Run("Failed bind create by user", func(t *testing.T) {
		e := echo.New()
		e.Validator = &CustomValidator{validator: validator.New()}

		reqBody, _ := json.Marshal(map[string]interface{}{
			"menu_category": "",
			"foods":         []entities.Food{},
		})

		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(reqBody))
		res := httptest.NewRecorder()

		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtTokenUser))

		req.Header.Set("Content-Type", "application/json")

		context := e.NewContext(req, res)
		context.SetPath("/menus")

		goalController := New(&MockFailedMenuRepository{})

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

	t.Run("Failed bind create by admin", func(t *testing.T) {
		e := echo.New()
		e.Validator = &CustomValidator{validator: validator.New()}

		reqBody, _ := json.Marshal(map[string]interface{}{
			"menu_category": "",
			"foods":         []entities.Food{},
		})

		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(reqBody))
		res := httptest.NewRecorder()

		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtTokenAdmin))

		req.Header.Set("Content-Type", "application/json")

		context := e.NewContext(req, res)
		context.SetPath("/menus")

		goalController := New(&MockFailedMenuRepository{})

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
}

func TestUpdate(t *testing.T) {

	t.Run("success", func(t *testing.T) {
		e := echo.New()
		e.Validator = &CustomValidator{validator: validator.New()}

		reqBody, _ := json.Marshal(map[string]interface{}{
			"menu_category": "food",
			"foods":         entities.Food{},
		})

		req := httptest.NewRequest(http.MethodPut, "/", bytes.NewBuffer(reqBody))
		res := httptest.NewRecorder()

		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtTokenAdmin))

		req.Header.Set("Content-Type", "application/json")

		context := e.NewContext(req, res)
		context.SetPath("/menus")
		context.SetParamNames("menu_uid")
		context.SetParamValues("xyz")

		goalController := New(&MockMenuRepository{})

		// goalController.Create()(context)

		err := middleware.JWT([]byte(configs.JWT_SECRET))(goalController.Update())(context)

		if err != nil {
			fmt.Println(err)
			return
		}

		var resp common.Response

		json.Unmarshal([]byte(res.Body.Bytes()), &resp)
		assert.Equal(t, float64(http.StatusOK), resp.Code)
		assert.Equal(t, "Success update menu", resp.Message)

	})
	t.Run("failed update menu", func(t *testing.T) {
		e := echo.New()
		e.Validator = &CustomValidator{validator: validator.New()}

		reqBody, _ := json.Marshal(map[string]interface{}{
			"menu_category": "food",
			"foods":         entities.Food{Food_uid: "xyz"},
		})

		req := httptest.NewRequest(http.MethodPut, "/", bytes.NewBuffer(reqBody))
		res := httptest.NewRecorder()

		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtTokenAdmin))

		req.Header.Set("Content-Type", "application/json")

		context := e.NewContext(req, res)
		context.SetPath("/menus")
		context.SetParamNames("menu_uid")
		context.SetParamValues("xyz")

		goalController := New(&MockFailedMenuRepository{})

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
	t.Run("failed bind update menu", func(t *testing.T) {
		e := echo.New()
		e.Validator = &CustomValidator{validator: validator.New()}

		reqBody, _ := json.Marshal(map[string]interface{}{
			"menu_category": "food@",
			"foods":         entities.Food{Food_uid: "xyz"},
		})

		req := httptest.NewRequest(http.MethodPut, "/", bytes.NewBuffer(reqBody))
		res := httptest.NewRecorder()

		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtTokenAdmin))

		req.Header.Set("Content-Type", "application/json")

		context := e.NewContext(req, res)
		context.SetPath("/menus")
		context.SetParamNames("menu_uid)")
		context.SetParamValues("xyz")

		goalController := New(&MockMenuRepository{})

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
			"menu_category": "food",
			"foods":         entities.Food{Food_uid: "xyz"},
		})
		req := httptest.NewRequest(http.MethodPut, "/", bytes.NewBuffer(reqBody))
		res := httptest.NewRecorder()

		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtTokenUser))

		req.Header.Set("Content-Type", "application/json")

		context := e.NewContext(req, res)
		context.SetPath("/menus")
		context.SetParamNames("menu_uid")
		context.SetParamValues("xyz")

		goalController := New(&MockFailedMenuRepository{})

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

		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtTokenAdmin))

		req.Header.Set("Content-Type", "application/json")

		context := e.NewContext(req, res)
		context.SetPath("/menus")
		context.SetParamNames("menu_uid")
		context.SetParamValues("xyz")

		goalController := New(&MockMenuRepository{})

		// goalController.Create()(context)

		err := middleware.JWT([]byte(configs.JWT_SECRET))(goalController.Delete())(context)

		if err != nil {
			fmt.Println(err)
			return
		}

		var resp common.Response

		json.Unmarshal([]byte(res.Body.Bytes()), &resp)
		assert.Equal(t, float64(http.StatusOK), resp.Code)
		assert.Equal(t, "Success delete menu", resp.Message)

	})
	t.Run("failed delete menu", func(t *testing.T) {
		e := echo.New()
		e.Validator = &CustomValidator{validator: validator.New()}

		req := httptest.NewRequest(http.MethodDelete, "/", nil)
		res := httptest.NewRecorder()

		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtTokenAdmin))

		req.Header.Set("Content-Type", "application/json")

		context := e.NewContext(req, res)
		context.SetPath("/menus")
		context.SetParamNames("menu_uid")
		context.SetParamValues("xyz")

		goalController := New(&MockFailedMenuRepository{})

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

		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtTokenUser))

		req.Header.Set("Content-Type", "application/json")

		context := e.NewContext(req, res)
		context.SetPath("/menus")
		context.SetParamNames("menu_uid")
		context.SetParamValues("xyz")

		goalController := New(&MockFailedMenuRepository{})

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
func TestGetAllMenu(t *testing.T) {

	t.Run("success", func(t *testing.T) {
		e := echo.New()
		e.Validator = &CustomValidator{validator: validator.New()}

		req := httptest.NewRequest(http.MethodGet, "/", nil)
		res := httptest.NewRecorder()

		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtTokenUser))

		req.Header.Set("Content-Type", "application/json")

		context := e.NewContext(req, res)
		context.SetPath("/menus")
		context.Set("ctegory", "food")
		context.Set("createdBy", "admin")

		goalController := New(&MockMenuRepository{})
		err := middleware.JWT([]byte(configs.JWT_SECRET))(goalController.GetAll())(context)

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
		assert.Equal(t, "Success Get All Menu ", response.Message)

	})

	t.Run("Failed get all menu", func(t *testing.T) {
		e := echo.New()
		e.Validator = &CustomValidator{validator: validator.New()}

		req := httptest.NewRequest(http.MethodGet, "/", nil)
		res := httptest.NewRecorder()

		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtTokenUser))

		req.Header.Set("Content-Type", "application/json")

		context := e.NewContext(req, res)
		context.SetPath("/menus")
		context.Set("ctegory", "food")
		context.Set("createdBy", "admin")

		goalController := New(&MockFailedMenuRepository{})

		err := middleware.JWT([]byte(configs.JWT_SECRET))(goalController.GetAll())(context)

		if err != nil {
			fmt.Println(err)
			return
		}

		var response common.Response

		json.Unmarshal([]byte(res.Body.Bytes()), &response)
		// data := (response.Data).(map[string]interface{})
		// log.Info(data)
		// log.Info(response)
		assert.Equal(t, float64(http.StatusInternalServerError), response.Code)
		assert.Equal(t, "There is some error on server", response.Message)

	})

}
func TestGetRecommendBreakfast(t *testing.T) {

	t.Run("success", func(t *testing.T) {
		e := echo.New()
		e.Validator = &CustomValidator{validator: validator.New()}

		req := httptest.NewRequest(http.MethodGet, "/", nil)
		res := httptest.NewRecorder()

		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtTokenUser))

		req.Header.Set("Content-Type", "application/json")

		context := e.NewContext(req, res)
		context.SetPath("/menus/recommend/breakfast")

		goalController := New(&MockMenuRepository{})
		err := middleware.JWT([]byte(configs.JWT_SECRET))(goalController.GetRecommendBreakfast())(context)

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
		assert.Equal(t, "Success Get Menu Recommended", response.Message)

	})

	t.Run("Failed get all menu recommended", func(t *testing.T) {
		e := echo.New()
		e.Validator = &CustomValidator{validator: validator.New()}

		req := httptest.NewRequest(http.MethodGet, "/", nil)
		res := httptest.NewRecorder()

		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtTokenUser))

		req.Header.Set("Content-Type", "application/json")

		context := e.NewContext(req, res)
		context.SetPath("/menus/recommend/breakfast")

		goalController := New(&MockFailedMenuRepository{})

		err := middleware.JWT([]byte(configs.JWT_SECRET))(goalController.GetRecommendBreakfast())(context)

		if err != nil {
			fmt.Println(err)
			return
		}

		var response common.Response

		json.Unmarshal([]byte(res.Body.Bytes()), &response)
		// data := (response.Data).(map[string]interface{})
		// log.Info(data)
		// log.Info(response)
		assert.Equal(t, float64(http.StatusInternalServerError), response.Code)
		assert.Equal(t, "There is some error on server", response.Message)

	})
	t.Run("Failed access get all menu recommended", func(t *testing.T) {
		e := echo.New()
		e.Validator = &CustomValidator{validator: validator.New()}

		req := httptest.NewRequest(http.MethodGet, "/", nil)
		res := httptest.NewRecorder()

		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtTokenAdmin))

		req.Header.Set("Content-Type", "application/json")

		context := e.NewContext(req, res)
		context.SetPath("/menus/recommend/breakfast")

		goalController := New(&MockFailedMenuRepository{})

		err := middleware.JWT([]byte(configs.JWT_SECRET))(goalController.GetRecommendBreakfast())(context)

		if err != nil {
			fmt.Println(err)
			return
		}

		var response common.Response

		json.Unmarshal([]byte(res.Body.Bytes()), &response)
		// data := (response.Data).(map[string]interface{})
		// log.Info(data)
		// log.Info(response)
		assert.Equal(t, float64(http.StatusBadRequest), response.Code)
		assert.Equal(t, "access denied", response.Message)

	})
	t.Run("impossible", func(t *testing.T) {
		e := echo.New()
		e.Validator = &CustomValidator{validator: validator.New()}

		req := httptest.NewRequest(http.MethodGet, "/", nil)
		res := httptest.NewRecorder()

		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtTokenUser))

		req.Header.Set("Content-Type", "application/json")

		context := e.NewContext(req, res)
		context.SetPath("/menus/recommend/breakfast")

		goalController := New(&MockImpossibleMenuRepository{})
		err := middleware.JWT([]byte(configs.JWT_SECRET))(goalController.GetRecommendBreakfast())(context)

		if err != nil {
			fmt.Println(err)
			return
		}

		var response common.Response

		json.Unmarshal([]byte(res.Body.Bytes()), &response)
		// data := (response.Data).(map[string]interface{})
		// log.Info(data)
		// log.Info(response)
		assert.Equal(t, float64(http.StatusBadRequest), response.Code)
		assert.Equal(t, "impossible", response.Message)

	})

}

func TestGetRecommendLunch(t *testing.T) {

	t.Run("success", func(t *testing.T) {
		e := echo.New()
		e.Validator = &CustomValidator{validator: validator.New()}

		req := httptest.NewRequest(http.MethodGet, "/", nil)
		res := httptest.NewRecorder()

		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtTokenUser))

		req.Header.Set("Content-Type", "application/json")

		context := e.NewContext(req, res)
		context.SetPath("/menus/recommend/breakfast")

		goalController := New(&MockMenuRepository{})
		err := middleware.JWT([]byte(configs.JWT_SECRET))(goalController.GetRecommendLunch())(context)

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
		assert.Equal(t, "Success Get Menu Recommended", response.Message)

	})

	t.Run("Failed get all menu recommended", func(t *testing.T) {
		e := echo.New()
		e.Validator = &CustomValidator{validator: validator.New()}

		req := httptest.NewRequest(http.MethodGet, "/", nil)
		res := httptest.NewRecorder()

		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtTokenUser))

		req.Header.Set("Content-Type", "application/json")

		context := e.NewContext(req, res)
		context.SetPath("/menus/recommend/breakfast")

		goalController := New(&MockFailedMenuRepository{})

		err := middleware.JWT([]byte(configs.JWT_SECRET))(goalController.GetRecommendLunch())(context)

		if err != nil {
			fmt.Println(err)
			return
		}

		var response common.Response

		json.Unmarshal([]byte(res.Body.Bytes()), &response)
		// data := (response.Data).(map[string]interface{})
		// log.Info(data)
		// log.Info(response)
		assert.Equal(t, float64(http.StatusInternalServerError), response.Code)
		assert.Equal(t, "There is some error on server", response.Message)

	})
	t.Run("Failed access get all menu recommended", func(t *testing.T) {
		e := echo.New()
		e.Validator = &CustomValidator{validator: validator.New()}

		req := httptest.NewRequest(http.MethodGet, "/", nil)
		res := httptest.NewRecorder()

		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtTokenAdmin))

		req.Header.Set("Content-Type", "application/json")

		context := e.NewContext(req, res)
		context.SetPath("/menus/recommend/breakfast")

		goalController := New(&MockFailedMenuRepository{})

		err := middleware.JWT([]byte(configs.JWT_SECRET))(goalController.GetRecommendLunch())(context)

		if err != nil {
			fmt.Println(err)
			return
		}

		var response common.Response

		json.Unmarshal([]byte(res.Body.Bytes()), &response)
		// data := (response.Data).(map[string]interface{})
		// log.Info(data)
		// log.Info(response)
		assert.Equal(t, float64(http.StatusBadRequest), response.Code)
		assert.Equal(t, "access denied", response.Message)

	})
	t.Run("impossible", func(t *testing.T) {
		e := echo.New()
		e.Validator = &CustomValidator{validator: validator.New()}

		req := httptest.NewRequest(http.MethodGet, "/", nil)
		res := httptest.NewRecorder()

		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtTokenUser))

		req.Header.Set("Content-Type", "application/json")

		context := e.NewContext(req, res)
		context.SetPath("/menus/recommend/breakfast")

		goalController := New(&MockImpossibleMenuRepository{})
		err := middleware.JWT([]byte(configs.JWT_SECRET))(goalController.GetRecommendLunch())(context)

		if err != nil {
			fmt.Println(err)
			return
		}

		var response common.Response

		json.Unmarshal([]byte(res.Body.Bytes()), &response)
		// data := (response.Data).(map[string]interface{})
		// log.Info(data)
		// log.Info(response)
		assert.Equal(t, float64(http.StatusBadRequest), response.Code)
		assert.Equal(t, "impossible", response.Message)

	})

}

func TestGetRecommendDinner(t *testing.T) {

	t.Run("success", func(t *testing.T) {
		e := echo.New()
		e.Validator = &CustomValidator{validator: validator.New()}

		req := httptest.NewRequest(http.MethodGet, "/", nil)
		res := httptest.NewRecorder()

		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtTokenUser))

		req.Header.Set("Content-Type", "application/json")

		context := e.NewContext(req, res)
		context.SetPath("/menus/recommend/dinner")

		goalController := New(&MockMenuRepository{})
		err := middleware.JWT([]byte(configs.JWT_SECRET))(goalController.GetRecommendDinner())(context)

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
		assert.Equal(t, "Success Get Menu Recommended", response.Message)

	})

	t.Run("Failed get all menu recommended", func(t *testing.T) {
		e := echo.New()
		e.Validator = &CustomValidator{validator: validator.New()}

		req := httptest.NewRequest(http.MethodGet, "/", nil)
		res := httptest.NewRecorder()

		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtTokenUser))

		req.Header.Set("Content-Type", "application/json")

		context := e.NewContext(req, res)
		context.SetPath("/menus/recommend/dinner")

		goalController := New(&MockFailedMenuRepository{})

		err := middleware.JWT([]byte(configs.JWT_SECRET))(goalController.GetRecommendDinner())(context)

		if err != nil {
			fmt.Println(err)
			return
		}

		var response common.Response

		json.Unmarshal([]byte(res.Body.Bytes()), &response)
		// data := (response.Data).(map[string]interface{})
		// log.Info(data)
		// log.Info(response)
		assert.Equal(t, float64(http.StatusInternalServerError), response.Code)
		assert.Equal(t, "There is some error on server", response.Message)

	})
	t.Run("Failed access get all menu recommended", func(t *testing.T) {
		e := echo.New()
		e.Validator = &CustomValidator{validator: validator.New()}

		req := httptest.NewRequest(http.MethodGet, "/", nil)
		res := httptest.NewRecorder()

		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtTokenAdmin))

		req.Header.Set("Content-Type", "application/json")

		context := e.NewContext(req, res)
		context.SetPath("/menus/recommend/dinner")

		goalController := New(&MockFailedMenuRepository{})

		err := middleware.JWT([]byte(configs.JWT_SECRET))(goalController.GetRecommendDinner())(context)

		if err != nil {
			fmt.Println(err)
			return
		}

		var response common.Response

		json.Unmarshal([]byte(res.Body.Bytes()), &response)
		// data := (response.Data).(map[string]interface{})
		// log.Info(data)
		// log.Info(response)
		assert.Equal(t, float64(http.StatusBadRequest), response.Code)
		assert.Equal(t, "access denied", response.Message)

	})
	t.Run("impossible", func(t *testing.T) {
		e := echo.New()
		e.Validator = &CustomValidator{validator: validator.New()}

		req := httptest.NewRequest(http.MethodGet, "/", nil)
		res := httptest.NewRecorder()

		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtTokenUser))

		req.Header.Set("Content-Type", "application/json")

		context := e.NewContext(req, res)
		context.SetPath("/menus/recommend/dinner")

		goalController := New(&MockImpossibleMenuRepository{})
		err := middleware.JWT([]byte(configs.JWT_SECRET))(goalController.GetRecommendDinner())(context)

		if err != nil {
			fmt.Println(err)
			return
		}

		var response common.Response

		json.Unmarshal([]byte(res.Body.Bytes()), &response)
		// data := (response.Data).(map[string]interface{})
		// log.Info(data)
		// log.Info(response)
		assert.Equal(t, float64(http.StatusBadRequest), response.Code)
		assert.Equal(t, "impossible", response.Message)

	})

}

func TestGetRecommendOverTime(t *testing.T) {

	t.Run("success", func(t *testing.T) {
		e := echo.New()
		e.Validator = &CustomValidator{validator: validator.New()}

		req := httptest.NewRequest(http.MethodGet, "/", nil)
		res := httptest.NewRecorder()

		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtTokenUser))

		req.Header.Set("Content-Type", "application/json")

		context := e.NewContext(req, res)
		context.SetPath("/menus/recommend/overtime")

		goalController := New(&MockMenuRepository{})
		err := middleware.JWT([]byte(configs.JWT_SECRET))(goalController.GetRecommendOverTime())(context)

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
		assert.Equal(t, "Success Get Menu Recommended", response.Message)

	})

	t.Run("Failed get all menu recommended", func(t *testing.T) {
		e := echo.New()
		e.Validator = &CustomValidator{validator: validator.New()}

		req := httptest.NewRequest(http.MethodGet, "/", nil)
		res := httptest.NewRecorder()

		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtTokenUser))

		req.Header.Set("Content-Type", "application/json")

		context := e.NewContext(req, res)
		context.SetPath("/menus/recommend/overtime")

		goalController := New(&MockFailedMenuRepository{})

		err := middleware.JWT([]byte(configs.JWT_SECRET))(goalController.GetRecommendOverTime())(context)

		if err != nil {
			fmt.Println(err)
			return
		}

		var response common.Response

		json.Unmarshal([]byte(res.Body.Bytes()), &response)
		// data := (response.Data).(map[string]interface{})
		// log.Info(data)
		// log.Info(response)
		assert.Equal(t, float64(http.StatusInternalServerError), response.Code)
		assert.Equal(t, "There is some error on server", response.Message)

	})
	t.Run("Failed access get all menu recommended", func(t *testing.T) {
		e := echo.New()
		e.Validator = &CustomValidator{validator: validator.New()}

		req := httptest.NewRequest(http.MethodGet, "/", nil)
		res := httptest.NewRecorder()

		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtTokenAdmin))

		req.Header.Set("Content-Type", "application/json")

		context := e.NewContext(req, res)
		context.SetPath("/menus/recommend/overtime")

		goalController := New(&MockFailedMenuRepository{})

		err := middleware.JWT([]byte(configs.JWT_SECRET))(goalController.GetRecommendOverTime())(context)

		if err != nil {
			fmt.Println(err)
			return
		}

		var response common.Response

		json.Unmarshal([]byte(res.Body.Bytes()), &response)
		// data := (response.Data).(map[string]interface{})
		// log.Info(data)
		// log.Info(response)
		assert.Equal(t, float64(http.StatusBadRequest), response.Code)
		assert.Equal(t, "access denied", response.Message)

	})
	t.Run("impossible", func(t *testing.T) {
		e := echo.New()
		e.Validator = &CustomValidator{validator: validator.New()}

		req := httptest.NewRequest(http.MethodGet, "/", nil)
		res := httptest.NewRecorder()

		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtTokenUser))

		req.Header.Set("Content-Type", "application/json")

		context := e.NewContext(req, res)
		context.SetPath("/menus/recommend/overtime")

		goalController := New(&MockImpossibleMenuRepository{})
		err := middleware.JWT([]byte(configs.JWT_SECRET))(goalController.GetRecommendOverTime())(context)

		if err != nil {
			fmt.Println(err)
			return
		}

		var response common.Response

		json.Unmarshal([]byte(res.Body.Bytes()), &response)
		// data := (response.Data).(map[string]interface{})
		// log.Info(data)
		// log.Info(response)
		assert.Equal(t, float64(http.StatusBadRequest), response.Code)
		assert.Equal(t, "impossible", response.Message)

	})

}

type MockMenuRepository struct{}

func (m *MockMenuRepository) CreateMenuAdmin(foods []entities.Food, menus entities.Menu) (entities.Menu, error) {

	return entities.Menu{
		User_uid:      "xyz",
		Menu_category: "food",
		Created_by:    "admin",
		Detail_menu: []entities.Detail_menu{
			{Food: entities.Food{Food_uid: "JSKJDSK"}},
			{Food: entities.Food{Food_uid: "sfsaafsd"}},
		},
	}, nil
}
func (m *MockMenuRepository) CreateMenuUser(foods []entities.Food, menus entities.Menu) (entities.Menu, error) {
	return entities.Menu{
		User_uid:      "xyz",
		Menu_category: "food",
		Created_by:    "user",
		Detail_menu: []entities.Detail_menu{
			{Food: entities.Food{Food_uid: "JSKJDSK"}},
			{Food: entities.Food{Food_uid: "sfsaafsd"}},
		},
	}, nil
}

func (m *MockMenuRepository) GetAllMenu(category string, createdBy string) ([]entities.Menu, error) {

	return []entities.Menu{{
		User_uid:      "xyz",
		Menu_category: "food",
		Created_by:    "user",
		Detail_menu: []entities.Detail_menu{
			{Food: entities.Food{Food_uid: "JSKJDSK"}},
			{Food: entities.Food{Food_uid: "sfsaafsd"}},
		},
	}}, nil
}
func (m *MockMenuRepository) GetRecommendBreakfast(user_uid string) ([]entities.Menu, int64, int, error) {

	return []entities.Menu{{
		User_uid:      "xyz",
		Menu_category: "food",
		Created_by:    "user",
		Detail_menu: []entities.Detail_menu{
			{Food: entities.Food{Food_uid: "JSKJDSK"}},
			{Food: entities.Food{Food_uid: "sfsaafsd"}},
		},
	}}, 0, 0, nil
}
func (m *MockMenuRepository) GetRecommendLunch(user_uid string) ([]entities.Menu, int64, int, error) {

	return []entities.Menu{{
		User_uid:      "xyz",
		Menu_category: "food",
		Created_by:    "user",
		Detail_menu: []entities.Detail_menu{
			{Food: entities.Food{Food_uid: "JSKJDSK"}},
			{Food: entities.Food{Food_uid: "sfsaafsd"}},
		},
	}}, 0, 0, nil
}
func (m *MockMenuRepository) GetRecommendDinner(user_uid string) ([]entities.Menu, int64, int, error) {
	return []entities.Menu{{
		User_uid:      "xyz",
		Menu_category: "food",
		Created_by:    "user",
		Detail_menu: []entities.Detail_menu{
			{Food: entities.Food{Food_uid: "JSKJDSK"}},
			{Food: entities.Food{Food_uid: "sfsaafsd"}},
		},
	}}, 0, 0, nil
}
func (m *MockMenuRepository) GetRecommendOverTime(user_uid string) ([]entities.Menu, int64, int, error) {

	return []entities.Menu{{
		User_uid:      "xyz",
		Menu_category: "food",
		Created_by:    "user",
		Detail_menu: []entities.Detail_menu{
			{Food: entities.Food{Food_uid: "JSKJDSK"}},
			{Food: entities.Food{Food_uid: "sfsaafsd"}},
		},
	}}, 0, 0, nil
}
func (m *MockMenuRepository) Update(menu_uid string, foods []entities.Food, updateMenu entities.Menu) (entities.Menu, error) {

	return entities.Menu{
		User_uid:      "xyz",
		Menu_category: "food",
		Created_by:    "admin",
		Detail_menu: []entities.Detail_menu{
			{Food: entities.Food{Food_uid: "JSKJDSK"}},
			{Food: entities.Food{Food_uid: "sfsaafsd"}},
		},
	}, nil
}
func (m *MockMenuRepository) Delete(menu_uid string) error {

	return nil
}

type MockMenuAdminRepository struct{}

func (m *MockMenuAdminRepository) CreateMenuAdmin(foods []entities.Food, menus entities.Menu) (entities.Menu, error) {

	return entities.Menu{
		User_uid:      "xyz",
		Menu_category: "food",
		Created_by:    "admin",
	}, nil
}
func (m *MockMenuAdminRepository) CreateMenuUser(foods []entities.Food, menus entities.Menu) (entities.Menu, error) {
	return entities.Menu{
		User_uid:      "xyz",
		Menu_category: "food",
		Created_by:    "admin",
	}, nil

}

func (m *MockMenuAdminRepository) GetAllMenu(category string, createdBy string) ([]entities.Menu, error) {

	return []entities.Menu{{
		User_uid:      "xyz",
		Menu_category: "food",
		Created_by:    "admin",
	}}, nil
}
func (m *MockMenuAdminRepository) GetRecommendBreakfast(user_uid string) ([]entities.Menu, int64, int, error) {

	return []entities.Menu{{
		User_uid:      "xyz",
		Menu_category: "food",
		Created_by:    "admin",
	}}, 0, 0, nil
}
func (m *MockMenuAdminRepository) GetRecommendLunch(user_uid string) ([]entities.Menu, int64, int, error) {

	return []entities.Menu{{
		User_uid:      "xyz",
		Menu_category: "food",
		Created_by:    "admin",
	}}, 0, 0, nil
}
func (m *MockMenuAdminRepository) GetRecommendDinner(user_uid string) ([]entities.Menu, int64, int, error) {

	return []entities.Menu{{
		User_uid:      "xyz",
		Menu_category: "food",
		Created_by:    "admin",
	}}, 0, 0, nil
}
func (m *MockMenuAdminRepository) GetRecommendOverTime(user_uid string) ([]entities.Menu, int64, int, error) {

	return []entities.Menu{{
		User_uid:      "xyz",
		Menu_category: "food",
		Created_by:    "admin",
	}}, 0, 0, nil
}
func (m *MockMenuAdminRepository) Update(menu_uid string, foods []entities.Food, updateMenu entities.Menu) (entities.Menu, error) {

	return entities.Menu{
		User_uid:      "xyz",
		Menu_category: "food",
		Created_by:    "admin",
	}, nil
}
func (m *MockMenuAdminRepository) Delete(menu_uid string) error {

	return nil
}

type MockFailedMenuRepository struct{}

func (m *MockFailedMenuRepository) CreateMenuAdmin(foods []entities.Food, menus entities.Menu) (entities.Menu, error) {

	return entities.Menu{}, errors.New("There is some error on server")
}
func (m *MockFailedMenuRepository) CreateMenuUser(foods []entities.Food, menus entities.Menu) (entities.Menu, error) {
	return entities.Menu{}, errors.New("There is some error on server")

}

func (m *MockFailedMenuRepository) GetAllMenu(category string, createdBy string) ([]entities.Menu, error) {

	return []entities.Menu{}, errors.New("There is some error on server")
}
func (m *MockFailedMenuRepository) GetRecommendBreakfast(user_uid string) ([]entities.Menu, int64, int, error) {

	return []entities.Menu{}, 0, 0, errors.New("There is some error on server")
}
func (m *MockFailedMenuRepository) GetRecommendLunch(user_uid string) ([]entities.Menu, int64, int, error) {

	return []entities.Menu{}, 0, 0, errors.New("There is some error on server")
}
func (m *MockFailedMenuRepository) GetRecommendDinner(user_uid string) ([]entities.Menu, int64, int, error) {

	return []entities.Menu{}, 0, 0, errors.New("There is some error on server")
}
func (m *MockFailedMenuRepository) GetRecommendOverTime(user_uid string) ([]entities.Menu, int64, int, error) {

	return []entities.Menu{}, 0, 0, errors.New("There is some error on server")
}
func (m *MockFailedMenuRepository) Update(menu_uid string, foods []entities.Food, updateMenu entities.Menu) (entities.Menu, error) {

	return entities.Menu{}, errors.New("There is some error on server")
}
func (m *MockFailedMenuRepository) Delete(menu_uid string) error {

	return errors.New("There is some error on server")
}

type MockImpossibleMenuRepository struct{}

func (m *MockImpossibleMenuRepository) CreateMenuAdmin(foods []entities.Food, menus entities.Menu) (entities.Menu, error) {

	return entities.Menu{}, errors.New("There is some error on server")
}
func (m *MockImpossibleMenuRepository) CreateMenuUser(foods []entities.Food, menus entities.Menu) (entities.Menu, error) {
	return entities.Menu{}, errors.New("There is some error on server")

}

func (m *MockImpossibleMenuRepository) GetAllMenu(category string, createdBy string) ([]entities.Menu, error) {

	return []entities.Menu{}, errors.New("There is some error on server")
}
func (m *MockImpossibleMenuRepository) GetRecommendBreakfast(user_uid string) ([]entities.Menu, int64, int, error) {

	return []entities.Menu{}, 0, 0, errors.New("impossible")
}
func (m *MockImpossibleMenuRepository) GetRecommendLunch(user_uid string) ([]entities.Menu, int64, int, error) {

	return []entities.Menu{}, 0, 0, errors.New("impossible")
}
func (m *MockImpossibleMenuRepository) GetRecommendDinner(user_uid string) ([]entities.Menu, int64, int, error) {

	return []entities.Menu{}, 0, 0, errors.New("impossible")
}
func (m *MockImpossibleMenuRepository) GetRecommendOverTime(user_uid string) ([]entities.Menu, int64, int, error) {

	return []entities.Menu{}, 0, 0, errors.New("impossible")
}
func (m *MockImpossibleMenuRepository) Update(menu_uid string, foods []entities.Food, updateMenu entities.Menu) (entities.Menu, error) {

	return entities.Menu{}, errors.New("There is some error on server")
}
func (m *MockImpossibleMenuRepository) Delete(menu_uid string) error {

	return errors.New("There is some error on server")
}
