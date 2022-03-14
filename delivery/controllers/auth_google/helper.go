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
