package food

import (
	"HealthFit/configs"
	"HealthFit/entities"
	utils "HealthFit/utils/mysql"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreate(t *testing.T) {
	config := configs.GetConfig()
	db := utils.InitDB(config)
	repo := New(db)
	db.Migrator().DropTable(&entities.Food{}, &entities.Menu{}, &entities.Detail_menu{})
	db.AutoMigrate(&entities.Food{}, &entities.Menu{}, &entities.Detail_menu{})

	t.Run("success run Create", func(t *testing.T) {

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

		assert.Nil(t, err)
		assert.Equal(t, "makanan", res.Name)

	})

}
