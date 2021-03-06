package route

import (
	"HealthFit/delivery/controllers/admin"
	"HealthFit/delivery/controllers/auth"
	"HealthFit/delivery/controllers/foods"
	"HealthFit/delivery/controllers/goal"
	"HealthFit/delivery/controllers/menu"
	"HealthFit/delivery/controllers/user"
	userhistory "HealthFit/delivery/controllers/user_history"
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
	mc *menu.MenuController,
	uh *userhistory.UserHistoryController,
) {

	//CORS
	e.Use(middleware.CORS())

	//LOGGER
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}",
	}))

	//ROUTE ADMIN
	e.POST("/admin/register", ac.Register())
	e.POST("/admin/login", aa.AdminLogin())
	e.GET("/admin", uc.GetById(), middlewares.JwtMiddleware())
	e.PUT("/admin", uc.Update(), middlewares.JwtMiddleware())
	e.DELETE("/admin", uc.Delete(), middlewares.JwtMiddleware())

	//ROUTE REGISTER - LOGIN USERS
	e.POST("users/register", uc.Register())
	e.POST("users/login", aa.Login())

	//ROUTE USERS
	e.GET("/users", uc.GetById(), middlewares.JwtMiddleware())
	e.PUT("/users", uc.Update(), middlewares.JwtMiddleware())
	e.DELETE("/users", uc.Delete(), middlewares.JwtMiddleware())

	//ROUTE GOALS
	e.POST("/users/goals", gc.Create(), middlewares.JwtMiddleware())
	e.GET("/users/goals", gc.GetAll(), middlewares.JwtMiddleware())
	e.GET("/users/goals/:goal_uid", gc.GetById(), middlewares.JwtMiddleware())
	e.PUT("/users/goals/:goal_uid", gc.Update(), middlewares.JwtMiddleware())
	e.DELETE("/users/goals/:goal_uid", gc.Delete(), middlewares.JwtMiddleware())
	e.PUT("/users/goals/cancel", gc.CencelGoal(), middlewares.JwtMiddleware())

	//ROUTE FOODS
	e.POST("/foods", fc.Create(), middlewares.JwtMiddleware())
	e.GET("/foods", fc.GetAll())
	e.GET("/foods/search", fc.Search())
	e.GET("/foods/:food_uid", fc.GetById())
	e.PUT("/foods/:food_uid", fc.Update(), middlewares.JwtMiddleware())
	e.DELETE("/foods/:food_uid", fc.Delete(), middlewares.JwtMiddleware())
	e.GET("/foods/be", fc.GetFromThirdPary(), middlewares.JwtMiddleware())

	//ROUTE MENU
	e.POST("/menus", mc.Create(), middlewares.JwtMiddleware())
	e.GET("/menus", mc.GetAll())
	e.GET("/menus/recommend/breakfast", mc.GetRecommendBreakfast(), middlewares.JwtMiddleware())
	e.GET("/menus/recommend/lunch", mc.GetRecommendLunch(), middlewares.JwtMiddleware())
	e.GET("/menus/recommend/dinner", mc.GetRecommendDinner(), middlewares.JwtMiddleware())
	e.GET("/menus/recommend/overtime", mc.GetRecommendOverTime(), middlewares.JwtMiddleware())
	e.PUT("/menus/:menu_uid", mc.Update(), middlewares.JwtMiddleware())
	e.DELETE("/menus/:menu_uid", mc.Delete(), middlewares.JwtMiddleware())

	//ROUTE USER HISTORY
	e.GET("/userhistories", uh.GetAll(), middlewares.JwtMiddleware())
	e.GET("/userhistories/:user_history_uid", uh.GetByUid(), middlewares.JwtMiddleware())
	e.POST("/userhistories", uh.Insert(), middlewares.JwtMiddleware())

}
