package entities

import "gorm.io/gorm"

type Menu struct {
	gorm.Model
	Menu_uid      string        `gorm:"index;type:varchar(22)"`
	Menu_category string        `gorm:"type:enum('breakfast','lunch','dinner','None');default:'None'"`
	Foods         []Detail_menu `gorm:"foreignKey:Menu_uid;references:Menu_uid"`
}
