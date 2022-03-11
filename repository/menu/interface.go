package menu

import "HealthFit/entities"

type Menu interface {
	Create(foods []entities.Food, menus entities.Menu) (entities.Menu, error)
	GetAllMenu(category string, createdBy string) ([]entities.Menu, error)
	GetMenuRecommend() ([]entities.Menu, error)
	Update(menu_uid string, foods []entities.Food, updateMenu entities.Menu) (entities.Menu, error)
	Delete(menu_uid string) error
}
