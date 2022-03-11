package user

import (
	"HealthFit/delivery/controllers/common"
	"HealthFit/delivery/middlewares"
	"HealthFit/entities"
	"HealthFit/repository/user"
	utils "HealthFit/utils/aws_S3"
	"net/http"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

type UserController struct {
	repo user.User
	conn *session.Session
}

func New(repository user.User, S3 *session.Session) *UserController {
	return &UserController{
		repo: repository,
		conn: S3,
	}
}

func (ac *UserController) Register() echo.HandlerFunc {
	return func(c echo.Context) error {
		user := CreateUserRequestFormat{}

		c.Bind(&user)
		err := c.Validate(&user)

		if err != nil {
			return c.JSON(http.StatusBadRequest, common.BadRequest(http.StatusBadRequest, "There is some problem from input", nil))
		}

		file, errO := c.FormFile("image")
		if errO != nil {
			log.Info(errO)
		}
		src, _ := file.Open()
		link, errU := utils.Upload(ac.conn, src, *file)
		if errU != nil {
			return c.JSON(http.StatusBadRequest, common.BadRequest(http.StatusBadRequest, "Upload Failed", nil))
		}

		res, err_repo := ac.repo.Register(entities.User{
			Name:     user.Name,
			Email:    user.Email,
			Password: user.Password,
			Gender:   user.Gender,
			Image:    link,
		})

		if err_repo != nil {
			return c.JSON(http.StatusConflict, common.InternalServerError(http.StatusConflict, err_repo.Error(), nil))
		}

		response := UserCreateResponse{}
		response.User_uid = res.User_uid
		response.Name = res.Name
		response.Email = res.Email
		response.Gender = res.Gender
		response.Roles = res.Roles
		response.Image = res.Image

		return c.JSON(http.StatusCreated, common.Success(http.StatusCreated, "Success Create User", response))

	}
}
func (ac *UserController) GetById() echo.HandlerFunc {
	return func(c echo.Context) error {
		user_uid := middlewares.ExtractTokenUserUid(c)

		res, err := ac.repo.GetById(user_uid)

		if err != nil {
			return c.JSON(http.StatusNotFound, common.InternalServerError(http.StatusNotFound, err.Error(), nil))
		}

		// response := UserCompleksResponse{}

		// response.User_uid = res.User_uid
		// response.Name = res.Name
		// response.Email = res.Email
		// response.Gender = res.Gender
		// response.Roles = res.Roles

		// responseGoal := []UserGoal{}
		// for i := 0; i < len(res.Goal); i++ {
		// 	user_goal := UserGoal{}
		// 	user_goal.Height = res.Goal[i].Height
		// 	user_goal.Weight = res.Goal[i].Weight
		// 	user_goal.Age = res.Goal[i].Age
		// 	user_goal.Daily_active = res.Goal[i].Daily_active
		// 	user_goal.Weight_target = res.Goal[i].Weight_target
		// 	user_goal.Range_time = res.Goal[i].Range_time
		// 	responseGoal = append(responseGoal, user_goal)
		// }
		// log.Info(responseGoal)
		// response.Goal = responseGoal
		// log.Info(response.Goal)

		// responseHistory := []UserHistoryResponse{}
		// for i := 0; i < len(res.Goal); i++ {
		// 	user_history := UserHistoryResponse{}
		// 	user_history.User_uid = res.History[i].User_uid
		// 	user_history.Menu_uid = res.History[i].Menu_uid

		// 	responseHistory = append(responseHistory, user_history)
		// }

		// response.History = responseHistory

		// log.Info(response)

		return c.JSON(http.StatusOK, common.Success(http.StatusOK, "Success get user", res))
	}
}

func (ac *UserController) Update() echo.HandlerFunc {
	return func(c echo.Context) error {
		user_uid := middlewares.ExtractTokenUserUid(c)
		var newUser = UpdateUserRequestFormat{}
		c.Bind(&newUser)

		err := c.Validate(&newUser)
		if err != nil {
			return c.JSON(http.StatusBadRequest, common.BadRequest(http.StatusBadRequest, "There is some problem from input", nil))
		}

		res, err_repo := ac.repo.Update(user_uid, entities.User{Name: newUser.Name, Email: newUser.Email, Password: newUser.Password, Gender: newUser.Gender})

		if err_repo != nil {
			return c.JSON(http.StatusInternalServerError, common.InternalServerError(http.StatusInternalServerError, "There is some error on server", nil))
		}

		response := UserUpdateResponse{}
		response.User_uid = res.User_uid
		response.Name = res.Name
		response.Email = res.Email
		response.Gender = res.Gender
		response.Roles = res.Roles

		return c.JSON(http.StatusOK, common.Success(http.StatusOK, "Success Update User", response))
	}
}

func (ac *UserController) Delete() echo.HandlerFunc {
	return func(c echo.Context) error {
		user_uid := middlewares.ExtractTokenUserUid(c)

		err := ac.repo.Delete(user_uid)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.InternalServerError(http.StatusInternalServerError, "There is some error on server", nil))
		}

		return c.JSON(http.StatusOK, common.Success(http.StatusOK, "Success Delete User", nil))
	}
}
