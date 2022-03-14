package entities

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model    `json:"-"`
	User_uid      string         `gorm:"index;unique;type:varchar(22)" json:"user_uid"`
	Name          string         `gorm:"type:varchar(100)" json:"name"`
	Email         string         `gorm:"unique" json:"email"`
	Password      string         `json:"-"`
	Gender        string         `gorm:"type:enum('male','female','none');default:'none'" json:"gender"`
	Roles         bool           `gorm:"type:bool" json:"roles"`
	Image         string         `json:"image"`
	Goal          []Goal         `gorm:"foreignKey:User_uid;references:User_uid" json:"goal"`
	History       []User_history `gorm:"foreignKey:User_uid;references:User_uid" json:"history"`
	Goal_active   bool           `gorm:"-" json:"-"`
	Goal_exspired bool           `gorm:"-" json:"-"`
}
