package entities

import "gorm.io/gorm"

type Detail_menu struct {
	gorm.Model
	Detail_menu_uid string
	Menu_uid        string `gorm:"index;type:varchar(22)"`
	Food_uid        string `gorm:"index;type:varchar(22)"`
	Food            Food   `gorm:"foreignKey:Food_uid;references:Food_uid"`
}
