package admin

import (
	"HealthFit/configs"
	"HealthFit/delivery/middlewares"
	"HealthFit/entities"
	utils "HealthFit/utils/mysql"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestCreateAdin(t *testing.T) {
	config := configs.GetConfig()
	db := utils.InitDB(config)
	repo := New(db)

	t.Run("fail run Create", func(t *testing.T) {
		db.Migrator().DropTable(&entities.User{})
		db.AutoMigrate(&entities.User{})
		mocUserP := entities.User{
			Name:     "anonim1",
			Email:    "anonim@1",
			Password: "anonim1",
		}
		if _, err := repo.Register(mocUserP); err != nil {
			t.Fatal()
		}
		mocUser := entities.User{Model: gorm.Model{ID: 1}, Name: "anonim123", Email: "anonim@123", Password: "anonim123"}
		_, err := repo.Register(mocUser)
		assert.NotNil(t, err)
	})

	t.Run("success run Create", func(t *testing.T) {
		db.Migrator().DropTable(&entities.User{})
		db.AutoMigrate(&entities.User{})
		mocUser := entities.User{Name: "anonim123", Email: "anonim@123", Password: "anonim123"}
		res, err := repo.Register(mocUser)

		res.Password, _ = middlewares.HashPassword(mocUser.Password)

		assert.Nil(t, err)
		assert.Equal(t, "anonim123", res.Name)
		assert.Equal(t, "anonim@123", res.Email)
		assert.Equal(t, res.Password, res.Password)

	})

}
