package entities

import (
	"gorm.io/gorm"
)

type Admin struct {
	gorm.Model
	Admin_uid string `gorm:"index;unique;type:varchar(22)"`
	Name      string `gorm:"type:varchar(100)"`
	Email     string `gorm:"unique"`
	Password  string
	Gender    string `gorm:"type:enum('Pria','Wanita','None');default:'None'"`
}
