package menu

import (
	"HealthFit/entities"
	"errors"
	"fmt"
	"math"

	"github.com/lithammer/shortuuid"

	"gorm.io/gorm"
)

type MenuRepository struct {
	database *gorm.DB
}

func New(db *gorm.DB) *MenuRepository {
	return &MenuRepository{
		database: db,
	}
}

func (mr *MenuRepository) CreateMenuAdmin(foods []entities.Food, newMenu entities.Menu) (entities.Menu, error) {

	uid := shortuuid.New()
	newMenu.Menu_uid = uid
	err := mr.database.Transaction(func(tx *gorm.DB) error {
		var total_calories int

		if err := tx.Preload("Detail_menu").Preload("Detail_menu.Food").Create(&newMenu).Error; err != nil {
			return err
		}
		for i := 0; i < len(foods); i++ {
			detail := entities.Detail_menu{
				Menu_uid: newMenu.Menu_uid,
				Food_uid: foods[i].Food_uid,
			}
			if err := tx.Model(entities.Detail_menu{}).Create(&detail).Error; err != nil {
				return err
			}

			var food entities.Food
			if err := tx.Debug().Model(entities.Food{}).Where("food_uid=?", foods[i].Food_uid).First(&food).Error; err != nil {
				return err
			}
			total_calories += food.Calories
		}

		if err := tx.Model(entities.Menu{}).Where("menu_uid = ?", uid).Updates(entities.Menu{Total_calories: total_calories}).Error; err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return newMenu, err
	}

	res := mr.database.Preload("Detail_menu").Preload("Detail_menu.Food").Where("menu_uid = ?", uid).Find(&newMenu)

	if err := res.Error; err != nil {
		return entities.Menu{}, err
	}

	return newMenu, nil
}
func (mr *MenuRepository) CreateMenuUser(foods []entities.Food, newMenu entities.Menu) (entities.Menu, error) {

	uid := shortuuid.New()
	newMenu.Menu_uid = uid

	err := mr.database.Transaction(func(tx *gorm.DB) error {
		var total_calories int

		if err := tx.Create(&newMenu).Error; err != nil {
			return err
		}
		for i := 0; i < len(foods); i++ {
			detail := entities.Detail_menu{
				Menu_uid: newMenu.Menu_uid,
				Food_uid: foods[i].Food_uid,
			}
			if err := tx.Model(entities.Detail_menu{}).Create(&detail).Error; err != nil {
				return err
			}

			var food entities.Food
			if err := tx.Debug().Model(entities.Food{}).Where("food_uid=?", foods[i].Food_uid).First(&food).Error; err != nil {
				return err
			}
			total_calories += food.Calories

		}
		var goal entities.Goal
		if err := tx.Debug().Model(entities.Goal{}).Where("user_uid=? AND status =?", newMenu.User_uid, "active").First(&goal).Error; err != nil {
			return err
		}

		var user_history entities.User_history
		user_history.Menu_uid = newMenu.Menu_uid
		user_history.User_uid = newMenu.User_uid
		user_history.Goal_uid = goal.Goal_uid
		uid := shortuuid.New()
		user_history.User_history_uid = uid

		if err := tx.Debug().Model(entities.User_history{}).Create(&user_history).Error; err != nil {
			return err
		}

		if err := tx.Model(entities.Menu{}).Where("menu_uid = ?", uid).Updates(entities.Menu{Total_calories: total_calories}).Error; err != nil {
			return err
		}
		res := mr.database.Preload("Detail_menu").Preload("Detail_menu.Food").Where("menu_uid = ?", uid).Find(&newMenu)

		if err := res.Error; err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return newMenu, err
	}

	return newMenu, nil

}

func (mr *MenuRepository) GetAllMenu(category string, createdBy string) ([]entities.Menu, error) {
	menus := []entities.Menu{}

	if category != "" && createdBy == "" {
		res := mr.database.Preload("Detail_menu").Preload("Detail_menu.Food").Where("menu_category = ?", category).Find(&menus)
		if err := res.Error; err != nil {
			return menus, err
		}
	} else if createdBy != "" && category == "" {
		res := mr.database.Preload("Detail_menu").Preload("Detail_menu.Food").Where("created_by = ?", createdBy).Find(&menus)

		if err := res.Error; err != nil {
			return menus, err
		}
	} else if createdBy != "" && category != "" {
		res := mr.database.Preload("Detail_menu").Preload("Detail_menu.Food").Where("created_by = ? AND menu_category = ?", createdBy, category).Find(&menus)

		if err := res.Error; err != nil {
			return menus, err
		}
	} else {
		res := mr.database.Preload("Detail_menu").Preload("Detail_menu.Food").Find(&menus)

		if err := res.Error; err != nil {
			return menus, err
		}
	}

	return menus, nil
}

func (mr *MenuRepository) GetMenuRecommendGoal(user_uid string) (int, int, int, int, error) {

	var goal entities.Goal
	var user entities.User

	var breakfast int
	var lunch int
	var dinner int
	var overtime int
	var bmrDay int
	var posible int

	err := mr.database.Transaction(func(tx *gorm.DB) error {

		resGoal := tx.Model(entities.Goal{}).Where("user_uid = ? AND status =?", user_uid, "active").First(&goal)

		if err := resGoal.Error; err != nil {
			return err
		}
		resUser := tx.Model(entities.User{}).Where("user_uid = ?", user_uid).First(&user)

		if err := resUser.Error; err != nil {
			return err
		}
		needed := math.Round(float64(goal.Weight_target * 7700 / goal.Range_time))

		var bmr int
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

		if goal.Target == "gain weight" {
			bmrDay = bmr + int(needed)
			posible = bmr * 2
			if int(bmrDay) > posible {
				return errors.New("impossible")
			}
		}
		if goal.Target == "lose weight" {
			bmrDay = bmr - int(needed)
			posible = bmr * 50 / 100
			if int(bmrDay) < posible {
				return errors.New("impossible")
			}

		}

		breakfast = bmrDay * 25 / 100
		lunch = bmrDay * 35 / 100
		dinner = bmrDay * 30 / 100
		overtime = bmrDay * 10 / 100
		fmt.Println(breakfast, lunch, dinner, overtime)
		return nil
	})
	if err != nil {
		return 0, 0, 0, 0, err
	}

	return breakfast, lunch, dinner, overtime, nil
}
func (mr *MenuRepository) GetRecommendBreakfast(user_uid string) ([]entities.Menu, int64, int, error) {

	breakfast, _, _, _, err := mr.GetMenuRecommendGoal(user_uid)
	if err != nil {
		return []entities.Menu{}, 0, 0, err
	}

	menus := []entities.Menu{}
	start := breakfast * 50 / 100

	res := mr.database.Debug().Preload("Detail_menu").Preload("Detail_menu.Food").Where("menu_category=? AND created_by = ? AND total_calories BETWEEN ? AND ?", "breakfast", "admin", start, breakfast).Order("count desc").Find(&menus)

	if err := res.Error; err != nil {
		return menus, 0, 0, err
	}
	row := res.RowsAffected

	return menus, row, breakfast, nil
}
func (mr *MenuRepository) GetRecommendLunch(user_uid string) ([]entities.Menu, int64, int, error) {

	_, lunch, _, _, err := mr.GetMenuRecommendGoal(user_uid)
	if err != nil {
		return []entities.Menu{}, 0, 0, err
	}

	menus := []entities.Menu{}

	start := lunch * 50 / 100

	res := mr.database.Debug().Preload("Detail_menu").Preload("Detail_menu.Food").Where("menu_category=? AND created_by = ? AND total_calories BETWEEN ? AND ?", "lunch", "admin", start, lunch).Order("count desc").Find(&menus)

	if err := res.Error; err != nil {
		return menus, 0, 0, err
	}
	row := res.RowsAffected

	return menus, row, lunch, nil
}
func (mr *MenuRepository) GetRecommendDinner(user_uid string) ([]entities.Menu, int64, int, error) {

	_, _, dinner, _, err := mr.GetMenuRecommendGoal(user_uid)
	if err != nil {
		return []entities.Menu{}, 0, 0, err
	}

	menus := []entities.Menu{}

	start := dinner * 50 / 100

	res := mr.database.Debug().Preload("Detail_menu").Preload("Detail_menu.Food").Where("menu_category=? AND created_by = ? AND total_calories BETWEEN ? AND ?", "dinner", "admin", start, dinner).Order("count desc").Find(&menus)

	if err := res.Error; err != nil {
		return menus, 0, 0, err
	}
	row := res.RowsAffected

	return menus, row, dinner, nil
}
func (mr *MenuRepository) GetRecommendOverTime(user_uid string) ([]entities.Menu, int64, int, error) {

	_, _, _, overtime, err := mr.GetMenuRecommendGoal(user_uid)
	if err != nil {
		return []entities.Menu{}, 0, 0, err
	}

	menus := []entities.Menu{}

	res := mr.database.Preload("Detail_menu").Preload("Detail_menu.Food").Where("menu_category=? AND created_by = ? AND total_calories <= ?", "overtime", "admin", overtime).Order("count desc").Find(&menus)

	if err := res.Error; err != nil {
		return menus, 0, 0, err
	}
	row := res.RowsAffected

	return menus, row, overtime, nil
}

func (mr *MenuRepository) Update(menu_uid string, foods []entities.Food, updateMenu entities.Menu) (entities.Menu, error) {
	//code baru masih belum work
	err := mr.database.Transaction(func(tx *gorm.DB) error {
		var total_calories int

		var menu entities.Menu

		if err := tx.Model(entities.Menu{}).Where("menu_uid = ?", menu_uid).Find(&menu).Error; err != nil {
			return err
		}

		// var detail entities.Detail_menu
		if err := tx.Where("menu_uid =?", menu_uid).Delete(&entities.Detail_menu{}).Error; err != nil {
			return err
		}

		if err := tx.Where("menu_uid =?", menu_uid).Delete(&entities.Menu{}).Error; err != nil {
			return err
		}

		uid := shortuuid.New()
		updateMenu.Menu_uid = uid

		if err := tx.Create(&updateMenu).Error; err != nil {
			return err
		}
		for i := 0; i < len(foods); i++ {
			detail := entities.Detail_menu{
				Menu_uid: updateMenu.Menu_uid,
				Food_uid: foods[i].Food_uid,
			}
			if err := tx.Model(entities.Detail_menu{}).Create(&detail).Error; err != nil {
				return err
			}
			var food entities.Food
			if err := tx.Debug().Model(entities.Food{}).Where("food_uid=?", foods[i].Food_uid).First(&food).Error; err != nil {
				return err
			}
			total_calories += food.Calories
		}

		if err := tx.Model(entities.Menu{}).Where("menu_uid = ?", uid).Updates(entities.Menu{Total_calories: total_calories}).Error; err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return updateMenu, err
	}

	res := mr.database.Preload("Detail_menu").Preload("Detail_menu.Food").Where("menu_uid = ?", updateMenu.Menu_uid).First(&updateMenu)

	if err := res.Error; err != nil {
		return entities.Menu{}, err
	}

	return updateMenu, nil
}

func (mr *MenuRepository) Delete(menu_uid string) error {

	menuss := menu_uid

	var menu entities.Menu
	err := mr.database.Transaction(func(tx *gorm.DB) error {

		if err := tx.Debug().Where("menu_uid = ?", menuss).Delete(&entities.Detail_menu{}).Error; err != nil {
			return err
		}

		if err := tx.Debug().Where("menu_uid =?", menuss).Delete(&menu).Error; err != nil {
			return err
		}
		return nil

	})

	if err != nil {
		return err
	}

	return nil

}
