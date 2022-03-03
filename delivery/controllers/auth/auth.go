package auth

import (
	"HealthFit/delivery/controllers/common"
	"HealthFit/delivery/middlewares"
	"HealthFit/repository/auth"
	"net/http"

	"github.com/labstack/echo/v4"
)

type AuthController struct {
	repo auth.Auth
}

func New(repo auth.Auth) *AuthController {
	return &AuthController{
		repo: repo,
	}
}

func (ac *AuthController) Login() echo.HandlerFunc {
	return func(c echo.Context) error {
		Userlogin := LoginReqFormat{}

		if err := c.Bind(&Userlogin); err != nil || Userlogin.Email == "" || Userlogin.Password == "" {
			return c.JSON(http.StatusBadRequest, common.BadRequest(nil, "error in input file", nil))
		}
		checkedUser, err := ac.repo.Login(Userlogin.Email, Userlogin.Password)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.InternalServerError(nil, "error in call database", nil))
		}
		token, err := middlewares.GenerateToken(checkedUser)

		if err != nil {
			return c.JSON(http.StatusNotAcceptable, common.BadRequest(http.StatusNotAcceptable, "error in process token", nil))
		}

		return c.JSON(http.StatusOK, common.Success(http.StatusOK, "USERS - berhasil masuk, mendapatkan token baru", token))

	}
}

func (ac *AuthController) AdminLogin() echo.HandlerFunc {
	return func(c echo.Context) error {
		Userlogin := LoginReqFormat{}

		if err := c.Bind(&Userlogin); err != nil || Userlogin.Email == "" || Userlogin.Password == "" {
			return c.JSON(http.StatusBadRequest, common.BadRequest(nil, "error in input file", nil))
		}
		checkedUser, err := ac.repo.Login(Userlogin.Email, Userlogin.Password)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.InternalServerError(nil, "error in call database", nil))
		}
		token, err := middlewares.GenerateToken(checkedUser)

		if err != nil {
			return c.JSON(http.StatusNotAcceptable, common.BadRequest(http.StatusNotAcceptable, "error in process token", nil))
		}

		return c.JSON(http.StatusOK, common.Success(http.StatusOK, "ADMIN - berhasil masuk, mendapatkan token baru", token))

	}
}
