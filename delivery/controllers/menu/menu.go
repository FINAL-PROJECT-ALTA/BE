package menu

import (
	"HealthFit/delivery/controllers/common"
	"HealthFit/delivery/middlewares"
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

		isAdmin := middlewares.ExtractRoles(c)
		if !isAdmin {
			return c.JSON(http.StatusBadRequest, common.BadRequest(http.StatusBadRequest, "access denied ", nil))
		}
		newMenu := MenuCreateRequestFormat{}

		c.Bind(&newMenu)
		errB := c.Validate(&newMenu)
		if errB != nil {
			return c.JSON(http.StatusBadRequest, common.BadRequest(http.StatusBadRequest, "access denied ", nil))
		}

		res, err := fc.repo.Create(entities.Menu{
			Menu_category: newMenu.Menu_category,
		})

		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.InternalServerError(http.StatusInternalServerError, "There is some error on server", nil))
		}

		response := MenuCreateResponse{}
		response.Menu_uid = res.Menu_uid
		response.Menu_category = res.Menu_category

		return c.JSON(http.StatusCreated, common.Success(http.StatusCreated, "Success create foods", response))

	}
}

func (fc *FoodsController) Update() echo.HandlerFunc {
	return func(c echo.Context) error {

		isAdmin := middlewares.ExtractRoles(c)
		if !isAdmin {
			return c.JSON(http.StatusBadRequest, common.BadRequest(http.StatusBadRequest, "access denied ", nil))
		}

		food_uid := c.Param("menu_uid")
		newMenu := MenuUpdateRequestFormat{}

		c.Bind(&newMenu)
		errB := c.Validate(&newMenu)
		if errB != nil {
			return c.JSON(http.StatusBadRequest, common.BadRequest(http.StatusBadRequest, "access denied ", nil))
		}

		res, err := fc.repo.Update(food_uid, entities.Menu{
			Menu_category: newMenu.Menu_category,
		})

		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.InternalServerError(http.StatusInternalServerError, "There is some error on server", nil))
		}

		response := MenuUpdateResponse{}
		response.Menu_uid = res.Menu_uid
		response.Menu_category = res.Menu_category

		return c.JSON(http.StatusCreated, common.Success(http.StatusCreated, "Success create foods", response))

	}
}

func (fc *FoodsController) Delete() echo.HandlerFunc {
	return func(c echo.Context) error {

		isAdmin := middlewares.ExtractRoles(c)
		if !isAdmin {
			return c.JSON(http.StatusBadRequest, common.BadRequest(http.StatusBadRequest, "access denied ", nil))
		}

		food_uid := c.Param("menu_uid")
		err := fc.repo.Delete(food_uid)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.InternalServerError(http.StatusInternalServerError, "There is some error on server", nil))
		}

		return c.JSON(http.StatusCreated, common.Success(http.StatusCreated, "Success create foods", nil))

	}
}
