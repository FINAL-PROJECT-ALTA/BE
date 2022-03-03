package user

import "HealthFit/entities"

type User interface {
	Register(user entities.User) (entities.User, error)
	GetById(userUid string) (entities.User, error)
	Update(userUid string, newUser entities.User) (entities.User, error)
	Delete(userUid string) error
	// GetAll() ([]entities.User, error)
}
