package foods

// import (
// 	"HealthFit/configs"
// 	"HealthFit/delivery/controllers/auth"
// 	"HealthFit/delivery/controllers/common"
// 	"HealthFit/entities"
// 	"bytes"
// 	"encoding/json"
// 	"errors"
// 	"fmt"
// 	"net/http"
// 	"net/http/httptest"
// 	"testing"

// 	"github.com/labstack/echo/v4"
// 	"github.com/labstack/echo/v4/middleware"

// 	"github.com/go-playground/assert"
// 	"github.com/go-playground/validator"
// 	"github.com/labstack/gommon/log"
// 	"gorm.io/gorm"
// 	// "github.com/labstack/gommon/log"
// )

// type CustomValidator struct {
// 	validator *validator.Validate
// }

// var jwtToken = ""

// func (cv *CustomValidator) Validate(i interface{}) error {
// 	return cv.validator.Struct(i)
// }

// //////
// type MockAuthRepository struct{}

// func (m *MockAuthRepository) Login(email, password string) (entities.User, error) {

// 	return entities.User{Model: gorm.Model{ID: 1}, Email: "test", Password: "test", Roles: true}, nil
// }
// func TestLogin(t *testing.T) {
// 	t.Run(
// 		"1. Success Login Test", func(t *testing.T) {
// 			e := echo.New()
// 			e.Validator = &CustomValidator{validator: validator.New()}

// 			requestBody, _ := json.Marshal(
// 				auth.LoginReqFormat{
// 					Email:    "test@gmail.com",
// 					Password: "test1",
// 				},
// 			)

// 			req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
// 			res := httptest.NewRecorder()
// 			req.Header.Set("Content-Type", "application/json")
// 			context := e.NewContext(req, res)
// 			context.SetPath("/admin/login")

// 			authController := auth.New(&MockAuthRepository{})
// 			authController.AdminLogin()(context)

// 			var response common.Response

// 			json.Unmarshal([]byte(res.Body.Bytes()), &response)
// 			data := (response.Data).(map[string]interface{})
// 			log.Info(data)
// 			log.Info(response)
// 			jwtToken = data["token"].(string)

// 			assert.Equal(t, "ADMIN - berhasil masuk, mendapatkan token baru", response.Message)
// 		},
// 	)
// }

// func TestCreate(t *testing.T) {

// 	t.Run(
// 		"1. Success Login Test", func(t *testing.T) {
// 			e := echo.New()
// 			e.Validator = &CustomValidator{validator: validator.New()}

// 			requestBody, _ := json.Marshal(
// 				auth.LoginReqFormat{
// 					Email:    "test@gmail.com",
// 					Password: "test1",
// 				},
// 			)

// 			req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
// 			res := httptest.NewRecorder()
// 			req.Header.Set("Content-Type", "application/json")
// 			context := e.NewContext(req, res)
// 			context.SetPath("/admin/login")

// 			authController := auth.New(&MockAuthRepository{})
// 			authController.AdminLogin()(context)

// 			var response common.Response

// 			json.Unmarshal([]byte(res.Body.Bytes()), &response)
// 			data := (response.Data).(map[string]interface{})
// 			log.Info(data)
// 			log.Info(response)
// 			jwtToken = data["token"].(string)

// 			assert.Equal(t, "ADMIN - berhasil masuk, mendapatkan token baru", response.Message)
// 		},
// 	)
// 	t.Run("success", func(t *testing.T) {
// 		e := echo.New()
// 		e.Validator = &CustomValidator{validator: validator.New()}

// 		reqBody, _ := json.Marshal(FoodsCreateRequestFormat{
// 			Name:          "makanan",
// 			Calories:      100,
// 			Energy:        200,
// 			Carbohidrate:  300,
// 			Protein:       400,
// 			Food_category: "snack",
// 			Unit:          "ons",
// 			Unit_value:    1,
// 		})

// 		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(reqBody))
// 		res := httptest.NewRecorder()

// 		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))

// 		req.Header.Set("Content-Type", "application/json")

// 		context := e.NewContext(req, res)
// 		context.SetPath("/foods")

// 		foodController := New(&MockFoodRepository{})

// 		// foodController.Create()(context)

// 		err := middleware.JWT([]byte(configs.JWT_SECRET))(foodController.Create())(context)

// 		if err != nil {
// 			fmt.Println(err)
// 			return
// 		}

// 		var resp common.Response

// 		json.Unmarshal([]byte(res.Body.Bytes()), &resp)
// 		assert.Equal(t, float64(http.StatusCreated), resp.Code)
// 		assert.Equal(t, "Success create foods", resp.Message)

// 	})

// 	t.Run("Failed create", func(t *testing.T) {
// 		e := echo.New()
// 		e.Validator = &CustomValidator{validator: validator.New()}

// 		reqBody, _ := json.Marshal(map[string]interface{}{
// 			"name":          "makanan",
// 			"calories":      100,
// 			"energy":        200,
// 			"carbohidrate":  "300",
// 			"protein":       400,
// 			"food_category": "snack",
// 			"unit":          "ons",
// 			"unit_value":    1,
// 		})

// 		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(reqBody))
// 		res := httptest.NewRecorder()

// 		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))

// 		req.Header.Set("Content-Type", "application/json")

// 		context := e.NewContext(req, res)
// 		context.SetPath("/foods")

// 		foodController := New(&MockFailedFoodRepository{})

// 		// foodController.Create()(context)

// 		err := middleware.JWT([]byte(configs.JWT_SECRET))(foodController.Create())(context)

