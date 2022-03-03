package entities

import "gorm.io/gorm"

type Detail_menu struct {
	gorm.Model
	Menu_uid string `gorm:"index;type:varchar(22)"`
	Food_uid string
}
