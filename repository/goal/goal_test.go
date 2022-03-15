package goal

import (
	"HealthFit/configs"
	"HealthFit/entities"
	up "HealthFit/repository/user"
	utils "HealthFit/utils/mysql"
	"testing"

	"github.com/labstack/gommon/log"
	"github.com/stretchr/testify/assert"
)

func TestCreate(t *testing.T) {
	config := configs.GetConfig()
	db := utils.InitDB(config)
	repo := New(db)

	t.Run("Failed to create goal", func(t *testing.T) {
		db.Migrator().DropTable(&entities.User{}, &entities.Goal{})
		db.AutoMigrate(&entities.User{}, &entities.Goal{})
		mocUser := entities.User{
			Name:     "arya",
			Email:    "arya@mail.com",
			Password: "arya",
			Gender:   "male",
		}
		resU, errU := up.New(db).Register(mocUser)
		if errU != nil {
			t.Fail()
		}

		mockGoal := entities.Goal{
			User_uid:      resU.User_uid,
			Height:        170,
			Weight:        75,
			Age:           25,
			Daily_active:  "not active",
			Weight_target: 10,
			Range_time:    365,
			Target:        "gain weight",
			Status:        "active",
		}

		_, errG := repo.Create(mockGoal)
		log.Info(errG)

		mockGoalU := entities.Goal{
			User_uid:      resU.User_uid,
			Height:        170,
			Weight:        75,
			Age:           25,
			Daily_active:  "not active",
			Weight_target: 10,
			Range_time:    365,
			Target:        "gain weight",
			Status:        "active",
		}

		_, err := repo.Create(mockGoalU)

		assert.NotNil(t, err)

	})

	t.Run("failed to create goal", func(t *testing.T) {
		db.Migrator().DropTable(&entities.User{}, &entities.Goal{})
		db.AutoMigrate(&entities.User{}, &entities.Goal{})
		mocUser := entities.User{
			Name:     "anonim123",
			Email:    "anonim@123",
			Password: "anonim123",
			Gender:   "male",
		}
		resU, errU := up.New(db).Register(mocUser)
		if errU != nil {
			t.Fail()
		}

		mockGoal := entities.Goal{
			User_uid:      resU.User_uid,
			Height:        170,
			Weight:        75,
			Age:           25,
			Daily_active:  "not active",
			Weight_target: 10,
			Range_time:    365,
			Target:        "gain weight",
		}

		res, err := repo.Create(mockGoal)
		log.Info(res.Status)

		assert.NotNil(t, err)

	})

	t.Run("fail to create goal", func(t *testing.T) {
		db.Migrator().DropTable(&entities.User{}, &entities.Goal{})
		db.AutoMigrate(&entities.User{}, &entities.Goal{})
		mocUser := entities.User{
			Name:     "anonim123",
			Email:    "anonim@123",
			Password: "anonim123",
		}
		resU, errU := up.New(db).Register(mocUser)
		if errU != nil {
			t.Fail()
		}

		mockGoal := entities.Goal{
			User_uid:      resU.User_uid,
			Height:        170,
			Weight:        75,
			Age:           25,
			Daily_active:  "not active",
			Weight_target: 10,
			Range_time:    365,
			Target:        "gain weight",
		}

		_, err := repo.Create(mockGoal)

		assert.NotNil(t, err)

	})

	t.Run("succes to create goal", func(t *testing.T) {
		db.Migrator().DropTable(&entities.User{}, &entities.Goal{})
		db.AutoMigrate(&entities.User{}, &entities.Goal{})
		mocUser := entities.User{
			Name:     "arya",
			Email:    "arya@mail.com",
			Password: "arya",
			Gender:   "male",
		}
		resU, errU := up.New(db).Register(mocUser)
		if errU != nil {
			t.Fail()
		}

		mockGoal := entities.Goal{
			User_uid:      resU.User_uid,
			Height:        170,
			Weight:        75,
			Age:           25,
			Daily_active:  "not active",
			Weight_target: 10,
			Range_time:    365,
			Target:        "gain weight",
			Status:        "active",
		}

		_, err := repo.Create(mockGoal)

		assert.Nil(t, err)

	})

}

