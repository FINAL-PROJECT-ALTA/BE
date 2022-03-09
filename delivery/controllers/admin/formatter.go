package admin

// "gorm.io/gorm"

type AdminResponse struct {
	Admin_uid string `json:"admin_uid"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Gender    string `json:"gender"`
	Roles     bool   `json:"roles"`
}

type AdminGetByIdResponse struct {
	Admin_uid string `json:"admin_uid"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Gender    string `json:"gender"`
	Roles     bool   `json:"roles"`
}

//=========================================================

// =================== Create User Request =======================
type CreateAdminRequestFormat struct {
	Name      string `json:"name" form:"name" validate:"required,min=3,max=20,excludesall=!@#?^#*()_+-=0123456789%&"`
	Admin_uid string
	Email     string `json:"email" form:"email" validate:"required,email"`
	Password  string `json:"password" form:"password" validate:"required,min=3,max=15"`
	Gender    string `json:"gender" form:"gender"`
}

// =================== Update User Request =======================
type UpdateAdminRequestFormat struct {
	Name     string `json:"name" form:"name" validate:"min=3,max=20,excludesall=!@#?^#*()_+-=0123456789%&"`
	Email    string `json:"email" form:"email" validate:"omitempty,email"`
	Password string `json:"password" form:"password"`
	Gender   string `json:"gender" form:"gender"`
}
