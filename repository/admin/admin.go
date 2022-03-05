package admin

import (
	"HealthFit/delivery/middlewares"
	"HealthFit/entities"

	"github.com/lithammer/shortuuid"
	"gorm.io/gorm"
)

type AdminRepository struct {
	database *gorm.DB
}

func New(db *gorm.DB) *AdminRepository {
	return &AdminRepository{
		database: db,
	}
}

func (ar *AdminRepository) Register(a entities.Admin) (entities.Admin, error) {

	a.Password, _ = middlewares.HashPassword(a.Password)
	uid := shortuuid.New()
	a.Admin_uid = uid

	if err := ar.database.Create(&a).Error; err != nil {
		return a, err
	}

	return a, nil
}

func (ar *AdminRepository) GetById(admin_uid string) (entities.Admin, error) {
	admin := entities.Admin{}

	result := ar.database.Where("admin_uid = ?", admin_uid).First(&admin)
	if err := result.Error; err != nil {
		return admin, err
	}

	return admin, nil
}

func (ar *AdminRepository) Update(admin_uid string, newAdmin entities.Admin) (entities.Admin, error) {

	var admin entities.Admin
	ar.database.Where("admin_uid =?", admin_uid).First(&admin)

	if err := ar.database.Model(&admin).Updates(&newAdmin).Error; err != nil {
		return admin, err
	}

	return admin, nil
}

func (ar *AdminRepository) Delete(admin_uid string) error {

	var admin entities.Admin

	if err := ar.database.Where("admin_uid =?", admin_uid).First(&admin).Error; err != nil {
		return err
	}
	ar.database.Delete(&admin, admin_uid)
	return nil

}
