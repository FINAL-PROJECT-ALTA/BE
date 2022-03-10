package detailmenu

import (
	"HealthFit/delivery/controllers/common"
	"HealthFit/entities"
	detailmenu "HealthFit/repository/detail_menu"
	"net/http"

	"github.com/labstack/echo/v4"
)

type DetailMenuController struct {
	repo detailmenu.Detail_menu
}

func New(repository detailmenu.Detail_menu) *DetailMenuController {
	return &DetailMenuController{
		repo: repository,
	}
}

func (dc *DetailMenuController) Create() echo.HandlerFunc {
	return func(c echo.Context) error {

		newDetailMenu := DetailMenuCreateRequestFormat{}
		c.Bind(&newDetailMenu)
		errB := c.Validate(&newDetailMenu)
		if errB != nil {
			return c.JSON(http.StatusBadRequest, common.BadRequest(http.StatusBadRequest, "There is some problem from input", nil))
		}

		res, err := dc.repo.Create(entities.Detail_menu{
			Menu_uid: newDetailMenu.Menu_uid,
			Food_uid: newDetailMenu.Food_uid,
		})

		if err != nil {
			return c.JSON(http.StatusBadRequest, common.BadRequest(http.StatusBadRequest, "There is some problem from input", nil))
		}

		response := DetailMenuCreateResponse{}
		response.Detail_menu_uid = res.Detail_menu_uid
		response.Menu_uid = res.Menu_uid
		response.Food_uid = res.Food_uid

		return c.JSON(http.StatusCreated, common.Success(http.StatusCreated, "succes to create detail maenu", response))
	}
}

func (dc *DetailMenuController) GetByUid() echo.HandlerFunc {
	return func(c echo.Context) error {
		detail_menu_uid := c.Param("detail_menu_uid")

		res, err := dc.repo.GetDetailMenuByUid(detail_menu_uid)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.InternalServerError(http.StatusInternalServerError, "There is some error on server", nil))
		}

		response := DetailMenuGetResponse{}
		response.Detail_menu_uid = res.Detail_menu_uid
		response.Menu_uid = res.Menu_uid
		response.Food_uid = res.Food_uid

		return c.JSON(http.StatusOK, common.Success(http.StatusOK, "Succes to Get detail menu by id", response))
	}
}

func (dc *DetailMenuController) GetAll() echo.HandlerFunc {
	return func(c echo.Context) error {

		res, err := dc.repo.GetAllMenu()

		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.InternalServerError(http.StatusInternalServerError, "There is some error on server", nil))
		}

		response := []DetailMenuGetResponse{}
		for _, result := range res {
			response = append(response, DetailMenuGetResponse{
				Detail_menu_uid: result.Detail_menu_uid,
				Menu_uid:        result.Menu_uid,
				Food_uid:        result.Food_uid,
			})
		}

		return c.JSON(http.StatusOK, common.Success(http.StatusOK, "Succes to Get detail menu by id", response))
	}
}

func (dc *DetailMenuController) Update() echo.HandlerFunc {
	return func(c echo.Context) error {
		detail_menu_uid := c.Param("detail_menu_uid")
		var updateDetailMenu = DetailMenuCreateRequestFormat{}

		c.Bind(&updateDetailMenu)

		errB := c.Validate(&updateDetailMenu)
		if errB != nil {
			return c.JSON(http.StatusBadRequest, common.BadRequest(http.StatusBadRequest, "There is some problem from input", nil))
		}
		res, err := dc.repo.Update(detail_menu_uid, entities.Detail_menu{
			Menu_uid: updateDetailMenu.Menu_uid,
			Food_uid: updateDetailMenu.Food_uid,
		})

		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.InternalServerError(http.StatusInternalServerError, "There is some error on server", nil))
		}

		response := DetailMenuCreateResponse{}
		response.Detail_menu_uid = res.Detail_menu_uid
		response.Menu_uid = res.Menu_uid
		response.Food_uid = res.Food_uid

		return c.JSON(http.StatusOK, common.Success(http.StatusOK, "Succes to Get detail menu by id", response))
	}
}

func (dc *DetailMenuController) Delete() echo.HandlerFunc {
	return func(c echo.Context) error {

		detail_menu_uid := c.Param("detail_menu_uid")
		err := dc.repo.Delete(detail_menu_uid)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.InternalServerError(http.StatusInternalServerError, "There is some error on server", nil))
		}

		return c.JSON(http.StatusOK, common.Success(http.StatusOK, "success delete foods", nil))
	}
}
