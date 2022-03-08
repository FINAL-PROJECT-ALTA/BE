package entities

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	User_uid string `gorm:"index;unique;type:varchar(22)"`
	Name     string `gorm:"type:varchar(100)"`
	Email    string `gorm:"unique"`
	Password string
	Gender   string         `gorm:"type:enum('Pria','Wanita','None');default:'None'"`
	Roles    bool           `gorm:"type:bool" json:"roles"`
	Goal     []Goal         `gorm:"foreignKey:User_uid;references:User_uid"`
	History  []User_history `gorm:"foreignKey:User_uid;references:User_uid"`
}
