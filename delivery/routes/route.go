package route

import (
	"HealthFit/delivery/controllers/admin"
	"HealthFit/delivery/controllers/auth"
	"HealthFit/delivery/controllers/foods"
	"HealthFit/delivery/controllers/goal"
	"HealthFit/delivery/controllers/user"
	"HealthFit/delivery/middlewares"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func RegisterPath(e *echo.Echo,
	uc *user.UserController,
	aa *auth.AuthController,
	gc *goal.GoalController,
	fc *foods.FoodsController,
	ac *admin.AdminController,
) {

	//CORS
	e.Use(middleware.CORS())

	//LOGGER
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}",
	}))

	//ROUTE REGISTER - LOGIN USERS
	e.POST("users/register", uc.Register())
	e.POST("users/login", aa.Login())

	//ROUTE ADMIN
	e.POST("/admin/register", ac.Register())
	e.POST("/admin/login", aa.AdminLogin())
	e.GET("/admin", ac.GetById(), middlewares.JwtMiddleware())
	e.PUT("/admin", ac.Update(), middlewares.JwtMiddleware())
	e.DELETE("/admin", ac.Delete(), middlewares.JwtMiddleware())

	//ROUTE USERS
	e.GET("/users", uc.GetById(), middlewares.JwtMiddleware())
	e.PUT("/users", uc.Update(), middlewares.JwtMiddleware())
	e.DELETE("/users", uc.Delete(), middlewares.JwtMiddleware())

	//ROUTE GOALS
	e.POST("/users/goals", gc.Create(), middlewares.JwtMiddleware())
	e.GET("/users/goals/:goal_uid", gc.GetById(), middlewares.JwtMiddleware())
	e.PUT("/users/goals/:goal_uid", gc.Update(), middlewares.JwtMiddleware())
	e.DELETE("/users/goals/:goal_uid", gc.Delete(), middlewares.JwtMiddleware())

	//ROUTE FOODS
	e.POST("/foods", fc.Create(), middlewares.JwtMiddleware())
	e.GET("/foods", fc.GetAll())
	e.GET("/foods/search", fc.Search())
	e.PUT("/foods/:food_uid", fc.Update(), middlewares.JwtMiddleware())
	e.DELETE("/foods/:food_uid", fc.Delete(), middlewares.JwtMiddleware())

}
