package route

import (
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

	//ROUTE USERS
	e.GET("/users", uc.GetById(), middlewares.JwtMiddleware())
	e.PUT("/users", uc.Update(), middlewares.JwtMiddleware())
	e.DELETE("/users", uc.Delete(), middlewares.JwtMiddleware())

	//ROUTE GOALS
	e.POST("/goals", gc.Create(), middlewares.JwtMiddleware())
	e.GET("/goals/:goal_uid", gc.GetById(), middlewares.JwtMiddleware())
	e.PUT("/goals/:goal_uid", gc.Update(), middlewares.JwtMiddleware())
	e.DELETE("/goals/:goal_uid", gc.Delete(), middlewares.JwtMiddleware())

	//ROUTE FOODS
	e.POST("/foods", fc.Create(), middlewares.JwtMiddleware())
	e.GET("foods", fc.GetAll())
}
