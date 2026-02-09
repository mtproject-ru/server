package services

import (
	"fmt"
	"log/slog"

	"github.com/mtproject-ru/server/database"
	"github.com/mtproject-ru/server/database/entities"
)

const USER_SERVICE_KEY = "user-service"

type UserService interface {
	GetUserByUsername(username string) (*entities.UserEntity, error)
}

type UserServiceImpl struct{}

func (u *UserServiceImpl) GetUserByUsername(username string) (*entities.UserEntity, error) {
	var user entities.UserEntity

	err := database.DB.Get(&user, "SELECT * FROM users WHERE username=$1", username)
	if err != nil {
		slog.Error(err.Error())
		return nil, fmt.Errorf("Пользователь с именем %s не найден", username)
	}

	return &user, nil
}
