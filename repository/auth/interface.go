package auth

import "HealthFit/entities"

type Auth interface {
	Login(email, password string) (entities.User, error)
}
