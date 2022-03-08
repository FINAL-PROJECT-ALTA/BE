package admin

import "HealthFit/entities"

type Admin interface {
	Register(admin entities.User) (entities.User, error)
	// GetById(adminUid string) (entities.Admin, error)
	// Update(adminUid string, newAdmin entities.Admin) (entities.Admin, error)
	// Delete(adminUid string) error
	// GetAll() ([]entities.Admin, error)
}
