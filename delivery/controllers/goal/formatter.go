package goal

//=============== Request Format ===================//
type CreateGoalRequest struct {
	Height     int    `json:"height" form:"height" validate:"required"`
	Weight     int    `json:"weight" form:"weight" validate:"required"`
	Age        int    `json:"age" form:"age" validate:"required"`
	Range_time int    `json:"range" form:"range" validate:"required"`
	Target     string `json:"target" form:"target" validate:"required"`
}

type UpdateGoalRequest struct {
	Height     int    `json:"height" form:"height" validate:"required"`
	Weight     int    `json:"weight" form:"weight" validate:"required"`
	Age        int    `json:"age" form:"age" validate:"required"`
	Range_time int    `json:"range" form:"range" validate:"required"`
	Target     string `json:"target" form:"target" validate:"required"`
}

//=============== Response Format ==================//

type GoalResponse struct {
	Goal_uid   string `json:"goal_uid"`
	Height     int    `json:"height"`
	Weight     int    `json:"weight"`
	Age        int    `json:"age"`
	Range_time int    `json:"range"`
	Target     string `json:"target"`
}

type GetByIdGoalResponse struct {
	Goal_uid   string `json:"goal_uid"`
	Height     int    `json:"height"`
	Weight     int    `json:"weight"`
	Age        int    `json:"age"`
	Range_time int    `json:"range"`
	Target     string `json:"target"`
}
