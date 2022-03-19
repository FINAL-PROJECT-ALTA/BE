package entities

import (
	"time"

	"gorm.io/gorm"
)

type User_history struct {
	ID uint `gorm:"primarykey" json:"-"`

	User_history_uid string         `json:"user_history_uid"`
	User_uid         string         `gorm:"index;type:varchar(32)" json:"user_uid"`
	Goal_uid         string         `gorm:"index;type:varchar(32)" json:"goal_uid"`
	Menu_uid         string         `gorm:"index;type:varchar(32)" json:"menu_uid"`
	Menu             Menu           `gorm:"foreignKey:Menu_uid;references:Menu_uid" json:"menu"`
	CreatedAt        time.Time      `json:"cretedAt"`
	UpdatedAt        time.Time      `json:"-"`
	DeletedAt        gorm.DeletedAt `gorm:"index" json:"-"`
}
