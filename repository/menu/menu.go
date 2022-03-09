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

func (mr *MenuRepository) GetMenuByCategory(category string) ([]entities.Menu, error) {
	menus := []entities.Menu{}

	res := mr.database.Preload("Detail_menu").Preload("Detail_menu.Food").Where("menu_category = ?", category).Find(&menus)

	if err := res.Error; err != nil {
		return menus, err
	}
	return menus, nil
}

func (mr *MenuRepository) GetAllMenu() ([]entities.Menu, error) {
	menu := []entities.Menu{}

	res := mr.database.Preload("Detail_menu").Preload("Detail_menu.Food").Find(&menu)

	if err := res.Error; err != nil {
		return menu, err
	}
	return menu, nil
}

func (mr *MenuRepository) Update(menu_uid string, foods []entities.Food, updateMenu entities.Menu) (entities.Menu, error) {
	//code baru masih belum work
	var menu entities.Menu
	mr.database.Model(entities.Menu{}).Where("menu_uid = ?", menu_uid).First(&menu)
	var detail entities.Detail_menu

	if err := mr.database.Model(entities.Detail_menu{}).Where("menu_uid = ?", menu_uid).Delete(&detail).Error; err != nil {
		return entities.Menu{}, err
	}

	mr.database.Model(entities.Menu{}).Where("menu_uid = ?", menu_uid).Delete(&menu)

	uid := shortuuid.New()
	updateMenu.Menu_uid = uid

	if err := mr.database.Preload("Detail_menu").Preload("Detail_menu.Food").Create(&updateMenu).Error; err != nil {
		return updateMenu, err
	}
	for i := 0; i < len(foods); i++ {
		detail := entities.Detail_menu{
			Menu_uid: updateMenu.Menu_uid,
			Food_uid: foods[i].Food_uid,
		}
		if err := mr.database.Model(entities.Detail_menu{}).Create(&detail).Error; err != nil {
			return updateMenu, err
		}
	}

	return updateMenu, nil
	//end code baru

	// var menu entities.Menu
	// mr.database.Where("menu_uid = ?", menu_uid).First(&menu)

	// for i := 0; i < len(foods); i++ {

	// 	if err := mr.database.Delete(entities.Detail_menu{}, "menu_uid =?", menu.Menu_uid).Error; err != nil {
	// 		return entities.Menu{}, err
	// 	}
	// }

	// if err := mr.database.Create(&updateMenu).Error; err != nil {
	// 	return entities.Menu{}, err
	// }
	// for i := 0; i < len(foods); i++ {
	// 	detail := entities.Detail_menu{
	// 		Menu_uid: updateMenu.Menu_uid,
	// 		Food_uid: foods[i].Food_uid,
	// 	}
	// 	if err := mr.database.Model(entities.Detail_menu{}).Create(&detail).Error; err != nil {
	// 		return entities.Menu{}, err
	// 	}
	// }

	// // code lama
	// // if err := mr.database.Model(&menu).Updates(&updateMenu).Error; err != nil {
	// // 	return updateMenu, err
	// // }

	// return updateMenu, nil
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
