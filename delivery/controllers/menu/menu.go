package menu

import (
	"HealthFit/delivery/controllers/common"
	"HealthFit/delivery/middlewares"
	"HealthFit/entities"
	"HealthFit/repository/menu"
	"net/http"

	"github.com/labstack/echo/v4"
)

type MenuController struct {
	repo menu.Menu
}

func New(repository menu.Menu) *MenuController {
	return &MenuController{
		repo: repository,
	}
}

func (mc *MenuController) Create() echo.HandlerFunc {
	return func(c echo.Context) error {

		isAdmin := middlewares.ExtractRoles(c)
		if !isAdmin {
			return c.JSON(http.StatusUnauthorized, common.BadRequest(http.StatusUnauthorized, "access denied ", nil))
		}
		newMenu := MenuCreateRequestFormat{}

		c.Bind(&newMenu)
		errB := c.Validate(&newMenu)
		if errB != nil {
			return c.JSON(http.StatusBadRequest, common.BadRequest(http.StatusBadRequest, "access denied ", nil))
		}

		res, err := mc.repo.Create(entities.Menu{
			Menu_category: newMenu.Menu_category,
		})

		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.InternalServerError(http.StatusInternalServerError, "There is some error on server", nil))
		}

		response := MenuCreateResponse{}
		response.Menu_uid = res.Menu_uid
		response.Menu_category = res.Menu_category

		return c.JSON(http.StatusCreated, common.Success(http.StatusCreated, "Success create menu", response))

	}
}

func (mc *MenuController) GetAll() echo.HandlerFunc {
	return func(c echo.Context) error {
		res, err := mc.repo.GetAllMenu()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.InternalServerError(http.StatusInternalServerError, "There is some error on server", nil))
		}

		response := []MenuGetResponse{}
		for _, result := range res {
			response = append(response, MenuGetResponse{
				Menu_uid:      result.Menu_uid,
				Menu_category: result.Menu_category,
				// Foods: result.Foods,
			})
		}

		return c.JSON(http.StatusOK, common.Success(http.StatusOK, "Success Get All Menu Category", response))
	}
}

func (mc *MenuController) GetMenuByMenuCategory() echo.HandlerFunc {
	return func(c echo.Context) error {
		menuCategory := c.Param("menu_category")

		res, err := mc.repo.GetMenuByCategory(menuCategory)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.InternalServerError(http.StatusInternalServerError, "There is some error on server", nil))
		}

		response := MenuGetResponse{}
		response.Menu_uid = res.Menu_uid
		response.Menu_category = res.Menu_category

		return c.JSON(http.StatusOK, common.Success(http.StatusOK, "Success Get Menu Category", response))
	}
}

func (mc *MenuController) Update() echo.HandlerFunc {
	return func(c echo.Context) error {

		isAdmin := middlewares.ExtractRoles(c)
		if !isAdmin {
			return c.JSON(http.StatusBadRequest, common.BadRequest(http.StatusBadRequest, "access denied ", nil))
		}

		menu_uid := c.Param("menu_uid")
		newMenu := MenuUpdateRequestFormat{}

		c.Bind(&newMenu)
		errB := c.Validate(&newMenu)
		if errB != nil {
			return c.JSON(http.StatusBadRequest, common.BadRequest(http.StatusBadRequest, "access denied ", nil))
		}

		res, err := mc.repo.Update(menu_uid, entities.Menu{
			Menu_category: newMenu.Menu_category,
		})

		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.InternalServerError(http.StatusInternalServerError, "There is some error on server", nil))
		}

		response := MenuUpdateResponse{}
		response.Menu_uid = res.Menu_uid
		response.Menu_category = res.Menu_category

		return c.JSON(http.StatusCreated, common.Success(http.StatusCreated, "Success update menu", response))

	}
}

func (mc *MenuController) Delete() echo.HandlerFunc {
	return func(c echo.Context) error {

		isAdmin := middlewares.ExtractRoles(c)
		if !isAdmin {
			return c.JSON(http.StatusBadRequest, common.BadRequest(http.StatusBadRequest, "access denied ", nil))
		}

		menu_uid := c.Param("menu_uid")
		err := mc.repo.Delete(menu_uid)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.InternalServerError(http.StatusInternalServerError, "There is some error on server", nil))
		}

		return c.JSON(http.StatusCreated, common.Success(http.StatusCreated, "Success delete menu", nil))

	}
}