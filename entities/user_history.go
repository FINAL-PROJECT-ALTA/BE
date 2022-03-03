package entities

import "gorm.io/gorm"

type User_history struct {
	gorm.Model
	User_uid string
	Menu_uid string
}
