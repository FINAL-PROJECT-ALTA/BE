package goal

import "time"

//=============== Request Format ===================//
type CreateGoalRequest struct {
	Height        int    `json:"height" form:"height" validate:"required"`
	Weight        int    `json:"weight" form:"weight" validate:"required"`
	Age           int    `json:"age" form:"age" validate:"required"`
	Daily_active  string `json:"daily_active" form:"daily_active" validate:"required,min=2,max=20,excludesall=!@#?^#*()_+-=0123456789%&"`
	Weight_target int    `json:"weight_target" form:"weight_target" validate:"required"`
	Range_time    int    `json:"range_time" form:"range_time" validate:"required"`
	Target        string `json:"target" form:"target" validate:"required,min=2,max=20,excludesall=!@#?^#*()_+-=0123456789%&"`
}

type UpdateGoalRequest struct {
	Height        int    `json:"height" form:"height" validate:"omitempty,required"`
	Weight        int    `json:"weight" form:"weight" validate:"omitempty,required"`
	Age           int    `json:"age" form:"age" validate:"omitempty,required"`
	Daily_active  string `json:"daily_active" form:"daily_active" validate:"omitempty,required,min=2,max=20,excludesall=!@#?^#*()_+-=0123456789%&"`
	Weight_target int    `json:"weight_target" form:"weight_target" validate:"omitempty,required"`
	Range_time    int    `json:"range_time" form:"range_time" validate:"omitempty,required"`
	Target        string `json:"target" form:"target" validate:"omitempty,required,min=2,max=20,excludesall=!@#?^#*()_+-=0123456789%&"`
}

//=============== Response Format ==================//

type GoalResponse struct {
	Goal_uid      string `json:"goal_uid"`
	Height        int    `json:"height"`
	Weight        int    `json:"weight"`
	Age           int    `json:"age"`
	Daily_active  string `json:"daily_active"`
	Weight_target int    `json:"weight_target"`
	Range_time    int    `json:"range_time"`
	Target        string `json:"target"`
}

type GetAllResponse struct {
	Goal_uid      string    `json:"goal_uid"`
	Height        int       `json:"height"`
	Weight        int       `json:"weight"`
	Age           int       `json:"age"`
	Daily_active  string    `json:"daily_active"`
	Weight_target int       `json:"weight_target"`
	Range_time    int       `json:"range_time"`
	Target        string    `json:"target"`
	Status        string    `json:"status"`
	CreatedAt     time.Time `json:"createdAt"`
	Count         int       `json:"count"`
}

type GetByIdGoalResponse struct {
	Goal_uid      string    `json:"goal_uid"`
	Height        int       `json:"height"`
	Weight        int       `json:"weight"`
	Age           int       `json:"age"`
	Daily_active  string    `json:"daily_active"`
	Weight_target int       `json:"weight_target"`
	Range_time    int       `json:"range_time"`
	Target        string    `json:"target"`
	Status        string    `json:"status"`
	CreatedAt     time.Time `json:"createdAt"`
	Count         int       `json:"count"`
}

type CreateResponseErrorGoal struct {
	Bmr                    int `json:"bmr"`
	Cut_calories_every_day int `json:"cut_calories_every_day"`
}
