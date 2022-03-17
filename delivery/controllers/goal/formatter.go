package goal

//=============== Request Format ===================//
type CreateGoalRequest struct {
	Height        int    `json:"height" form:"height" validate:"required"`
	Weight        int    `json:"weight" form:"weight" validate:"required"`
	Age           int    `json:"age" form:"age" validate:"required"`
	Daily_active  string `json:"daily_active" form:"daily_active" validate:"required"`
	Weight_target int    `json:"weight_target" form:"weight_target" validate:"required"`
	Range_time    int    `json:"range_time" form:"range_time" validate:"required"`
	Target        string `json:"target" form:"target" validate:"required"`
}

type UpdateGoalRequest struct {
	Height        int    `json:"height" form:"height" validate:"required"`
	Weight        int    `json:"weight" form:"weight" validate:"required"`
	Age           int    `json:"age" form:"age" validate:"required"`
	Daily_active  string `json:"daily_active" form:"daily_active" validate:"required"`
	Weight_target int    `json:"weight_target" form:"weight_target" validate:"required"`
	Range_time    int    `json:"range_time" form:"range_time" validate:"required"`
	Target        string `json:"target" form:"target" validate:"required"`
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

type GetByIdGoalResponse struct {
	Goal_uid      string `json:"goal_uid"`
	Height        int    `json:"height"`
	Weight        int    `json:"weight"`
	Age           int    `json:"age"`
	Daily_active  string `json:"daily_active"`
	Weight_target int    `json:"weight_target"`
	Range_time    int    `json:"range_time"`
	Target        string `json:"target"`
	Status        string `json:"status"`
	Count         int    `json:"count"`
}
type CreateResponseErrorGoal struct {
	Bmr                    int `json:"bmr"`
	Cut_calories_every_day int `json:"cut_calories_every_day"`
}
