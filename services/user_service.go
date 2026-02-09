package services

import (
	"fmt"

	"github.com/mtproject-ru/server/database/entities"
	"github.com/mtproject-ru/server/database/repositories"
)

const USER_SERVICE_KEY = "user-service"

type UserService interface {
	GetUserByUsername(username string) (*entities.UserEntity, error)
}

type UserServiceImpl struct {
	repo repositories.UserRepository
}

func (u *UserServiceImpl) GetUserByUsername(username string) (*entities.UserEntity, error) {
	user, err := u.repo.GetUserByUsername(username)
	if err != nil {
		return nil, fmt.Errorf("Пользователь с именем %s не найден", username)
	}
	return user, nil
}
