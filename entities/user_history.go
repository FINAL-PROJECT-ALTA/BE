package entities

import "gorm.io/gorm"

type User_history struct {
	gorm.Model
	User_history_uid string
	User_uid         string `gorm:"index;type:varchar(22)"`
	Goal_uid         string `gorm:"index;type:varchar(22)"`
	Menu_uid         string `gorm:"index;type:varchar(22)"`
	Menu             []Menu `gorm:"-"`
}
