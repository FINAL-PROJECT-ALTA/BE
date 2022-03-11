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

func (mr *MenuRepository) CreateMenuAdmin(foods []entities.Food, newMenu entities.Menu) (entities.Menu, error) {

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
func (mr *MenuRepository) CreateMenuUser(foods []entities.Food, newMenu entities.Menu) (entities.Menu, error) {

	uid := shortuuid.New()
	newMenu.Menu_uid = uid
	err := mr.database.Transaction(func(tx *gorm.DB) error {

		if err := tx.Create(&newMenu).Error; err != nil {
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
		var goal entities.Goal
		if err := tx.Model(entities.Goal{}).Where("user_uid=? AND status =?", newMenu.User_uid, "active").Find(&goal).Error; err != nil {
			return err
		}
		var user_history entities.User_history
		user_history.Menu_uid = newMenu.Menu_uid
		user_history.User_uid = newMenu.User_uid

		if err := tx.Model(entities.User_history{}).Create(&user_history).Error; err != nil {
			return err
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

func (mr *MenuRepository) GetAllMenu(category string, createdBy string) ([]entities.Menu, error) {
	menus := []entities.Menu{}

	if category != "" && createdBy == "" {
		res := mr.database.Preload("Detail_menu").Preload("Detail_menu.Food").Where("menu_category = ?", category).Find(&menus)
		if err := res.Error; err != nil {
			return menus, err
		}
	} else if createdBy != "" && category == "" {
		res := mr.database.Preload("Detail_menu").Preload("Detail_menu.Food").Where("created_by = ?", createdBy).Find(&menus)

		if err := res.Error; err != nil {
			return menus, err
		}
	} else if createdBy != "" && category != "" {
		res := mr.database.Preload("Detail_menu").Preload("Detail_menu.Food").Where("created_by = ? AND menu_category = ?", createdBy, category).Find(&menus)

		if err := res.Error; err != nil {
			return menus, err
		}
	} else {
		res := mr.database.Preload("Detail_menu").Preload("Detail_menu.Food").Find(&menus)

		if err := res.Error; err != nil {
			return menus, err
		}
	}

	return menus, nil
}

func (mr *MenuRepository) GetMenuRecommend() ([]entities.Menu, error) {
	menus := []entities.Menu{}

	res := mr.database.Preload("Detail_menu").Preload("Detail_menu.Food").Where("Created_by = ?", "createdby").Find(&menus)

	if err := res.Error; err != nil {
		return menus, err
	}
	return menus, nil
}

func (mr *MenuRepository) Update(menu_uid string, foods []entities.Food, updateMenu entities.Menu) (entities.Menu, error) {
	//code baru masih belum work
	err := mr.database.Transaction(func(tx *gorm.DB) error {
		var menu entities.Menu

		if err := tx.Model(entities.Menu{}).Where("menu_uid = ?", menu_uid).Find(&menu).Error; err != nil {
			return err
		}

		// var detail entities.Detail_menu
		if err := tx.Where("menu_uid =?", menu_uid).Delete(&entities.Detail_menu{}).Error; err != nil {
			return err
		}

		if err := tx.Where("menu_uid =?", menu_uid).Delete(&entities.Menu{}).Error; err != nil {
			return err
		}

		uid := shortuuid.New()
		updateMenu.Menu_uid = uid

		if err := tx.Create(&updateMenu).Error; err != nil {
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

	menuss := menu_uid

	var menu entities.Menu
	err := mr.database.Transaction(func(tx *gorm.DB) error {

		if err := tx.Debug().Where("menu_uid = ?", menuss).Delete(&entities.Detail_menu{}).Error; err != nil {
			return err
		}

		if err := tx.Debug().Where("menu_uid =?", menuss).Delete(&menu).Error; err != nil {
			return err
		}
		return nil

	})

	if err != nil {
		return err
	}

	return nil

}
