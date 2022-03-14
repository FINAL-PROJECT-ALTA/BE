package userhistory

import (
	"HealthFit/entities"
	"time"
)

type CreateUserHistoryRequestFormat struct {
	User_uid string
	Menu_uid string `json:"menu_uid" form:"menu_uid" validate:"required"`
	Goal_uid string `json:"goal_uid" form:"goal_uid" validate:"required"`
}

type UpdateUserHistoryRequestFormat struct {
	User_uid string
	Menu_uid string `json:"menu_uid" form:"menu_uid"`
	Goal_uid string `json:"goal_uid" form:"goal_uid"`
}

type CreateUserHistoryResponse struct {
	User_history_uid string `json:"user_history_uid"`
	User_uid         string `json:"-"`
	Goal_uid         string `json:"goal_uid"`
	Menu_uid         string `json:"menu_uid"`
}

type Menu struct {
	Menu_uid      string                 `json:"menu_uid"`
	Menu_category string                 `json:"menu_category"`
	Created_by    string                 `json:"created_by"`
	Detail_menu   []entities.Detail_menu `json:"foods"`
}

type GetAllUserHistoryResponse struct {
	User_history_uid string          `json:"user_history_uid"`
	Goal_uid         string          `json:"goal_uid"`
	CreatedAt        time.Time       `json:"created_at"`
	Menu             []entities.Menu `json:"menu"`
}

type GetUserHistoryResponse struct {
	User_history_uid string          `json:"user_history_uid"`
	Goal_uid         string          `json:"goal_uid"`
	Menu             []entities.Menu `json:"menu"`
}

type UpdateUserHistoryResponse struct {
	Menu_uid string `json:"menu_uid"`
	Goal_uid string `json:"goal_uid"`
}
