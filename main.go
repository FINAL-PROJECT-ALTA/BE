package main

import (
	config "HealthFit/configs"
	route "HealthFit/delivery/routes"
	utils "HealthFit/utils/mysql"

	_authController "HealthFit/delivery/controllers/auth"
	_foodsController "HealthFit/delivery/controllers/foods"
	_goalController "HealthFit/delivery/controllers/goal"
	_userController "HealthFit/delivery/controllers/user"

	_authRepo "HealthFit/repository/auth"
	_foodsRepo "HealthFit/repository/foods"
	_goalRepo "HealthFit/repository/goal"
	_userRepo "HealthFit/repository/user"

	"fmt"
	"log"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/coreapi"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}
func main() {
	config := config.GetConfig()
	db := utils.InitDB(config)
	// awsConn := awss3.InitS3(config.S3_KEY, config.S3_SECRET, config.S3_REGION)
	midtransConfig := coreapi.Client{}
	midtransConfig.New(config.Midtrans, midtrans.Sandbox)

	//REPOSITORY-DATABASE
	authRepo := _authRepo.New(db)
	userRepo := _userRepo.New(db)
	goalRepo := _goalRepo.New(db)
	foodsRepo := _foodsRepo.New(db)

	//CONTROLLER
	authController := _authController.New(authRepo)
	userController := _userController.New(userRepo)
	goalController := _goalController.New(goalRepo)
	foodsController := _foodsController.New(foodsRepo)

	e := echo.New()
	e.Validator = &CustomValidator{validator: validator.New()}

	route.RegisterPath(e,
		userController,
		authController,
		goalController,
		foodsController,
	)

	log.Fatal(e.Start(fmt.Sprintf(":%d", config.Port)))

}
