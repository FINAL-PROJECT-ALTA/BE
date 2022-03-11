package entities

import "gorm.io/gorm"

type Goal struct {
	gorm.Model
	Goal_uid      string
	User_uid      string `gorm:"index;type:varchar(22)"`
	Height        int
	Weight        int
	Age           int
	Daily_active  string `gorm:"type:enum('not active','little active','quite active','active','very active')"`
	Weight_target int
	Range_time    int
	Target        string `gorm:"type:enum('gain weight','lose weight')"`
	Status        string `gorm:"type:enum('active', 'not active','cencel')"`
}
