package auth

type LoginReqFormat struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginRespFormat struct {
	Code    int                    `json:"code"`
	Message string                 `json:"message"`
	Data    map[string]interface{} `json:"data"`
}
