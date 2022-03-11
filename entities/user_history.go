package entities

import "gorm.io/gorm"

type User_history struct {
	gorm.Model
	User_history_uid string `gorm:"index;type:varchar(22)"`
	User_uid         string
	Goal_uid         string
	Menu_uid         string
	Menu             []Menu `gorm:"foreignKey:Menu_uid;references:Menu_uid"`
}
