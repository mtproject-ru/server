package entities

type UserEntity struct {
	Id       uint   `db:"id"`
	Username string `db:"username"`
	Password string `db:"password"`
}

// DTO
type UserResponse struct {
	Id       uint   `json:"id"`
	Username string `json:"username"`
}

// Mappers
func UserEntityToResponse(entity *UserEntity) UserResponse {
	return UserResponse{
		Id:       entity.Id,
		Username: entity.Username,
	}
}
