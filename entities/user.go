package entities

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `gorm:"type:varchar(100)" json:"name"`
	Email    string `gorm:"unique" json:"email"`
	Password string `json:"-"`
	Gender   string `gorm:"type:enum('Pria','Wanita','None');default:'None'"`
}
