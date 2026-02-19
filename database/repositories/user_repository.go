package repositories

import (
	"github.com/mtproject-ru/server/database/entities"
)

type UserRepository interface {
	// CreateOrUpdate Если пользователя не существовало (id == nil), то функция создаст его.
	// Иначе обновит существующего пользователя с указанным id
	CreateOrUpdate(
		id *uint,
		username string,
		firstName string,
		lastName *string,
		password string,
	) (*entities.UserEntity, error)

	// AddEvent Добавляет указанное событие (eventId) к указанному пользоветлю (userId)
	AddEvent(userId uint, eventId uint) error

	// GetByID Возвращает пользователя с указанным айди. Если не найдено, то nil и ошибку
	GetByID(id int64) (*entities.UserEntity, error)

	// GetByUsername Возвращает пользователя с указанным юзернеймом. Если не найдено, то nil и ошибку
	GetByUsername(username string) (*entities.UserEntity, error)

	// ExistsByID Возвращает true, если пользователь с указанным айди существует, иначе false
	ExistsByID(id int64) (bool, error)

	// ExistsByUsername Возвращает true, если пользователь с указанным юзернеймом существует, иначе false
	ExistsByUsername(username string) (bool, error)

	// Delete Удаляет пользователя по указанному id
	Delete(id uint)
}

type UserRepositoryImpl struct{}
