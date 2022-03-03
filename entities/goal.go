package entities

import "gorm.io/gorm"

type Goal struct {
	gorm.Model
	User_uid   string `gorm:"index;type:varchar(22)"`
	Height     int
	Weight     int
	Age        int
	Range_time int
}
