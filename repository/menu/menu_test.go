package menu

import (
	"HealthFit/configs"
	"HealthFit/repository/admin"
	food "HealthFit/repository/foods"

	"HealthFit/repository/goal"
	"HealthFit/repository/user"
	"fmt"

	utils "HealthFit/utils/mysql"

	"HealthFit/entities"
	"testing"

	"github.com/labstack/gommon/log"

	"github.com/stretchr/testify/assert"
)

func TestCreateMenuAdmin(t *testing.T) {
	config := configs.GetConfig()
	db := utils.InitDB(config)
	repo := New(db)

	t.Run("success Create menu", func(t *testing.T) {
		db.Migrator().DropTable(&entities.User{}, &entities.User_history{}, &entities.Detail_menu{}, &entities.Food{}, &entities.Menu{})
		db.Migrator().DropTable(&entities.Food{}, &entities.Menu{})
		db.AutoMigrate(&entities.User{})
		db.AutoMigrate(&entities.Food{})
		db.AutoMigrate(&entities.Menu{})
		db.AutoMigrate(&entities.Detail_menu{})
		db.AutoMigrate(&entities.User_history{})

		mocFood1 := entities.Food{
			Name:          "makanan",
			Calories:      100,
			Energy:        200,
			Carbohidrate:  300,
			Protein:       400,
			Food_category: "snack",
			Unit:          "ons",
			Unit_value:    1,
		}
		f1, _ := food.New(db).Create(mocFood1)
		mocUser := entities.User{
			Name:     "test",
			Email:    "test@mail.com",
			Password: "test",
		}
		adminres, _ := admin.New(db).Register(mocUser)

		mocFood := []entities.Food{
			{
				Food_uid: f1.Food_uid,
			},
			{
				Food_uid: f1.Food_uid,
			},
		}
		mocMenu := entities.Menu{
			User_uid:      adminres.User_uid,
			Menu_category: "breakfast",
			Created_by:    "admin",
		}
		_, err := repo.CreateMenuAdmin(mocFood, mocMenu)

		assert.Nil(t, err)

	})

	t.Run("fail Create menu", func(t *testing.T) {
		db.Migrator().DropTable(&entities.User{}, &entities.User_history{}, &entities.Detail_menu{}, &entities.Food{}, &entities.Menu{})
		db.Migrator().DropTable(&entities.Food{}, &entities.Menu{})
		db.AutoMigrate(&entities.User{})
		db.AutoMigrate(&entities.Food{})
		db.AutoMigrate(&entities.Menu{})
		db.AutoMigrate(&entities.Detail_menu{})
		db.AutoMigrate(&entities.User_history{})

		mocFood1 := entities.Food{
			Name:          "makanan",
			Calories:      100,
			Energy:        200,
			Carbohidrate:  300,
			Protein:       400,
			Food_category: "food",
			Unit:          "ons",
			Unit_value:    1,
		}
		f1, _ := food.New(db).Create(mocFood1)
		mocUser := entities.User{
			Name:     "test",
			Email:    "test@mail.com",
			Password: "test",
		}
		adminres, _ := admin.New(db).Register(mocUser)

		mocFood := []entities.Food{
			{
				Food_uid: f1.Food_uid,
			},
			{
				Food_uid: f1.Food_uid,
			},
		}
		mocMenu := entities.Menu{
			User_uid:      adminres.User_uid,
			Menu_category: "food",
			Created_by:    "admin",
		}
		_, err := repo.CreateMenuAdmin(mocFood, mocMenu)

		assert.NotNil(t, err)

	})

	t.Run("fail Create detail menu", func(t *testing.T) {
		db.Migrator().DropTable(&entities.User{}, &entities.User_history{}, &entities.Detail_menu{}, &entities.Food{}, &entities.Menu{})
		// db.Migrator().DropTable(&entities.Food{}, &entities.Menu{})
		db.AutoMigrate(&entities.User{})
		// db.AutoMigrate(&entities.Food{})
		db.AutoMigrate(&entities.Menu{})
		db.AutoMigrate(&entities.Detail_menu{})
		db.AutoMigrate(&entities.User_history{})

		mocFood1 := entities.Food{
			Name:          "makanan",
			Calories:      100,
			Energy:        200,
			Carbohidrate:  300,
			Protein:       400,
			Food_category: "snack",
			Unit:          "ons",
			Unit_value:    1,
		}
		f1, _ := food.New(db).Create(mocFood1)
		mocUser := entities.User{
			Name:     "test",
			Email:    "test@mail.com",
			Password: "test",
		}
		adminres, _ := admin.New(db).Register(mocUser)

		mocFood := []entities.Food{
			{
				Food_uid: f1.Food_uid,
			},
			{
				Food_uid: f1.Food_uid,
			},
		}
		mocMenu := entities.Menu{
			User_uid:      adminres.User_uid,
			Menu_category: "breakfast",
			Created_by:    "admin",
		}
		_, err := repo.CreateMenuAdmin(mocFood, mocMenu)

		assert.NotNil(t, err)

	})

	t.Run("fail Create detail menu", func(t *testing.T) {
		db.Migrator().DropTable(&entities.User{}, &entities.User_history{}, &entities.Detail_menu{}, &entities.Food{}, &entities.Menu{})
		// db.Migrator().DropTable(&entities.Food{}, &entities.Menu{})
		db.AutoMigrate(&entities.User{})
		db.AutoMigrate(&entities.Food{})
		db.AutoMigrate(&entities.Menu{})
		db.AutoMigrate(&entities.User_history{})
		db.Migrator().DropTable(&entities.Detail_menu{})

		mocFood1 := entities.Food{
			Name:          "makanan",
			Calories:      100,
			Energy:        200,
			Carbohidrate:  300,
			Protein:       400,
			Food_category: "snack",
			Unit:          "ons",
			Unit_value:    1,
		}
		f1, _ := food.New(db).Create(mocFood1)
		mocUser := entities.User{
			Name:     "test",
			Email:    "test@mail.com",
			Password: "test",
		}
		fmt.Println(f1)
		adminres, _ := admin.New(db).Register(mocUser)

		mocFood := []entities.Food{
			{
				Food_uid: "f1Fooduid",
			},
			{
				Food_uid: "f1Foodid",
			},
		}
		mocMenu := entities.Menu{
			User_uid:      adminres.User_uid,
			Menu_category: "breakfast",
			Created_by:    "admin",
		}
		_, err := repo.CreateMenuAdmin(mocFood, mocMenu)

		assert.NotNil(t, err)

	})

	t.Run("fail update menu add total calories", func(t *testing.T) {
		db.Migrator().DropTable(&entities.User{}, &entities.User_history{}, &entities.Detail_menu{}, &entities.Food{}, &entities.Menu{})
		db.AutoMigrate(&entities.User{})
		db.AutoMigrate(&entities.Food{})
		db.AutoMigrate(&entities.Menu{})
		db.AutoMigrate(&entities.Detail_menu{})
		db.AutoMigrate(&entities.User_history{})
		// db.Migrator().DropColumn(&entities.Menu{}, "total_calories")

		mocFood1 := entities.Food{
			Name:          "makanan",
			Calories:      2147483646,
			Energy:        200,
			Carbohidrate:  300,
			Protein:       400,
			Food_category: "snack",
			Unit:          "ons",
			Unit_value:    1,
		}
		f1, _ := food.New(db).Create(mocFood1)
		mocUser := entities.User{
			Name:     "test",
			Email:    "test@mail.com",
			Password: "test",
		}
		adminres, _ := admin.New(db).Register(mocUser)

		mocFood := []entities.Food{
			{
				Food_uid: f1.Food_uid,
			},
			{
				Food_uid: f1.Food_uid,
			},
			{
				Food_uid: f1.Food_uid,
			},
			{
				Food_uid: f1.Food_uid,
			},
			{
				Food_uid: f1.Food_uid,
			},
		}
		mocMenu := entities.Menu{
			User_uid:      adminres.User_uid,
			Menu_category: "breakfast",
			Created_by:    "admin",
		}
		_, err := repo.CreateMenuAdmin(mocFood, mocMenu)

		assert.NotNil(t, err)

	})

}

