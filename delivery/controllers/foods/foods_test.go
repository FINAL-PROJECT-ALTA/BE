package foods

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

	t.Run("success", func(t *testing.T) {
		e := echo.New()
		e.Validator = &CustomValidator{validator: validator.New()}

		reqBody, _ := json.Marshal(FoodsCreateRequestFormat{
			Name:          "makanan",
			Calories:      100,
			Energy:        200,
			Carbohidrate:  300,
			Protein:       400,
			Food_category: "snack",
			Unit:          "ons",
			Unit_value:    1,
		})

		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(reqBody))
		res := httptest.NewRecorder()

		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtTokenAdmin))

		req.Header.Set("Content-Type", "application/json")

		context := e.NewContext(req, res)
		context.SetPath("/foods")

		foodController := New(&MockFoodRepository{})

		// foodController.Create()(context)

		err := middleware.JWT([]byte(configs.JWT_SECRET))(foodController.Create())(context)

		if err != nil {
			fmt.Println(err)
			return
		}

		var resp common.Response

		json.Unmarshal([]byte(res.Body.Bytes()), &resp)
		assert.Equal(t, float64(http.StatusCreated), resp.Code)
		assert.Equal(t, "Success create foods", resp.Message)

	})

	t.Run("Failed create", func(t *testing.T) {
		e := echo.New()
		e.Validator = &CustomValidator{validator: validator.New()}

		reqBody, _ := json.Marshal(map[string]interface{}{
			"name":          "makanan",
			"calories":      100,
			"energy":        200,
			"carbohidrate":  "300",
			"protein":       400,
			"food_category": "snack",
			"unit":          "ons",
			"unit_value":    1,
		})

		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(reqBody))
		res := httptest.NewRecorder()

		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtTokenAdmin))

		req.Header.Set("Content-Type", "application/json")

		context := e.NewContext(req, res)
		context.SetPath("/foods")

		foodController := New(&MockFailedFoodRepository{})

		// foodController.Create()(context)

		err := middleware.JWT([]byte(configs.JWT_SECRET))(foodController.Create())(context)

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
			"name":          "makanan@",
			"calories":      100,
			"energy":        200,
			"carbohidrate":  "300",
			"protein":       400,
			"food_category": "snack",
			"unit":          "ons",
			"unit_value":    1,
		})

		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(reqBody))
		res := httptest.NewRecorder()

		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtTokenAdmin))

		req.Header.Set("Content-Type", "application/json")

		context := e.NewContext(req, res)
		context.SetPath("/foods")

		foodController := New(&MockFailedFoodRepository{})

		// foodController.Create()(context)

		err := middleware.JWT([]byte(configs.JWT_SECRET))(foodController.Create())(context)

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
			"name":          "makanan",
			"calories":      100,
			"energy":        200,
			"carbohidrate":  "300",
			"protein":       400,
			"food_category": "snack",
			"unit":          "ons",
			"unit_value":    1,
		})

		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(reqBody))
		res := httptest.NewRecorder()

		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtTokenUser))

		req.Header.Set("Content-Type", "application/json")

		context := e.NewContext(req, res)
		context.SetPath("/foods")

		foodController := New(&MockFailedFoodRepository{})

		err := middleware.JWT([]byte(configs.JWT_SECRET))(foodController.Create())(context)

		if err != nil {
			fmt.Println(err)
			return
		}

		var resp common.Response

		json.Unmarshal([]byte(res.Body.Bytes()), &resp)
		assert.Equal(t, float64(http.StatusBadRequest), resp.Code)
		assert.Equal(t, "Access Denied", resp.Message)

	})
}

