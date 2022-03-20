package userhistory

import (
	"HealthFit/delivery/controllers/common"
	"HealthFit/delivery/middlewares"
	"HealthFit/entities"
	userhistory "HealthFit/repository/user_history"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserHistoryController struct {
	repo userhistory.UserHistory
}

func New(repository userhistory.UserHistory) *UserHistoryController {
	return &UserHistoryController{
		repo: repository,
	}
}

func (uh *UserHistoryController) Insert() echo.HandlerFunc {
	return func(c echo.Context) error {

		isAdmin := middlewares.ExtractRoles(c)
		if isAdmin {
			return c.JSON(http.StatusUnauthorized, common.BadRequest(http.StatusUnauthorized, "access denied", nil))

		}

		user := middlewares.ExtractTokenUserUid(c)
		userHistory := CreateUserHistoryRequestFormat{}
		userHistory.User_uid = user

		c.Bind(&userHistory)
		err := c.Validate(&userHistory)

		if err != nil {
			return c.JSON(http.StatusBadRequest, common.BadRequest(http.StatusBadRequest, "There is some problem from input", nil))
		}

		res, errH := uh.repo.Insert(entities.User_history{
			User_uid: userHistory.User_uid,
			Menu_uid: userHistory.Menu_uid,
			Goal_uid: userHistory.Goal_uid,
		})

		if errH != nil {
			return c.JSON(http.StatusInternalServerError, common.InternalServerError(http.StatusInternalServerError, "Internal Server Error", nil))
		}
		response := CreateUserHistoryResponse{}
		response.User_history_uid = res.User_history_uid
		response.Goal_uid = res.Goal_uid
		response.Menu_uid = res.Menu_uid

		return c.JSON(http.StatusCreated, common.Success(http.StatusCreated, "Success Create User History", response))

	}
}

func (uh *UserHistoryController) GetAll() echo.HandlerFunc {
	return func(c echo.Context) error {

		isAdmin := middlewares.ExtractRoles(c)
		if isAdmin {
			return c.JSON(http.StatusUnauthorized, common.BadRequest(http.StatusUnauthorized, "access denied", nil))

		}

		user := middlewares.ExtractTokenUserUid(c)

		res, err := uh.repo.GetAll(user)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.InternalServerError(http.StatusInternalServerError, "There is some error on server", nil))
		}
		allHistory := []GetUserHistoryByIdResponse{}
		for i := 0; i < len(res); i++ {
			response := GetUserHistoryByIdResponse{}
			response.User_history_uid = res[i].User_history_uid
			response.Goal_uid = res[i].Goal_uid
			response.Menu.Menu_uid = res[i].Menu_uid
			response.Menu.Menu_category = res[i].Menu.Menu_category
			response.Menu.Created_by = res[i].Menu.Created_by
			response.Menu.Total_calories = res[i].Menu.Total_calories
			for j := 0; j < len(res[i].Menu.Detail_menu); j++ {
				response.Menu.Foods = append(response.Menu.Foods, res[i].Menu.Detail_menu[j].Food)
			}
			allHistory = append(allHistory, response)
		}

		return c.JSON(http.StatusOK, common.Success(http.StatusOK, "Success get user histories", allHistory))
	}
}

func (uh *UserHistoryController) GetByUid() echo.HandlerFunc {
	return func(c echo.Context) error {

		isAdmin := middlewares.ExtractRoles(c)
		if isAdmin {
			return c.JSON(http.StatusUnauthorized, common.BadRequest(http.StatusUnauthorized, "access denied", nil))

		}

		user := middlewares.ExtractTokenUserUid(c)
		userHistory_uid := c.Param("user_history_uid")

		res, err := uh.repo.GetById(user, userHistory_uid)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.InternalServerError(http.StatusInternalServerError, "Internal Server Error", nil))
		}

		response := GetUserHistoryByIdResponse{}
		response.User_history_uid = res.User_history_uid
		response.Goal_uid = res.Goal_uid
		response.Menu.Menu_uid = res.Menu_uid
		response.Menu.Menu_category = res.Menu.Menu_category
		response.Menu.Created_by = res.Menu.Created_by
		response.Menu.Total_calories = res.Menu.Total_calories
		for i := 0; i < len(res.Menu.Detail_menu); i++ {
			response.Menu.Foods = append(response.Menu.Foods, res.Menu.Detail_menu[i].Food)
		}

		return c.JSON(http.StatusOK, common.Success(http.StatusOK, "Success get user", response))
	}
}
