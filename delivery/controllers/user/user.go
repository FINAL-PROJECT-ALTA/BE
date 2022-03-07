package user

import (
	"HealthFit/delivery/controllers/common"
	"HealthFit/delivery/middlewares"
	"HealthFit/entities"
	"HealthFit/repository/user"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserController struct {
	repo user.User
}

func New(repository user.User) *UserController {
	return &UserController{
		repo: repository,
	}
}

func (ac *UserController) Register() echo.HandlerFunc {
	return func(c echo.Context) error {
		user := CreateUserRequestFormat{}

		c.Bind(&user)
		err := c.Validate(&user)

		if err != nil {
			return c.JSON(http.StatusBadRequest, common.BadRequest(http.StatusBadRequest, "There is some problem from input", nil))
		}

		res, err := ac.repo.Register(entities.User{Name: user.Name, Email: user.Email, Password: user.Password, Gender: user.Gender})

		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.InternalServerError(http.StatusInternalServerError, "There is some error on server", nil))
		}

		response := UserCreateResponse{}
		response.User_uid = res.User_uid
		response.Name = res.Name
		response.Email = res.Email
		response.Gender = res.Gender

		return c.JSON(http.StatusCreated, common.Success(http.StatusCreated, "Success Create User", response))

	}
}
func (ac *UserController) GetById() echo.HandlerFunc {
	return func(c echo.Context) error {
		user_uid := middlewares.ExtractTokenUserUid(c)

		res, err := ac.repo.GetById(user_uid)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.InternalServerError(http.StatusInternalServerError, "There is some error on server", nil))
		}

		response := UserCompleksResponse{}

		response.User_uid = res.User_uid
		response.Name = res.Name
		response.Email = res.Email
		response.Gender = res.Gender

		responseGoal := []UserGoal{}
		for i := 0; i < len(res.Goal); i++ {
			user_goal := UserGoal{}
			user_goal.Height = res.Goal[i].Height
			user_goal.Weight = res.Goal[i].Weight
			user_goal.Age = res.Goal[i].Age
			user_goal.Range_time = res.Goal[i].Range_time
			responseGoal = append(responseGoal, user_goal)
		}
		responseHistory := []UserHistoryResponse{}
		for i := 0; i < len(res.Goal); i++ {
			user_history := UserHistoryResponse{}
			user_history.User_uid = res.History[i].User_uid
			user_history.Menu_uid = res.History[i].Menu_uid

			responseHistory = append(responseHistory, user_history)
		}
		response.Goal = responseGoal
		response.History = responseHistory

		return c.JSON(http.StatusOK, common.Success(http.StatusOK, "Success get user", response))
	}
}

func (ac *UserController) Update() echo.HandlerFunc {
	return func(c echo.Context) error {
		user_uid := middlewares.ExtractTokenUserUid(c)
		var newUser = UpdateUserRequestFormat{}
		c.Bind(&newUser)

		err := c.Validate(&newUser)
		if err != nil {
			return c.JSON(http.StatusBadRequest, common.BadRequest(http.StatusBadRequest, "There is some problem from input", nil))
		}

		res, err := ac.repo.Update(user_uid, entities.User{Name: newUser.Name, Email: newUser.Email, Password: newUser.Password, Gender: newUser.Gender})

		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.InternalServerError(http.StatusInternalServerError, "There is some error on server", nil))
		}

		response := UserUpdateResponse{}
		response.User_uid = res.User_uid
		response.Name = res.Name
		response.Email = res.Email
		response.Gender = res.Gender

		return c.JSON(http.StatusOK, common.Success(http.StatusOK, "Success Update User", response))
	}
}

func (ac *UserController) Delete() echo.HandlerFunc {
	return func(c echo.Context) error {
		user_uid := middlewares.ExtractTokenUserUid(c)

		err := ac.repo.Delete(user_uid)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.InternalServerError(http.StatusInternalServerError, "There is some error on server", nil))
		}

		return c.JSON(http.StatusOK, common.Success(http.StatusOK, "Success Delete User", nil))
	}
}
