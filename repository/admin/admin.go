package admin

import (
	"HealthFit/delivery/middlewares"
	"HealthFit/entities"
	"errors"

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

func (ar *AdminRepository) Register(a entities.User) (entities.User, error) {

	a.Password, _ = middlewares.HashPassword(a.Password)
	uid := shortuuid.New()
	a.User_uid = uid
	a.Roles = true

	if err := ar.database.Create(&a).Error; err != nil {
		return a, errors.New("email already exist")
	}

	return a, nil
}
