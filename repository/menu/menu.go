package menu

import (
	"HealthFit/entities"

	"github.com/lithammer/shortuuid"

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

func (mr *MenuRepository) Create(foods []entities.Food, newMenu entities.Menu) (entities.Menu, error) {

	uid := shortuuid.New()
	newMenu.Menu_uid = uid
	err := mr.database.Transaction(func(tx *gorm.DB) error {

		if err := tx.Preload("Detail_menu").Preload("Detail_menu.Food").Create(&newMenu).Error; err != nil {
			return err
		}
		for i := 0; i < len(foods); i++ {
			detail := entities.Detail_menu{
				Menu_uid: newMenu.Menu_uid,
				Food_uid: foods[i].Food_uid,
			}
			if err := tx.Model(entities.Detail_menu{}).Create(&detail).Error; err != nil {
				return err
			}
		}
		return nil
	})

	if err != nil {
		return newMenu, err
	}

	res := mr.database.Preload("Detail_menu").Preload("Detail_menu.Food").Where("menu_uid = ?", uid).First(&newMenu)

	if err := res.Error; err != nil {
		return entities.Menu{}, err
	}

	return newMenu, nil
}

func (mr *MenuRepository) GetAllMenu() ([]entities.Menu, error) {
	menus := []entities.Menu{}

	res := mr.database.Preload("Detail_menu").Preload("Detail_menu.Food").Find(&menus)

	if err := res.Error; err != nil {
		return menus, err
	}
	return menus, nil
}

func (mr *MenuRepository) GetMenuByCategory(category string) ([]entities.Menu, error) {
	menus := []entities.Menu{}

	res := mr.database.Preload("Detail_menu").Preload("Detail_menu.Food").Where("menu_category = ?", category).Find(&menus)

	if err := res.Error; err != nil {
		return menus, err
	}
	return menus, nil
}

func (mr *MenuRepository) GetMenuRecom(createdby string) ([]entities.Menu, error) {
	menus := []entities.Menu{}

	res := mr.database.Preload("Detail_menu").Preload("Detail_menu.Food").Where("Created_by = ?", createdby).Find(&menus)

	if err := res.Error; err != nil {
		return menus, err
	}
	return menus, nil
}

func (mr *MenuRepository) GetMenuUser(createdby string, user_uid string) ([]entities.Menu, error) {
	menus := []entities.Menu{}

	res := mr.database.Preload("Detail_menu").Preload("Detail_menu.Food").Where("Created_by = ? AND User_uid", createdby, user_uid).Find(&menus)

	if err := res.Error; err != nil {
		return menus, err
	}
	return menus, nil
}

func (mr *MenuRepository) Update(menu_uid string, foods []entities.Food, updateMenu entities.Menu) (entities.Menu, error) {
	//code baru masih belum work
	var menu entities.Menu
	err := mr.database.Transaction(func(tx *gorm.DB) error {

		if err := tx.Model(entities.Menu{}).Where("menu_uid = ?", menu_uid).First(&menu).Error; err != nil {
			return err
		}

		var detail entities.Detail_menu
		if err := tx.Model(entities.Detail_menu{}).Where("menu_uid = ?", menu_uid).Delete(&detail).Error; err != nil {
			return err
		}

		if err := tx.Model(entities.Menu{}).Where("menu_uid = ?", menu_uid).Delete(&menu).Error; err != nil {
			return err
		}

		uid := shortuuid.New()
		updateMenu.Menu_uid = uid

		if err := tx.Preload("Detail_menu").Preload("Detail_menu.Food").Create(&updateMenu).Error; err != nil {
			return err
		}
		for i := 0; i < len(foods); i++ {
			detail := entities.Detail_menu{
				Menu_uid: updateMenu.Menu_uid,
				Food_uid: foods[i].Food_uid,
			}
			if err := tx.Model(entities.Detail_menu{}).Create(&detail).Error; err != nil {
				return err
			}
		}
		return nil
	})

	if err != nil {
		return updateMenu, err
	}

	res := mr.database.Preload("Detail_menu").Preload("Detail_menu.Food").Where("menu_uid = ?", updateMenu.Menu_uid).First(&updateMenu)

	if err := res.Error; err != nil {
		return entities.Menu{}, err
	}

	return updateMenu, nil
}

func (mr *MenuRepository) Delete(menu_uid string) error {

	var menu entities.Menu
	err := mr.database.Transaction(func(tx *gorm.DB) error {

		if err := tx.Where("menu_uid = ?", menu_uid).Delete(&entities.Detail_menu{}).Error; err != nil {
			return err
		}

		if err := tx.Where("menu_uid =?", menu_uid).Delete(&menu).Error; err != nil {
			return err
		}
		return nil

	})

	if err != nil {
		return err
	}

	return nil

}
