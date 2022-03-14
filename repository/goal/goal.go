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

	if bmr, cutCaloriesDay, err := ur.CheckRecommendGoal(g); err != nil {
		return entities.Goal{Weight: bmr, Height: cutCaloriesDay}, err
	}

	uid := shortuuid.New()
	g.Goal_uid = uid
	if err := ur.database.Create(&g).Error; err != nil {
		return g, err
	}
	return g, nil
}

func (ur *GoalRepository) GetAll(user_uid string) ([]entities.Goal, error) {
	goals := []entities.Goal{}

	res := ur.database.Where("user_uid", user_uid).Find(&goals)

	if err := res.Error; err != nil {
		return []entities.Goal{}, err
	}

	return goals, nil
}

func (ur *GoalRepository) GetById(goal_uid string, user_uid string) (entities.Goal, error) {
	goal := entities.Goal{}

	result := ur.database.Where("goal_uid = ? AND user_uid=?", goal_uid).First(&goal)
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

func (ur *GoalRepository) Delete(goal_uid string, user_uid string) error {

	var goal entities.Goal

	if err := ur.database.Where("goal_uid =? AND user_uid =?", goal_uid, user_uid).First(&goal).Error; err != nil {
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
	if int(days) > goal.Range_time {
		status := "not active"
		if err := ur.database.Model(&entities.Goal{}).Where("goal_uid = ?", goal.Goal_uid).Update("status", status).Error; err != nil {
			return false, err
		}

	}

	return true, nil

}
func (ur *GoalRepository) CancelGoal(user_uid string) (entities.Goal, error) {

	var goal entities.Goal

	if err := ur.database.Model(entities.Goal{}).Where("user_uid =? AND status=?", user_uid, "active").First(&goal).Error; err != nil {
		return entities.Goal{}, errors.New("failed to cancel goal")
	}

	if err := ur.database.Model(&goal).Where("goal_uid =?", goal.Goal_uid).Update("status", "cancel").Error; err != nil {
		return entities.Goal{}, errors.New("failed to cancel goal")
	}

	return goal, nil

}
func (ur *GoalRepository) CheckGoal(user_uid string) error {

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

func (ur *GoalRepository) CheckRecommendGoal(goal entities.Goal) (int, int, error) {

	var user entities.User
	var bmr int
	var cutCaloriesDay int

	resUser := ur.database.Model(entities.User{}).Where("user_uid = ?", goal.User_uid).First(&user)

	if err := resUser.Error; err != nil {
		return 0, 0, err
	}
	cutCaloriesDay = int(math.Round(float64(goal.Weight_target * 7700 / goal.Range_time)))

	var daily_active float32
	switch goal.Daily_active {
	case "not active":
		daily_active = 1.2
	case "little active":
		daily_active = 1.37
	case "quite active":
		daily_active = 1.5
	case "active":
		daily_active = 1.72
	case "very active":
		daily_active = 1.9
	}
	if user.Gender == "male" {
		bmr = int(daily_active) * (66 + (14 * goal.Weight) + (5 * goal.Height) - (7 * goal.Age))

	}
	if user.Gender == "female" {
		bmr = int(daily_active) * (655 + (9 * goal.Weight) + (2 * goal.Height) - (5 * goal.Age))
	}
	bmrDay := bmr - cutCaloriesDay

	posible := bmr * 50 / 100
	if int(bmrDay) < posible {
		return bmr, cutCaloriesDay, errors.New("impossible")
	}

	return bmr, cutCaloriesDay, nil
}
