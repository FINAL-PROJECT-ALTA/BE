package detailmenu

import (
	"HealthFit/entities"

	"github.com/lithammer/shortuuid"
	"gorm.io/gorm"
)

type DetailMenuRepository struct {
	database *gorm.DB
}

func New(db *gorm.DB) *DetailMenuRepository {
	return &DetailMenuRepository{
		database: db,
	}
}

func (dm *DetailMenuRepository) Create(newDetailMenu entities.Detail_menu) (entities.Detail_menu, error) {

	uid := shortuuid.New()
	newDetailMenu.Detail_menu_uid = uid

	if err := dm.database.Create(&newDetailMenu).Error; err != nil {
		return newDetailMenu, err
	}

	return newDetailMenu, nil
}

func (dm *DetailMenuRepository) GetDetailMenuByUid(detail_menu_uid string) (entities.Detail_menu, error) {
	detailMenu := entities.Detail_menu{}

	res := dm.database.Where("detail_menu_uid = ?", detail_menu_uid).First(&detailMenu)

	if err := res.Error; err != nil {
		return detailMenu, err
	}
	return detailMenu, nil
}

func (dm *DetailMenuRepository) GetAllMenu() ([]entities.Detail_menu, error) {
	detailMenu := []entities.Detail_menu{}

	res := dm.database.Find(&detailMenu)

	if err := res.Error; err != nil {
		return detailMenu, err
	}
	return detailMenu, nil
}

func (dm *DetailMenuRepository) Update(detail_menu_uid string, updateDetailMenu entities.Detail_menu) (entities.Detail_menu, error) {
	var detailMenu entities.Detail_menu
	dm.database.Where("detail_menu_uid = ?", detail_menu_uid).First(&detailMenu)

	if err := dm.database.Model(&detailMenu).Updates(&updateDetailMenu).Error; err != nil {
		return updateDetailMenu, err
	}

	return updateDetailMenu, nil
}

func (dm *DetailMenuRepository) Delete(detail_menu_uid string) error {

	var detailMenu entities.Detail_menu

	if err := dm.database.Where("detail_menu_uid =?", detail_menu_uid).First(&detailMenu).Error; err != nil {
		return err
	}
	if err := dm.database.Delete(&detailMenu, detail_menu_uid).Error; err != nil {
		return err
	}
	return nil

}
