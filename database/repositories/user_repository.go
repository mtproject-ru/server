package repositories

import (
	"github.com/mtproject-ru/server/database"
	"github.com/mtproject-ru/server/database/entities"
)

type UserRepository interface {
	GetUserByUsername(username string) (*entities.UserEntity, error)
}

type UserRepositoryImpl struct{}

func (u *UserRepositoryImpl) GetUserByUsername(username string) (*entities.UserEntity, error) {
	var user entities.UserEntity

	err := database.DB.Get(&user, "SELECT * FROM users WHERE username=$1", username)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