func TestCreateMenuUser(t *testing.T) {
	config := configs.GetConfig()
	db := utils.InitDB(config)
	repo := New(db)

	t.Run("success create menu", func(t *testing.T) {
		db.Migrator().DropTable(&entities.Goal{}, &entities.User{}, &entities.User_history{}, &entities.Detail_menu{}, &entities.Food{}, &entities.Menu{})
		db.AutoMigrate(&entities.User{})
		db.AutoMigrate(&entities.Goal{})
		db.AutoMigrate(&entities.Food{})
		db.AutoMigrate(&entities.Menu{})
		db.AutoMigrate(&entities.User_history{})
		db.AutoMigrate(&entities.Detail_menu{})

		mocFood1 := entities.Food{
			Name:          "makanan",
			Calories:      100,
			Energy:        200,
			Carbohidrate:  300,
			Protein:       400,
			Food_category: "snack",
			Unit:          "ons",
			Unit_value:    1,
		}
		f1, _ := food.New(db).Create(mocFood1)

		mocUser := entities.User{
			Name:     "test",
			Email:    "testuser@mail.com",
			Password: "test",
			Gender:   "Male",
		}

		resUser, _ := user.New(db).Register(mocUser)
		mocAdmin := entities.User{
			Name:     "test",
			Email:    "testadmin@mail.com",
			Password: "test",
		}

		resAdmin, _ := admin.New(db).Register(mocAdmin)
		fmt.Println(resAdmin)
		mockGoal := entities.Goal{
			User_uid:      resUser.User_uid,
			Height:        160,
			Weight:        70,
			Age:           24,
			Daily_active:  "not active",
			Weight_target: 1,
			Range_time:    30,
			Target:        "lose weight",
			Status:        "active",
		}
		resGoal, _ := goal.New(db).Create(mockGoal)
		fmt.Println(resGoal)

		mocFood := []entities.Food{
			{
				Food_uid: f1.Food_uid,
			},
			{
				Food_uid: f1.Food_uid,
			},
		}
		mocMenu := entities.Menu{
			User_uid:      resAdmin.User_uid,
			Menu_category: "breakfast",
			Created_by:    "admin",
		}
		resMenuAdmin, _ := repo.CreateMenuAdmin(mocFood, mocMenu)
		fmt.Println(resMenuAdmin)

		mocMenuUser := entities.Menu{
			User_uid:      resUser.User_uid,
			Menu_category: "breakfast",
			Created_by:    "user",
		}
		_, errMenuUser := repo.CreateMenuUser(mocFood, mocMenuUser)
		fmt.Println(errMenuUser)

		assert.Nil(t, errMenuUser)

	})

	///===================new======================//

	t.Run("fail Create menu user because err database create menu", func(t *testing.T) {
		db.Migrator().DropTable(&entities.User{}, &entities.User_history{}, &entities.Detail_menu{}, &entities.Food{}, &entities.Menu{})
		db.Migrator().DropTable(&entities.Food{}, &entities.Menu{})
		db.AutoMigrate(&entities.User{})
		db.AutoMigrate(&entities.Food{})
		db.AutoMigrate(&entities.Menu{})
		db.AutoMigrate(&entities.Detail_menu{})
		db.AutoMigrate(&entities.User_history{})
		mocFood1 := entities.Food{
			Name:          "makanan",
			Calories:      100,
			Energy:        200,
			Carbohidrate:  300,
			Protein:       400,
			Food_category: "snack",
			Unit:          "ons",
			Unit_value:    1,
		}
		f1, _ := food.New(db).Create(mocFood1)
		mocUser := entities.User{
			Name:     "test",
			Email:    "test@mail.com",
			Password: "test",
		}
		adminres, _ := user.New(db).Register(mocUser)

		mocFood := []entities.Food{
			{
				Food_uid: f1.Food_uid,
			},
			{
				Food_uid: f1.Food_uid,
			},
		}
		mocMenu := entities.Menu{
			User_uid:      adminres.User_uid,
			Menu_category: "zzz",
			Created_by:    "user",
		}
		_, err := repo.CreateMenuUser(mocFood, mocMenu)

		assert.NotNil(t, err)
	})

	t.Run("fail Create menu user because dont have goal active", func(t *testing.T) {
		db.Migrator().DropTable(&entities.User{}, &entities.User_history{}, &entities.Detail_menu{}, &entities.Food{}, &entities.Menu{})
		db.Migrator().DropTable(&entities.Food{}, &entities.Menu{})
		db.AutoMigrate(&entities.User{})
		db.AutoMigrate(&entities.Food{})
		db.AutoMigrate(&entities.Menu{})
		db.AutoMigrate(&entities.Detail_menu{})
		db.AutoMigrate(&entities.User_history{})

		mocFood1 := entities.Food{
			Name:          "makanan",
			Calories:      100,
			Energy:        200,
			Carbohidrate:  300,
			Protein:       400,
			Food_category: "snack",
			Unit:          "ons",
			Unit_value:    1,
		}
		f1, _ := food.New(db).Create(mocFood1)
		mocUser := entities.User{
			Name:     "test",
			Email:    "test@mail.com",
			Password: "test",
		}
		adminres, _ := user.New(db).Register(mocUser)

		mocFood := []entities.Food{
			{
				Food_uid: f1.Food_uid,
			},
			{
				Food_uid: f1.Food_uid,
			},
		}
		mocMenu := entities.Menu{
			User_uid:      adminres.User_uid,
			Menu_category: "breakfast",
			Created_by:    "user",
		}
		_, err := repo.CreateMenuUser(mocFood, mocMenu)

		assert.NotNil(t, err)
	})

	t.Run("fail Create get food", func(t *testing.T) {
		db.Migrator().DropTable(&entities.User{}, &entities.User_history{}, &entities.Detail_menu{}, &entities.Food{}, &entities.Menu{})
		// db.Migrator().DropTable(&entities.Food{}, &entities.Menu{})
		db.AutoMigrate(&entities.User{})
		// db.AutoMigrate(&entities.Food{})
		db.AutoMigrate(&entities.Menu{})
		db.AutoMigrate(&entities.Detail_menu{})
		db.AutoMigrate(&entities.User_history{})

		mocFood1 := entities.Food{
			Name:          "makanan",
			Calories:      100,
			Energy:        200,
			Carbohidrate:  300,
			Protein:       400,
			Food_category: "snack",
			Unit:          "ons",
			Unit_value:    1,
		}
		f1, _ := food.New(db).Create(mocFood1)
		mocUser := entities.User{
			Name:     "test",
			Email:    "test@mail.com",
			Password: "test",
		}
		adminres, _ := user.New(db).Register(mocUser)

		mocFood := []entities.Food{
			{
				Food_uid: f1.Food_uid,
			},
			{
				Food_uid: f1.Food_uid,
			},
		}
		mocMenu := entities.Menu{
			User_uid:      adminres.User_uid,
			Menu_category: "breakfast",
			Created_by:    "user",
		}
		_, err := repo.CreateMenuUser(mocFood, mocMenu)

		assert.NotNil(t, err)

	})

	t.Run("fail Create detail menu", func(t *testing.T) {
		db.Migrator().DropTable(&entities.User{}, &entities.User_history{}, &entities.Detail_menu{}, &entities.Food{}, &entities.Menu{})
		db.AutoMigrate(&entities.User{})
		db.AutoMigrate(&entities.Food{})
		db.AutoMigrate(&entities.Menu{})
		db.AutoMigrate(&entities.User_history{})
		db.Migrator().DropTable(&entities.Detail_menu{})

		mocFood1 := entities.Food{
			Name:          "makanan",
			Calories:      100,
			Energy:        200,
			Carbohidrate:  300,
			Protein:       400,
			Food_category: "snack",
			Unit:          "ons",
			Unit_value:    1,
		}
		f1, _ := food.New(db).Create(mocFood1)
		mocUser := entities.User{
			Name:     "test",
			Email:    "test@mail.com",
			Password: "test",
		}
		fmt.Println(f1)
		adminres, _ := user.New(db).Register(mocUser)

		mocFood := []entities.Food{
			{
				Food_uid: "f1Fooduid",
			},
			{
				Food_uid: "f1Foodid",
			},
		}
		mocMenu := entities.Menu{
			User_uid:      adminres.User_uid,
			Menu_category: "breakfast",
			Created_by:    "admin",
		}
		_, err := repo.CreateMenuUser(mocFood, mocMenu)

		assert.NotNil(t, err)

	})

	t.Run("fail find goal active", func(t *testing.T) {
		db.Migrator().DropTable(&entities.Goal{}, &entities.User{}, &entities.User_history{}, &entities.Detail_menu{}, &entities.Food{}, &entities.Menu{})
		db.AutoMigrate(&entities.User{})
		db.AutoMigrate(&entities.Goal{})
		db.AutoMigrate(&entities.Food{})
		db.AutoMigrate(&entities.Menu{})
		db.AutoMigrate(&entities.User_history{})
		db.AutoMigrate(&entities.Detail_menu{})

		mocFood1 := entities.Food{
			Name:          "makanan",
			Calories:      100,
			Energy:        200,
			Carbohidrate:  300,
			Protein:       400,
			Food_category: "snack",
			Unit:          "ons",
			Unit_value:    1,
		}
		f1, _ := food.New(db).Create(mocFood1)

		mocUser := entities.User{
			Name:     "test",
			Email:    "testuser@mail.com",
			Password: "test",
			Gender:   "Male",
		}

		resUser, _ := user.New(db).Register(mocUser)
		mocAdmin := entities.User{
			Name:     "test",
			Email:    "testadmin@mail.com",
			Password: "test",
		}

		resAdmin, _ := admin.New(db).Register(mocAdmin)
		fmt.Println(resAdmin)
		mockGoal := entities.Goal{
			User_uid:      resUser.User_uid,
			Height:        160,
			Weight:        70,
			Age:           24,
			Daily_active:  "not active",
			Weight_target: 1,
			Range_time:    30,
			Target:        "lose weight",
			Status:        "active",
		}
		resGoal, _ := goal.New(db).Create(mockGoal)
		fmt.Println(resGoal)

		mocFood := []entities.Food{
			{
				Food_uid: f1.Food_uid,
			},
			{
				Food_uid: f1.Food_uid,
			},
		}
		mocMenu := entities.Menu{
			User_uid:      resAdmin.User_uid,
			Menu_category: "breakfast",
			Created_by:    "admin",
		}
		resMenuAdmin, _ := repo.CreateMenuAdmin(mocFood, mocMenu)
		fmt.Println(resMenuAdmin)

		// mockUserhistory := entities.User_history{User_uid: resUser.User_uid, Goal_uid: resGoal.Goal_uid, Menu_uid: resMenuAdmin.Menu_uid}
		// _, errUserHis := userhistory.New(db).Insert(mockUserhistory)
		// fmt.Println(errUserHis)

		mocMenuUser := entities.Menu{
			User_uid:      resUser.User_uid,
			Menu_category: "breakfast",
			Created_by:    "user",
		}
		_, errMenuUser := repo.CreateMenuUser(mocFood, mocMenuUser)
		fmt.Println(errMenuUser)

		assert.Nil(t, errMenuUser)

	})

	t.Run("fail create user history", func(t *testing.T) {
		db.Migrator().DropTable(&entities.Goal{}, &entities.User{}, &entities.User_history{}, &entities.Detail_menu{}, &entities.Food{}, &entities.Menu{})
		db.AutoMigrate(&entities.User{})
		db.AutoMigrate(&entities.Goal{})
		db.AutoMigrate(&entities.Food{})
		db.AutoMigrate(&entities.Menu{})
		db.AutoMigrate(&entities.User_history{})
		db.AutoMigrate(&entities.Detail_menu{})
		db.Migrator().DropColumn(&entities.User_history{}, "goal_uid")

		mocFood1 := entities.Food{
			Name:          "makanan",
			Calories:      100,
			Energy:        200,
			Carbohidrate:  300,
			Protein:       400,
			Food_category: "snack",
			Unit:          "ons",
			Unit_value:    1,
		}
		f1, _ := food.New(db).Create(mocFood1)

		mocUser := entities.User{
			Name:     "test",
			Email:    "testuser@mail.com",
			Password: "test",
			Gender:   "Male",
		}

		resUser, _ := user.New(db).Register(mocUser)
		mocAdmin := entities.User{
			Name:     "test",
			Email:    "testadmin@mail.com",
			Password: "test",
		}

		resAdmin, _ := admin.New(db).Register(mocAdmin)
		fmt.Println(resAdmin)
		mockGoal := entities.Goal{
			User_uid:      resUser.User_uid,
			Height:        160,
			Weight:        70,
			Age:           24,
			Daily_active:  "not active",
			Weight_target: 1,
			Range_time:    30,
			Target:        "lose weight",
			Status:        "active",
		}
		resGoal, _ := goal.New(db).Create(mockGoal)
		fmt.Println(resGoal)

		mocFood := []entities.Food{
			{
				Food_uid: f1.Food_uid,
			},
			{
				Food_uid: f1.Food_uid,
			},
		}
		mocMenu := entities.Menu{
			User_uid:      resAdmin.User_uid,
			Menu_category: "breakfast",
			Created_by:    "admin",
		}
		resMenuAdmin, _ := repo.CreateMenuAdmin(mocFood, mocMenu)
		fmt.Println(resMenuAdmin)

		// mockUserhistory := entities.User_history{User_uid: resUser.User_uid, Goal_uid: resGoal.Goal_uid, Menu_uid: resMenuAdmin.Menu_uid}
		// _, errUserHis := userhistory.New(db).Insert(mockUserhistory)
		// fmt.Println(errUserHis)

		mocMenuUser := entities.Menu{
			User_uid:      resUser.User_uid,
			Menu_category: "breakfast",
			Created_by:    "user",
		}
		_, errMenuUser := repo.CreateMenuUser(mocFood, mocMenuUser)
		fmt.Println(errMenuUser)

		assert.Nil(t, errMenuUser)

	})

	t.Run("fail update menu", func(t *testing.T) {
		db.Migrator().DropTable(&entities.Goal{}, &entities.User{}, &entities.User_history{}, &entities.Detail_menu{}, &entities.Food{}, &entities.Menu{})
		db.AutoMigrate(&entities.User{})
		db.AutoMigrate(&entities.Goal{})
		db.AutoMigrate(&entities.Food{})
		db.AutoMigrate(&entities.Menu{})
		db.AutoMigrate(&entities.User_history{})
		db.AutoMigrate(&entities.Detail_menu{})

		mocFood1 := entities.Food{
			Name:          "makanan",
			Calories:      2147483646,
			Energy:        200,
			Carbohidrate:  300,
			Protein:       400,
			Food_category: "snack",
			Unit:          "ons",
			Unit_value:    1,
		}
		f1, _ := food.New(db).Create(mocFood1)

		mocUser := entities.User{
			Name:     "testuser",
			Email:    "testuser@mail.com",
			Password: "test",
			Gender:   "Male",
		}

		resUser, _ := user.New(db).Register(mocUser)
		mocAdmin := entities.User{
			Name:     "test",
			Email:    "testadmin@mail.com",
			Password: "test",
		}

		resAdmin, _ := admin.New(db).Register(mocAdmin)
		fmt.Println(resAdmin)
		mockGoal := entities.Goal{
			User_uid:      resUser.User_uid,
			Height:        160,
			Weight:        70,
			Age:           24,
			Daily_active:  "not active",
			Weight_target: 1,
			Range_time:    30,
			Target:        "lose weight",
			Status:        "active",
		}
		resGoal, _ := goal.New(db).Create(mockGoal)
		fmt.Println(resGoal)

		mocFood := []entities.Food{
			{
				Food_uid: f1.Food_uid,
			},
			{
				Food_uid: f1.Food_uid,
			},
		}

		mocMenuUser := entities.Menu{
			User_uid:      resUser.User_uid,
			Menu_category: "breakfast",
			Created_by:    "user",
		}
		_, errMenuUser := repo.CreateMenuUser(mocFood, mocMenuUser)
		fmt.Println(errMenuUser)

		assert.NotNil(t, errMenuUser)

	})

}

