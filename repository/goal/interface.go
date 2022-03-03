package goal

import "HealthFit/entities"

type Goal interface {
	Create(goal entities.Goal) (entities.Goal, error)
	GetById(goal_uid string) (entities.Goal, error)
	Update(goal_uid string, newGoal entities.Goal) (entities.Goal, error)
	Delete(goal_uid string) error
	// GetAll() ([]entities.Goal, error)
}
