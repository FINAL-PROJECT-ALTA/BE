package foods

import "HealthFit/entities"

type Foods interface {
	Create(f entities.Foods) (entities.Foods, error)
	Search(input, category string) ([]entities.Foods, error)
	Update(food_uid string, newFoods entities.Foods) (entities.Foods, error)
	Delete(food_uid string) error
	GetAll() ([]entities.Foods, error)
}
