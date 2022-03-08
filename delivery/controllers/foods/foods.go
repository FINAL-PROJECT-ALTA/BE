package foods

import (
	"HealthFit/delivery/controllers/common"
	"HealthFit/delivery/middlewares"
	"HealthFit/entities"
	food "HealthFit/repository/foods"
	"net/http"

	"github.com/labstack/echo/v4"
)

type FoodsController struct {
	repo food.Food
}

func New(repository food.Food) *FoodsController {
	return &FoodsController{
		repo: repository,
	}
}

func (fc *FoodsController) Create() echo.HandlerFunc {
	return func(c echo.Context) error {

		isAdmin, errA := middlewares.ExtractTokenAdminUid(c)
		if errA != nil {
			return c.JSON(http.StatusBadRequest, common.BadRequest(http.StatusBadRequest, "There is some problem from input", nil))
		}
		newFoods := FoodsCreateRequestFormat{}

		c.Bind(&newFoods)
		errB := c.Validate(&newFoods)
		if errB != nil {
			return c.JSON(http.StatusBadRequest, common.BadRequest(http.StatusBadRequest, "There is some problem from input", nil))
		}

		res, err := fc.repo.Create(entities.Food{
			Admin_uid:     isAdmin,
			Name:          newFoods.Name,
			Calories:      newFoods.Calories,
			Energy:        newFoods.Energy,
			Carbohidrate:  newFoods.Carbohidrate,
			Protein:       newFoods.Protein,
			Food_category: newFoods.Food_category,
		})

		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.InternalServerError(http.StatusInternalServerError, "There is some error on server", nil))
		}

		response := FoodsCreateResponse{}
		response.Food_uid = res.Food_uid
		response.Name = res.Name
		response.Calories = res.Calories
		response.Energy = res.Energy
		response.Carbohidrate = res.Carbohidrate
		response.Protein = res.Protein
		response.Food_category = res.Food_category

		return c.JSON(http.StatusCreated, common.Success(http.StatusCreated, "Success create foods", response))

	}
}

func (fc *FoodsController) Search() echo.HandlerFunc {
	return func(c echo.Context) error {
		input := c.QueryParam("input")
		category := c.QueryParam("category")

		res, err := fc.repo.Search(input, category)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.InternalServerError(http.StatusInternalServerError, "There is some error on server", nil))
		}

		response := []FoodsSearchResponse{}
		for i := 0; i < len(res); i++ {
			resObj := FoodsSearchResponse{}
			resObj.Food_uid = res[i].Food_uid
			resObj.Name = res[i].Name
			resObj.Calories = res[i].Calories
			resObj.Energy = res[i].Energy
			resObj.Carbohidrate = res[i].Carbohidrate
			resObj.Protein = res[i].Protein
			resObj.Food_category = res[i].Food_category

			resObj.Images = res[i].Image

			response = append(response, resObj)

		}

		return c.JSON(http.StatusOK, common.Success(http.StatusOK, "Success get foods", response))
	}
}

func (fc *FoodsController) Update() echo.HandlerFunc {
	return func(c echo.Context) error {

		isAdmin, errA := middlewares.ExtractTokenAdminUid(c) // jangan lupa ganti extract token admin
		if errA != nil {
			return c.JSON(http.StatusUnauthorized, common.BadRequest(http.StatusUnauthorized, "access denied", nil))
		}

		food_uid := c.Param("food_uid")
		var updateFoods = FoodsUpdateRequestFormat{}
		c.Bind(&updateFoods)

		errB := c.Validate(&updateFoods)
		if errB != nil {
			return c.JSON(http.StatusBadRequest, common.BadRequest(http.StatusBadRequest, "There is some problem from input", nil))
		}

		res, err := fc.repo.Update(food_uid, entities.Food{
			Admin_uid:     isAdmin,
			Name:          updateFoods.Name,
			Calories:      updateFoods.Calories,
			Energy:        updateFoods.Energy,
			Carbohidrate:  updateFoods.Carbohidrate,
			Protein:       updateFoods.Protein,
			Food_category: updateFoods.Food_category,
		})

		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.InternalServerError(http.StatusInternalServerError, "There is some error on server", nil))
		}

		response := FoodsCreateResponse{}
		response.Food_uid = res.Food_uid
		response.Name = res.Name
		response.Calories = res.Calories
		response.Energy = res.Energy
		response.Carbohidrate = res.Carbohidrate
		response.Protein = res.Protein
		response.Food_category = res.Food_category

		return c.JSON(http.StatusCreated, common.Success(http.StatusCreated, "Success update foods", response))
	}
}

func (fc *FoodsController) Delete() echo.HandlerFunc {
	return func(c echo.Context) error {
		_, errA := middlewares.ExtractTokenAdminUid(c)
		if errA != nil {
			return c.JSON(http.StatusUnauthorized, common.BadRequest(http.StatusUnauthorized, "access denied", nil))
		}

		food_uid := c.Param("food_uid")
		err := fc.repo.Delete(food_uid)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.InternalServerError(http.StatusInternalServerError, "There is some error on server", nil))
		}

		return c.JSON(http.StatusOK, common.Success(http.StatusOK, "success delete foods", err))
	}
}

func (fc *FoodsController) GetAll() echo.HandlerFunc {
	return func(c echo.Context) error {
		res, err := fc.repo.GetAll()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.InternalServerError(http.StatusInternalServerError, "There is some error on server", nil))
		}

		response := []FoodsGetAllResponse{}
		for _, result := range res {
			response = append(response, FoodsGetAllResponse{
				Food_uid:      result.Food_uid,
				Name:          result.Name,
				Calories:      result.Calories,
				Energy:        result.Energy,
				Carbohidrate:  result.Carbohidrate,
				Protein:       result.Protein,
				Food_category: result.Food_category,
				// Images: result.Images,
			})
		}

		return c.JSON(http.StatusOK, common.Success(http.StatusOK, "Success get all foods", response))
	}
}
