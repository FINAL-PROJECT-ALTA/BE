package authgoogle

import (
	"context"
	"encoding/json"
	"os"

	"github.com/labstack/gommon/log"
	"golang.org/x/oauth2"
)

func TokenFromFile(file string, conf *oauth2.Config) (*oauth2.Token, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	tok := &oauth2.Token{}
	err = json.NewDecoder(f).Decode(tok)
	if err != nil {
		log.Warn(err)
	}

	tokenSource := conf.TokenSource(context.Background(), tok)
	newToken, err := tokenSource.Token()
	if err != nil {
		log.Warn(err)
	}

	return newToken, err
}

// func generateStateOauthCookie() string {
// 	var expiration = time.Now().Add(20 * time.Minute)

// 	b := make([]byte, 16)
// 	rand.Read(b)
// 	state := base64.URLEncoding.EncodeToString(b)
// 	cookie := http.Cookie{Name: "oauthstate", Value: state, Expires: expiration}
// 	http.SetCookie(w, &cookie)

// 	return state
// }

// func getUserDataFromGoogle(code string) ([]byte, error) {
// 	// Use code to get token and get user info from Google.

// 	var conf = &oauth2.Config{}

// 	token, err := conf.Exchange(context.Background(), code)
// 	if err != nil {
// 		return nil, fmt.Errorf("code exchange wrong: %s", err.Error())
// 	}
// 	response, err := http.Get(configs.OauthGoogleUrlAPI + token.AccessToken)
// 	if err != nil {
// 		return nil, fmt.Errorf("failed getting user info: %s", err.Error())
// 	}
// 	defer response.Body.Close()
// 	contents, err := ioutil.ReadAll(response.Body)
// 	if err != nil {
// 		return nil, fmt.Errorf("failed read response: %s", err.Error())
// 	}
// 	return contents, nil
// }
