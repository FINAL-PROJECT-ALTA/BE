package admin

// "gorm.io/gorm"

type AdminResponse struct {
	Admin_uid string `json:"admin_uid"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Gender    string `json:"gender"`
}

type AdminGetByIdResponse struct {
	Admin_uid string `json:"admin_uid"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Gender    string `json:"gender"`
}

//=========================================================

// =================== Create User Request =======================
type CreateAdminRequestFormat struct {
	Name      string `json:"name" form:"name" validate:"required"`
	Admin_uid string
	Email     string `json:"email" form:"email" validate:"required,email"`
	Password  string `json:"password" form:"password" validate:"required"`
	Gender    string `json:"gender" form:"gender" validate:"required"`
}

// =================== Update User Request =======================
type UpdateAdminRequestFormat struct {
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email" validate:"required,email"`
	Password string `json:"password" form:"password"`
	Gender   string `json:"gender" form:"gender"`
}
