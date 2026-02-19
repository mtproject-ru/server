package repositories

import (
	"time"

	"github.com/mtproject-ru/server/database/entities"
)

type EventRepository interface {
	// CreateOrUpdate Если события не существовало (id == nil), то функция создаст его.
	// Иначе обновит существующее событие с указанным id
	CreateOrUpdate(
		id *uint,
		visibility entities.Visibility,
		eventTitle string,
		eventDescription *string,
		eventStart *time.Time,
		eventEnd *time.Time,
	) entities.EventEntity

	// GetByID Возвращает события с указанным айди. Если не найдено, то nil и ошибку
	GetByID(id int64) (*entities.EventEntity, error)

	// GetByEventCreatedAt Возвращает события, созданные в указанное время. Если не найдено, то nil и ошибку
	GetByEventCreatedAt(createdAt time.Time) ([]entities.EventEntity, error)

	// GetByEventStart Получает все события, начинающиеся в указанную дату
	GetByEventStart(eventStart time.Time) ([]entities.EventEntity, error)

	// Delete Удаляет событие по указанному id
	Delete(id uint)
}

type EventRepositoryImpl struct{}
