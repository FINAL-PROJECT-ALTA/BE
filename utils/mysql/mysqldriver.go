package utils

import (
	config "HealthFit/configs"
	"HealthFit/entities"
	"fmt"

	"github.com/labstack/gommon/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB(config *config.AppConfig) *gorm.DB {

	connectionString := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8&parseTime=True&loc=%v",
		config.Username,
		config.Password,
		config.Address,
		config.DB_Port,
		config.Name,
		config.LOC,
	)
	fmt.Println(connectionString)
	DB, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{})

	if err != nil {
		log.Info("error in connect database ", err)
		panic(err)
	}

	AutoMigrate(DB)
	return DB
}

func AutoMigrate(DB *gorm.DB) {
	DB.AutoMigrate(&entities.User{})
	DB.AutoMigrate(&entities.Admin{})
	DB.AutoMigrate(&entities.Goal{})
	DB.AutoMigrate(&entities.User_history{})
	DB.AutoMigrate(&entities.Foods{})
}