func TestGetAll(t *testing.T) {
	config := configs.GetConfig()
	db := utils.InitDB(config)
	repo := New(db)

	t.Run("get all menu category tidak = kosong dan createdBy = kosong", func(t *testing.T) {
		db.Migrator().DropTable(&entities.User{}, &entities.User_history{}, &entities.Detail_menu{}, &entities.Food{}, &entities.Menu{})
		db.AutoMigrate(&entities.User{})
		db.AutoMigrate(&entities.Food{})
		db.AutoMigrate(&entities.Menu{})
		db.AutoMigrate(&entities.Detail_menu{})
		db.AutoMigrate(&entities.User_history{})
		mocFood1 := entities.Food{
			Name:          "makanan",
			Calories:      100,
			Energy:        200,
			Carbohidrate:  300,
			Protein:       400,
			Food_category: "snack",
			Unit:          "ons",
			Unit_value:    1,
		}
		f1, _ := food.New(db).Create(mocFood1)
		mocAdmin := entities.User{
			Name:     "test",
			Email:    "testadmin@mail.com",
			Password: "test",
		}

		resAdmin, _ := admin.New(db).Register(mocAdmin)

		mocFood := []entities.Food{
			{
				Food_uid: f1.Food_uid,
			},
			{
				Food_uid: f1.Food_uid,
			},
		}
		mocMenu := entities.Menu{
			User_uid:      resAdmin.User_uid,
			Menu_category: "zzz",
			Created_by:    "admin",
		}
		_, errAdmin := repo.CreateMenuAdmin(mocFood, mocMenu)
		fmt.Println(errAdmin)
		_, errGet := repo.GetAllMenu("breakfast", "")

		assert.Nil(t, errGet)
	})

	t.Run("fail get all menu category tidak = kosong dan createdBy = kosong", func(t *testing.T) {
		db.Migrator().DropTable(&entities.User{}, &entities.User_history{}, &entities.Detail_menu{}, &entities.Food{}, &entities.Menu{})
		db.AutoMigrate(&entities.User{})
		db.AutoMigrate(&entities.Food{})
		db.AutoMigrate(&entities.Menu{})
		db.Migrator().DropColumn(&entities.Menu{}, "created_at")

		_, errGet := repo.GetAllMenu("breakfast", "")

		assert.NotNil(t, errGet)
	})

	t.Run("get all menu createdBy != kosong dan category = kosong ", func(t *testing.T) {
		db.Migrator().DropTable(&entities.User{}, &entities.User_history{}, &entities.Detail_menu{}, &entities.Food{}, &entities.Menu{})
		db.AutoMigrate(&entities.User{})
		db.AutoMigrate(&entities.Food{})
		db.AutoMigrate(&entities.Menu{})
		db.AutoMigrate(&entities.Detail_menu{})
		db.AutoMigrate(&entities.User_history{})
		mocFood1 := entities.Food{
			Name:          "makanan",
			Calories:      100,
			Energy:        200,
			Carbohidrate:  300,
			Protein:       400,
			Food_category: "snack",
			Unit:          "ons",
			Unit_value:    1,
		}
		f1, _ := food.New(db).Create(mocFood1)
		mocAdmin := entities.User{
			Name:     "test",
			Email:    "testadmin@mail.com",
			Password: "test",
		}

		resAdmin, _ := admin.New(db).Register(mocAdmin)

		mocFood := []entities.Food{
			{
				Food_uid: f1.Food_uid,
			},
			{
				Food_uid: f1.Food_uid,
			},
		}
		mocMenu := entities.Menu{
			User_uid:      resAdmin.User_uid,
			Menu_category: "zzz",
			Created_by:    "admin",
		}
		_, errAdmin := repo.CreateMenuAdmin(mocFood, mocMenu)
		fmt.Println(errAdmin)
		_, errGet := repo.GetAllMenu("", "admin")

		assert.Nil(t, errGet)
	})

	t.Run("fail get all menu createdBy != kosong dan category = kosong", func(t *testing.T) {
		db.Migrator().DropTable(&entities.User{}, &entities.User_history{}, &entities.Detail_menu{}, &entities.Food{}, &entities.Menu{})
		db.AutoMigrate(&entities.User{})
		db.AutoMigrate(&entities.Food{})
		db.AutoMigrate(&entities.Menu{})
		db.Migrator().DropColumn(&entities.Menu{}, "created_by")

		_, errGet := repo.GetAllMenu("", "admin")

		assert.NotNil(t, errGet)
	})

	t.Run("get all menu createdBy != kosong dan category != kosong ", func(t *testing.T) {
		db.Migrator().DropTable(&entities.User{}, &entities.User_history{}, &entities.Detail_menu{}, &entities.Food{}, &entities.Menu{})
		db.AutoMigrate(&entities.User{})
		db.AutoMigrate(&entities.Food{})
		db.AutoMigrate(&entities.Menu{})
		db.AutoMigrate(&entities.Detail_menu{})
		db.AutoMigrate(&entities.User_history{})
		mocFood1 := entities.Food{
			Name:          "makanan",
			Calories:      100,
			Energy:        200,
			Carbohidrate:  300,
			Protein:       400,
			Food_category: "snack",
			Unit:          "ons",
			Unit_value:    1,
		}
		f1, _ := food.New(db).Create(mocFood1)
		mocAdmin := entities.User{
			Name:     "test",
			Email:    "testadmin@mail.com",
			Password: "test",
		}

		resAdmin, _ := admin.New(db).Register(mocAdmin)

		mocFood := []entities.Food{
			{
				Food_uid: f1.Food_uid,
			},
			{
				Food_uid: f1.Food_uid,
			},
		}
		mocMenu := entities.Menu{
			User_uid:      resAdmin.User_uid,
			Menu_category: "breakfast",
			Created_by:    "admin",
		}
		_, errAdmin := repo.CreateMenuAdmin(mocFood, mocMenu)
		fmt.Println(errAdmin)
		_, errGet := repo.GetAllMenu("breakfast", "admin")

		assert.Nil(t, errGet)
	})

	t.Run("fail get all menu createdBy != kosong dan category != kosong", func(t *testing.T) {
		db.Migrator().DropTable(&entities.User{}, &entities.User_history{}, &entities.Detail_menu{}, &entities.Food{}, &entities.Menu{})
		db.AutoMigrate(&entities.User{})
		db.AutoMigrate(&entities.Food{})
		db.AutoMigrate(&entities.Menu{})
		db.Migrator().DropColumn(&entities.Menu{}, "created_by")
		db.Migrator().DropColumn(&entities.Menu{}, "menu_category")

		_, errGet := repo.GetAllMenu("breakfast", "admin")

		assert.NotNil(t, errGet)
	})

	t.Run("get all menu createdBy = kosong dan category = kosong ", func(t *testing.T) {
		db.Migrator().DropTable(&entities.User{}, &entities.User_history{}, &entities.Detail_menu{}, &entities.Food{}, &entities.Menu{})
		db.AutoMigrate(&entities.User{})
		db.AutoMigrate(&entities.Food{})
		db.AutoMigrate(&entities.Menu{})
		db.AutoMigrate(&entities.Detail_menu{})
		db.AutoMigrate(&entities.User_history{})
		mocFood1 := entities.Food{
			Name:          "makanan",
			Calories:      100,
			Energy:        200,
			Carbohidrate:  300,
			Protein:       400,
			Food_category: "snack",
			Unit:          "ons",
			Unit_value:    1,
		}
		f1, _ := food.New(db).Create(mocFood1)
		mocAdmin := entities.User{
			Name:     "test",
			Email:    "testadmin@mail.com",
			Password: "test",
		}

		resAdmin, _ := admin.New(db).Register(mocAdmin)

		mocFood := []entities.Food{
			{
				Food_uid: f1.Food_uid,
			},
			{
				Food_uid: f1.Food_uid,
			},
		}
		mocMenu := entities.Menu{
			User_uid:      resAdmin.User_uid,
			Menu_category: "breakfast",
			Created_by:    "admin",
		}
		_, errAdmin := repo.CreateMenuAdmin(mocFood, mocMenu)
		fmt.Println(errAdmin)
		_, errGet := repo.GetAllMenu("", "")

		assert.Nil(t, errGet)
	})

	t.Run("fail get all menu createdBy = kosong dan category = kosong", func(t *testing.T) {
		db.Migrator().DropTable(&entities.User{}, &entities.User_history{}, &entities.Detail_menu{}, &entities.Food{}, &entities.Menu{})

		_, errGet := repo.GetAllMenu("", "")

		assert.NotNil(t, errGet)
	})

}

