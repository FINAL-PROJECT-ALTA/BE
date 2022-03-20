package userhistory

import (
	"HealthFit/configs"
	"HealthFit/entities"

	dp "HealthFit/repository/detail_menu"
	fp "HealthFit/repository/foods"
	gp "HealthFit/repository/goal"
	mp "HealthFit/repository/menu"
	up "HealthFit/repository/user"
	utils "HealthFit/utils/mysql"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreate(t *testing.T) {
	config := configs.GetConfig()
	db := utils.InitDB(config)
	repo := New(db)

	t.Run("fail run Create", func(t *testing.T) {
		db.Migrator().DropTable(&entities.User{}, &entities.Goal{}, &entities.Food{}, &entities.Menu{}, &entities.User_history{})
		db.AutoMigrate(&entities.User{}, &entities.Goal{}, &entities.Food{}, &entities.Menu{}, &entities.User_history{})

		mocUser := entities.User{
			Name:     "anonim123",
			Email:    "anonim@1",
			Password: "anonim1",
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

		resG, errG := gp.New(db).Create(mockGoal)
		if errG != nil {
			t.Fail()
		}

		mockMenu := entities.Menu{Menu_uid: "bkaebvkaebvk"}

		mockHistory := entities.User_history{
			User_uid: resU.User_uid,
			Goal_uid: resG.Goal_uid,
			Menu_uid: mockMenu.Menu_uid,
		}

		_, err := repo.Insert(mockHistory)

		assert.Nil(t, err)

	})

	t.Run("Fail run Create", func(t *testing.T) {
		db.Migrator().DropTable(&entities.User{}, &entities.Goal{}, &entities.Food{}, &entities.Menu{}, &entities.User_history{})
		db.AutoMigrate(&entities.User{}, &entities.Goal{}, &entities.Food{}, &entities.Menu{}, &entities.User_history{})

		mockGoal := entities.Goal{
			User_uid:      "resU.User_uid",
			Height:        170,
			Weight:        75,
			Age:           25,
			Daily_active:  "not active",
			Weight_target: 10,
			Range_time:    365,
			Target:        "gain weight",
			Status:        "active",
		}

		_, errG := gp.New(db).Create(mockGoal)
		if errG != nil {
			t.Fail()
		}
		mockHistory := entities.User_history{}

		_, err := repo.Insert(mockHistory)

		assert.NotNil(t, err)

	})

	t.Run("Fail run Create", func(t *testing.T) {
		db.Migrator().DropTable(&entities.User{}, &entities.Goal{}, &entities.Food{}, &entities.Menu{}, &entities.User_history{})
		db.AutoMigrate(&entities.User{}, &entities.Goal{}, &entities.Food{}, &entities.Menu{}, &entities.User_history{})

		mocUser := entities.User{
			Name:     "anonim123",
			Email:    "anonim@1",
			Password: "anonim1",
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

		resG, errG := gp.New(db).Create(mockGoal)
		if errG != nil {
			t.Fail()
		}

		mockFood := entities.Food{
			Food_uid: "kenfkak",
		}

		resF, errF := fp.New(db).Create(mockFood)
		if errF != nil {
			t.Fail()
		}

		mockMenu := entities.Menu{
			User_uid:      resU.User_uid,
			Menu_category: "lunch",
			Created_by:    "admin",
		}
		resM, errM := mp.New(db).CreateMenuUser([]entities.Food{{Food_uid: resF.Food_uid}}, mockMenu)
		if errM != nil {
			t.Fail()
		}

		mockDetail := entities.Detail_menu{
			Menu_uid: resM.Menu_uid,
			Food_uid: resF.Food_uid,
		}

		_, errD := dp.New(db).Create(mockDetail)
		if errD != nil {
			t.Fail()
		}

		mockHistory := entities.User_history{
			User_uid: resU.User_uid,
			Goal_uid: resG.Goal_uid,
			Menu_uid: "akjbn alkbalvma akv av l",
		}

		_, err := repo.Insert(mockHistory)

		assert.NotNil(t, err)

	})
}

func TestGetAll(t *testing.T) {
	config := configs.GetConfig()
	db := utils.InitDB(config)
	repo := New(db)

	t.Run("success run Create", func(t *testing.T) {
		db.Migrator().DropTable(&entities.User{}, &entities.Goal{}, &entities.Food{}, &entities.Menu{}, &entities.User_history{})
		db.AutoMigrate(&entities.User{}, &entities.Goal{}, &entities.Food{}, &entities.Menu{}, &entities.User_history{})

		mocUser := entities.User{
			Name:     "anonim123",
			Email:    "anonim@1",
			Password: "anonim1",
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

		resG, errG := gp.New(db).Create(mockGoal)
		if errG != nil {
			t.Fail()
		}

		mockMenu := entities.Menu{Menu_uid: "bkaebvkaebvk"}

		mockHistory := entities.User_history{
			User_uid: resU.User_uid,
			Goal_uid: resG.Goal_uid,
			Menu_uid: mockMenu.Menu_uid,
		}

		_, errH := repo.Insert(mockHistory)
		if errH != nil {
			t.Fail()
		}

		_, err := repo.GetAll(resU.User_uid)

		assert.Nil(t, err)

	})

	t.Run("success run Create", func(t *testing.T) {
		db.Migrator().DropTable(&entities.User{}, &entities.Goal{}, &entities.Food{}, &entities.Menu{}, &entities.User_history{})
		db.AutoMigrate(&entities.User{}, &entities.Goal{}, &entities.Food{}, &entities.Menu{}, &entities.User_history{})

		mocUser := entities.User{
			Name:     "anonim123",
			Email:    "anonim@1",
			Password: "anonim1",
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

		_, errG := gp.New(db).Create(mockGoal)
		if errG != nil {
			t.Fail()
		}

		_, err := repo.GetAll(resU.User_uid)

		assert.NotNil(t, err)
	})

}

func TestGetById(t *testing.T) {
	config := configs.GetConfig()
	db := utils.InitDB(config)
	repo := New(db)

	t.Run("success run Create", func(t *testing.T) {
		db.Migrator().DropTable(&entities.User{}, &entities.Goal{}, &entities.Food{}, &entities.Menu{}, &entities.User_history{})
		db.AutoMigrate(&entities.User{}, &entities.Goal{}, &entities.Food{}, &entities.Menu{}, &entities.User_history{})

		mocUser := entities.User{
			Name:     "anonim123",
			Email:    "anonim@1",
			Password: "anonim1",
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

		resG, errG := gp.New(db).Create(mockGoal)
		if errG != nil {
			t.Fail()
		}

		mockMenu := entities.Menu{Menu_uid: "bkaebvkaebvk"}

		mockHistory := entities.User_history{
			User_uid: resU.User_uid,
			Goal_uid: resG.Goal_uid,
			Menu_uid: mockMenu.Menu_uid,
		}

		resH, errH := repo.Insert(mockHistory)
		if errH != nil {
			t.Fail()
		}

		_, err := repo.GetById(resU.User_uid, resH.User_history_uid)

		assert.Nil(t, err)

	})

	t.Run("success run Create", func(t *testing.T) {
		db.Migrator().DropTable(&entities.User{}, &entities.Goal{}, &entities.Food{}, &entities.Menu{}, &entities.User_history{})
		db.AutoMigrate(&entities.User{}, &entities.Goal{}, &entities.Food{}, &entities.Menu{}, &entities.User_history{})

		mocUser := entities.User{
			Name:     "anonim123",
			Email:    "anonim@1",
			Password: "anonim1",
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
			Status:        "cancel",
		}

		resG, errG := gp.New(db).Create(mockGoal)
		if errG != nil {
			t.Fail()
		}

		mockMenu := entities.Menu{Menu_uid: "bkaebvkaebvk"}

		mockHistory := entities.User_history{
			User_uid: resU.User_uid,
			Goal_uid: resG.Goal_uid,
			Menu_uid: mockMenu.Menu_uid,
		}

		_, errH := repo.Insert(mockHistory)
		if errH != nil {
			t.Fail()
		}

		_, err := repo.GetById(resU.User_uid, "ekfjankfjanv")

		assert.NotNil(t, err)

	})

}