func TestGetById(t *testing.T) {

	t.Run("success", func(t *testing.T) {
		e := echo.New()
		e.Validator = &CustomValidator{validator: validator.New()}

		req := httptest.NewRequest(http.MethodGet, "/", nil)
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")

		context := e.NewContext(req, res)
		context.SetPath("/foods")
		context.SetParamNames("food_uid")
		context.SetParamValues("xyz")

		foodController := New(&MockFoodRepository{})

		foodController.GetById()(context)

		var response common.Response

		json.Unmarshal([]byte(res.Body.Bytes()), &response)
		// data := (response.Data).(map[string]interface{})
		// log.Info(data)
		// log.Info(response)
		assert.Equal(t, float64(http.StatusOK), response.Code)
		assert.Equal(t, "Success get food by food_uid", response.Message)

	})

	t.Run("Failed get food ById", func(t *testing.T) {
		e := echo.New()
		e.Validator = &CustomValidator{validator: validator.New()}

		req := httptest.NewRequest(http.MethodGet, "/", nil)
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")

		context := e.NewContext(req, res)
		context.SetPath("/foods")
		context.SetParamNames("food_uid")
		context.SetParamValues("xyz")

		foodController := New(&MockFailedFoodRepository{})

		foodController.GetById()(context)

		var response common.Response

		json.Unmarshal([]byte(res.Body.Bytes()), &response)
		// data := (response.Data).(map[string]interface{})
		// log.Info(data)
		// log.Info(response)
		assert.Equal(t, float64(http.StatusInternalServerError), response.Code)
		assert.Equal(t, "There is some error on server", response.Message)

	})
	t.Run("Failed not found get ById", func(t *testing.T) {
		e := echo.New()
		e.Validator = &CustomValidator{validator: validator.New()}

		req := httptest.NewRequest(http.MethodGet, "/", nil)
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")

		context := e.NewContext(req, res)
		context.SetPath("/foods")
		context.SetParamNames("food_uid")
		context.SetParamValues("xyz")

		foodController := New(&MockNotFoundFoodRepository{})

		foodController.GetById()(context)

		var response common.Response

		json.Unmarshal([]byte(res.Body.Bytes()), &response)
		// data := (response.Data).(map[string]interface{})
		// log.Info(data)
		// log.Info(response)
		assert.Equal(t, float64(http.StatusNotFound), response.Code)
		assert.Equal(t, "Not found", response.Message)

	})

}
func TestSearch(t *testing.T) {

	t.Run("success", func(t *testing.T) {
		e := echo.New()
		e.Validator = &CustomValidator{validator: validator.New()}

		req := httptest.NewRequest(http.MethodGet, "/", nil)
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")

		context := e.NewContext(req, res)
		context.SetPath("/foods/search")
		context.Set("input", "a")
		context.Set("category", "foods")

		foodController := New(&MockFoodRepository{})

		foodController.Search()(context)

		var response common.Response

		json.Unmarshal([]byte(res.Body.Bytes()), &response)
		// data := (response.Data).(map[string]interface{})
		// log.Info(data)
		// log.Info(response)
		assert.Equal(t, float64(http.StatusOK), response.Code)
		assert.Equal(t, "Success get foods", response.Message)

	})

	t.Run("Failed search food", func(t *testing.T) {
		e := echo.New()
		e.Validator = &CustomValidator{validator: validator.New()}

		req := httptest.NewRequest(http.MethodGet, "/", nil)
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")

		context := e.NewContext(req, res)
		context.SetPath("/foods/search")
		context.Set("input", "a")
		context.Set("category", "foods")

		foodController := New(&MockFailedFoodRepository{})

		foodController.Search()(context)

		var response common.Response

		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		assert.Equal(t, float64(http.StatusInternalServerError), response.Code)
		assert.Equal(t, "There is some error on server", response.Message)
	})
}
func TestUpdate(t *testing.T) {

	t.Run("success", func(t *testing.T) {
		e := echo.New()
		e.Validator = &CustomValidator{validator: validator.New()}

		reqBody, _ := json.Marshal(FoodsUpdateRequestFormat{
			Name:          "makananbaru",
			Calories:      100,
			Energy:        200,
			Carbohidrate:  300,
			Protein:       400,
			Food_category: "snack",
			Unit:          "ons",
			Unit_value:    1,
		})

		req := httptest.NewRequest(http.MethodPut, "/", bytes.NewBuffer(reqBody))
		res := httptest.NewRecorder()

		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtTokenAdmin))

		req.Header.Set("Content-Type", "application/json")

		context := e.NewContext(req, res)
		context.SetPath("/foods")
		context.SetParamNames("food_uid")
		context.SetParamValues("xyz")

		foodController := New(&MockFoodRepository{})

		// foodController.Create()(context)

		err := middleware.JWT([]byte(configs.JWT_SECRET))(foodController.Update())(context)

		if err != nil {
			fmt.Println(err)
			return
		}

		var resp common.Response

		json.Unmarshal([]byte(res.Body.Bytes()), &resp)
		assert.Equal(t, float64(http.StatusCreated), resp.Code)
		assert.Equal(t, "Success update foods", resp.Message)

	})
	t.Run("failed update food", func(t *testing.T) {
		e := echo.New()
		e.Validator = &CustomValidator{validator: validator.New()}

		reqBody, _ := json.Marshal(FoodsUpdateRequestFormat{
			Name:          "makananbaru",
			Calories:      100,
			Energy:        200,
			Carbohidrate:  300,
			Protein:       400,
			Food_category: "snack",
			Unit:          "ons",
			Unit_value:    1,
		})

		req := httptest.NewRequest(http.MethodPut, "/", bytes.NewBuffer(reqBody))
		res := httptest.NewRecorder()

		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtTokenAdmin))

		req.Header.Set("Content-Type", "application/json")

		context := e.NewContext(req, res)
		context.SetPath("/foods")
		context.SetParamNames("food_uid")
		context.SetParamValues("xyz")

		foodController := New(&MockFailedFoodRepository{})

		// foodController.Create()(context)

		err := middleware.JWT([]byte(configs.JWT_SECRET))(foodController.Update())(context)

		if err != nil {
			fmt.Println(err)
			return
		}

		var resp common.Response

		json.Unmarshal([]byte(res.Body.Bytes()), &resp)
		assert.Equal(t, float64(http.StatusInternalServerError), resp.Code)
		assert.Equal(t, "There is some error on server", resp.Message)

	})
	t.Run("failed bind update food", func(t *testing.T) {
		e := echo.New()
		e.Validator = &CustomValidator{validator: validator.New()}

		reqBody, _ := json.Marshal(map[string]interface{}{
			"name":          "makanan@",
			"calories":      100,
			"energy":        200,
			"carbohidrate":  "300",
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
		context.SetPath("/foods")
		context.SetParamNames("food_uid")
		context.SetParamValues("xyz")

		foodController := New(&MockFoodRepository{})

		// foodController.Create()(context)

		err := middleware.JWT([]byte(configs.JWT_SECRET))(foodController.Update())(context)

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

		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtTokenUser))

		req.Header.Set("Content-Type", "application/json")

		context := e.NewContext(req, res)
		context.SetPath("/foods")
		context.SetParamNames("food_uid")
		context.SetParamValues("xyz")

		foodController := New(&MockFailedFoodRepository{})

		err := middleware.JWT([]byte(configs.JWT_SECRET))(foodController.Update())(context)

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
func TestDelete(t *testing.T) {

	t.Run("success", func(t *testing.T) {
		e := echo.New()
		e.Validator = &CustomValidator{validator: validator.New()}

		req := httptest.NewRequest(http.MethodDelete, "/", nil)
		res := httptest.NewRecorder()

		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtTokenAdmin))

		req.Header.Set("Content-Type", "application/json")

		context := e.NewContext(req, res)
		context.SetPath("/foods")
		context.SetParamNames("food_uid")
		context.SetParamValues("xyz")

		foodController := New(&MockFoodRepository{})

		// foodController.Create()(context)

		err := middleware.JWT([]byte(configs.JWT_SECRET))(foodController.Delete())(context)

		if err != nil {
			fmt.Println(err)
			return
		}

		var resp common.Response

		json.Unmarshal([]byte(res.Body.Bytes()), &resp)
		assert.Equal(t, float64(http.StatusOK), resp.Code)
		assert.Equal(t, "success delete foods", resp.Message)

	})
	t.Run("failed delete food", func(t *testing.T) {
		e := echo.New()
		e.Validator = &CustomValidator{validator: validator.New()}

		req := httptest.NewRequest(http.MethodDelete, "/", nil)
		res := httptest.NewRecorder()

		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtTokenAdmin))

		req.Header.Set("Content-Type", "application/json")

		context := e.NewContext(req, res)
		context.SetPath("/foods")
		context.SetParamNames("food_uid")
		context.SetParamValues("xyz")

		foodController := New(&MockFailedFoodRepository{})

		// foodController.Create()(context)

		err := middleware.JWT([]byte(configs.JWT_SECRET))(foodController.Delete())(context)

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
		context.SetPath("/foods")
		context.SetParamNames("food_uid")
		context.SetParamValues("xyz")

		foodController := New(&MockFailedFoodRepository{})

		err := middleware.JWT([]byte(configs.JWT_SECRET))(foodController.Delete())(context)

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
func TestGetAll(t *testing.T) {

	t.Run("success", func(t *testing.T) {
		e := echo.New()
		e.Validator = &CustomValidator{validator: validator.New()}

		req := httptest.NewRequest(http.MethodGet, "/", nil)
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")

		context := e.NewContext(req, res)
		context.SetPath("/foods")
		context.Set("category", "a")

		foodController := New(&MockFoodRepository{})

		foodController.GetAll()(context)

		var response common.Response

		json.Unmarshal([]byte(res.Body.Bytes()), &response)
		// data := (response.Data).(map[string]interface{})
		// log.Info(data)
		// log.Info(response)
		assert.Equal(t, float64(http.StatusOK), response.Code)
		assert.Equal(t, "Success get all foods", response.Message)

	})

	t.Run("Failed get all foods", func(t *testing.T) {
		e := echo.New()
		e.Validator = &CustomValidator{validator: validator.New()}

		req := httptest.NewRequest(http.MethodGet, "/", nil)
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")

		context := e.NewContext(req, res)
		context.SetPath("/foods")
		context.Set("category", "a")

		foodController := New(&MockFailedFoodRepository{})

		foodController.GetAll()(context)

		var response common.Response

		json.Unmarshal([]byte(res.Body.Bytes()), &response)
		// data := (response.Data).(map[string]interface{})
		// log.Info(data)
		// log.Info(response)
		assert.Equal(t, float64(http.StatusInternalServerError), response.Code)
		assert.Equal(t, "There is some error on server", response.Message)

	})

}

//mock success food

type MockFoodRepository struct{}

func (m *MockFoodRepository) Create(food entities.Food) (entities.Food, error) {

	return entities.Food{
		Name:          "makanan",
		Calories:      100,
		Energy:        200,
		Carbohidrate:  300,
		Protein:       400,
		Food_category: "snack",
		Unit:          "ons",
		Unit_value:    1,
	}, nil
}
func (m *MockFoodRepository) GetById(food_uid string) (entities.Food, error) {
	return entities.Food{
		Food_uid:      "xyz",
		Name:          "makanan",
		Calories:      100,
		Energy:        200,
		Carbohidrate:  300,
		Protein:       400,
		Food_category: "snack",
		Unit:          "ons",
		Unit_value:    1,
	}, nil

}
func (m *MockFoodRepository) Search(input, category string) ([]entities.Food, error) {

	return []entities.Food{{
		Name:          "makanan",
		Calories:      100,
		Energy:        200,
		Carbohidrate:  300,
		Protein:       400,
		Food_category: "snack",
		Unit:          "ons",
		Unit_value:    1,
	}}, nil
}
func (m *MockFoodRepository) Update(food_uid string, newSood entities.Food) (entities.Food, error) {

	return entities.Food{
		Name:          "makananbaru",
		Calories:      100,
		Energy:        200,
		Carbohidrate:  300,
		Protein:       400,
		Food_category: "snack",
		Unit:          "ons",
		Unit_value:    1,
	}, nil
}
func (m *MockFoodRepository) Delete(food_uid string) error {

	return nil
}
func (m *MockFoodRepository) GetAll(category string) ([]entities.Food, error) {

	return []entities.Food{{
		Name:          "makanan",
		Calories:      100,
		Energy:        200,
		Carbohidrate:  300,
		Protein:       400,
		Food_category: "snack",
		Unit:          "ons",
		Unit_value:    1,
	}}, nil
}
func (m *MockFoodRepository) CreateFoodThirdParty(f entities.Food) (entities.Food, error) {

	return entities.Food{
		Name:          "makanan",
		Calories:      100,
		Energy:        200,
		Carbohidrate:  300,
		Protein:       400,
		Food_category: "snack",
		Unit:          "ons",
		Unit_value:    1,
	}, nil
}

type MockFailedFoodRepository struct{}

func (m *MockFailedFoodRepository) Create(food entities.Food) (entities.Food, error) {

	return entities.Food{}, errors.New("There is some error on server")
}

func (m *MockFailedFoodRepository) GetById(food_uid string) (entities.Food, error) {

	return entities.Food{}, errors.New("There is some error on server")
}
func (m *MockFailedFoodRepository) Search(inpit, category string) ([]entities.Food, error) {

	return []entities.Food{}, errors.New("There is some error on server")
}
func (m *MockFailedFoodRepository) Update(food_uid string, newSood entities.Food) (entities.Food, error) {

	return entities.Food{}, errors.New("There is some error on server")
}
func (m *MockFailedFoodRepository) Delete(food_uid string) error {

	return errors.New("")
}
func (m *MockFailedFoodRepository) GetAll(category string) ([]entities.Food, error) {

	return []entities.Food{}, errors.New("There is some error on server")
}
func (m *MockFailedFoodRepository) CreateFoodThirdParty(f entities.Food) (entities.Food, error) {

	return entities.Food{
		Name:          "makanan",
		Calories:      100,
		Energy:        200,
		Carbohidrate:  300,
		Protein:       400,
		Food_category: "snack",
		Unit:          "ons",
		Unit_value:    1,
	}, nil
}

//mock notfound

type MockNotFoundFoodRepository struct{}

func (m *MockNotFoundFoodRepository) Create(food entities.Food) (entities.Food, error) {

	return entities.Food{}, errors.New("There is some error on server")
}

func (m *MockNotFoundFoodRepository) GetById(food_uid string) (entities.Food, error) {

	return entities.Food{}, nil
}

func (m *MockNotFoundFoodRepository) Search(inpit, category string) ([]entities.Food, error) {

	return []entities.Food{{
		Name:          "makanan",
		Calories:      100,
		Energy:        200,
		Carbohidrate:  300,
		Protein:       400,
		Food_category: "snack",
		Unit:          "ons",
		Unit_value:    1,
	}}, nil
}
func (m *MockNotFoundFoodRepository) Update(food_uid string, newSood entities.Food) (entities.Food, error) {

	return entities.Food{
		Name:          "makanan",
		Calories:      100,
		Energy:        200,
		Carbohidrate:  300,
		Protein:       400,
		Food_category: "snack",
		Unit:          "ons",
		Unit_value:    1,
	}, nil
}
func (m *MockNotFoundFoodRepository) Delete(food_uid string) error {

	return nil
}
func (m *MockNotFoundFoodRepository) GetAll(category string) ([]entities.Food, error) {

	return []entities.Food{{
		Name:          "makanan",
		Calories:      100,
		Energy:        200,
		Carbohidrate:  300,
		Protein:       400,
		Food_category: "snack",
		Unit:          "ons",
		Unit_value:    1,
	}}, nil
}

func (m *MockNotFoundFoodRepository) CreateFoodThirdParty(f entities.Food) (entities.Food, error) {

	return entities.Food{
		Name:          "makanan",
		Calories:      100,
		Energy:        200,
		Carbohidrate:  300,
		Protein:       400,
		Food_category: "snack",
		Unit:          "ons",
		Unit_value:    1,
	}, nil
}
