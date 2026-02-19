package repositories

import (
	"time"

	"github.com/mtproject-ru/server/database/entities"
)

type RoomRepository interface {
	// CreateOrUpdate Если комнаты не существовало (id == nil), то функция создаст его.
	// Иначе обновит существующую комнату с указанным id.
	//
	// Если при обновлении комнаты ее название не изменилось, то указывается то же название.
	//
	// Если при обновлении комнаты ее владелец не изменился, то указывается айди того же пользователя.
	CreateOrUpdate(
		id *uint,
		visibility entities.Visibility,
		roomTitle string,
		roomDescription *string,
		roomDates []time.Time,
		createdBy uint,
	) entities.EventEntity

	// AddRoomDate Добавляет указанную дату в список общих дат комнаты
	AddRoomDate(id uint, date time.Time) (*entities.RoomEntity, error)

	// AddUser Добавляет участника в комнату
	AddUser(roomId uint, userId uint) error

	// GetRoomDates Получает все даты комнаты, имеющей указанный id
	GetRoomDates(id uint) ([]time.Time, error)

	// GetByID Возвращает события с указанным айди. Если не найдено, то nil и ошибку
	GetByID(id int64) (*entities.RoomEntity, error)

	// Delete Удаляет событие по указанному id
	Delete(id uint)
}

type RoomRepositoryImpl struct{}
