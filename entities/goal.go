package entities

import (
	"time"

	"gorm.io/gorm"
)

type Goal struct {
	ID uint `gorm:"primarykey" json:"-"`

	Goal_uid      string `gorm:"index;type:varchar(22)" json:"goal_uid"`
	User_uid      string `gorm:"index;type:varchar(22)" json:"-"`
	Height        int    `json:"height"`
	Weight        int    `json:"weight"`
	Age           int    `json:"age"`
	Daily_active  string `gorm:"type:enum('not active','little active','quite active','active','very active')" json:"daily_active"`
	Weight_target int    `json:"weight_target"`
	Range_time    int    `json:"range_time"`
	Target        string `gorm:"type:enum('gain weight','lose weight')" json:"target"`
	Status        string `gorm:"type:enum('active','not active','cancel')" json:"status"`

	CreatedAt time.Time      `json:"cretedAt"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