func TestGetRecommedBreakfast(t *testing.T) {
	config := configs.GetConfig()
	db := utils.InitDB(config)
	repo := New(db)

	t.Run("Success to get recommed breakfast (lose weight)", func(t *testing.T) {
		db.Migrator().DropTable(&entities.User{}, &entities.User_history{}, &entities.Detail_menu{}, &entities.Food{}, &entities.Menu{})
		db.AutoMigrate(&entities.User{})
		db.AutoMigrate(&entities.Food{})
		db.AutoMigrate(&entities.Menu{})
		db.AutoMigrate(&entities.Detail_menu{})
		db.AutoMigrate(&entities.User_history{})
		mocFood1 := entities.Food{
			Name:          "makanan",
			Calories:      100,
			Energy:        200,
			Carbohidrate:  300,
			Protein:       400,
			Food_category: "snack",
			Unit:          "ons",
			Unit_value:    1,
		}
		f1, _ := food.New(db).Create(mocFood1)
		mocAdmin := entities.User{
			Name:     "test",
			Email:    "testadmin@mail.com",
			Password: "test",
		}

		resAdmin, _ := admin.New(db).Register(mocAdmin)

		mocFood := []entities.Food{
			{
				Food_uid: f1.Food_uid,
			},
			{
				Food_uid: f1.Food_uid,
			},
		}
		mocMenu := entities.Menu{
			User_uid:      resAdmin.User_uid,
			Menu_category: "breakfast",
			Created_by:    "admin",
		}
		_, errAdmin := repo.CreateMenuAdmin(mocFood, mocMenu)
		fmt.Println(errAdmin)

		mocUser := entities.User{
			Name:     "arya",
			Email:    "arya@mail.com",
			Password: "arya",
			Gender:   "male",
		}
		resU, errU := user.New(db).Register(mocUser)
		if errU != nil {
			t.Fail()
		}

		mockGoal := entities.Goal{
			User_uid:      resU.User_uid,
			Height:        160,
			Weight:        50,
			Age:           25,
			Daily_active:  "quite active",
			Weight_target: 2,
			Range_time:    30,
			Target:        "lose weight",
			Status:        "active",
		}

		_, errG := goal.New(db).Create(mockGoal)
		if errG != nil {
			t.Fail()
		}

		_, _, _, err := repo.GetRecommendBreakfast(resU.User_uid)

		assert.Nil(t, err)
	})

	t.Run("Success to get recommed breakfast (gain weight)", func(t *testing.T) {
		db.Migrator().DropTable(&entities.User{}, &entities.User_history{}, &entities.Detail_menu{}, &entities.Food{}, &entities.Menu{})
		db.AutoMigrate(&entities.User{})
		db.AutoMigrate(&entities.Food{})
		db.AutoMigrate(&entities.Menu{})
		db.AutoMigrate(&entities.Detail_menu{})
		db.AutoMigrate(&entities.User_history{})
		mocFood1 := entities.Food{
			Name:          "makanan",
			Calories:      100,
			Energy:        200,
			Carbohidrate:  300,
			Protein:       400,
			Food_category: "snack",
			Unit:          "ons",
			Unit_value:    1,
		}
		f1, _ := food.New(db).Create(mocFood1)
		mocAdmin := entities.User{
			Name:     "test",
			Email:    "testadmin@mail.com",
			Password: "test",
		}

		resAdmin, _ := admin.New(db).Register(mocAdmin)

		mocFood := []entities.Food{
			{
				Food_uid: f1.Food_uid,
			},
			{
				Food_uid: f1.Food_uid,
			},
		}
		mocMenu := entities.Menu{
			User_uid:      resAdmin.User_uid,
			Menu_category: "breakfast",
			Created_by:    "admin",
		}
		_, errAdmin := repo.CreateMenuAdmin(mocFood, mocMenu)
		fmt.Println(errAdmin)

		mocUser := entities.User{
			Name:     "arya",
			Email:    "arya@mail.com",
			Password: "arya",
			Gender:   "male",
		}
		resU, errU := user.New(db).Register(mocUser)
		if errU != nil {
			t.Fail()
		}

		mockGoal := entities.Goal{
			User_uid:      resU.User_uid,
			Height:        160,
			Weight:        50,
			Age:           25,
			Daily_active:  "active",
			Weight_target: 2,
			Range_time:    30,
			Target:        "gain weight",
			Status:        "active",
		}

		_, errG := goal.New(db).Create(mockGoal)
		if errG != nil {
			t.Fail()
		}

		_, _, _, err := repo.GetRecommendBreakfast(resU.User_uid)

		assert.Nil(t, err)
	})

	t.Run("Failed to get recommed breakfast (no goal)", func(t *testing.T) {
		db.Migrator().DropTable(&entities.User{}, &entities.User_history{}, &entities.Detail_menu{}, &entities.Food{}, &entities.Menu{})
		db.AutoMigrate(&entities.User{})
		db.AutoMigrate(&entities.Food{})
		db.AutoMigrate(&entities.Menu{})
		db.AutoMigrate(&entities.Detail_menu{})
		db.AutoMigrate(&entities.User_history{})
		mocFood1 := entities.Food{
			Name:          "makanan",
			Calories:      100,
			Energy:        200,
			Carbohidrate:  300,
			Protein:       400,
			Food_category: "snack",
			Unit:          "ons",
			Unit_value:    1,
		}
		f1, _ := food.New(db).Create(mocFood1)
		mocAdmin := entities.User{
			Name:     "test",
			Email:    "testadmin@mail.com",
			Password: "test",
		}

		resAdmin, _ := admin.New(db).Register(mocAdmin)

		mocFood := []entities.Food{
			{
				Food_uid: f1.Food_uid,
			},
			{
				Food_uid: f1.Food_uid,
			},
		}
		mocMenu := entities.Menu{
			User_uid:      resAdmin.User_uid,
			Menu_category: "breakfast",
			Created_by:    "admin",
		}
		_, errAdmin := repo.CreateMenuAdmin(mocFood, mocMenu)
		fmt.Println(errAdmin)

		mocUser := entities.User{
			Name:     "arya",
			Email:    "arya@mail.com",
			Password: "arya",
			Gender:   "male",
		}
		resU, errU := user.New(db).Register(mocUser)
		if errU != nil {
			t.Fail()
		}

		_, _, _, err := repo.GetRecommendBreakfast(resU.User_uid)

		assert.NotNil(t, err)
	})

	t.Run("Failed to get recommed breakfast (no menus)", func(t *testing.T) {
		db.Migrator().DropTable(&entities.User{}, &entities.User_history{}, &entities.Detail_menu{}, &entities.Food{}, &entities.Menu{})
		db.AutoMigrate(&entities.User{})
		db.AutoMigrate(&entities.Food{})
		db.AutoMigrate(&entities.Menu{})
		db.AutoMigrate(&entities.Detail_menu{})
		db.AutoMigrate(&entities.User_history{})

		mocFood1 := entities.Food{
			Name:          "makanan",
			Calories:      100,
			Energy:        200,
			Carbohidrate:  300,
			Protein:       400,
			Food_category: "snack",
			Unit:          "ons",
			Unit_value:    1,
		}
		f1, _ := food.New(db).Create(mocFood1)
		mocAdmin := entities.User{
			Name:     "test",
			Email:    "testadmin@mail.com",
			Password: "test",
		}

		resAdmin, _ := admin.New(db).Register(mocAdmin)

		mocFood := []entities.Food{
			{
				Food_uid: f1.Food_uid,
			},
			{
				Food_uid: f1.Food_uid,
			},
		}
		mocMenu := entities.Menu{
			User_uid:      resAdmin.User_uid,
			Menu_category: "breakfast",
			Created_by:    "admin",
		}
		_, errAdmin := repo.CreateMenuAdmin(mocFood, mocMenu)
		log.Info(errAdmin)

		mocUser := entities.User{
			Name:     "arya",
			Email:    "arya@mail.com",
			Password: "arya",
			Gender:   "female",
		}
		resU, errU := user.New(db).Register(mocUser)
		if errU != nil {
			t.Fail()
		}

		_, row, breakfast, err := repo.GetRecommendBreakfast(resU.User_uid)

		assert.NotNil(t, err)
		assert.Equal(t, int64(0), row)
		assert.Equal(t, 0, breakfast)
	})

}

