package goal

import (
	"HealthFit/entities"
	"math"
	"time"

	"github.com/lithammer/shortuuid"

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

	uid := shortuuid.New()
	u.Goal_uid = uid
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
	time := time.Now()
	different := goal.CreatedAt.Sub(time)

	days := math.Abs(float64(int(different.Hours() / 24)))
	diff := goal.Range_time - int(days)
	if diff <= 0 && goal.Status == "active" {
		goal.Status = "not active"
		if err := ur.database.Model(entities.Goal{}).Where("goal_uid =?", goal.Goal_uid).Updates(&goal).Error; err != nil {
			return goal, err
		}

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
