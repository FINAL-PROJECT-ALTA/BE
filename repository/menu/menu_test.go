package menu

import (
	"HealthFit/configs"
	food "HealthFit/repository/foods"
	"fmt"

	// user "HealthFit/repository/user"
	admin "HealthFit/repository/admin"
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
