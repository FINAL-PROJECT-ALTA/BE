package entities

import "gorm.io/gorm"

type Image struct {
	gorm.Model
	Image    string `gorm:"type:varchar(100)"`
	Food_uid string `gorm:"type:varchar(100)"`
}
