package middlewares

import (
	config "HealthFit/configs"
	"HealthFit/entities"
	"errors"

	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func GenerateToken(u entities.User) (string, error) {
	if u.ID == 0 {
		return "cannot Generate token", errors.New("id == 0")
	}

	codes := jwt.MapClaims{
		"user_uid": u.User_uid,
		"email":    u.Email,
		"password": u.Password,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
		"auth":     true,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, codes)
	// fmt.Println(token)
	return token.SignedString([]byte(config.JWT_SECRET))
}
func ExtractTokenUserUid(e echo.Context) string {
	user := e.Get("user").(*jwt.Token) //convert to jwt token from interface
	if user.Valid {
		codes := user.Claims.(jwt.MapClaims)
		id := codes["user_uid"].(string)
		return id
	}
	return ""
}
func GenerateTokenAdmin(a entities.Admin) (string, error) {
	if a.ID == 0 {
		return "cannot Generate token", errors.New("id == 0")
	}

	codes := jwt.MapClaims{
		"admin_uid": a.Admin_uid,
		"email":     a.Email,
		"password":  a.Password,
		"exp":       time.Now().Add(time.Hour * 24).Unix(),
		"auth":      true,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, codes)
	// fmt.Println(token)
	return token.SignedString([]byte(config.JWT_SECRET))
}

func ExtractTokenAdminUid(e echo.Context) string {
	admin := e.Get("admin").(*jwt.Token) //convert to jwt token from interface
	if admin.Valid {
		codes := admin.Claims.(jwt.MapClaims)
		id := codes["admin_uid"].(string)
		return id
	}
	return ""
}

func ExtractRoles(e echo.Context) bool {
	user := e.Get("user").(*jwt.Token) //convert to jwt token from interface
	if user.Valid {
		codes := user.Claims.(jwt.MapClaims)
		id := codes["roles"].(bool)
		return id
	}
	return false
}

// func ExtractTokenAdmin(e echo.Context) (result [2]string) {
// 	user := e.Get("user").(*jwt.Token) //convert to jwt token from interface
// 	if user.Valid {
// 		codes := user.Claims.(jwt.MapClaims)
// 		result[0] = codes["email"].(string)
// 		result[1] = codes["password"].(string)
// 		return result
// 	}
// 	return [2]string{}
// }
