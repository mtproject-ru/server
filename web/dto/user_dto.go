package dto

import (
	"github.com/mtproject-ru/server/database/entities"
)

type UserResponse struct {
	Id       uint   `json:"id"`
	Username string `json:"username"`
}

// Mappers
func ToUserResponse(entity *entities.UserEntity) UserResponse {
	return UserResponse{
		Id:       entity.Id,
		Username: entity.Username,
	}
}
