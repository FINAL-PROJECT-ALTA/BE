package entities

import (
	"time"

	"gorm.io/gorm"
)

type Detail_menu struct {
	ID uint `gorm:"primarykey" json:"-"`

	Detail_menu_uid string         `json:"-"`
	Menu_uid        string         `gorm:"index;type:varchar(22)" json:"-"`
	Food_uid        string         `gorm:"index;type:varchar(22)" json:"-"`
	Food            Food           `gorm:"foreignKey:Food_uid;references:Food_uid" json:"food"`
	CreatedAt       time.Time      `json:"-"`
	UpdatedAt       time.Time      `json:"-"`
	DeletedAt       gorm.DeletedAt `gorm:"index" json:"-"`
}
