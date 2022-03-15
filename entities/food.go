package entities

import (
	"time"

	"gorm.io/gorm"
)

type Food struct {
	ID            uint   `gorm:"primarykey" json:"-"`
	Food_uid      string `gorm:"index,unique;type:varchar(22)" json:"food_uid"`
	Name          string `gorm:"type:varchar(100)" json:"name"`
	Calories      int    `gorm:"type:int(100)" json:"calories"`
	Energy        int    `gorm:"type:int(100)" json:"energy"`
	Carbohidrate  int    `gorm:"type:int(100)" json:"carbohidrate"`
	Protein       int    `gorm:"type:int(100)" json:"protein"`
	Unit          string `gorm:"type:varchar(100)" json:"unit"`
	Unit_value    int    `gorm:"type:int(100)" json:"unit_value"`
	Food_category string `gorm:"type:enum('fruit','drink','junk food','food','snack','None');default:'None'" json:"food_category"`
	Image         string `gorm:"type:varchar(100)" json:"image"`

	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
