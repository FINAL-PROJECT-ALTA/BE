package auth

import (
	"HealthFit/delivery/middlewares"
	"HealthFit/entities"
	"errors"

	"gorm.io/gorm"
)

type AuthDb struct {
	db *gorm.DB
}

func New(db *gorm.DB) *AuthDb {
	return &AuthDb{
		db: db,
	}
}

func (ad *AuthDb) Login(email, password string) (entities.User, error) {
	user := entities.User{}

	if err := ad.db.Model(&user).Where("email = ?", email).First(&user).Error; err != nil {
		return user, errors.New("email not found")
	}

	match := middlewares.CheckPasswordHash(password, user.Password)

	if !match {
		return user, errors.New("incorrect password")
	}

	return user, nil
}

// func (ad *AuthDb) LoginAdmin(email, password string) (entities.Admin, error) {
// 	admin := entities.Admin{}

// 	ad.db.Model(&admin).Where("email = ?", email).First(&admin)

// 	match := middlewares.CheckPasswordHash(password, admin.Password)

// 	if !match {
// 		return entities.Admin{}, errors.New("incorrect password")
// 	}

// 	return admin, nil
// }
