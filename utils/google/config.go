package google

import (
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

func GoogleConfig(db, clientID, clientSecret string) *oauth2.Config {

	googleOauthConfig := &oauth2.Config{}

	if db == "root" {
		googleOauthConfig = &oauth2.Config{
			RedirectURL:  "http://localhost:8000/auth/google/callback",
			ClientID:     clientID,
			ClientSecret: clientSecret,
			Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email"},
			Endpoint:     google.Endpoint,
		}
	} else if db == "admin" {
		googleOauthConfig = &oauth2.Config{
			RedirectURL:  "https://aaryadewangga.cloud.okteto.net/auth/google/callback",
			ClientID:     clientID,
			ClientSecret: clientSecret,
			Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email"},
			Endpoint:     google.Endpoint,
		}
	}
	return googleOauthConfig

}
