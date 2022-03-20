package food

import (
	"HealthFit/configs"
	"HealthFit/entities"
	utils "HealthFit/utils/mysql"
	"testing"

	"github.com/labstack/gommon/log"
	"github.com/stretchr/testify/assert"
)

func TestCreate(t *testing.T) {
	config := configs.GetConfig()
	db := utils.InitDB(config)
	repo := New(db)

	t.Run("success run Create", func(t *testing.T) {
		db.Migrator().DropTable(&entities.Detail_menu{}, &entities.Food{}, &entities.Menu{})
		db.Migrator().DropTable(&entities.Food{}, &entities.Menu{})
		db.AutoMigrate(&entities.Food{})
		db.AutoMigrate(&entities.Menu{})
		db.AutoMigrate(&entities.Detail_menu{})

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
	t.Run("failed run Create", func(t *testing.T) {

		db.Migrator().DropTable(&entities.Detail_menu{})
		db.Migrator().DropTable(&entities.Food{})
		db.Migrator().DropTable(&entities.Menu{})

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

			Food_uid: "asdfslkdajlkfd",
		}
		_, errA := repo.Create(mocFoodU)

		assert.NotNil(t, errA)

	})

}

func TestGetById(t *testing.T) {
	config := configs.GetConfig()
	db := utils.InitDB(config)
	repo := New(db)

	t.Run("success GetById", func(t *testing.T) {
		db.Migrator().DropTable(&entities.Detail_menu{})
		db.Migrator().DropTable(&entities.Food{})
		db.Migrator().DropTable(&entities.Menu{})
		db.AutoMigrate(&entities.Food{})
		db.AutoMigrate(&entities.Menu{})
		db.AutoMigrate(&entities.Detail_menu{})

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

	t.Run("fail getById", func(t *testing.T) {
		db.Migrator().DropTable(&entities.Detail_menu{})
		db.Migrator().DropTable(&entities.Food{})
		db.Migrator().DropTable(&entities.Menu{})
		db.AutoMigrate(&entities.Food{})
		db.AutoMigrate(&entities.Menu{})
		db.AutoMigrate(&entities.Detail_menu{})

		resG, errG := repo.GetById("")
		log.Info(resG)

		assert.NotNil(t, errG)

	})

}

func TestUpdateById(t *testing.T) {
	config := configs.GetConfig()
	db := utils.InitDB(config)
	repo := New(db)

	t.Run("success UpdateById", func(t *testing.T) {
		db.Migrator().DropTable(&entities.Detail_menu{})
		db.Migrator().DropTable(&entities.Food{})
		db.Migrator().DropTable(&entities.Menu{})
		db.AutoMigrate(&entities.Food{})
		db.AutoMigrate(&entities.Menu{})
		db.AutoMigrate(&entities.Detail_menu{})

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

	t.Run("fail updateBy Id", func(t *testing.T) {
		db.Migrator().DropTable(&entities.Detail_menu{})
		db.Migrator().DropTable(&entities.Food{})
		db.Migrator().DropTable(&entities.Menu{})
		db.AutoMigrate(&entities.Food{})
		db.AutoMigrate(&entities.Menu{})
		db.AutoMigrate(&entities.Detail_menu{})

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

	t.Run("success Delete ById ", func(t *testing.T) {
		db.Migrator().DropTable(&entities.Detail_menu{})
		db.Migrator().DropTable(&entities.Food{})
		db.Migrator().DropTable(&entities.Menu{})
		db.AutoMigrate(&entities.Food{})
		db.AutoMigrate(&entities.Menu{})
		db.AutoMigrate(&entities.Detail_menu{})

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

	t.Run("fail delete byId", func(t *testing.T) {
		db.Migrator().DropTable(&entities.Detail_menu{})
		db.Migrator().DropTable(&entities.Food{})
		db.Migrator().DropTable(&entities.Menu{})
		db.AutoMigrate(&entities.Food{})
		db.AutoMigrate(&entities.Menu{})
		db.AutoMigrate(&entities.Detail_menu{})

		errK := repo.Delete("dj")

		assert.NotNil(t, errK)

	})
}

func TestGetAll(t *testing.T) {
	config := configs.GetConfig()
	db := utils.InitDB(config)
	repo := New(db)

	t.Run("success get all if category != string kosong", func(t *testing.T) {
		db.Migrator().DropTable(&entities.Detail_menu{})
		db.Migrator().DropTable(&entities.Food{})
		db.Migrator().DropTable(&entities.Menu{})
		db.AutoMigrate(&entities.Food{})
		db.AutoMigrate(&entities.Menu{})
		db.AutoMigrate(&entities.Detail_menu{})

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

	t.Run("fail get all if category != string kosong", func(t *testing.T) {
		db.Migrator().DropTable(&entities.Detail_menu{})
		db.Migrator().DropTable(&entities.Food{})
		db.Migrator().DropTable(&entities.Menu{})

		_, errG := repo.GetAll("food")

		assert.NotNil(t, errG)

	})

	t.Run("fail get all if category == string kosong", func(t *testing.T) {
		db.Migrator().DropTable(&entities.Detail_menu{})
		db.Migrator().DropTable(&entities.Food{})
		db.Migrator().DropTable(&entities.Menu{})
		db.AutoMigrate(&entities.Food{})
		db.AutoMigrate(&entities.Menu{})
		db.AutoMigrate(&entities.Detail_menu{})

		_, errG := repo.GetAll("")

		assert.NotNil(t, errG)

	})

}

func TestSearch(t *testing.T) {
	config := configs.GetConfig()
	db := utils.InitDB(config)
	repo := New(db)

	t.Run("success run Create", func(t *testing.T) {
		db.Migrator().DropTable(&entities.Detail_menu{})
		db.Migrator().DropTable(&entities.Food{})
		db.Migrator().DropTable(&entities.Menu{})
		db.AutoMigrate(&entities.Food{})
		db.AutoMigrate(&entities.Menu{})
		db.AutoMigrate(&entities.Detail_menu{})

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
		db.Migrator().DropTable(&entities.Detail_menu{})
		db.Migrator().DropTable(&entities.Food{})
		db.Migrator().DropTable(&entities.Menu{})
		db.AutoMigrate(&entities.Food{})
		db.AutoMigrate(&entities.Menu{})
		db.AutoMigrate(&entities.Detail_menu{})

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
	t.Run("success run Create", func(t *testing.T) {
		db.Migrator().DropTable(&entities.Detail_menu{})
		db.Migrator().DropTable(&entities.Food{})
		db.Migrator().DropTable(&entities.Menu{})

		_, errG := repo.Search("food", "food")

		assert.Nil(t, errG)

	})
	t.Run("success category == foods && input != string kosong", func(t *testing.T) {
		db.Migrator().DropTable(&entities.Detail_menu{})
		db.Migrator().DropTable(&entities.Food{})
		db.Migrator().DropTable(&entities.Menu{})

		_, errG := repo.Search("a", "foods")

		assert.NotNil(t, errG)

	})
	t.Run("success category == calories && input != string kosong", func(t *testing.T) {
		db.Migrator().DropTable(&entities.Detail_menu{})
		db.Migrator().DropTable(&entities.Food{})
		db.Migrator().DropTable(&entities.Menu{})

		_, errG := repo.Search("a", "calories")

		assert.NotNil(t, errG)

	})
}

func TestThird(t *testing.T) {
	config := configs.GetConfig()
	db := utils.InitDB(config)
	repo := New(db)

	t.Run("success create from thirdParty", func(t *testing.T) {
		db.Migrator().DropTable(&entities.Detail_menu{})
		db.Migrator().DropTable(&entities.Food{})
		db.Migrator().DropTable(&entities.Menu{})
		db.AutoMigrate(&entities.Food{})
		db.AutoMigrate(&entities.Menu{})
		db.AutoMigrate(&entities.Detail_menu{})

		mocFood := entities.Food{
			Name:          "makanan",
			Calories:      100,
			Energy:        200,
			Carbohidrate:  300,
			Protein:       400,
			Food_category: "snack",
			Unit:          "ons",
			Unit_value:    1,
			Image:         "",
		}

		res, err := repo.CreateFoodThirdParty(mocFood)
		log.Info(res)

		assert.NotNil(t, err)

	})

	t.Run("fail create from ThirdpARTY", func(t *testing.T) {
		db.Migrator().DropTable(&entities.Detail_menu{})
		db.Migrator().DropTable(&entities.Food{})
		db.Migrator().DropTable(&entities.Menu{})
		db.AutoMigrate(&entities.Food{})
		db.AutoMigrate(&entities.Menu{})
		db.AutoMigrate(&entities.Detail_menu{})

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

		resC, _ := repo.Create(mocFood)
		log.Info(resC)

		mocFoodI := entities.Food{
			ID:            1,
			Name:          "makanan",
			Calories:      100,
			Energy:        200,
			Carbohidrate:  300,
			Protein:       400,
			Food_category: "snack",
			Unit:          "ons",
			Unit_value:    1,
			Image:         "image.jpg",
		}

		res, err := repo.CreateFoodThirdParty(mocFoodI)
		log.Info(res)

		assert.NotNil(t, err)

	})
	t.Run("fail create from thirdParty because foods already exist", func(t *testing.T) {
		db.Migrator().DropTable(&entities.Detail_menu{})
		db.Migrator().DropTable(&entities.Food{})
		db.Migrator().DropTable(&entities.Menu{})
		db.AutoMigrate(&entities.Food{})
		db.AutoMigrate(&entities.Menu{})
		db.AutoMigrate(&entities.Detail_menu{})

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

		resC, _ := repo.Create(mocFood)

		mocFoodNew := entities.Food{
			Food_uid:      resC.Food_uid,
			Name:          "makanan",
			Calories:      100,
			Energy:        200,
			Carbohidrate:  300,
			Protein:       400,
			Food_category: "snack",
			Unit:          "ons",
			Unit_value:    1,
			Image:         "",
		}

		res, err := repo.CreateFoodThirdParty(mocFoodNew)
		log.Info(res)

		assert.NotNil(t, err)

	})
}
