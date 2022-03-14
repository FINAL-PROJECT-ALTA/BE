package auth

type LoginReqFormat struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type LoginRespFormat struct {
	Code    int                    `json:"code"`
	Message string                 `json:"message"`
	Data    map[string]interface{} `json:"data"`
}

type UserLoginResponse struct {
	User_uid      string `json:"user_uid"`
	Name          string `json:"name"`
	Email         string `json:"email"`
	Roles         bool   `json:"roles"`
	Token         string `json:"token"`
	Goal_active   bool   `json:"goal_active"`
	Goal_exspired bool   `json:"goal_exspired"`
}
type AdminLoginResponse struct {
	User_uid string `json:"user_uid"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Roles    bool   `json:"roles"`
	Token    string `json:"token"`
}
