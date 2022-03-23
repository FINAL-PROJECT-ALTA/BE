package menu

import (
	"HealthFit/configs"
	"HealthFit/repository/admin"
	food "HealthFit/repository/foods"
	"fmt"

	"HealthFit/repository/goal"

	"HealthFit/repository/user"

	utils "HealthFit/utils/mysql"

	"HealthFit/entities"
	"testing"

	// "github.com/labstack/gommon/log"
	"github.com/stretchr/testify/assert"
)

func TestCreate(t *testing.T) {
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
		// db.Migrator().DropColumn(&entities.Menu{}, "total_calories")

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
		// db.Migrator().DropColumn(&entities.Menu{}, "total_calories")

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

		assert.NotNil(t, errMenuUser)

	})

}
func TestGetAll(t *testing.T) {
	config := configs.GetConfig()
	db := utils.InitDB(config)
	repo := New(db)

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
		_, errGet := repo.GetAllMenu("breakfast", "admin")

		assert.Nil(t, errGet)
	})

}
