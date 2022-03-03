package user

import (
	"HealthFit/delivery/middlewares"
	"HealthFit/entities"

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

	if err := ur.database.Create(&u).Error; err != nil {
		return u, err
	}

	return u, nil
}

func (ur *UserRepository) GetById(userId int) (entities.User, error) {
	arrUser := entities.User{}

	result := ur.database.Where("ID = ?", userId).First(&arrUser)
	if err := result.Error; err != nil {
		return arrUser, err
	}

	return arrUser, nil
}

func (ur *UserRepository) Update(userId int, newUser entities.User) (entities.User, error) {

	var user entities.User
	ur.database.First(&user, userId)

	if err := ur.database.Model(&user).Updates(&newUser).Error; err != nil {
		return user, err
	}

	return user, nil
}

func (ur *UserRepository) Delete(userId int) error {

	var user entities.User

	if err := ur.database.First(&user, userId).Error; err != nil {
		return err
	}
	ur.database.Delete(&user, userId)
	return nil

}
