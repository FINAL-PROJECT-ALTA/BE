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

	t.Run("failed run Create", func(t *testing.T) {
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

		_, err := repo.Create(mocFood)
		if err != nil {
			t.Fail()
		}

		mocFoodU := entities.Food{
			ID: 1,
		}
		_, errA := repo.Create(mocFoodU)

		assert.NotNil(t, errA)

	})

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

func TestGetById(t *testing.T) {
	config := configs.GetConfig()
	db := utils.InitDB(config)
	repo := New(db)

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
		res, _ := repo.Create(mocFood)

		resG, errG := repo.GetById(res.Food_uid)

		assert.Nil(t, errG)
		assert.Equal(t, "makanan", resG.Name)

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
		_, err := repo.Create(mocFood)
		if err != nil {
			t.Fail()
		}

		resG, _ := repo.GetById("kn hb")

		assert.NotEqual(t, "dadwa", resG.Name)

	})

}

func TestUpdateById(t *testing.T) {
	config := configs.GetConfig()
	db := utils.InitDB(config)
	repo := New(db)

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
		res, _ := repo.Create(mocFood)

		mockUp := entities.Food{
			Name: "testUp",
		}

		resU, errU := repo.Update(res.Food_uid, mockUp)

		assert.Nil(t, errU)
		assert.Equal(t, "testUp", resU.Name)

	})

	t.Run("fail run Create", func(t *testing.T) {
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
		_, err := repo.Create(mocFood)
		if err != nil {
			t.Fail()
		}

		mockUp := entities.Food{
			Name: "testUp",
		}

		_, errU := repo.Update("khadbk", mockUp)

		assert.NotNil(t, errU)

	})
}

func TestDeleteById(t *testing.T) {
	config := configs.GetConfig()
	db := utils.InitDB(config)
	repo := New(db)

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
		res, _ := repo.Create(mocFood)

		errD := repo.Delete(res.Food_uid)

		assert.Nil(t, errD)

	})

	// t.Run("fail run Create", func(t *testing.T) {
	// 	db.Migrator().DropTable(&entities.Food{}, &entities.Menu{}, &entities.Detail_menu{})
	// 	db.AutoMigrate(&entities.Food{}, &entities.Menu{}, &entities.Detail_menu{})

	// 	mocFood := entities.Food{
	// 		Name:          "herbal",
	// 		Calories:      100,
	// 		Energy:        200,
	// 		Carbohidrate:  300,
	// 		Protein:       400,
	// 		Food_category: "snack",
	// 		Unit:          "ons",
	// 		Unit_value:    1,
	// 	}
	// 	res, err := repo.Create(mocFood)
	// 	if err != nil {
	// 		t.Fail()
	// 	}

	// 	errD := repo.Delete(res.Food_uid)
	// 	if errD != nil {
	// 		t.Fail()
	// 	}
	// 	errK := repo.Delete(res.Food_uid)

	// 	assert.NotNil(t, errK)

	// })
}

func TestGetAll(t *testing.T) {
	config := configs.GetConfig()
	db := utils.InitDB(config)
	repo := New(db)

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
		res, _ := repo.Create(mocFood)

		_, errG := repo.GetAll(res.Food_category)

		assert.Nil(t, errG)

	})

	t.Run("success run Create", func(t *testing.T) {
		db.Migrator().DropTable(&entities.Food{}, &entities.Menu{}, &entities.Detail_menu{})
		db.AutoMigrate(&entities.Food{}, &entities.Menu{}, &entities.Detail_menu{})

		// mocFood := entities.Food{
		// 	Name:          "makanan",
		// 	Calories:      100,
		// 	Energy:        200,
		// 	Carbohidrate:  300,
		// 	Protein:       400,
		// 	Food_category: "snack",
		// 	Unit:          "ons",
		// 	Unit_value:    1,
		// }
		// res, _ := repo.Create(mocFood)

		_, errG := repo.GetAll("res.Food_category")

		assert.NotNil(t, errG)

	})

	t.Run("success run Create", func(t *testing.T) {
		db.Migrator().DropTable(&entities.Food{}, &entities.Menu{}, &entities.Detail_menu{})
		db.AutoMigrate(&entities.Food{}, &entities.Menu{}, &entities.Detail_menu{})

		// mocFood := entities.Food{
		// 	Name:          "makanan",
		// 	Calories:      100,
		// 	Energy:        200,
		// 	Carbohidrate:  300,
		// 	Protein:       400,
		// 	Food_category: "snack",
		// 	Unit:          "ons",
		// 	Unit_value:    1,
		// }
		// res, _ := repo.Create(mocFood)

		_, errG := repo.GetAll("res.Food_category")

		assert.NotNil(t, errG)

	})
}

func TestSearch(t *testing.T) {
	config := configs.GetConfig()
	db := utils.InitDB(config)
	repo := New(db)

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
		res, _ := repo.Create(mocFood)

		_, errG := repo.Search("snack", res.Food_category)

		assert.Nil(t, errG)

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
		_, err := repo.Create(mocFood)
		if err != nil {
			t.Fail()
		}

		_, errG := repo.Search("", "")

		assert.Nil(t, errG)

	})
}
