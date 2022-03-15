package food

import (
	"HealthFit/configs"
	"HealthFit/entities"
	detailMenu "HealthFit/repository/detail_menu"
	utils "HealthFit/utils/mysql"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreate(t *testing.T) {
	config := configs.GetConfig()
	db := utils.InitDB(config)
	repo := New(db)

	t.Run("fail run Create", func(t *testing.T) {
		db.Migrator().DropTable(&entities.Food{}, &entities.Menu{}, &entities.Detail_menu{})
		db.AutoMigrate(&entities.Food{}, &entities.Menu{}, &entities.Detail_menu{})
		mocFoodFail := entities.Food{
			Name:          "makanan",
			Calories:      100,
			Energy:        200,
			Carbohidrate:  300,
			Protein:       400,
			Food_category: "snack",
			Unit:          "ons",
			Unit_value:    1,
		}
		res, err := repo.Create(mocFoodFail)
		if err != nil {
			t.Fatal()
		}
		mocUser := entities.Food{Food_uid: res.Food_uid, Name: "makananku"}
		_, errA := repo.Create(mocUser)
		assert.NotNil(t, errA)
	})

	t.Run("success run Create", func(t *testing.T) {
		db.Migrator().DropTable(&entities.Food{}, &entities.Menu{}, &entities.Detail_menu{})
		db.AutoMigrate(&entities.Food{}, &entities.Menu{}, &entities.Detail_menu{})
		mocFood := entities.Food{
			Name:          "makanan",
			Calories:      100,
			Energy:        200,
			Carbohidrate:  300,
			Protein:       400,
			Food_category: "snack",
			Unit:          "ons",
			Unit_value:    1,
		}
		res, err := repo.Create(mocFood)

		// mockMenu := entities.Menu{

		// }

		mockDetailMenu := entities.Detail_menu{
			Food_uid: res.Food_uid,
		}
		_, errDM := detailMenu.New(db).Create(mockDetailMenu)
		if errDM != nil {
			t.Fail()
		}

		assert.Nil(t, err)
		assert.Equal(t, "makanan", res.Name)

	})

}
