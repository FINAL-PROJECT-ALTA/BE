package menu

import (
	"HealthFit/delivery/controllers/common"
	"HealthFit/entities"
	"HealthFit/repository/menu"
	"net/http"

	"github.com/labstack/echo/v4"
)

type FoodsController struct {
	repo menu.Menu
}

func New(repository menu.Menu) *FoodsController {
	return &FoodsController{
		repo: repository,
	}
}

func (fc *FoodsController) Create() echo.HandlerFunc {
	return func(c echo.Context) error {

		// isAdmin := middlewares.ExtractTokenAdminUid(c)
		newMenu := MenuCreateRequestFormat{}

		c.Bind(&newMenu)
		err := c.Validate(&newMenu)
		if err != nil {
			return c.JSON(http.StatusBadRequest, common.BadRequest(http.StatusBadRequest, "access denied ", nil))
		}

		res, err := fc.repo.Create(entities.Menu{})

		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.InternalServerError(http.StatusInternalServerError, "There is some error on server", nil))
		}

		response := MenuCreateResponse{}
		response.Menu_uid = res.Menu_uid
		response.Menu_category = res.Menu_category
		response.Total_calories = res.Total_calories

		return c.JSON(http.StatusCreated, common.Success(http.StatusCreated, "Success create foods", response))

	}
}
