package user

// "gorm.io/gorm"

type UserCreateResponse struct {
	User_uid string `json:"user_uid"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Gender   string `json:"gender" form:"gender"`
}
type UserUpdateResponse struct {
	User_uid string `json:"user_uid"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Gender   string `json:"gender" form:"gender"`
}
type UserGetByIdResponse struct {
	User_uid string `json:"user_uid"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Gender   string `json:"gender" form:"gender"`
}

type UserGoal struct {
	Height     int `json:"height"`
	Weight     int `json:"weight"`
	Age        int `json:"age"`
	Range_time int `json:"range"`
}
type UserHistoryResponse struct {
	User_uid string `json:"user_uid"`
	Menu_uid string `json:"menu_uid"`
}

type UserCompleksResponse struct {
	User_uid string                `json:"user_uid"`
	Name     string                `json:"name"`
	Email    string                `json:"email"`
	Gender   string                `json:"gender" form:"gender"`
	Goal     []UserGoal            `json:"goal"`
	History  []UserHistoryResponse `json:"history"`
}

//=========================================================

// =================== Create User Request =======================
type CreateUserRequestFormat struct {
	Name     string `json:"name" form:"name" validate:"required"`
	User_uid string
	Email    string `json:"email" form:"email" validate:"required,email"`
	Password string `json:"password" form:"password" validate:"required"`
	Gender   string `json:"gender" form:"gender"`
}

// =================== Update User Request =======================
type UpdateUserRequestFormat struct {
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email" validate:"required,email"`
	Password string `json:"password" form:"password"`
	Gender   string `json:"gender" form:"gender"`
}
