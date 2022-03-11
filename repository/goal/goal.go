package goal

import (
	"HealthFit/entities"
	"errors"
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

func (ur *GoalRepository) Create(g entities.Goal) (entities.Goal, error) {
	var goal entities.Goal
	result := ur.database.Model(entities.Goal{}).Where("user_uid = ? AND status =?", g.User_uid, "active").First(&goal)
	if res := result.RowsAffected; res == 1 {
		return entities.Goal{}, errors.New("vailed to create goal")
	}

	uid := shortuuid.New()
	g.Goal_uid = uid
	if err := ur.database.Create(&g).Error; err != nil {
		return g, err
	}
	return g, nil
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

func (ur *GoalRepository) RefreshGoal(user_uid string) (bool, error) {

	var goal entities.Goal

	if err := ur.database.Model(entities.Goal{}).Where("user_uid =? AND status=?", user_uid, "active").First(&goal).Error; err != nil {
		return false, err
	}

	time := time.Now()
	different := goal.CreatedAt.Sub(time)

	days := math.Abs(float64(int(different.Hours() / 24)))
	diff := goal.Range_time - int(days)
	if diff <= 0 && goal.Status == "active" {
		goal.Status = "not active"
		if err := ur.database.Model(entities.Goal{}).Where("goal_uid =?", goal.Goal_uid).Updates(&goal).Error; err != nil {
			return false, err
		}

	}

	return true, nil

}
func (ur *GoalRepository) CencelGoal(user_uid string) error {

	var goal entities.Goal

	if err := ur.database.Model(entities.Goal{}).Where("user_uid =? AND status=?", user_uid, "active").First(&goal).Error; err != nil {
		return errors.New("failed to cencel goal")
	}

	goal.Status = "cencel"
	if err := ur.database.Model(entities.Goal{}).Where("goal_uid =?", goal.Goal_uid).Updates(&goal).Error; err != nil {
		return errors.New("failed to cencel goal")
	}

	return nil

}
