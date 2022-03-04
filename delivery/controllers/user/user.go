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

		return c.JSON(http.StatusCreated, common.Success(http.StatusCreated, "Success Create User", ToCreateUserResponseFormat(res)))
	}
}

func (ac *UserController) GetById() echo.HandlerFunc {
	return func(c echo.Context) error {
		user_uid := middlewares.ExtractTokenUserUid(c)

		res, err := ac.repo.GetById(user_uid)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.InternalServerError(http.StatusInternalServerError, "Not Found", nil))
		}

		return c.JSON(http.StatusOK, common.Success(http.StatusOK, "Success Get User", ToGetUserByIdResponseFormat(res)))
	}
}

func (ac *UserController) Update() echo.HandlerFunc {
	return func(c echo.Context) error {
		user_uid := middlewares.ExtractTokenUserUid(c)
		var newUser = UpdateUserRequestFormat{}
		c.Bind(&newUser)

		if newUser.Email != "" {
			err := c.Validate(&newUser)
			if err != nil {
				return c.JSON(http.StatusBadRequest, common.BadRequest(http.StatusBadRequest, "There is some problem from input", nil))
			}
		}

		res, err := ac.repo.Update(user_uid, entities.User{Name: newUser.Name, Email: newUser.Email, Password: newUser.Password, Gender: newUser.Gender})

		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.InternalServerError(http.StatusInternalServerError, "There is some error on server", nil))
		}

		return c.JSON(http.StatusOK, common.Success(http.StatusOK, "Success Update User", ToUpdateUserResponseFormat(res)))
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
