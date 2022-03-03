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

	ad.db.Model(&user).Where("email = ?", email).First(&user)

	match := middlewares.CheckPasswordHash(password, user.Password)

	if !match {
		return entities.User{}, errors.New("incorrect password")
	}

	return user, nil
}
