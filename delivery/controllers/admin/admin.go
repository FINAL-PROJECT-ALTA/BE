package admin

import (
	"HealthFit/delivery/controllers/common"
	"HealthFit/delivery/middlewares"
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

		res, err := ac.repo.Register(entities.Admin{Name: admin.Name, Email: admin.Email, Password: admin.Password, Gender: admin.Gender})

		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.InternalServerError(http.StatusInternalServerError, "There is some error on server", nil))
		}

		response := AdminResponse{}
		response.Admin_uid = res.Admin_uid
		response.Name = res.Name
		response.Email = res.Email
		response.Gender = res.Gender

		return c.JSON(http.StatusCreated, common.Success(http.StatusCreated, "Success create admin", response))

	}
}
func (ac *AdminController) GetById() echo.HandlerFunc {
	return func(c echo.Context) error {
		admin_uid, _ := middlewares.ExtractTokenAdminUid(c)

		res, err := ac.repo.GetById(admin_uid)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.InternalServerError(http.StatusInternalServerError, "There is some error on server", nil))
		}

		response := AdminResponse{}

		response.Admin_uid = res.Admin_uid
		response.Name = res.Name
		response.Email = res.Email
		response.Gender = res.Gender

		return c.JSON(http.StatusOK, common.Success(http.StatusOK, "Success get admin", response))
	}
}

func (ac *AdminController) Update() echo.HandlerFunc {
	return func(c echo.Context) error {
		admin_uid, _ := middlewares.ExtractTokenAdminUid(c)
		var newAdmin = UpdateAdminRequestFormat{}
		c.Bind(&newAdmin)

		if newAdmin.Email != "" {
			err := c.Validate(&newAdmin)
			if err != nil {
				return c.JSON(http.StatusBadRequest, common.BadRequest(http.StatusBadRequest, "There is some problem from input", nil))
			}
		}

		res, err := ac.repo.Update(admin_uid, entities.Admin{Name: newAdmin.Name, Email: newAdmin.Email, Password: newAdmin.Password, Gender: newAdmin.Gender})

		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.InternalServerError(http.StatusInternalServerError, "There is some error on server", nil))
		}

		response := AdminResponse{}
		response.Admin_uid = res.Admin_uid
		response.Name = res.Name
		response.Email = res.Email
		response.Gender = res.Gender

		return c.JSON(http.StatusOK, common.Success(http.StatusOK, "Success update admin", response))
	}
}

func (ac *AdminController) Delete() echo.HandlerFunc {
	return func(c echo.Context) error {
		admin_uid, _ := middlewares.ExtractTokenAdminUid(c)

		err := ac.repo.Delete(admin_uid)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.InternalServerError(http.StatusInternalServerError, "There is some error on server", nil))
		}

		return c.JSON(http.StatusOK, common.Success(http.StatusOK, "Success delete admin", nil))
	}
}
