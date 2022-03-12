package userhistory

import (
	"HealthFit/entities"
	"errors"

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

func (uh *UserHistoryRepository) Register(newHistory entities.User_history) (entities.User_history, error) {

	if err := uh.database.Create(&newHistory).Error; err != nil {
		return newHistory, errors.New("invalid input")
	}

	return newHistory, nil
}
