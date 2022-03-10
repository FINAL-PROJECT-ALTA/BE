package detailmenu

import "HealthFit/entities"

type Detail_menu interface {
	Create(newDetailMenu entities.Detail_menu) (entities.Detail_menu, error)
	GetDetailMenuByUid(detail_menu_uid string) (entities.Detail_menu, error)
	GetAllMenu() ([]entities.Detail_menu, error)
	Update(detail_menu_uid string, updateDetailMenu entities.Detail_menu) (entities.Detail_menu, error)
	Delete(detail_menu_uid string) error
}
