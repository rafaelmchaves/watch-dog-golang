package service

import (
	"errors"
	"log"

	"watchdog-go.com/internal/infrastructure"
)

type LoginService interface {
	Login(login string, password string) (string, error)
}

type LoginServiceImpl struct{}

func (loginService *LoginServiceImpl) Login(login string, password string) (string, error) {
	user := infrastructure.CheckLogin(login, password)

	if user == nil || user.Id == "" {
		return "", errors.New("user not found")
	}

	token, err := GenerateToken(*user)
	if err != nil {
		log.Fatal("Login failed - ", err)
		return "", err
	}

	return token, nil
}
