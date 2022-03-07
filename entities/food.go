package entities

import "gorm.io/gorm"

type Food struct {
	gorm.Model
	Admin_uid     string `gorm:"index;type:varchar(22)"`
	Food_uid      string `gorm:"index;type:varchar(22)"`
	Name          string `gorm:"type:varchar(100)"`
	Calories      int    `gorm:"type:int(100)"`
	Energy        int    `gorm:"type:int(100)"`
	Carbohidrate  int    `gorm:"type:int(100)"`
	Protein       int    `gorm:"type:int(100)"`
	Food_category string `gorm:"type:enum('fruit','drink','junk food','food','snack','None');default:'None'"`
	Image         Image  `gorm:"foreignKey:Food_uid_uid;references:Food_uid"`
}
