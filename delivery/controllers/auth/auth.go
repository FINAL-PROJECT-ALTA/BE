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

		c.Bind(&Userlogin)
		err := c.Validate(&Userlogin)

		if err != nil {
			return c.JSON(http.StatusBadRequest, common.BadRequest(http.StatusBadRequest, "There is some problem from input", nil))
		}

		checkedUser, err := ac.repo.Login(Userlogin.Email, Userlogin.Password)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.InternalServerError(nil, "error in call database", nil))
		}
		token, err := middlewares.GenerateToken(checkedUser)
		response := UserLoginResponse{
			User_uid: checkedUser.User_uid,
			Name:     checkedUser.Name,
			Email:    checkedUser.Email,
			Token:    token,
		}

		if err != nil {
			return c.JSON(http.StatusNotAcceptable, common.BadRequest(http.StatusNotAcceptable, "error in process token", nil))
		}

		return c.JSON(http.StatusOK, common.Success(http.StatusOK, "USERS - berhasil masuk, mendapatkan token baru", response))

	}
}

func (ac *AuthController) AdminLogin() echo.HandlerFunc {
	return func(c echo.Context) error {
		AdminLogin := LoginReqFormat{}

		c.Bind(&AdminLogin)
		err := c.Validate(&AdminLogin)

		if err != nil {
			return c.JSON(http.StatusBadRequest, common.BadRequest(http.StatusBadRequest, "There is some problem from input", nil))
		}
		checkedAdmin, err := ac.repo.LoginAdmin(AdminLogin.Email, AdminLogin.Password)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.InternalServerError(nil, "error in call database", nil))
		}
		token, err := middlewares.GenerateTokenAdmin(checkedAdmin)
		response := AdminLoginResponse{
			Admin_uid: checkedAdmin.Admin_uid,
			Name:      checkedAdmin.Name,
			Email:     checkedAdmin.Email,
			Token:     token,
		}

		if err != nil {
			return c.JSON(http.StatusNotAcceptable, common.BadRequest(http.StatusNotAcceptable, "error in process token", nil))
		}

		return c.JSON(http.StatusOK, common.Success(http.StatusOK, "ADMIN - berhasil masuk, mendapatkan token baru", response))

	}
}
