package entities

import "gorm.io/gorm"

type Food struct {
	gorm.Model
	Food_uid      string `gorm:"index;type:varchar(22)"`
	Name          string `gorm:"type:varchar(100)"`
	Kalori        int    `gorm:"type:int(100)"`
	Food_category string `gorm:"type:enum('fruit','minuman','junk food','food', 'snack');default:'None'"`
	Images        Image  `gorm:"foreignKey:Food_uid_uid;references:Food_uid"`
}
