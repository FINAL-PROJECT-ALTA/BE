package goal

import (
	"HealthFit/delivery/controllers/common"
	"HealthFit/delivery/middlewares"
	"HealthFit/entities"
	"HealthFit/repository/goal"
	"net/http"

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
		user_uid := middlewares.ExtractTokenUserUid(c)
		goal := CreateGoalRequest{}

		c.Bind(&goal)
		err := c.Validate(&goal)

		if err != nil {
			return c.JSON(http.StatusBadRequest, common.BadRequest(http.StatusBadRequest, "There is some problem from input", nil))
		}

		res, err := ac.repo.Create(entities.Goal{
			User_uid:      user_uid,
			Height:        goal.Height,
			Weight:        goal.Weight,
			Age:           goal.Age,
			Daily_active:  goal.Daily_active,
			Weight_target: goal.Weight_target,
			Range_time:    goal.Range_time,
			Target:        goal.Target,
		})

		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.InternalServerError(http.StatusInternalServerError, "There is some error on server", nil))
		}

		response := GoalResponse{}
		response.Goal_uid = res.Goal_uid
		response.Weight = res.Weight
		response.Age = res.Age
		response.Daily_active = res.Daily_active
		response.Weight_target = res.Weight_target
		response.Range_time = res.Range_time
		response.Target = res.Target

		return c.JSON(http.StatusCreated, common.Success(http.StatusCreated, "Success create goal", response))

	}
}
func (ac *GoalController) GetById() echo.HandlerFunc {
	return func(c echo.Context) error {
		goal_uid := c.Param("goal_uid")

		res, err := ac.repo.GetById(goal_uid)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.InternalServerError(http.StatusInternalServerError, "There is some error on server", nil))
		}

		response := GetByIdGoalResponse{}
		response.Goal_uid = res.Goal_uid
		response.Weight = res.Weight
		response.Age = res.Age
		response.Daily_active = res.Daily_active
		response.Weight_target = res.Weight_target
		response.Range_time = res.Range_time
		response.Target = res.Target

		return c.JSON(http.StatusOK, common.Success(http.StatusOK, "Success get goal", response))
	}
}

func (ac *GoalController) Update() echo.HandlerFunc {
	return func(c echo.Context) error {
		user_uid := middlewares.ExtractTokenUserUid(c)
		goal_uid := c.Param("goal_uid")
		var newGoal = UpdateGoalRequest{}
		c.Bind(&newGoal)

		err := c.Validate(&newGoal)
		if err != nil {
			return c.JSON(http.StatusBadRequest, common.BadRequest(http.StatusBadRequest, "There is some problem from input", nil))
		}

		res, err := ac.repo.Update(goal_uid, entities.Goal{User_uid: user_uid, Height: newGoal.Height, Weight: newGoal.Weight, Age: newGoal.Age, Range_time: newGoal.Range_time, Target: newGoal.Target})

		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.InternalServerError(http.StatusInternalServerError, "There is some error on server", nil))
		}

		response := GoalResponse{}
		response.Goal_uid = res.Goal_uid
		response.Weight = res.Weight
		response.Age = res.Age
		response.Daily_active = res.Daily_active
		response.Weight_target = res.Weight_target
		response.Range_time = res.Range_time
		response.Target = res.Target

		return c.JSON(http.StatusOK, common.Success(http.StatusOK, "Success update goal", response))
	}
}

func (ac *GoalController) Delete() echo.HandlerFunc {
	return func(c echo.Context) error {
		goal_uid := middlewares.ExtractTokenUserUid(c)

		err := ac.repo.Delete(goal_uid)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.InternalServerError(http.StatusInternalServerError, "There is some error on server", nil))
		}

		return c.JSON(http.StatusOK, common.Success(http.StatusOK, "Success delete goal", nil))
	}
}
