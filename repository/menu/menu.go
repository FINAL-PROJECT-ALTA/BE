package menu

import (
	"HealthFit/entities"
	"errors"

	"gorm.io/gorm"
)

type MenuRepository struct {
	database *gorm.DB
}

func New(db *gorm.DB) *MenuRepository {
	return &MenuRepository{
		database: db,
	}
}

func (mr *MenuRepository) Create(newMenu entities.Menu) (entities.Menu, error) {

	if err := mr.database.Create(&newMenu).Error; err != nil {
		return newMenu, err
	}

	return newMenu, nil
}

func (mr *MenuRepository) GetMenuByCategory(category string) (entities.Menu, error) {
	menu := entities.Menu{}

	res := mr.database.Preload("Foods").Where("menu_category = ?", category).First(&menu)

	if err := res.Error; err != nil {
		return menu, err
	}
	return menu, nil
}

func (mr *MenuRepository) Update(menu_uid string, updateMenu entities.Menu) (entities.Menu, error) {
	var menu entities.Menu
	mr.database.Where("menu_uid = ?", menu_uid).First(&menu)

	if err := mr.database.Model(&menu).Updates(&updateMenu).Error; err != nil {
		return updateMenu, err
	}

	return updateMenu, nil
}

func (mr *MenuRepository) Delete(menu_uid string) error {

	var menu entities.Menu

	if err := mr.database.Where("menu_uid =?", menu_uid).First(&menu).Error; err != nil {
		return err
	}
	if err := mr.database.Delete(&menu, menu_uid).Error; err != nil {
		return err
	}
	return nil

}

func (mr *MenuRepository) GetAll() ([]entities.Menu, error) {
	menu := []entities.Menu{}

	mr.database.Find(&menu)
	if len(menu) < 1 {
		return nil, errors.New("nil value")
	}
	return menu, nil
}
