package food

import (
	"HealthFit/entities"
	"errors"
	"fmt"
	"strconv"

	"github.com/lithammer/shortuuid"

	"gorm.io/gorm"
)

type FoodRepository struct {
	database *gorm.DB
}

func New(db *gorm.DB) *FoodRepository {
	return &FoodRepository{
		database: db,
	}
}

func (fr *FoodRepository) Create(f entities.Food) (entities.Food, error) {

	uid := shortuuid.New()
	f.Food_uid = uid
	if err := fr.database.Create(&f).Error; err != nil {
		return f, err
	}

	return f, nil
}

func (fr *FoodRepository) Search(input, category string) ([]entities.Food, error) {

	foods := []entities.Food{}
	sql := "SELECT * FROM foods"

	if input == "fruit" || input == "drink" || input == "junk food" || input == "food" || input == "snack" {
		sql = fmt.Sprintf("%s WHERE food_category =%s ", sql, input)
	}

	if category == "foods" && input != "" {
		sql = fmt.Sprintf("%s WHERE name LIKE '%%%s%%'", sql, input)
	}
	if category == "calories" && input != "" {
		input, _ := strconv.Atoi(input)
		sql = fmt.Sprintf("%s WHERE calories < %d", sql, input)
	}

	if err := fr.database.Raw(sql).Scan(&foods).Error; err != nil {
		return foods, err
	}

	// result := fr.database.Where("name = ?", input).Or("calories", input).Find(&foods)

	// if err := result.Error; err != nil {
	// 	return foods, err
	// }

	return foods, nil
}

func (fr *FoodRepository) Update(food_uid string, newFoods entities.Food) (entities.Food, error) {

	var foods entities.Food
	fr.database.Where("food_uid =?", food_uid).First(&foods)

	if err := fr.database.Model(&foods).Updates(&newFoods).Error; err != nil {
		return foods, err
	}

	return foods, nil
}

func (fr *FoodRepository) Delete(food_uid string) error {

	var foods entities.Food

	if err := fr.database.Model(&foods).Where("food_uid = ?", food_uid).Delete(&foods).Error; err != nil {
		return err
	}

	return nil
}

func (fr *FoodRepository) GetAll(category string) ([]entities.Food, error) {
	foods := []entities.Food{}

	if category != "" {
		fr.database.Where("food_category=?", category).Find(&foods)
		if len(foods) < 1 {
			return nil, errors.New("nil value")
		}
	} else {
		fr.database.Find(&foods)
		if len(foods) < 1 {
			return nil, errors.New("nil value")
		}
	}

	return foods, nil
}
