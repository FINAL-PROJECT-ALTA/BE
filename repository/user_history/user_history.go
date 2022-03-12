package userhistory

import (
	"HealthFit/entities"
	"errors"

	"github.com/lithammer/shortuuid"
	"gorm.io/gorm"
)

type UserHistoryRepository struct {
	database *gorm.DB
}

func New(db *gorm.DB) *UserHistoryRepository {
	return &UserHistoryRepository{
		database: db,
	}
}

func (uh *UserHistoryRepository) Insert(newHistory entities.User_history) (entities.User_history, error) {

	uid := shortuuid.New()
	newHistory.User_history_uid = uid

	if err := uh.database.Preload("Menu").Create(&newHistory).Error; err != nil {
		return newHistory, err
	}
	return newHistory, nil
}

func (uh *UserHistoryRepository) GetAll(user_uid string) ([]entities.User_history, error) {
	userHistory := []entities.User_history{}

	if err := uh.database.Preload("Menu").Where("user_uid", user_uid).Find(&userHistory).Error; err != nil {
		return userHistory, err
	}

	return userHistory, nil
}

func (uh *UserHistoryRepository) GetById(user_uid, user_history_uid string) (entities.User_history, error) {
	user_history := entities.User_history{}

	result := uh.database.Preload("Menu").Where("user_uid = ? AND user_history_uid = ?", user_uid, user_history_uid).First(&user_history)
	if err := result.Error; err != nil {
		return user_history, errors.New("record not found")
	}

	return user_history, nil
}

func (uh *UserHistoryRepository) Update(user_uid, user_history_uid string, updateHistory entities.User_history) (entities.User_history, error) {
	userHistory := entities.User_history{}

	if err := uh.database.Where("user_uid = ? AND user_history_uid = ?", user_uid, user_history_uid).First(&userHistory).Updates(&updateHistory).Error; err != nil {
		return updateHistory, err
	}

	return updateHistory, nil
}

func (uh *UserHistoryRepository) Delete(user_uid, user_history_uid string) error {
	userHistory := entities.User_history{}

	if err := uh.database.Where("user_uid = ? AND user_history_uid = ?", user_uid, user_history_uid).First(&userHistory).Delete(&userHistory).Error; err != nil {
		return err
	}
	return nil
}
