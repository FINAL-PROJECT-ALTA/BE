package authgoogle

type Response struct {
	Id             string `json:"id"`
	Email          string `json:"email"`
	Verified_email string `json:"verified_email"`
	Name           string `json:"name"`
	Given_name     string `json:"given_name"`
	Family_name    string `json:"family_name"`
	Link           string `json:"link"`
	Picture        string `json:"picture"`
}
