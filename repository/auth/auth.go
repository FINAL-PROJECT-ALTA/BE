package auth

import (
	"HealthFit/delivery/middlewares"
	"HealthFit/entities"
	"math"
	"time"

	"errors"

	"gorm.io/gorm"
)

type AuthDb struct {
	db *gorm.DB
}

func New(db *gorm.DB) *AuthDb {
	return &AuthDb{
		db: db,
	}
}

func (ad *AuthDb) Login(email, password string) (entities.User, error) {
	user := entities.User{}

	if err := ad.db.Model(&user).Where("email = ?", email).First(&user).Error; err != nil {
		return user, errors.New("email not found")
	}

	match := middlewares.CheckPasswordHash(password, user.Password)

	if !match {
		return user, errors.New("incorrect password")
	}

	if !user.Roles {
		message, err := ad.RefreshGoalAuth(user.User_uid)
		if message == "updated" && err == nil {
			user.Goal_exspired = true
			return user, nil
		} else if message == "have goal active and nothing to update" {
			user.Goal_active = true
			return user, nil
		} else {
			return user, nil
		}
	}

	return user, nil
}
func (ad *AuthDb) RefreshGoalAuth(user_uid string) (interface{}, error) {

	goal := entities.Goal{}
	res := ad.db.Model(entities.Goal{}).Where("user_uid =? AND status=?", user_uid, "active").First(&goal)

	if res.Error != nil {
		return "failed get goal active", res.Error
	}
	time := time.Now()
	different := goal.CreatedAt.Sub(time)
	days := math.Abs(float64(int(different.Hours() / 24)))

	if int(days) > goal.Range_time {
		status := "not active"
		if err := ad.db.Model(&goal).Where("goal_uid = ?", goal.Goal_uid).Update("status", status).Error; err != nil {
			return "failed updated", err
		}
		return "updated", nil

	}

	return "have goal active and nothing to update", nil

}
