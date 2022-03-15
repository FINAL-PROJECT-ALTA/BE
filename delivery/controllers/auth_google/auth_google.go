package authgoogle

import (
	"HealthFit/configs"
	"HealthFit/delivery/controllers/common"
	"context"
	"fmt"
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

		var url = gl.config.AuthCodeURL("randomstate", oauth2.AccessTypeOffline)
		res := c.Redirect(http.StatusSeeOther, url)
		if res != nil {
			return c.JSON(http.StatusBadRequest, common.BadRequest(http.StatusBadRequest, "failed to redirect google oauth", res.Error()))
		}
		return c.JSON(http.StatusOK, common.Success(http.StatusOK, "succes redirect to google oauth", nil))
	}
}

// func (gl *GoogleController) LoginGoogle(w http.ResponseWriter, r *http.Request) {
// 	oauthState := generateStateOauthCookie(w)
// 	u := gl.config.AuthCodeURL(oauthState)
// 	http.Redirect(w, r, u, http.StatusTemporaryRedirect)
// }

// func (gl *GoogleController) Callback(w http.ResponseWriter, r *http.Request) {

// 	oauthState, _ := r.Cookie("oauthstate")

// 	if r.FormValue("state") != oauthState.Value {
// 		log.Println("invalid oauth google state")
// 		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
// 		return
// 	}

// 	data, err := getUserDataFromGoogle(r.FormValue("code"))
// 	if err != nil {
// 		log.Println(err.Error())
// 		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
// 		return
// 	}

// 	fmt.Fprintf(w, "UserInfo: %s\n", data)
// }

func (gl *GoogleController) Callback() echo.HandlerFunc {
	return func(c echo.Context) error {
		if c.FormValue("state") != "randomstate" {
			return c.JSON(http.StatusBadRequest, common.BadRequest(http.StatusBadRequest, "state is not valid", nil))
		}
		// token, err := TokenFromFile("./utils/google/secret.json", gl.config)
		response, err := gl.getUserDataFromGoogle(c.FormValue("code"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, common.BadRequest(http.StatusBadRequest, "error to get token %s\n", err))
		}

		return c.JSON(http.StatusOK, common.Success(http.StatusOK, "success login", response))
	}
}

func (gl *GoogleController) getUserDataFromGoogle(code string) ([]byte, error) {
	// Use code to get token and get user info from Google.

	token, err := gl.config.Exchange(context.Background(), code)
	if err != nil {
		return nil, fmt.Errorf("code exchange wrong: %s", err.Error())
	}
	response, err := http.Get(configs.OauthGoogleUrlAPI + token.AccessToken)
	if err != nil {
		return nil, fmt.Errorf("failed getting user info: %s", err.Error())
	}

	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("failed read response: %s", err.Error())
	}
	return contents, nil
}
