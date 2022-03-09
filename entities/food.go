package entities

import (
	"time"

	"gorm.io/gorm"
)

type Food struct {
	ID            uint   `gorm:"primarykey" json:"-"`
	Food_uid      string `gorm:"index;type:varchar(22)"`
	Name          string `gorm:"type:varchar(100)"`
	Calories      int    `gorm:"type:int(100)"`
	Energy        int    `gorm:"type:int(100)"`
	Carbohidrate  int    `gorm:"type:int(100)"`
	Protein       int    `gorm:"type:int(100)"`
	Unit          string `gorm:"type:varchar(100)"`
	Unit_value    int    `gorm:"type:int(100)"`
	Food_category string `gorm:"type:enum('fruit','drink','junk food','food','snack','None');default:'None'"`
	Image         string `gorm:"type:varchar(100)"`

	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
