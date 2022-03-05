package foods

import (
	"HealthFit/entities"
	"errors"

	"gorm.io/gorm"
)

type FoodsRepository struct {
	database *gorm.DB
}

func New(db *gorm.DB) *FoodsRepository {
	return &FoodsRepository{
		database: db,
	}
}

func (fr *FoodsRepository) Create(f entities.Foods) (entities.Foods, error) {
	if err := fr.database.Create(&f).Error; err != nil {
		return f, err
	}

	return f, nil
}

func (fr *FoodsRepository) Search(input entities.Foods) (entities.Foods, error) {
	foods := entities.Foods{}

	result := fr.database.Where("name = ?", input).Or("calories", input).First(&foods)

	if err := result.Error; err != nil {
		return foods, err
	}

	return foods, nil
}

func (fr *FoodsRepository) Update(food_uid string, newFoods entities.Foods) (entities.Foods, error) {

	var foods entities.Foods
	fr.database.Where("food_uid =?", food_uid).First(&foods)

	if err := fr.database.Model(&foods).Updates(&newFoods).Error; err != nil {
		return foods, err
	}

	return foods, nil
}

func (fr *FoodsRepository) Delete(food_uid string) error {

	var foods entities.Foods

	if err := fr.database.Where("food_uid =?", food_uid).First(&foods).Error; err != nil {
		return err
	}

	if err := fr.database.Delete(&foods, food_uid).Error; err != nil {
		return err
	}

	return nil
}

func (fr *FoodsRepository) GetAll() ([]entities.Foods, error) {
	foods := []entities.Foods{}
	fr.database.Find(&foods)
	if len(foods) < 1 {
		return nil, errors.New("nil value")
	}
	return foods, nil
}
