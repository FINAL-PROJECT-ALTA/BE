package goal

import (
	"HealthFit/delivery/controllers/common"
	"HealthFit/delivery/middlewares"
	"HealthFit/entities"
	"HealthFit/repository/goal"
	"math"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

type GoalController struct {
	repo goal.Goal
}

func New(repository goal.Goal) *GoalController {
	return &GoalController{
		repo: repository,
	}
}

func (ac *GoalController) Create() echo.HandlerFunc {
	return func(c echo.Context) error {
		goal := CreateGoalRequest{}
		isAdmin := middlewares.ExtractRoles(c)
		if isAdmin {
			return c.JSON(http.StatusBadRequest, common.BadRequest(http.StatusBadRequest, "access denied ", nil))
		}
		user_uid := middlewares.ExtractTokenUserUid(c)

		c.Bind(&goal)
		err := c.Validate(&goal)

		if err != nil {
			return c.JSON(http.StatusBadRequest, common.BadRequest(http.StatusBadRequest, "There is some problem from input", nil))
		}

		res, errRepo := ac.repo.Create(entities.Goal{
			User_uid:      user_uid,
			Height:        goal.Height,
			Weight:        goal.Weight,
			Age:           goal.Age,
			Daily_active:  goal.Daily_active,
			Weight_target: goal.Weight_target,
			Range_time:    goal.Range_time,
			Target:        goal.Target,
			Status:        "active",
		})

		if errRepo != nil {
			if errRepo.Error() == "impossible" {
				resErr := CreateResponseErrorGoal{Bmr: res.Weight, Cut_calories_every_day: res.Height}
				return c.JSON(http.StatusInternalServerError, common.InternalServerError(http.StatusInternalServerError, errRepo.Error(), resErr))
			}
			return c.JSON(http.StatusBadRequest, common.BadRequest(http.StatusBadRequest, "There is some error on server", nil))
		}

		return c.JSON(http.StatusCreated, common.Success(http.StatusCreated, "Success create goal", res))

	}
}

func (ac *GoalController) GetAll() echo.HandlerFunc {
	return func(c echo.Context) error {

		isAdmin := middlewares.ExtractRoles(c)

		if isAdmin {
			return c.JSON(http.StatusBadRequest, common.BadRequest(http.StatusBadRequest, "access denied ", nil))

		}

		user_uid := middlewares.ExtractTokenUserUid(c)

		res, err := ac.repo.GetAll(user_uid)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.InternalServerError(http.StatusInternalServerError, "There is some error on server", nil))
		}

		response := []GetAllResponse{}
		for _, result := range res {
			var count int
			time := time.Now()
			different := result.CreatedAt.Sub(time)
			days := math.Abs(float64(int(different.Hours() / 24)))
			diff := result.Range_time - int(days)
			if diff <= 0 {
				count = 0
			} else {
				count = diff
			}

			response = append(response, GetAllResponse{
				Goal_uid:      result.Goal_uid,
				Height:        result.Height,
				Weight:        result.Weight,
				Age:           result.Age,
				Daily_active:  result.Daily_active,
				Weight_target: result.Weight_target,
				Range_time:    result.Range_time,
				Status:        result.Status,
				Target:        result.Target,
				CreatedAt:     result.CreatedAt,
				Count:         count,
			})
		}

		return c.JSON(http.StatusOK, common.Success(http.StatusOK, "Success get all goal", response))
	}
}

func (ac *GoalController) GetById() echo.HandlerFunc {
	return func(c echo.Context) error {

		isAdmin := middlewares.ExtractRoles(c)

		if isAdmin {
			return c.JSON(http.StatusBadRequest, common.BadRequest(http.StatusBadRequest, "access denied ", nil))

		}
		user_uid := middlewares.ExtractTokenUserUid(c)
		goal_uid := c.Param("goal_uid")

		res, err := ac.repo.GetById(goal_uid, user_uid)

		response := GetByIdGoalResponse{}
		response.Goal_uid = res.Goal_uid
		response.Height = res.Height
		response.Weight = res.Weight
		response.Age = res.Age
		response.Daily_active = res.Daily_active
		response.Weight_target = res.Weight_target
		response.Range_time = res.Range_time
		response.Status = res.Status
		response.Target = res.Target

		time := time.Now()
		different := res.CreatedAt.Sub(time)
		days := math.Abs(float64(int(different.Hours() / 24)))
		diff := res.Range_time - int(days)
		if diff <= 0 {
			response.Count = 0
		} else {
			response.Count = diff
		}

		if err != nil {
			statusCode := 500
			message := "There is some error on server"
			if err.Error() == "not found" {
				statusCode = 404
				message = " Goal is not found"
			}
			return c.JSON(statusCode, common.InternalServerError(statusCode, message, nil))
		}

		return c.JSON(http.StatusOK, common.Success(http.StatusOK, "Success get goal", response))
	}
}

func (ac *GoalController) Update() echo.HandlerFunc {
	return func(c echo.Context) error {

		isAdmin := middlewares.ExtractRoles(c)

		if isAdmin {
			return c.JSON(http.StatusBadRequest, common.BadRequest(http.StatusBadRequest, "access denied ", nil))

		}

		user_uid := middlewares.ExtractTokenUserUid(c)
		goal_uid := c.Param("goal_uid")
		var newGoal = UpdateGoalRequest{}
		c.Bind(&newGoal)

		err := c.Validate(&newGoal)
		if err != nil {
			return c.JSON(http.StatusBadRequest, common.BadRequest(http.StatusBadRequest, "There is some problem from input", nil))
		}

		res, err := ac.repo.Update(goal_uid, entities.Goal{
			User_uid:      user_uid,
			Height:        newGoal.Height,
			Weight:        newGoal.Weight,
			Age:           newGoal.Age,
			Daily_active:  newGoal.Daily_active,
			Weight_target: newGoal.Weight_target,
			Range_time:    newGoal.Range_time,
			Target:        newGoal.Target,
		})

		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.InternalServerError(http.StatusInternalServerError, "There is some error on server", nil))
		}

		return c.JSON(http.StatusOK, common.Success(http.StatusOK, "Success update goal", res))
	}
}

func (ac *GoalController) Delete() echo.HandlerFunc {
	return func(c echo.Context) error {
		isAdmin := middlewares.ExtractRoles(c)

		if isAdmin {
			return c.JSON(http.StatusBadRequest, common.BadRequest(http.StatusBadRequest, "access denied ", nil))

		}
		user_uid := middlewares.ExtractTokenUserUid(c)

		goal_uid := c.Param("goal_uid")

		err := ac.repo.Delete(goal_uid, user_uid)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.InternalServerError(http.StatusInternalServerError, "There is some error on server", nil))
		}

		return c.JSON(http.StatusOK, common.Success(http.StatusOK, "Success delete goal", nil))
	}
}

func (ac *GoalController) CencelGoal() echo.HandlerFunc {
	return func(c echo.Context) error {
		user_uid := middlewares.ExtractTokenUserUid(c)

		isAdmin := middlewares.ExtractRoles(c)

		if isAdmin {
			return c.JSON(http.StatusBadRequest, common.BadRequest(http.StatusBadRequest, "access denied ", nil))
		}

		_, err := ac.repo.CancelGoal(user_uid)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.InternalServerError(http.StatusInternalServerError, "There is some error on server", nil))
		}

		return c.JSON(http.StatusOK, common.Success(http.StatusOK, "Success cancel goal", nil))
	}
}
