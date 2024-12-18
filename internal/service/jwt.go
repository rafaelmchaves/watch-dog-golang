package service

import (
	"time"

	"github.com/golang-jwt/jwt"
	"watchdog-go.com/internal/infrastructure"
)

func GenerateToken(user infrastructure.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":    user.Email,
		"id":       user.Id,
		"name":     user.Name,
		"phone":    user.Phone,
		"city":     user.City,
		"state":    user.State,
		"fields":   user.Fields,
		"exp":      time.Now().Add(time.Minute * 40).Unix(),
		"userType": user.UserType,
	})

	return token.SignedString([]byte("SECRET"))
}
