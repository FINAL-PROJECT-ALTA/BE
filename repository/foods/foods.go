package food

import (
	"HealthFit/entities"
	"errors"
	"fmt"
	"strconv"

	"github.com/labstack/gommon/log"
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

func (fr *FoodRepository) GetById(food_uid string) (entities.Food, error) {

	var food entities.Food
	res := fr.database.Raw("SELECT * FROM foods WHERE food_uid = ? AND deleted_at IS NULL", food_uid).Scan(&food)
	if res.RowsAffected == 0 {
		return food, errors.New("not found")
	}
	log.Info(food)
	return food, nil
}

func (fr *FoodRepository) Search(input, category string) ([]entities.Food, error) {

	foods := []entities.Food{}
	sql := "SELECT * FROM foods"

	if input == "food" || input == "fruit" || input == "drink" || input == "junk food" || input == "snack" {
		if err := fr.database.Where("food_category =?", input).Find(&foods).Error; err != nil {
			return foods, nil
		}
		return foods, nil
	}

	if category == "foods" && input != "" {
		sql = fmt.Sprintf("%s WHERE name LIKE '%%%s%%' AND deleted_at IS NULL", sql, input)
	}
	if category == "calories" && input != "" {
		input, _ := strconv.Atoi(input)
		sql = fmt.Sprintf("%s WHERE calories < %d AND deleted_at IS NULL", sql, input)
	}

	if err := fr.database.Raw(sql).Scan(&foods).Error; err != nil {
		return foods, err
	}

	return foods, nil
}

func (fr *FoodRepository) Update(food_uid string, newFoods entities.Food) (entities.Food, error) {

	var foods entities.Food

	res := fr.database.Model(&foods).Where("food_uid =?", food_uid).Updates(&newFoods)
	if res.RowsAffected == 0 {
		return foods, errors.New("")
	}

	return foods, nil
}

func (fr *FoodRepository) Delete(food_uid string) error {

	res := fr.database.Where("food_uid = ?", food_uid).Delete(&entities.Food{})
	if res.RowsAffected == 0 {
		return errors.New("")
	}

	return nil
}

func (fr *FoodRepository) GetAll(category string) ([]entities.Food, error) {
	foods := []entities.Food{}

	if category != "" {
		fr.database.Where("food_category=?", category).Order("created_at desc").Find(&foods)
		if len(foods) < 1 {
			return nil, errors.New("nil value")
		}
	} else {
		res := fr.database.Order("name").Find(&foods)
		if res.RowsAffected == 0 {
			return nil, errors.New("nil value")
		}
	}

	return foods, nil
}

func (fr *FoodRepository) CreateFoodThirdParty(foodNew entities.Food) (entities.Food, error) {

	res, err := fr.GetById(foodNew.Food_uid)
	log.Info(res)
	if res.Food_uid == "" || err != nil {
		if foodNew.Image == "" {
			foodNew.Image = "https://raw.githubusercontent.com/FINAL-PROJECT-ALTA/FE/development/image/logo-white.png"
		}
		if err := fr.database.Create(&foodNew).Error; err != nil {
			return foodNew, errors.New("failed to create food from third party")
		}
		return foodNew, errors.New("succes to create")
	}
	return foodNew, errors.New("food is found")
}
