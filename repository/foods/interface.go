package food

import "HealthFit/entities"

type Food interface {
	Create(f entities.Food) (entities.Food, error)
	Search(input, category string) ([]entities.Food, error)
	Update(food_uid string, newFood entities.Food) (entities.Food, error)
	Delete(food_uid string) error
	GetAll() ([]entities.Food, error)
}
