package auth

import (
	"HealthFit/configs"
	"HealthFit/delivery/middlewares"
	"HealthFit/entities"
	gr "HealthFit/repository/goal"
	"HealthFit/repository/user"
	utils "HealthFit/utils/mysql"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestLogin(t *testing.T) {
	config := configs.GetConfig()
	db := utils.InitDB(config)
	repo := New(db)

	t.Run("invalid email", func(t *testing.T) {
		db.Migrator().DropTable(&entities.User{})
		db.AutoMigrate(&entities.User{})

		mockUser := entities.User{Name: "test", Email: "test@mail.com", Password: "test"}
		res, err := user.New(db).Register(mockUser)
		if err != nil {
			t.Fail()
		}

		mockLogin := entities.User{Email: "dordd", Password: res.Password}
		_, errL := repo.Login(mockLogin.Email, mockLogin.Password)

		assert.NotNil(t, errL)
	})

	t.Run("incorrect password", func(t *testing.T) {
		db.Migrator().DropTable(&entities.User{})
		db.AutoMigrate(&entities.User{})

		mockUser := entities.User{Name: "test", Email: "test@mail.com", Password: "test"}
		res, err := user.New(db).Register(mockUser)
		if err != nil {
			t.Fail()
		}

		mockLogin := entities.User{Email: res.Email, Password: res.Password}
		_, errL := repo.Login(mockLogin.Email, mockLogin.Password)

		assert.NotNil(t, errL)
	})

	t.Run("roles error", func(t *testing.T) {
		db.Migrator().DropTable(&entities.User{})
		db.AutoMigrate(&entities.User{})

		mockUser := entities.User{Name: "test", Email: "test@mail.com", Password: "test"}
		res, err := user.New(db).Register(mockUser)
		if err != nil {
			t.Fail()
		}

		resPass := middlewares.CheckPasswordHash("test", res.Password)
		if resPass == false {
			t.Fail()
		}

		var pass string = "$2a$14$LArsepWalaQWifUH7x7uyecACDMg6w8a6k/LmgGg5BQy79/VAwREq"

		mockLogin := entities.User{Email: res.Email, Password: pass}
		_, errL := repo.Login(mockLogin.Email, mockLogin.Password)

		assert.NotNil(t, errL)
	})

	t.Run("success", func(t *testing.T) {
		db.Migrator().DropTable(&entities.User{}, &entities.Goal{})
		db.AutoMigrate(&entities.User{}, &entities.Goal{})

		mockUser := entities.User{Name: "test", Email: "test@mail.com", Password: "test"}
		res, err := user.New(db).Register(mockUser)
		if err != nil {
			t.Fail()
		}

		pas, _ := middlewares.HashPassword(res.Password)

		resPass := middlewares.CheckPasswordHash(mockUser.Password, pas)
		if resPass == false {
			t.Fail()
		}

		mockGoal := entities.Goal{
			User_uid:      res.User_uid,
			Height:        150,
			Weight:        55,
			Age:           24,
			Daily_active:  "not active",
			Weight_target: 50,
			Range_time:    14,
			Target:        "lose weight",
		}

		resG, errG := gr.New(db).Create(mockGoal)
		if errG != nil {
			t.Fail()
		}

		_, errRef := repo.RefreshGoalAuth(resG.User_uid)
		if errRef != nil {
			t.Fail()
		}

		mockLogin := entities.User{Email: res.Email, Password: res.Password}
		_, errL := repo.Login(mockLogin.Email, mockLogin.Password)

		assert.NotNil(t, errL)
	})

	//new
	t.Run("roles user not have goal", func(t *testing.T) {
		db.Migrator().DropTable(&entities.User{})
		db.AutoMigrate(&entities.User{})

		mockUser := entities.User{Name: "test", Email: "test@mail.com", Password: "test", Roles: false}
		res, err := user.New(db).Register(mockUser)
		if err != nil {
			t.Fail()
		}

		resPass := middlewares.CheckPasswordHash("test", res.Password)
		if resPass == false {
			t.Fail()
		}

		mockLogin := entities.User{Email: res.Email, Password: res.Password}
		resLogin, errL := repo.Login(mockLogin.Email, mockUser.Password)

		assert.Nil(t, errL)
		assert.Equal(t, false, resLogin.Goal_active)
	})
	t.Run("roles user have goal", func(t *testing.T) {
		db.Migrator().DropTable(&entities.Goal{})
		db.Migrator().DropTable(&entities.User{})
		db.AutoMigrate(&entities.User{})
		db.AutoMigrate(&entities.Goal{})

		mockUser := entities.User{Name: "test", Email: "test@mail.com", Password: "test", Gender: "male", Roles: false}
		res, err := user.New(db).Register(mockUser)
		if err != nil {
			t.Fail()
		}
		mockGoal := entities.Goal{
			User_uid:      res.User_uid,
			Height:        150,
			Weight:        55,
			Age:           24,
			Daily_active:  "not active",
			Weight_target: 2,
			Range_time:    30,
			Target:        "lose weight",
			Status:        "active",
		}

		_, errG := gr.New(db).Create(mockGoal)
		if errG != nil {
			t.Fail()
		}

		resPass := middlewares.CheckPasswordHash("test", res.Password)
		if resPass == false {
			t.Fail()
		}

		mockLogin := entities.User{Email: res.Email, Password: res.Password}
		resLogin, errL := repo.Login(mockLogin.Email, mockUser.Password)

		assert.Nil(t, errL)
		assert.Equal(t, true, resLogin.Goal_active)
		assert.Equal(t, false, resLogin.Goal_exspired)
	})
	t.Run("roles user have goal exspired", func(t *testing.T) {
		db.Migrator().DropTable(&entities.Goal{})
		db.Migrator().DropTable(&entities.User{})
		db.AutoMigrate(&entities.User{})
		db.AutoMigrate(&entities.Goal{})

		mockUser := entities.User{Name: "test", Email: "test@mail.com", Password: "test", Gender: "male", Roles: false}
		res, err := user.New(db).Register(mockUser)
		if err != nil {
			t.Fail()
		}
		february := time.Date(2022, time.February, 11, 0, 0, 0, 0, time.UTC)
		// var timer time.Time = "2022-03-13 00:48:23.026"
		mockGoal := entities.Goal{
			User_uid:      res.User_uid,
			Height:        150,
			Weight:        55,
			Age:           24,
			Daily_active:  "not active",
			Weight_target: 2,
			Range_time:    30,
			Target:        "lose weight",
			Status:        "active",
			CreatedAt:     february,
		}

		_, errG := gr.New(db).Create(mockGoal)
		if errG != nil {
			t.Fail()
		}

		resPass := middlewares.CheckPasswordHash("test", res.Password)
		if resPass == false {
			t.Fail()
		}

		mockLogin := entities.User{Email: res.Email, Password: res.Password}
		resLogin, errL := repo.Login(mockLogin.Email, mockUser.Password)

		assert.Nil(t, errL)
		assert.Equal(t, false, resLogin.Goal_active)
		assert.Equal(t, true, resLogin.Goal_exspired)
	})
}
