package entities

import "gorm.io/gorm"

type User_history struct {
	gorm.Model
	User_history_uid string
	User_uid         string
	Menu_uid         string
	Goal_uid         string
}
