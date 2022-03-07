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
	User_uid string `json:"user_uid"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Token    string `json:"token"`
}
type AdminLoginResponse struct {
	Admin_uid string `json:"admin_uid"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Token     string `json:"token"`
}
