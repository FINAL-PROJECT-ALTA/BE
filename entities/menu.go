package entities

import "gorm.io/gorm"

type Menu struct {
	gorm.Model
	Menu_uid      string        `gorm:"index;type:varchar(22)"`
	User_uid      string        `gorm:"type:varchar(22)"`
	Menu_category string        `gorm:"type:enum('breakfast','lunch','dinner','None');default:'None'"`
	Detail_menu   []Detail_menu `gorm:"foreignKey:Menu_uid;references:Menu_uid"`
}
