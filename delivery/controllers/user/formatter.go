package user

import (
	"HealthFit/entities"
	// "gorm.io/gorm"
)

// =================== Create User =======================
type CreateUserRequestFormat struct {
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
	Gender   string `json:"gender" form:"gender"`
}

type CreateUserResponseFormat struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Gender string `json:"gender"`
}

func ToCreateUserResponseFormat(UserResponse entities.User) CreateUserResponseFormat {
	return CreateUserResponseFormat{
		ID:     int(UserResponse.ID),
		Name:   UserResponse.Name,
		Email:  UserResponse.Email,
		Gender: UserResponse.Gender,
	}
}

// =================== Update User =======================
type UpdateUserRequestFormat struct {
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
	Gender   string `json:"gender" form:"gender"`
}

// func (UURF UpdateUserRequestFormat) ToUpdateUserRequestFormat(ID uint) entities.User {
// 	return entities.User{
// 		Model:    gorm.Model{ID: ID},
// 		Name:     UURF.Name,
// 		Email:    UURF.Email,
// 		Password: UURF.Password,
// 	}
// }

type UpdateUserResponseFormat struct {
	Name   string `json:"name"`
	Email  string `json:"email"`
	Gender string `json:"gender" form:"gender"`
}

func ToUpdateUserResponseFormat(UserResponse entities.User) UpdateUserResponseFormat {
	return UpdateUserResponseFormat{
		Name:   UserResponse.Name,
		Email:  UserResponse.Email,
		Gender: UserResponse.Gender,
	}
}

type GetUserByIdResponseFormat struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Gender string `json:"gender" form:"gender"`
}

func ToGetUserByIdResponseFormat(UserResponse entities.User) GetUserByIdResponseFormat {
	return GetUserByIdResponseFormat{
		ID:     int(UserResponse.ID),
		Name:   UserResponse.Name,
		Email:  UserResponse.Email,
		Gender: UserResponse.Gender,
	}
}

type InsertUserResponseFormat struct {
	Code    int           `json:"code"`
	Message string        `json:"message"`
	Data    entities.User `json:"data"`
}

type GetUsersResponseFormat struct {
	Code    int             `json:"code"`
	Message string          `json:"message"`
	Data    []entities.User `json:"data"`
}

type GetAllUserResponseFormat struct {
	Code    int           `json:"code"`
	Message string        `json:"message"`
	Data    entities.User `json:"data"`
}

type UpdateResponseFormat struct {
	Code    int           `json:"code"`
	Message string        `json:"message"`
	Data    entities.User `json:"data"`
}

type DeleteUserResponseFormat struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
