package admin

import (
	"HealthFit/delivery/controllers/common"
	"HealthFit/entities"
	"HealthFit/repository/admin"
	"net/http"

	"github.com/labstack/echo/v4"
)

type AdminController struct {
	repo admin.Admin
}

func New(repository admin.Admin) *AdminController {
	return &AdminController{
		repo: repository,
	}
}

func (ac *AdminController) Register() echo.HandlerFunc {
	return func(c echo.Context) error {
		admin := CreateAdminRequestFormat{}

		c.Bind(&admin)
		err := c.Validate(&admin)

		if err != nil {
			return c.JSON(http.StatusBadRequest, common.BadRequest(http.StatusBadRequest, "There is some problem from input", nil))
		}

		res, err_repo := ac.repo.Register(entities.User{Name: admin.Name, Email: admin.Email, Password: admin.Password, Gender: admin.Gender, Roles: true})

		if err_repo != nil {
			return c.JSON(http.StatusConflict, common.InternalServerError(http.StatusConflict, err_repo.Error(), nil))
		}

		response := AdminResponse{}
		response.Admin_uid = res.User_uid
		response.Name = res.Name
		response.Email = res.Email
		response.Gender = res.Gender
		response.Roles = res.Roles

		return c.JSON(http.StatusCreated, common.Success(http.StatusCreated, "Success create admin", response))

	}
}
