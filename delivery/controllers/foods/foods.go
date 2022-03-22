package foods

import (
	"HealthFit/delivery/controllers/common"
	"HealthFit/delivery/middlewares"
	"HealthFit/entities"
	food "HealthFit/repository/foods"
	utils "HealthFit/utils/aws_S3"
	edamam "HealthFit/utils/edamam"
	"strings"

	"math"
	"net/http"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

type FoodsController struct {
	repo food.Food
	conn *session.Session
}

func New(repository food.Food, S3 *session.Session) *FoodsController {
	return &FoodsController{
		repo: repository,
		conn: S3,
	}
}

func (fc *FoodsController) Create() echo.HandlerFunc {
	return func(c echo.Context) error {

		isAdmin := middlewares.ExtractRoles(c)
		if !isAdmin {
			return c.JSON(http.StatusBadRequest, common.BadRequest(http.StatusBadRequest, "Access Denied", nil))
		}

		newFoods := FoodsCreateRequestFormat{}
		c.Bind(&newFoods)
		errB := c.Validate(&newFoods)
		if errB != nil {
			return c.JSON(http.StatusBadRequest, common.BadRequest(http.StatusBadRequest, "There is some problem from input", nil))
		}

		file, errO := c.FormFile("image")
		if errO != nil {
			log.Info(errO)
		}

		if file != nil {
			src, _ := file.Open()
			link, errU := utils.Upload(fc.conn, src, *file)
			if errU != nil {
				return c.JSON(http.StatusBadRequest, common.BadRequest(http.StatusBadRequest, "Upload Failed", nil))
			}
			newFoods.Image = link
		} else if file == nil {
			newFoods.Image = "https://raw.githubusercontent.com/FINAL-PROJECT-ALTA/FE/main/image/logo-white.png"
		}

		res, err := fc.repo.Create(entities.Food{
			Name:          newFoods.Name,
			Calories:      newFoods.Calories,
			Energy:        newFoods.Energy,
			Carbohidrate:  newFoods.Carbohidrate,
			Protein:       newFoods.Protein,
			Unit:          newFoods.Unit,
			Unit_value:    newFoods.Unit_value,
			Food_category: newFoods.Food_category,
			Image:         newFoods.Image,
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
		response.Unit = res.Unit
		response.Unit_value = res.Unit_value
		response.Food_category = res.Food_category
		response.Image = res.Image

		return c.JSON(http.StatusCreated, common.Success(http.StatusCreated, "Success create foods", response))

	}
}

func (fc *FoodsController) GetById() echo.HandlerFunc {
	return func(c echo.Context) error {

		food_uid := c.Param("food_uid")
		res, err := fc.repo.GetById(food_uid)

		if err != nil {
			var statusCode int = 500
			var message string = "There is some error on server"
			if err.Error() == "not found" {
				statusCode = 404
				message = "food is not found"
			}
			return c.JSON(statusCode, common.InternalServerError(statusCode, message, nil))
		}

		response := FoodsCreateResponse{}
		response.Food_uid = res.Food_uid
		response.Name = res.Name
		response.Calories = res.Calories
		response.Energy = res.Energy
		response.Carbohidrate = res.Carbohidrate
		response.Protein = res.Protein
		response.Unit = res.Unit
		response.Unit_value = res.Unit_value
		response.Food_category = res.Food_category
		response.Image = res.Image

		return c.JSON(http.StatusCreated, common.Success(http.StatusCreated, "Success get food by food_uid", response))
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
			resObj.Unit = res[i].Unit
			resObj.Unit_value = res[i].Unit_value
			resObj.Food_category = res[i].Food_category

			resObj.Image = res[i].Image

			response = append(response, resObj)

		}

		return c.JSON(http.StatusOK, common.Success(http.StatusOK, "Success get foods", response))
	}
}

func (fc *FoodsController) Update() echo.HandlerFunc {
	return func(c echo.Context) error {

		isAdmin := middlewares.ExtractRoles(c)
		if !isAdmin {
			return c.JSON(http.StatusUnauthorized, common.BadRequest(http.StatusUnauthorized, "access denied", nil))
		}

		food_uid := c.Param("food_uid")
		var updateFoods = FoodsUpdateRequestFormat{}
		c.Bind(&updateFoods)

		errB := c.Validate(&updateFoods)
		if errB != nil {
			return c.JSON(http.StatusBadRequest, common.BadRequest(http.StatusBadRequest, "There is some problem from input", nil))
		}

		resGet, errGet := fc.repo.GetById(food_uid)
		if errGet != nil {
			log.Info(resGet)
		}

		file, errO := c.FormFile("image")
		if errO != nil {
			log.Info(errO)
		} else if errO == nil {
			src, _ := file.Open()
			link := "https://raw.githubusercontent.com/FINAL-PROJECT-ALTA/FE/main/image/logo-white.png"
			if resGet.Image == link {

				var resUp, errUp = utils.Upload(fc.conn, src, *file)
				if errUp != nil {
					return c.JSON(http.StatusBadRequest, common.BadRequest(http.StatusBadRequest, "Upload Failed", nil))
				}
				updateFoods.Image = resUp

			} else if resGet.Image != link {

				var updateImage = resGet.Image
				updateImage = strings.Replace(updateImage, "https://airbnb-app.s3.ap-southeast-1.amazonaws.com/", "", -1)

				var resUp = utils.UpdateUpload(fc.conn, updateImage, src, *file)
				if resUp != "success to update image" {
					return c.JSON(http.StatusInternalServerError, common.InternalServerError(http.StatusInternalServerError, "There is some error on server"+resUp, nil))
				}

			} else if file == nil {
				updateFoods.Image = resGet.Image
			}
		}

		res, err := fc.repo.Update(food_uid, entities.Food{
			Name:          updateFoods.Name,
			Calories:      updateFoods.Calories,
			Energy:        updateFoods.Energy,
			Carbohidrate:  updateFoods.Carbohidrate,
			Protein:       updateFoods.Protein,
			Unit:          updateFoods.Unit,
			Unit_value:    updateFoods.Unit_value,
			Food_category: updateFoods.Food_category,
			Image:         updateFoods.Image,
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
		response.Unit = res.Unit
		response.Unit_value = res.Unit_value
		response.Food_category = res.Food_category
		response.Image = res.Image

		return c.JSON(http.StatusCreated, common.Success(http.StatusCreated, "Success update foods", response))
	}
}

func (fc *FoodsController) Delete() echo.HandlerFunc {
	return func(c echo.Context) error {
		isAdmin := middlewares.ExtractRoles(c)
		if !isAdmin {
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

		category := c.QueryParam("category")

		res, err := fc.repo.GetAll(category)
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
				Unit:          result.Unit,
				Unit_value:    result.Unit_value,
				Food_category: result.Food_category,
				Image:         result.Image,
			})
		}

		return c.JSON(http.StatusOK, common.Success(http.StatusOK, "Success get all foods", response))
	}
}

func (fc *FoodsController) GetFromThirdPary() echo.HandlerFunc {
	return func(c echo.Context) error {
		isAdmin := middlewares.ExtractRoles(c)
		if !isAdmin {
			return c.JSON(http.StatusUnauthorized, common.BadRequest(http.StatusUnauthorized, "access denied", nil))
		}

		s := c.QueryParam("s")
		count := 0

		req := FoodsCreateRequestFormatEdamam{}
		response, err := edamam.FoodThirdParty(s)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.InternalServerError(http.StatusInternalServerError, "There is some error on thirdparty", nil))
		}
		for i := 0; i < len(response.Hints); i++ {
			req.Name = response.Hints[i].Food.Label
			req.Food_uid = strings.Replace(response.Hints[i].Food.FoodId, "_", "", -1)
			req.Unit = response.Hints[i].Measures[1].Label
			req.Unit_value = int(math.Round(response.Hints[i].Measures[1].Weight))
			req.Calories = int(math.Round(response.Hints[i].Food.Nutrients.Enerc_kcal))
			req.Protein = int(math.Round(response.Hints[i].Food.Nutrients.Procnt))
			req.Carbohidrate = int(math.Round(response.Hints[i].Food.Nutrients.Chocdf))
			req.Energy = int(math.Round(response.Hints[i].Food.Nutrients.Enerc_kcal))
			req.Image = response.Hints[i].Food.Image
			req.Food_category = response.Hints[i].Food.CategoryLabel

			// _, errGet := fc.repo.GetFoodThirdParty(req.Food_uid)
			// if errGet == nil {
			_, err := fc.repo.CreateFoodThirdParty(entities.Food{
				Food_uid:      req.Food_uid,
				Name:          req.Name,
				Unit:          req.Unit,
				Unit_value:    req.Unit_value,
				Food_category: req.Food_category,
				Image:         req.Image,
				Calories:      req.Calories,
				Protein:       req.Protein,
				Carbohidrate:  req.Carbohidrate,
				Energy:        req.Energy,
			})
			if err.Error() == "succes to create" {
				count++
			}

			// }

			// time.Sleep(time.Second * 4)

			// }
		}

		return c.JSON(http.StatusOK, common.Success(http.StatusOK, "Success add foods from", count))

	}
}