// 		if err != nil {
// 			fmt.Println(err)
// 			return
// 		}

// 		var resp common.Response

// 		json.Unmarshal([]byte(res.Body.Bytes()), &resp)
// 		assert.Equal(t, float64(http.StatusInternalServerError), resp.Code)
// 		assert.Equal(t, "There is some error on server", resp.Message)

// 	})
// }

// //mock success food

// type MockFoodRepository struct{}

// func (m *MockFoodRepository) Create(food entities.Food) (entities.Food, error) {

// 	return entities.Food{
// 		Name:          "makanan",
// 		Calories:      100,
// 		Energy:        200,
// 		Carbohidrate:  300,
// 		Protein:       400,
// 		Food_category: "snack",
// 		Unit:          "ons",
// 		Unit_value:    1,
// 	}, nil
// }
// func (m *MockFoodRepository) GetById(food_uid string) (entities.Food, error) {

// 	return entities.Food{
// 		Name:          "makanan",
// 		Calories:      100,
// 		Energy:        200,
// 		Carbohidrate:  300,
// 		Protein:       400,
// 		Food_category: "snack",
// 		Unit:          "ons",
// 		Unit_value:    1,
// 	}, nil
// }
// func (m *MockFoodRepository) Search(inpit, category string) ([]entities.Food, error) {

// 	return []entities.Food{{
// 		Name:          "makanan",
// 		Calories:      100,
// 		Energy:        200,
// 		Carbohidrate:  300,
// 		Protein:       400,
// 		Food_category: "snack",
// 		Unit:          "ons",
// 		Unit_value:    1,
// 	}}, nil
// }
// func (m *MockFoodRepository) Update(food_uid string, newSood entities.Food) (entities.Food, error) {

// 	return entities.Food{
// 		Name:          "makanan",
// 		Calories:      100,
// 		Energy:        200,
// 		Carbohidrate:  300,
// 		Protein:       400,
// 		Food_category: "snack",
// 		Unit:          "ons",
// 		Unit_value:    1,
// 	}, nil
// }
// func (m *MockFoodRepository) Delete(food_uid string) error {

// 	return nil
// }
// func (m *MockFoodRepository) GetAll(category string) ([]entities.Food, error) {

// 	return []entities.Food{{
// 		Name:          "makanan",
// 		Calories:      100,
// 		Energy:        200,
// 		Carbohidrate:  300,
// 		Protein:       400,
// 		Food_category: "snack",
// 		Unit:          "ons",
// 		Unit_value:    1,
// 	}}, nil
// }
// func (m *MockFoodRepository) CreateFoodThirdParty(f entities.Food) (entities.Food, error) {

// 	return entities.Food{
// 		Name:          "makanan",
// 		Calories:      100,
// 		Energy:        200,
// 		Carbohidrate:  300,
// 		Protein:       400,
// 		Food_category: "snack",
// 		Unit:          "ons",
// 		Unit_value:    1,
// 	}, nil
// }

// type MockFailedFoodRepository struct{}

// func (m *MockFailedFoodRepository) Create(food entities.Food) (entities.Food, error) {

// 	return entities.Food{}, errors.New("There is some error on server")
// }

// func (m *MockFailedFoodRepository) GetById(food_uid string) (entities.Food, error) {

// 	return entities.Food{
// 		Name:          "makanan",
// 		Calories:      100,
// 		Energy:        200,
// 		Carbohidrate:  300,
// 		Protein:       400,
// 		Food_category: "snack",
// 		Unit:          "ons",
// 		Unit_value:    1,
// 	}, nil
// }
// func (m *MockFailedFoodRepository) Search(inpit, category string) ([]entities.Food, error) {

// 	return []entities.Food{{
// 		Name:          "makanan",
// 		Calories:      100,
// 		Energy:        200,
// 		Carbohidrate:  300,
// 		Protein:       400,
// 		Food_category: "snack",
// 		Unit:          "ons",
// 		Unit_value:    1,
// 	}}, nil
// }
// func (m *MockFailedFoodRepository) Update(food_uid string, newSood entities.Food) (entities.Food, error) {

// 	return entities.Food{
// 		Name:          "makanan",
// 		Calories:      100,
// 		Energy:        200,
// 		Carbohidrate:  300,
// 		Protein:       400,
// 		Food_category: "snack",
// 		Unit:          "ons",
// 		Unit_value:    1,
// 	}, nil
// }
// func (m *MockFailedFoodRepository) Delete(food_uid string) error {

// 	return nil
// }
// func (m *MockFailedFoodRepository) GetAll(category string) ([]entities.Food, error) {

// 	return []entities.Food{{
// 		Name:          "makanan",
// 		Calories:      100,
// 		Energy:        200,
// 		Carbohidrate:  300,
// 		Protein:       400,
// 		Food_category: "snack",
// 		Unit:          "ons",
// 		Unit_value:    1,
// 	}}, nil
// }
// func (m *MockFailedFoodRepository) CreateFoodThirdParty(f entities.Food) (entities.Food, error) {

// 	return entities.Food{
// 		Name:          "makanan",
// 		Calories:      100,
// 		Energy:        200,
// 		Carbohidrate:  300,
// 		Protein:       400,
// 		Food_category: "snack",
// 		Unit:          "ons",
// 		Unit_value:    1,
// 	}, nil
// }

// type MockFailedAdminRepository struct{}

// func (m *MockFailedAdminRepository) Register(user entities.Food) (entities.Food, error) {

// 	return entities.Food{}, errors.New("There is some problem from server")
// }
