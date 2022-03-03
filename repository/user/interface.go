package user

import "HealthFit/entities"

type User interface {
	Register(user entities.User) (entities.User, error)
	GetById(userId int) (entities.User, error)
	Update(userId int, newUser entities.User) (entities.User, error)
	Delete(userId int) error
	// GetAll() ([]entities.User, error)
}
