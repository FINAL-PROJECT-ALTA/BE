package user

import (
	"HealthFit/delivery/middlewares"
	"HealthFit/entities"
	"errors"

	"github.com/lithammer/shortuuid"
	"gorm.io/gorm"
)

type UserRepository struct {
	database *gorm.DB
}

func New(db *gorm.DB) *UserRepository {
	return &UserRepository{
		database: db,
	}
}

func (ur *UserRepository) Register(u entities.User) (entities.User, error) {

	u.Password, _ = middlewares.HashPassword(u.Password)
	uid := shortuuid.New()
	u.User_uid = uid
	u.Roles = false

	if err := ur.database.Create(&u).Error; err != nil {
		return u, errors.New("email already exist")
	}

	return u, nil
}

func (ur *UserRepository) GetById(user_uid string) (entities.User, error) {
	arrUser := entities.User{}

	result := ur.database.Preload("Goal").Preload("History").Where("user_uid = ?", user_uid).First(&arrUser)
	if err := result.Error; err != nil {
		return arrUser, err
	}

	return arrUser, nil
}

func (ur *UserRepository) Update(user_uid string, newUser entities.User) (entities.User, error) {

	var user entities.User
	ur.database.Where("user_uid =?", user_uid).First(&user)

	if err := ur.database.Model(&user).Updates(&newUser).Error; err != nil {
		return user, err
	}

	return user, nil
}

func (ur *UserRepository) Delete(user_uid string) error {

	var user entities.User

	if err := ur.database.Where("user_uid =?", user_uid).First(&user).Error; err != nil {
		return err
	}
	ur.database.Delete(&user, user_uid)
	return nil

}