func TestGetRecommedLunch(t *testing.T) {
	config := configs.GetConfig()
	db := utils.InitDB(config)
	repo := New(db)

	t.Run("Success to get recommed Lunch (lose weight)", func(t *testing.T) {
		db.Migrator().DropTable(&entities.User{}, &entities.User_history{}, &entities.Detail_menu{}, &entities.Food{}, &entities.Menu{})
		db.AutoMigrate(&entities.User{})
		db.AutoMigrate(&entities.Food{})
		db.AutoMigrate(&entities.Menu{})
		db.AutoMigrate(&entities.Detail_menu{})
		db.AutoMigrate(&entities.User_history{})
		mocFood1 := entities.Food{
			Name:          "makanan",
			Calories:      100,
			Energy:        200,
			Carbohidrate:  300,
			Protein:       400,
			Food_category: "snack",
			Unit:          "ons",
			Unit_value:    1,
		}
		f1, _ := food.New(db).Create(mocFood1)
		mocAdmin := entities.User{
			Name:     "test",
			Email:    "testadmin@mail.com",
			Password: "test",
		}

		resAdmin, _ := admin.New(db).Register(mocAdmin)

		mocFood := []entities.Food{
			{
				Food_uid: f1.Food_uid,
			},
			{
				Food_uid: f1.Food_uid,
			},
		}
		mocMenu := entities.Menu{
			User_uid:      resAdmin.User_uid,
			Menu_category: "lunch",
			Created_by:    "admin",
		}
		_, errAdmin := repo.CreateMenuAdmin(mocFood, mocMenu)
		fmt.Println(errAdmin)

		mocUser := entities.User{
			Name:     "arya",
			Email:    "arya@mail.com",
			Password: "arya",
			Gender:   "male",
		}
		resU, errU := user.New(db).Register(mocUser)
		if errU != nil {
			t.Fail()
		}

		mockGoal := entities.Goal{
			User_uid:      resU.User_uid,
			Height:        160,
			Weight:        50,
			Age:           25,
			Daily_active:  "little active",
			Weight_target: 2,
			Range_time:    30,
			Target:        "lose weight",
			Status:        "active",
		}

		_, errG := goal.New(db).Create(mockGoal)
		if errG != nil {
			t.Fail()
		}

		_, _, _, err := repo.GetRecommendLunch(resU.User_uid)

		assert.Nil(t, err)
	})

	t.Run("Success to get recommed Lunch (gain weight)", func(t *testing.T) {
		db.Migrator().DropTable(&entities.User{}, &entities.User_history{}, &entities.Detail_menu{}, &entities.Food{}, &entities.Menu{})
		db.AutoMigrate(&entities.User{})
		db.AutoMigrate(&entities.Food{})
		db.AutoMigrate(&entities.Menu{})
		db.AutoMigrate(&entities.Detail_menu{})
		db.AutoMigrate(&entities.User_history{})
		mocFood1 := entities.Food{
			Name:          "makanan",
			Calories:      100,
			Energy:        200,
			Carbohidrate:  300,
			Protein:       400,
			Food_category: "snack",
			Unit:          "ons",
			Unit_value:    1,
		}
		f1, _ := food.New(db).Create(mocFood1)
		mocAdmin := entities.User{
			Name:     "test",
			Email:    "testadmin@mail.com",
			Password: "test",
		}

		resAdmin, _ := admin.New(db).Register(mocAdmin)

		mocFood := []entities.Food{
			{
				Food_uid: f1.Food_uid,
			},
			{
				Food_uid: f1.Food_uid,
			},
		}
		mocMenu := entities.Menu{
			User_uid:      resAdmin.User_uid,
			Menu_category: "lunch",
			Created_by:    "admin",
		}
		_, errAdmin := repo.CreateMenuAdmin(mocFood, mocMenu)
		fmt.Println(errAdmin)

		mocUser := entities.User{
			Name:     "arya",
			Email:    "arya@mail.com",
			Password: "arya",
			Gender:   "male",
		}
		resU, errU := user.New(db).Register(mocUser)
		if errU != nil {
			t.Fail()
		}

		mockGoal := entities.Goal{
			User_uid:      resU.User_uid,
			Height:        160,
			Weight:        50,
			Age:           25,
			Daily_active:  "not active",
			Weight_target: 2,
			Range_time:    30,
			Target:        "gain weight",
			Status:        "active",
		}

		_, errG := goal.New(db).Create(mockGoal)
		if errG != nil {
			t.Fail()
		}

		_, _, _, err := repo.GetRecommendLunch(resU.User_uid)

		assert.Nil(t, err)
	})

	t.Run("Failed to get recommed Lunch (no goal)", func(t *testing.T) {
		db.Migrator().DropTable(&entities.User{}, &entities.User_history{}, &entities.Detail_menu{}, &entities.Food{}, &entities.Menu{})
		db.AutoMigrate(&entities.User{})
		db.AutoMigrate(&entities.Food{})
		db.AutoMigrate(&entities.Menu{})
		db.AutoMigrate(&entities.Detail_menu{})
		db.AutoMigrate(&entities.User_history{})
		mocFood1 := entities.Food{
			Name:          "makanan",
			Calories:      100,
			Energy:        200,
			Carbohidrate:  300,
			Protein:       400,
			Food_category: "snack",
			Unit:          "ons",
			Unit_value:    1,
		}
		f1, _ := food.New(db).Create(mocFood1)
		mocAdmin := entities.User{
			Name:     "test",
			Email:    "testadmin@mail.com",
			Password: "test",
		}

		resAdmin, _ := admin.New(db).Register(mocAdmin)

		mocFood := []entities.Food{
			{
				Food_uid: f1.Food_uid,
			},
			{
				Food_uid: f1.Food_uid,
			},
		}
		mocMenu := entities.Menu{
			User_uid:      resAdmin.User_uid,
			Menu_category: "lunch",
			Created_by:    "admin",
		}
		_, errAdmin := repo.CreateMenuAdmin(mocFood, mocMenu)
		fmt.Println(errAdmin)

		mocUser := entities.User{
			Name:     "arya",
			Email:    "arya@mail.com",
			Password: "arya",
			Gender:   "male",
		}
		resU, errU := user.New(db).Register(mocUser)
		if errU != nil {
			t.Fail()
		}

		_, _, _, err := repo.GetRecommendLunch(resU.User_uid)

		assert.NotNil(t, err)
	})
}

