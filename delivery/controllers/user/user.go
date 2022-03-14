package user

import (
	"HealthFit/delivery/controllers/common"
	"HealthFit/delivery/middlewares"
	"HealthFit/entities"
	"HealthFit/repository/user"
	utils "HealthFit/utils/aws_S3"
	"net/http"
	"strings"

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

		if file != nil {
			src, _ := file.Open()
			link, errU := utils.Upload(ac.conn, src, *file)
			if errU != nil {
				return c.JSON(http.StatusBadRequest, common.BadRequest(http.StatusBadRequest, "Upload Failed", nil))
			}
			user.Image = link
		} else if file == nil {
			user.Image = ""
		}

		res, err_repo := ac.repo.Register(entities.User{
			Name:     user.Name,
			Email:    user.Email,
			Password: user.Password,
			Gender:   user.Gender,
			Image:    user.Image,
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

		resGet, errGet := ac.repo.GetById(user_uid)
		if errGet != nil {
			log.Info(resGet)
		}

		file, errO := c.FormFile("image")
		if errO != nil {
			log.Info(errO)
		} else if errO == nil {
			src, _ := file.Open()
			if resGet.Image != "" {
				var updateImage = resGet.Image
				updateImage = strings.Replace(updateImage, "https://airbnb-app.s3.ap-southeast-1.amazonaws.com/", "", -1)

				var resUp = utils.UpdateUpload(ac.conn, updateImage, src, *file)
				if resUp != "success to update image" {
					return c.JSON(http.StatusInternalServerError, common.InternalServerError(http.StatusInternalServerError, "There is some error on server"+resUp, nil))
				}
			} else if resGet.Image == "" {
				var image, errUp = utils.Upload(ac.conn, src, *file)
				if errUp != nil {
					return c.JSON(http.StatusBadRequest, common.BadRequest(http.StatusBadRequest, "Upload Failed", nil))
				}
				newUser.Image = image
			}
		}

		res, err_repo := ac.repo.Update(user_uid, entities.User{
			Name:     newUser.Name,
			Email:    newUser.Email,
			Password: newUser.Password,
			Gender:   newUser.Gender,
			Image:    newUser.Image,
		})

		if err_repo != nil {
			return c.JSON(http.StatusInternalServerError, common.InternalServerError(http.StatusInternalServerError, "There is some error on server", nil))
		}

		response := UserUpdateResponse{}
		response.User_uid = res.User_uid
		response.Name = res.Name
		response.Email = res.Email
		response.Gender = res.Gender
		response.Roles = res.Roles
		response.Image = res.Image

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
