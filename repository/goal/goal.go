package goal

import (
	"HealthFit/entities"

	"gorm.io/gorm"
)

type GoalRepository struct {
	database *gorm.DB
}

func New(db *gorm.DB) *GoalRepository {
	return &GoalRepository{
		database: db,
	}
}

func (ur *GoalRepository) Create(u entities.Goal) (entities.Goal, error) {

	if err := ur.database.Create(&u).Error; err != nil {
		return u, err
	}

	return u, nil
}

func (ur *GoalRepository) GetById(goal_uid string) (entities.Goal, error) {
	goal := entities.Goal{}

	result := ur.database.Where("goal_uid = ?", goal_uid).First(&goal)
	if err := result.Error; err != nil {
		return goal, err
	}

	return goal, nil
}

func (ur *GoalRepository) Update(goal_uid string, newGoal entities.Goal) (entities.Goal, error) {

	var goal entities.Goal
	ur.database.Where("goal_uid =?", goal_uid).First(&goal)

	if err := ur.database.Model(&goal).Updates(&newGoal).Error; err != nil {
		return goal, err
	}

	return goal, nil
}

func (ur *GoalRepository) Delete(goal_uid string) error {

	var goal entities.Goal

	if err := ur.database.Where("goal_uid =?", goal_uid).First(&goal).Error; err != nil {
		return err
	}
	if err := ur.database.Delete(&goal, goal_uid).Error; err != nil {
		return err
	}
	return nil

}