func TestGetRecommedDinner(t *testing.T) {
	config := configs.GetConfig()
	db := utils.InitDB(config)
	repo := New(db)

	t.Run("Success to get recommed Dinner (lose weight)", func(t *testing.T) {
		db.Migrator().DropTable(&entities.User{}, &entities.User_history{}, &entities.Detail_menu{}, &entities.Food{}, &entities.Menu{})
		db.AutoMigrate(&entities.User{})
		db.AutoMigrate(&entities.Food{})
		db.AutoMigrate(&entities.Menu{})
		db.AutoMigrate(&entities.Detail_menu{})
		db.AutoMigrate(&entities.User_history{})
		mocFood1 := entities.Food{
			Name:          "makanan",
			Calories:      100,
			Energy:        200,
			Carbohidrate:  300,
			Protein:       400,
			Food_category: "snack",
			Unit:          "ons",
			Unit_value:    1,
		}
		f1, _ := food.New(db).Create(mocFood1)
		mocAdmin := entities.User{
			Name:     "test",
			Email:    "testadmin@mail.com",
			Password: "test",
		}

		resAdmin, _ := admin.New(db).Register(mocAdmin)

		mocFood := []entities.Food{
			{
				Food_uid: f1.Food_uid,
			},
			{
				Food_uid: f1.Food_uid,
			},
		}
		mocMenu := entities.Menu{
			User_uid:      resAdmin.User_uid,
			Menu_category: "dinner",
			Created_by:    "admin",
		}
		_, errAdmin := repo.CreateMenuAdmin(mocFood, mocMenu)
		fmt.Println(errAdmin)

		mocUser := entities.User{
			Name:     "arya",
			Email:    "arya@mail.com",
			Password: "arya",
			Gender:   "male",
		}
		resU, errU := user.New(db).Register(mocUser)
		if errU != nil {
			t.Fail()
		}

		mockGoal := entities.Goal{
			User_uid:      resU.User_uid,
			Height:        160,
			Weight:        50,
			Age:           25,
			Daily_active:  "little active",
			Weight_target: 2,
			Range_time:    30,
			Target:        "lose weight",
			Status:        "active",
		}

		_, errG := goal.New(db).Create(mockGoal)
		if errG != nil {
			t.Fail()
		}

		_, _, _, err := repo.GetRecommendDinner(resU.User_uid)

		assert.Nil(t, err)
	})

	t.Run("Success to get recommed Dinner (gain weight)", func(t *testing.T) {
		db.Migrator().DropTable(&entities.User{}, &entities.User_history{}, &entities.Detail_menu{}, &entities.Food{}, &entities.Menu{})
		db.AutoMigrate(&entities.User{})
		db.AutoMigrate(&entities.Food{})
		db.AutoMigrate(&entities.Menu{})
		db.AutoMigrate(&entities.Detail_menu{})
		db.AutoMigrate(&entities.User_history{})
		mocFood1 := entities.Food{
			Name:          "makanan",
			Calories:      100,
			Energy:        200,
			Carbohidrate:  300,
			Protein:       400,
			Food_category: "snack",
			Unit:          "ons",
			Unit_value:    1,
		}
		f1, _ := food.New(db).Create(mocFood1)
		mocAdmin := entities.User{
			Name:     "test",
			Email:    "testadmin@mail.com",
			Password: "test",
		}

		resAdmin, _ := admin.New(db).Register(mocAdmin)

		mocFood := []entities.Food{
			{
				Food_uid: f1.Food_uid,
			},
			{
				Food_uid: f1.Food_uid,
			},
		}
		mocMenu := entities.Menu{
			User_uid:      resAdmin.User_uid,
			Menu_category: "dinner",
			Created_by:    "admin",
		}
		_, errAdmin := repo.CreateMenuAdmin(mocFood, mocMenu)
		fmt.Println(errAdmin)

		mocUser := entities.User{
			Name:     "arya",
			Email:    "arya@mail.com",
			Password: "arya",
			Gender:   "male",
		}
		resU, errU := user.New(db).Register(mocUser)
		if errU != nil {
			t.Fail()
		}

		mockGoal := entities.Goal{
			User_uid:      resU.User_uid,
			Height:        160,
			Weight:        50,
			Age:           25,
			Daily_active:  "not active",
			Weight_target: 2,
			Range_time:    30,
			Target:        "gain weight",
			Status:        "active",
		}

		_, errG := goal.New(db).Create(mockGoal)
		if errG != nil {
			t.Fail()
		}

		_, _, _, err := repo.GetRecommendDinner(resU.User_uid)

		assert.Nil(t, err)
	})

	t.Run("Failed to get recommed Dinner (no goal)", func(t *testing.T) {
		db.Migrator().DropTable(&entities.User{}, &entities.User_history{}, &entities.Detail_menu{}, &entities.Food{}, &entities.Menu{})
		db.AutoMigrate(&entities.User{})
		db.AutoMigrate(&entities.Food{})
		db.AutoMigrate(&entities.Menu{})
		db.AutoMigrate(&entities.Detail_menu{})
		db.AutoMigrate(&entities.User_history{})
		mocFood1 := entities.Food{
			Name:          "makanan",
			Calories:      100,
			Energy:        200,
			Carbohidrate:  300,
			Protein:       400,
			Food_category: "snack",
			Unit:          "ons",
			Unit_value:    1,
		}
		f1, _ := food.New(db).Create(mocFood1)
		mocAdmin := entities.User{
			Name:     "test",
			Email:    "testadmin@mail.com",
			Password: "test",
		}

		resAdmin, _ := admin.New(db).Register(mocAdmin)

		mocFood := []entities.Food{
			{
				Food_uid: f1.Food_uid,
			},
			{
				Food_uid: f1.Food_uid,
			},
		}
		mocMenu := entities.Menu{
			User_uid:      resAdmin.User_uid,
			Menu_category: "dinner",
			Created_by:    "admin",
		}
		_, errAdmin := repo.CreateMenuAdmin(mocFood, mocMenu)
		fmt.Println(errAdmin)

		mocUser := entities.User{
			Name:     "arya",
			Email:    "arya@mail.com",
			Password: "arya",
			Gender:   "male",
		}
		resU, errU := user.New(db).Register(mocUser)
		if errU != nil {
			t.Fail()
		}

		_, _, _, err := repo.GetRecommendDinner(resU.User_uid)

		assert.NotNil(t, err)
	})
}

