package user

// "gorm.io/gorm"

type UserCreateResponse struct {
	User_uid string `json:"user_uid"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Gender   string `json:"gender" form:"gender"`
	Roles    bool   `json:"roles" form:"roles"`
}
type UserUpdateResponse struct {
	User_uid string `json:"user_uid"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Gender   string `json:"gender" form:"gender"`
	Roles    bool   `json:"roles" form:"roles"`
}
type UserGetByIdResponse struct {
	User_uid string `json:"user_uid"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Gender   string `json:"gender" form:"gender"`
	Roles    bool   `json:"roles" form:"roles"`
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
	Name     string `json:"name" form:"name" validate:"required,min=3,max=20,excludesall=!@#?^#*()_+-=0123456789%&"`
	User_uid string
	Email    string `json:"email" form:"email" validate:"required,email"`
	Password string `json:"password" form:"password" validate:"required,min=3,max=15"`
	Gender   string `json:"gender" form:"gender" validate:"required"`
}

// =================== Update User Request =======================
type UpdateUserRequestFormat struct {
	Name     string `json:"name" form:"name" validate:"omitempty,min=3,max=20,excludesall=!@#?^#*()_+-=0123456789%&"`
	Email    string `json:"email" form:"email" validate:"omitempty,email"`
	Password string `json:"password" form:"password" validate:"omitempty,required,min=3,max=15"`
	Gender   string `json:"gender" form:"gender"`
}
