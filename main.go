package main

import (
	config "HealthFit/configs"
	route "HealthFit/delivery/routes"
	utils "HealthFit/utils/mysql"

	_adminController "HealthFit/delivery/controllers/admin"
	_authController "HealthFit/delivery/controllers/auth"

	_foodsController "HealthFit/delivery/controllers/foods"
	_goalController "HealthFit/delivery/controllers/goal"
	_menuController "HealthFit/delivery/controllers/menu"
	_userController "HealthFit/delivery/controllers/user"
	_userHistoryController "HealthFit/delivery/controllers/user_history"

	_adminRepo "HealthFit/repository/admin"
	_authRepo "HealthFit/repository/auth"
	_foodsRepo "HealthFit/repository/foods"
	_goalRepo "HealthFit/repository/goal"
	_menuRepo "HealthFit/repository/menu"
	_userRepo "HealthFit/repository/user"
	_userHistoryRepo "HealthFit/repository/user_history"

	awsS3 "HealthFit/utils/aws_S3"

	"fmt"
	"log"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
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
	awsConn := awsS3.InitS3(config.S3_KEY, config.S3_SECRET, config.S3_REGION)

	//REPOSITORY-DATABASE
	authRepo := _authRepo.New(db)
	adminRepo := _adminRepo.New(db)
	userRepo := _userRepo.New(db)
	goalRepo := _goalRepo.New(db)
	foodsRepo := _foodsRepo.New(db)
	menuRepo := _menuRepo.New(db)
	userHistoryRepo := _userHistoryRepo.New(db)

	//CONTROLLER
	authController := _authController.New(authRepo)
	adminController := _adminController.New(adminRepo)
	userController := _userController.New(userRepo, awsConn)
	goalController := _goalController.New(goalRepo)
	foodsController := _foodsController.New(foodsRepo, awsConn)
	menuController := _menuController.New(menuRepo)
	userHistoryController := _userHistoryController.New(userHistoryRepo)

	e := echo.New()
	e.Validator = &CustomValidator{validator: validator.New()}

	route.RegisterPath(e,
		userController,
		authController,
		goalController,
		foodsController,
		adminController,
		menuController,
		userHistoryController,
	)

	log.Fatal(e.Start(fmt.Sprintf(":%d", config.Port)))

}
