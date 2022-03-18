package userhistory

import "HealthFit/entities"

type UserHistory interface {
	Insert(newHistory entities.User_history) (entities.User_history, error)
	GetAll(user_uid string) ([]entities.User_history, error)
	GetById(user_uid, user_history_uid string) (entities.User_history, error)
	// Update(user_uid, user_history_uid string, updateHistory entities.User_history)
	// Delete(user_uid, user_history_uid string) error
	//dlsahakjhkfhalkjfelkwjf
}
