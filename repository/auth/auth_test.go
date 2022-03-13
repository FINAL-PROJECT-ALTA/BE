package auth

import (
	"HealthFit/configs"
	// "HealthFit/delivery/middlewares"
	"HealthFit/entities"
	"HealthFit/repository/user"
	utils "HealthFit/utils/mysql"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLogin(t *testing.T) {
	config := configs.GetConfig()
	db := utils.InitDB(config)
	repo := New(db)
	db.Migrator().DropTable(&entities.User{})
	db.AutoMigrate(&entities.User{})

	t.Run("success run login", func(t *testing.T) {

		mockUser := entities.User{Name: "test", Email: "test@mail.com", Password: "test"}
		res, err := user.New(db).Register(mockUser)
		if err != nil {
			t.Fail()
		}

		// check := middlewares.CheckPasswordHash(mockUser.Password, res.Password)

		mockLogin := entities.User{Email: res.Email, Password: res.Password}
		resL, errL := repo.Login(mockLogin.Email, mockLogin.Password)

		assert.Nil(t, errL)
		assert.Equal(t, res.Email, resL.Email)
	})

	// t.Run("email not found", func(t *testing.T) {
	// 	mockUser := entities.User{Name: "test", Email: "test@mail.com", Password: "test"}
	// 	if _, err := user.New(db).Register(mockUser); err != nil {
	// 		t.Fail()
	// 	}

	// 	mockLogin := entities.User{Email: mockUser.Email, Password: mockUser.Password}
	// 	res, err := repo.Login(mockLogin.Email, mockLogin.Password)

	// 	assert.Nil(t, err)
	// 	assert.Equal(t, "test", res.Email)
	// })

}
