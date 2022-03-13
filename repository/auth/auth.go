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
		_, err := ad.RefreshGoalAuth(user.User_uid)
		if err != nil {
			return entities.User{}, err
		}
		return user, nil
	}

	return user, nil
}
func (ad *AuthDb) RefreshGoalAuth(user_uid string) (bool, error) {

	var goal entities.Goal
	res := ad.db.Model(entities.Goal{}).Where("user_uid =? AND status=?", user_uid, "active").First(&goal)

	if res.Error != nil {
		return false, res.Error
	}
	if res.RowsAffected == 0 {
		return false, nil
	}

	time := time.Now()
	different := goal.CreatedAt.Sub(time)

	days := math.Abs(float64(int(different.Hours() / 24)))
	if int(days) > goal.Range_time {
		status := "not active"
		if err := ad.db.Model(&entities.Goal{}).Where("goal_uid = ?", goal.Goal_uid).Update("status", status).Error; err != nil {
			return false, err
		}

	}

	return true, nil

}

// func (ad *AuthDb) LoginAdmin(email, password string) (entities.Admin, error) {
// 	admin := entities.Admin{}

// 	ad.db.Model(&admin).Where("email = ?", email).First(&admin)

// 	match := middlewares.CheckPasswordHash(password, admin.Password)

// 	if !match {
// 		return entities.Admin{}, errors.New("incorrect password")
// 	}

// 	return admin, nil
// }
