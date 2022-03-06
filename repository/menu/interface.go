package menu

import "HealthFit/entities"

type Menu interface {
	Create(newMenu entities.Menu) (entities.Menu, error)
	GetMenuByCategory(category string) (entities.Menu, error)
	Update(menu_uid string, updateMenu entities.Menu) (entities.Menu, error)
	Delete(menu_uid string) error
	GetAll() ([]entities.Menu, error)
}
