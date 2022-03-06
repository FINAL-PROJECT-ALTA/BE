package foods

import (
	"HealthFit/entities"
	"errors"
	"fmt"
	"strconv"

	"github.com/lithammer/shortuuid"

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

	uid := shortuuid.New()
	f.Food_uid = uid
	if err := fr.database.Create(&f).Error; err != nil {
		return f, err
	}

	return f, nil
}

func (fr *FoodsRepository) Search(input, category string) ([]entities.Foods, error) {

	foods := []entities.Foods{}
	sql := "SELECT * FROM foods"

	if category == "name" {
		sql = fmt.Sprintf("%s WHERE name LIKE '%%%s%%'", sql, input)
	}
	if category == "calories" {
		input, _ := strconv.Atoi(input)
		sql = fmt.Sprintf("%s WHERE calories < %d", sql, input)
	}

	if err := fr.database.Preload(("Image")).Raw(sql).Scan(&foods).Error; err != nil {
		return foods, err
	}

	// result := fr.database.Where("name = ?", input).Or("calories", input).Find(&foods)

	// if err := result.Error; err != nil {
	// 	return foods, err
	// }

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
