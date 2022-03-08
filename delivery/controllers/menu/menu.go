package menu

import (
	"HealthFit/delivery/controllers/common"
	"HealthFit/delivery/middlewares"
	"HealthFit/entities"
	"HealthFit/repository/menu"
	"fmt"
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

		user := middlewares.ExtractTokenUserUid(c)

		newMenu := MenuCreateRequestFormat{}
		newMenu.User_uid = user

		c.Bind(&newMenu)
		errB := c.Validate(&newMenu)
		if errB != nil {
			return c.JSON(http.StatusBadRequest, common.BadRequest(http.StatusBadRequest, "access denied ", nil))
		}

		res, err := mc.repo.Create(newMenu.Foods, entities.Menu{User_uid: newMenu.User_uid,
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

		response := []MenuGetAllResponse{}
		for i, result := range res {
			var foods []entities.Food
			for _, resultfood := range res[i].Detail_menu {
				foods = append(foods, resultfood.Food)
			}
			response = append(response, MenuGetAllResponse{
				Menu_uid:      result.Menu_uid,
				Menu_category: result.Menu_category,
				Foods:         foods,
			})
		}
		fmt.Println(res)

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

		response := []MenuGetAllResponse{}
		for i, result := range res {
			var foods []entities.Food
			for _, resultfood := range res[i].Detail_menu {
				foods = append(foods, resultfood.Food)
			}
			response = append(response, MenuGetAllResponse{
				Menu_uid:      result.Menu_uid,
				Menu_category: result.Menu_category,
				Foods:         foods,
			})
		}

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
		var foods []entities.Food
		for _, result := range res.Detail_menu {
			foods = append(foods, result.Food)
		}
		response.Foods = foods

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
