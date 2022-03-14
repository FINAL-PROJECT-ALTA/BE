package authgoogle

import (
	"HealthFit/configs"
	"HealthFit/delivery/controllers/common"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/labstack/echo/v4"
	"golang.org/x/oauth2"
)

type GoogleController struct {
	config *oauth2.Config
}

func New(config *oauth2.Config) *GoogleController {
	return &GoogleController{
		config: config,
	}
}

func (gl *GoogleController) LoginGoogle() echo.HandlerFunc {
	return func(c echo.Context) error {
		var url = gl.config.AuthCodeURL("random")

		res := c.Redirect(http.StatusTemporaryRedirect, url)
		if res != nil {
			return c.JSON(http.StatusBadRequest, common.BadRequest(http.StatusBadRequest, "state is not valid", nil))
		}
		return c.JSON(http.StatusOK, common.Success(http.StatusOK, "state is not valid", nil))
	}
}

func (gl *GoogleController) Callback() echo.HandlerFunc {
	return func(c echo.Context) error {
		if c.FormValue("state") != "random" {
			return c.JSON(http.StatusBadRequest, common.BadRequest(http.StatusBadRequest, "state is not valid", nil))
		}
		token, err := TokenFromFile("./utils/google/secret.json", gl.config)
		if err != nil {
			return c.JSON(http.StatusBadRequest, common.BadRequest(http.StatusBadRequest, "error to get token %s\n", err))
		}

		response, err := http.Get(configs.OauthGoogleUrlAPI + token.AccessToken)
		if err != nil {
			return c.JSON(http.StatusBadRequest, common.BadRequest(http.StatusBadRequest, "error to parse token", err))
		}

		content, err := ioutil.ReadAll(response.Body)
		if err != nil {
			return c.JSON(http.StatusBadRequest, common.BadRequest(http.StatusBadRequest, "error to parse token", err))
		}

		resLogin := Response{}
		errLogin := json.Unmarshal(content, &resLogin)
		if errLogin != nil {
			return c.JSON(http.StatusBadRequest, common.BadRequest(http.StatusBadRequest, "error to get the struct response", err))
		}
		return c.JSON(http.StatusOK, common.Success(http.StatusOK, "state is not valid", resLogin))
	}
}
