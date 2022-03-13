package entities

import (
	"time"

	"gorm.io/gorm"
)

type Menu struct {
	ID uint `gorm:"primarykey" json:"-"`

	Menu_uid       string        `gorm:"index;type:varchar(22)" json:"menu_uid"`
	User_uid       string        `gorm:"type:varchar(22)" json:"user_uid"`
	Menu_category  string        `gorm:"type:enum('breakfast','lunch','dinner','overtime')" json:"menu_category"`
	Created_by     string        `gorm:"type:enum('admin','user');default:'admin'" json:"created_by"`
	Count          int           `json:"-"`
	Total_calories int           `json:"total_calories"`
	Detail_menu    []Detail_menu `gorm:"foreignKey:Menu_uid;references:Menu_uid" json:"detail_menu"`

	CreatedAt time.Time      `json:"cretedAt"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
