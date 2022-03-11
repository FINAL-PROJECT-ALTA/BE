package menu

import "HealthFit/entities"

type Menu interface {
	CreateMenuAdmin(foods []entities.Food, menus entities.Menu) (entities.Menu, error)
	CreateMenuUser(foods []entities.Food, menus entities.Menu) (entities.Menu, error)
	GetAllMenu(category string, createdBy string) ([]entities.Menu, error)
	GetRecommendBreakfast(user_uid string) ([]entities.Menu, error)
	GetRecommendLunch(user_uid string) ([]entities.Menu, error)
	GetRecommendDinner(user_uid string) ([]entities.Menu, error)
	GetRecommendOverTime(user_uid string) ([]entities.Menu, error)
	Update(menu_uid string, foods []entities.Food, updateMenu entities.Menu) (entities.Menu, error)
	Delete(menu_uid string) error
}
