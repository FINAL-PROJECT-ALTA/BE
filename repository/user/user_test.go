package user

import (
	"HealthFit/configs"
	"HealthFit/delivery/middlewares"
	"HealthFit/entities"
	utils "HealthFit/utils/mysql"
	"testing"

	"github.com/labstack/gommon/log"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestCreate(t *testing.T) {
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

func TestGetById(t *testing.T) {
	config := configs.GetConfig()
	db := utils.InitDB(config)
	repo := New(db)
	db.Migrator().DropTable(&entities.User{})
	db.AutoMigrate(&entities.User{})

	t.Run("success run GetById", func(t *testing.T) {
		mocUser := entities.User{
			Name:     "anonim123",
			Email:    "anonim@1",
			Password: "anonim1",
			Gender:   "male",
		}

		res, err := repo.Register(mocUser)
		if err != nil {
			t.Fatal()
		}
		uid := res.User_uid

		_, errA := repo.GetById(uid)
		assert.Nil(t, errA)

	})

	t.Run("fail run GetById", func(t *testing.T) {
		mocUser := entities.User{
			Name:     "test",
			Email:    "test",
			Password: "test",
			Gender:   "male",
		}

		res, err := repo.Register(mocUser)
		if err != nil {
			t.Fatal()
		}
		log.Info(res)

		resA, _ := repo.GetById("eaf")

		log.Info(resA)
		assert.NotEqual(t, "hefb", res.User_uid)

	})
}

func TestUpdateById(t *testing.T) {
	config := configs.GetConfig()
	db := utils.InitDB(config)
	repo := New(db)
	db.Migrator().DropTable(&entities.User{})
	db.AutoMigrate(&entities.User{})

	t.Run("success run UpdateById", func(t *testing.T) {
		mocUser := entities.User{Name: "anonim123", Email: "anonim@123", Password: "anonim123"}
		res, err := repo.Register(mocUser)
		if err != nil {
			t.Fatal()
		}
		uid := res.User_uid

		mockUser := entities.User{Name: "anonim321", Email: "anonim@321", Password: "anonim321"}
		res, errA := repo.Update(uid, mockUser)
		assert.Nil(t, errA)
		assert.Equal(t, "anonim321", res.Name)
		assert.Equal(t, "anonim@321", res.Email)
		assert.Equal(t, "anonim321", res.Password)
	})

	t.Run("fail run UpdateById", func(t *testing.T) {
		mockUser := entities.User{Name: "anonim456", Email: "anonim@456", Password: "456"}
		_, err := repo.Update("10", mockUser)
		assert.NotNil(t, err)
	})
}

func TestDeleteById(t *testing.T) {
	config := configs.GetConfig()
	db := utils.InitDB(config)
	repo := New(db)
	db.Migrator().DropTable(&entities.User{})
	db.AutoMigrate(&entities.User{})

	t.Run("success run DeleteById", func(t *testing.T) {
		mocUser := entities.User{Name: "anonim123", Email: "anonim@123", Password: "anonim123"}
		res, err := repo.Register(mocUser)
		if err != nil {
			t.Fatal()
		}
		uid := res.User_uid

		errA := repo.Delete(uid)
		assert.Nil(t, errA)
	})

	t.Run("fail run DeleteById", func(t *testing.T) {
		err := repo.Delete("10")
		assert.NotNil(t, err)
	})
}