func TestGetRecommedOverTime(t *testing.T) {
	config := configs.GetConfig()
	db := utils.InitDB(config)
	repo := New(db)

	t.Run("Success to get recommed OverTime (lose weight)", func(t *testing.T) {
		db.Migrator().DropTable(&entities.User{}, &entities.User_history{}, &entities.Detail_menu{}, &entities.Food{}, &entities.Menu{})
		db.AutoMigrate(&entities.User{})
		db.AutoMigrate(&entities.Food{})
		db.AutoMigrate(&entities.Menu{})
		db.AutoMigrate(&entities.Detail_menu{})
		db.AutoMigrate(&entities.User_history{})
		mocFood1 := entities.Food{
			Name:          "makanan",
			Calories:      100,
			Energy:        200,
			Carbohidrate:  300,
			Protein:       400,
			Food_category: "snack",
			Unit:          "ons",
			Unit_value:    1,
		}
		f1, _ := food.New(db).Create(mocFood1)
		mocAdmin := entities.User{
			Name:     "test",
			Email:    "testadmin@mail.com",
			Password: "test",
		}

		resAdmin, _ := admin.New(db).Register(mocAdmin)

		mocFood := []entities.Food{
			{
				Food_uid: f1.Food_uid,
			},
			{
				Food_uid: f1.Food_uid,
			},
		}
		mocMenu := entities.Menu{
			User_uid:      resAdmin.User_uid,
			Menu_category: "overtime",
			Created_by:    "admin",
		}
		_, errAdmin := repo.CreateMenuAdmin(mocFood, mocMenu)
		fmt.Println(errAdmin)

		mocUser := entities.User{
			Name:     "arya",
			Email:    "arya@mail.com",
			Password: "arya",
			Gender:   "male",
		}
		resU, errU := user.New(db).Register(mocUser)
		if errU != nil {
			t.Fail()
		}

		mockGoal := entities.Goal{
			User_uid:      resU.User_uid,
			Height:        160,
			Weight:        50,
			Age:           25,
			Daily_active:  "little active",
			Weight_target: 2,
			Range_time:    30,
			Target:        "lose weight",
			Status:        "active",
		}

		_, errG := goal.New(db).Create(mockGoal)
		if errG != nil {
			t.Fail()
		}

		_, _, _, err := repo.GetRecommendOverTime(resU.User_uid)

		assert.Nil(t, err)
	})

	t.Run("Success to get recommed OverTime (gain weight)", func(t *testing.T) {
		db.Migrator().DropTable(&entities.User{}, &entities.User_history{}, &entities.Detail_menu{}, &entities.Food{}, &entities.Menu{})
		db.AutoMigrate(&entities.User{})
		db.AutoMigrate(&entities.Food{})
		db.AutoMigrate(&entities.Menu{})
		db.AutoMigrate(&entities.Detail_menu{})
		db.AutoMigrate(&entities.User_history{})
		mocFood1 := entities.Food{
			Name:          "makanan",
			Calories:      100,
			Energy:        200,
			Carbohidrate:  300,
			Protein:       400,
			Food_category: "snack",
			Unit:          "ons",
			Unit_value:    1,
		}
		f1, _ := food.New(db).Create(mocFood1)
		mocAdmin := entities.User{
			Name:     "test",
			Email:    "testadmin@mail.com",
			Password: "test",
		}

		resAdmin, _ := admin.New(db).Register(mocAdmin)

		mocFood := []entities.Food{
			{
				Food_uid: f1.Food_uid,
			},
			{
				Food_uid: f1.Food_uid,
			},
		}
		mocMenu := entities.Menu{
			User_uid:      resAdmin.User_uid,
			Menu_category: "overtime",
			Created_by:    "admin",
		}
		_, errAdmin := repo.CreateMenuAdmin(mocFood, mocMenu)
		fmt.Println(errAdmin)

		mocUser := entities.User{
			Name:     "arya",
			Email:    "arya@mail.com",
			Password: "arya",
			Gender:   "male",
		}
		resU, errU := user.New(db).Register(mocUser)
		if errU != nil {
			t.Fail()
		}

		mockGoal := entities.Goal{
			User_uid:      resU.User_uid,
			Height:        160,
			Weight:        50,
			Age:           25,
			Daily_active:  "not active",
			Weight_target: 2,
			Range_time:    30,
			Target:        "gain weight",
			Status:        "active",
		}

		_, errG := goal.New(db).Create(mockGoal)
		if errG != nil {
			t.Fail()
		}

		_, _, _, err := repo.GetRecommendOverTime(resU.User_uid)

		assert.Nil(t, err)
	})

	t.Run("Failed to get recommed OverTime (no goal)", func(t *testing.T) {
		db.Migrator().DropTable(&entities.User{}, &entities.User_history{}, &entities.Detail_menu{}, &entities.Food{}, &entities.Menu{})
		db.AutoMigrate(&entities.User{})
		db.AutoMigrate(&entities.Food{})
		db.AutoMigrate(&entities.Menu{})
		db.AutoMigrate(&entities.Detail_menu{})
		db.AutoMigrate(&entities.User_history{})
		mocFood1 := entities.Food{
			Name:          "makanan",
			Calories:      100,
			Energy:        200,
			Carbohidrate:  300,
			Protein:       400,
			Food_category: "snack",
			Unit:          "ons",
			Unit_value:    1,
		}
		f1, _ := food.New(db).Create(mocFood1)
		mocAdmin := entities.User{
			Name:     "test",
			Email:    "testadmin@mail.com",
			Password: "test",
		}

		resAdmin, _ := admin.New(db).Register(mocAdmin)

		mocFood := []entities.Food{
			{
				Food_uid: f1.Food_uid,
			},
			{
				Food_uid: f1.Food_uid,
			},
		}
		mocMenu := entities.Menu{
			User_uid:      resAdmin.User_uid,
			Menu_category: "overtime",
			Created_by:    "admin",
		}
		_, errAdmin := repo.CreateMenuAdmin(mocFood, mocMenu)
		fmt.Println(errAdmin)

		mocUser := entities.User{
			Name:     "arya",
			Email:    "arya@mail.com",
			Password: "arya",
			Gender:   "male",
		}
		resU, errU := user.New(db).Register(mocUser)
		if errU != nil {
			t.Fail()
		}

		_, _, _, err := repo.GetRecommendOverTime(resU.User_uid)

		assert.NotNil(t, err)
	})
}