func TestGetByID(t *testing.T) {
	config := configs.GetConfig()
	db := utils.InitDB(config)
	repo := New(db)

	t.Run("succes to create goal", func(t *testing.T) {
		db.Migrator().DropTable(&entities.User{}, &entities.Goal{})
		db.AutoMigrate(&entities.User{}, &entities.Goal{})
		mocUser := entities.User{
			Name:     "arya",
			Email:    "arya@mail.com",
			Password: "arya",
			Gender:   "male",
		}
		resU, errU := up.New(db).Register(mocUser)
		if errU != nil {
			t.Fail()
		}

		mockGoal := entities.Goal{
			User_uid:      resU.User_uid,
			Height:        170,
			Weight:        75,
			Age:           25,
			Daily_active:  "not active",
			Weight_target: 10,
			Range_time:    365,
			Target:        "gain weight",
			Status:        "active",
		}

		resG, errG := repo.Create(mockGoal)
		if errG != nil {
			t.Fail()
		}

		_, err := repo.GetById(resG.Goal_uid, resG.User_uid)
		log.Info(err)

		assert.NotNil(t, err)

	})
}

func TestGetAll(t *testing.T) {
	config := configs.GetConfig()
	db := utils.InitDB(config)
	repo := New(db)

	t.Run("succes to create goal", func(t *testing.T) {
		db.Migrator().DropTable(&entities.User{}, &entities.Goal{})
		db.AutoMigrate(&entities.User{}, &entities.Goal{})
		mocUser := entities.User{
			Name:     "arya",
			Email:    "arya@mail.com",
			Password: "arya",
			Gender:   "male",
		}
		resU, errU := up.New(db).Register(mocUser)
		if errU != nil {
			t.Fail()
		}

		mockGoal := entities.Goal{
			User_uid:      resU.User_uid,
			Height:        170,
			Weight:        75,
			Age:           25,
			Daily_active:  "not active",
			Weight_target: 10,
			Range_time:    365,
			Target:        "gain weight",
			Status:        "active",
		}

		resG, errG := repo.Create(mockGoal)
		if errG != nil {
			t.Fail()
		}

		_, err := repo.GetAll(resG.Goal_uid)
		log.Info(err)

		assert.Nil(t, err)

	})

}

func TestUpdate(t *testing.T) {
	config := configs.GetConfig()
	db := utils.InitDB(config)
	repo := New(db)
	db.Migrator().DropTable(&entities.User{}, &entities.Goal{})
	db.AutoMigrate(&entities.User{}, &entities.Goal{})

	t.Run("succes to create goal", func(t *testing.T) {

		mocUser := entities.User{
			Name:     "arya",
			Email:    "arya@mail.com",
			Password: "arya",
			Gender:   "male",
		}
		resU, errU := up.New(db).Register(mocUser)
		if errU != nil {
			t.Fail()
		}

		mockGoal := entities.Goal{
			User_uid:      resU.User_uid,
			Height:        170,
			Weight:        75,
			Age:           25,
			Daily_active:  "not active",
			Weight_target: 10,
			Range_time:    365,
			Target:        "gain weight",
			Status:        "active",
		}

		resG, errG := repo.Create(mockGoal)
		if errG != nil {
			t.Fail()
		}

		upGoal := entities.Goal{
			User_uid: resU.User_uid,
			Height:   200,
		}

		_, err := repo.Update(resG.Goal_uid, upGoal)

		assert.Nil(t, err)

	})

	t.Run("Fail to update goal", func(t *testing.T) {

		upGoal := entities.Goal{
			Height: 200,
		}

		_, err := repo.Update("ajnjanl", upGoal)

		assert.NotNil(t, err)

	})

}

func TestDelete(t *testing.T) {
	config := configs.GetConfig()
	db := utils.InitDB(config)
	repo := New(db)
	db.Migrator().DropTable(&entities.User{}, &entities.Goal{})
	db.AutoMigrate(&entities.User{}, &entities.Goal{})

	t.Run("succes to create goal", func(t *testing.T) {

		mocUser := entities.User{
			Name:     "arya",
			Email:    "arya@mail.com",
			Password: "arya",
			Gender:   "male",
		}
		resU, errU := up.New(db).Register(mocUser)
		if errU != nil {
			t.Fail()
		}

		mockGoal := entities.Goal{
			User_uid:      resU.User_uid,
			Height:        170,
			Weight:        75,
			Age:           25,
			Daily_active:  "not active",
			Weight_target: 10,
			Range_time:    365,
			Target:        "gain weight",
			Status:        "active",
		}

		resG, errG := repo.Create(mockGoal)
		if errG != nil {
			t.Fail()
		}

		err := repo.Delete(resG.Goal_uid, resG.User_uid)

		assert.Nil(t, err)

	})

}