func TestUpdate(t *testing.T) {
	config := configs.GetConfig()
	db := utils.InitDB(config)
	repo := New(db)

	t.Run("success Update", func(t *testing.T) {
		db.Migrator().DropTable(&entities.User{}, &entities.User_history{}, &entities.Detail_menu{}, &entities.Food{}, &entities.Menu{})
		db.AutoMigrate(&entities.User{})
		db.AutoMigrate(&entities.Food{})
		db.AutoMigrate(&entities.Menu{})
		db.AutoMigrate(&entities.Detail_menu{})
		db.AutoMigrate(&entities.User_history{})
		mocFood1 := entities.Food{
			Name:          "makanan",
			Calories:      100,
			Energy:        200,
			Carbohidrate:  300,
			Protein:       400,
			Food_category: "snack",
			Unit:          "ons",
			Unit_value:    1,
		}
		f1, _ := food.New(db).Create(mocFood1)
		mocAdmin := entities.User{
			Name:     "test",
			Email:    "testadmin@mail.com",
			Password: "test",
		}

		resAdmin, _ := admin.New(db).Register(mocAdmin)

		mocFood := []entities.Food{
			{
				Food_uid: f1.Food_uid,
			},
			{
				Food_uid: f1.Food_uid,
			},
		}
		mocMenu := entities.Menu{
			User_uid:      resAdmin.User_uid,
			Menu_category: "dinner",
			Created_by:    "admin",
		}
		resMenuAdmin, errAdmin := repo.CreateMenuAdmin(mocFood, mocMenu)
		fmt.Println(errAdmin)
		mocUpdateMenu := entities.Menu{
			User_uid:      resAdmin.User_uid,
			Menu_category: "breakfast",
			Created_by:    "admin",
		}
		_, errGet := repo.Update(resMenuAdmin.Menu_uid, mocFood, mocUpdateMenu)

		assert.Nil(t, errGet)
	})

	t.Run("fail Update", func(t *testing.T) {
		db.Migrator().DropTable(&entities.User{}, &entities.User_history{}, &entities.Detail_menu{}, &entities.Food{}, &entities.Menu{})
		db.AutoMigrate(&entities.User{})
		db.AutoMigrate(&entities.Food{})
		db.AutoMigrate(&entities.Menu{})
		db.AutoMigrate(&entities.Detail_menu{})
		db.AutoMigrate(&entities.User_history{})
		mocFood1 := entities.Food{
			Name:          "makanan",
			Calories:      100,
			Energy:        200,
			Carbohidrate:  300,
			Protein:       400,
			Food_category: "snack",
			Unit:          "ons",
			Unit_value:    1,
		}
		f1, _ := food.New(db).Create(mocFood1)
		mocAdmin := entities.User{
			Name:     "test",
			Email:    "testadmin@mail.com",
			Password: "test",
		}

		resAdmin, _ := admin.New(db).Register(mocAdmin)

		mocFood := []entities.Food{
			{
				Food_uid: f1.Food_uid,
			},
			{
				Food_uid: f1.Food_uid,
			},
		}
		mocMenu := entities.Menu{
			User_uid:      resAdmin.User_uid,
			Menu_category: "dinner",
			Created_by:    "admin",
		}
		_, errAdmin := repo.CreateMenuAdmin(mocFood, mocMenu)
		fmt.Println(errAdmin)
		mocUpdateMenu := entities.Menu{
			User_uid:      resAdmin.User_uid,
			Menu_category: "breakfast",
			Created_by:    "admin",
		}
		db.Migrator().DropColumn(&entities.Menu{}, "menu_uid")

		_, errGet := repo.Update("dfdsjs", mocFood, mocUpdateMenu)

		assert.NotNil(t, errGet)
	})

	t.Run("failed create new menu", func(t *testing.T) {
		db.Migrator().DropTable(&entities.User{}, &entities.User_history{}, &entities.Detail_menu{}, &entities.Food{}, &entities.Menu{})
		db.AutoMigrate(&entities.User{})
		db.AutoMigrate(&entities.Food{})
		db.AutoMigrate(&entities.Menu{})
		db.AutoMigrate(&entities.Detail_menu{})
		db.AutoMigrate(&entities.User_history{})
		mocFood1 := entities.Food{
			Name:          "makanan",
			Calories:      100,
			Energy:        200,
			Carbohidrate:  300,
			Protein:       400,
			Food_category: "snack",
			Unit:          "ons",
			Unit_value:    1,
		}
		f1, _ := food.New(db).Create(mocFood1)
		mocAdmin := entities.User{
			Name:     "test",
			Email:    "testadmin@mail.com",
			Password: "test",
		}

		resAdmin, _ := admin.New(db).Register(mocAdmin)

		mocFood := []entities.Food{
			{
				Food_uid: f1.Food_uid,
			},
			{
				Food_uid: f1.Food_uid,
			},
		}
		mocMenu := entities.Menu{
			User_uid:      resAdmin.User_uid,
			Menu_category: "dinner",
			Created_by:    "admin",
		}
		resMenuAdmin, errAdmin := repo.CreateMenuAdmin(mocFood, mocMenu)
		fmt.Println(errAdmin)
		mocUpdateMenu := entities.Menu{
			User_uid:      resAdmin.User_uid,
			Menu_category: "saa",
			Created_by:    "admin",
		}
		_, errGet := repo.Update(resMenuAdmin.Menu_uid, mocFood, mocUpdateMenu)

		assert.NotNil(t, errGet)
	})

	t.Run("failed create create new detail or get food data", func(t *testing.T) {
		db.Migrator().DropTable(&entities.User{}, &entities.User_history{}, &entities.Detail_menu{}, &entities.Food{}, &entities.Menu{})
		db.AutoMigrate(&entities.User{})
		db.AutoMigrate(&entities.Food{})
		db.AutoMigrate(&entities.Menu{})
		db.AutoMigrate(&entities.Detail_menu{})
		db.AutoMigrate(&entities.User_history{})
		mocFood1 := entities.Food{
			Name:          "makanan",
			Calories:      100,
			Energy:        200,
			Carbohidrate:  300,
			Protein:       400,
			Food_category: "snack",
			Unit:          "ons",
			Unit_value:    1,
		}
		f1, _ := food.New(db).Create(mocFood1)
		mocAdmin := entities.User{
			Name:     "test",
			Email:    "testadmin@mail.com",
			Password: "test",
		}

		resAdmin, _ := admin.New(db).Register(mocAdmin)

		mocFood := []entities.Food{
			{
				Food_uid: f1.Food_uid,
			},
			{
				Food_uid: f1.Food_uid,
			},
		}
		mocFoodupdate := []entities.Food{
			{
				Food_uid: "dsjds",
			},
			{
				Food_uid: f1.Food_uid,
			},
		}
		mocMenu := entities.Menu{
			User_uid:      resAdmin.User_uid,
			Menu_category: "dinner",
			Created_by:    "admin",
		}
		resMenuAdmin, errAdmin := repo.CreateMenuAdmin(mocFood, mocMenu)
		fmt.Println(errAdmin)
		mocUpdateMenu := entities.Menu{
			User_uid:      resAdmin.User_uid,
			Menu_category: "breakfast",
			Created_by:    "admin",
		}
		_, errGet := repo.Update(resMenuAdmin.Menu_uid, mocFoodupdate, mocUpdateMenu)

		assert.NotNil(t, errGet)
	})

	t.Run("failed update menu total calories", func(t *testing.T) {
		db.Migrator().DropTable(&entities.User{}, &entities.User_history{}, &entities.Detail_menu{}, &entities.Food{}, &entities.Menu{})
		db.AutoMigrate(&entities.User{})
		db.AutoMigrate(&entities.Food{})
		db.AutoMigrate(&entities.Menu{})
		db.AutoMigrate(&entities.Detail_menu{})
		db.AutoMigrate(&entities.User_history{})
		mocFood1 := entities.Food{
			Name:          "makanan",
			Calories:      100,
			Energy:        200,
			Carbohidrate:  300,
			Protein:       400,
			Food_category: "snack",
			Unit:          "ons",
			Unit_value:    1,
		}
		mocFood2 := entities.Food{
			Name:          "makanan",
			Calories:      2147483646,
			Energy:        200,
			Carbohidrate:  300,
			Protein:       400,
			Food_category: "snack",
			Unit:          "ons",
			Unit_value:    1,
		}
		f1, _ := food.New(db).Create(mocFood1)
		f2, _ := food.New(db).Create(mocFood2)
		mocAdmin := entities.User{
			Name:     "test",
			Email:    "testadmin@mail.com",
			Password: "test",
		}

		resAdmin, _ := admin.New(db).Register(mocAdmin)

		mocFood := []entities.Food{
			{
				Food_uid: f1.Food_uid,
			},
			{
				Food_uid: f1.Food_uid,
			},
		}
		mocFoodupdate := []entities.Food{
			{
				Food_uid: f2.Food_uid,
			},
			{
				Food_uid: f2.Food_uid,
			},
		}
		mocMenu := entities.Menu{
			User_uid:      resAdmin.User_uid,
			Menu_category: "dinner",
			Created_by:    "admin",
		}
		resMenuAdmin, errAdmin := repo.CreateMenuAdmin(mocFood, mocMenu)
		fmt.Println(errAdmin)
		mocUpdateMenu := entities.Menu{
			User_uid:      resAdmin.User_uid,
			Menu_category: "breakfast",
			Created_by:    "admin",
		}
		_, errGet := repo.Update(resMenuAdmin.Menu_uid, mocFoodupdate, mocUpdateMenu)

		assert.NotNil(t, errGet)
	})

}

func TestDelete(t *testing.T) {
	config := configs.GetConfig()
	db := utils.InitDB(config)
	repo := New(db)

	t.Run("success delete", func(t *testing.T) {
		db.Migrator().DropTable(&entities.User{}, &entities.User_history{}, &entities.Detail_menu{}, &entities.Food{}, &entities.Menu{})
		db.AutoMigrate(&entities.User{})
		db.AutoMigrate(&entities.Food{})
		db.AutoMigrate(&entities.Menu{})
		db.AutoMigrate(&entities.Detail_menu{})
		db.AutoMigrate(&entities.User_history{})
		mocFood1 := entities.Food{
			Name:          "makanan",
			Calories:      100,
			Energy:        200,
			Carbohidrate:  300,
			Protein:       400,
			Food_category: "snack",
			Unit:          "ons",
			Unit_value:    1,
		}
		f1, _ := food.New(db).Create(mocFood1)
		mocAdmin := entities.User{
			Name:     "test",
			Email:    "testadmin@mail.com",
			Password: "test",
		}

		resAdmin, _ := admin.New(db).Register(mocAdmin)

		mocFood := []entities.Food{
			{
				Food_uid: f1.Food_uid,
			},
			{
				Food_uid: f1.Food_uid,
			},
		}
		mocMenu := entities.Menu{
			User_uid:      resAdmin.User_uid,
			Menu_category: "zzz",
			Created_by:    "admin",
		}
		resMenuAdmin, errAdmin := repo.CreateMenuAdmin(mocFood, mocMenu)
		fmt.Println(errAdmin)
		errGet := repo.Delete(resMenuAdmin.Menu_uid)

		assert.Nil(t, errGet)
	})

	t.Run("failed delete detail menu", func(t *testing.T) {
		db.Migrator().DropTable(&entities.User{}, &entities.User_history{}, &entities.Detail_menu{}, &entities.Food{}, &entities.Menu{})

		errGet := repo.Delete("hfdkhfs")

		assert.NotNil(t, errGet)
	})

	t.Run("failed delete menu", func(t *testing.T) {
		db.Migrator().DropTable(&entities.User{}, &entities.User_history{}, &entities.Detail_menu{}, &entities.Food{}, &entities.Menu{})
		db.AutoMigrate(&entities.User{})
		db.AutoMigrate(&entities.Food{})
		db.AutoMigrate(&entities.Menu{})
		db.AutoMigrate(&entities.Detail_menu{})
		db.AutoMigrate(&entities.User_history{})
		mocFood1 := entities.Food{
			Name:          "makanan",
			Calories:      100,
			Energy:        200,
			Carbohidrate:  300,
			Protein:       400,
			Food_category: "snack",
			Unit:          "ons",
			Unit_value:    1,
		}
		f1, _ := food.New(db).Create(mocFood1)
		mocAdmin := entities.User{
			Name:     "test",
			Email:    "testadmin@mail.com",
			Password: "test",
		}

		resAdmin, _ := admin.New(db).Register(mocAdmin)

		mocFood := []entities.Food{
			{
				Food_uid: f1.Food_uid,
			},
			{
				Food_uid: f1.Food_uid,
			},
		}
		mocMenu := entities.Menu{
			User_uid:      resAdmin.User_uid,
			Menu_category: "zzz",
			Created_by:    "admin",
		}
		resMenuAdmin, errAdmin := repo.CreateMenuAdmin(mocFood, mocMenu)
		fmt.Println(errAdmin)
		db.Migrator().DropColumn(&entities.Menu{}, "menu_uid")

		errGet := repo.Delete(resMenuAdmin.Menu_uid)

		assert.NotNil(t, errGet)
	})

}
